package repository_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/atomic/atr/configs"
	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/app/api/user/repository"
	"github.com/atomic/atr/src/helpers"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type UserSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository repository.UserRepository
	user       *models.User

	VariableThatShouldStartAtFive int
}

func (s *UserSuite) SetupTest() {

	db1, mock, errmock, errconnection := configs.DBInitTest()
	require.NoError(s.T(), errmock)
	require.NoError(s.T(), errconnection)

	s.mock = mock
	s.DB = db1.DB

	s.repository = repository.NewUserRepository(db1)
}

func (suite *UserSuite) TestFindAll() {
	mockDatas := []models.User{
		models.User{
			ID: 1, Username: "usersip", Password: "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==", Name: "atr", Email: "atr@gmail.com", Phone: "080989999", Pin: "123456", LoginType: "Reguler", UID: "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1", Type: "Sales", Location: "-6.177949, 106.766092", Longitude: "-6.175110", Latitude: "106.865036", SessionToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE", DeviceToken: "-", NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT", RefreshToken: "-", Description: "Description User", BusinessID: 1, RoleID: 1,
			UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
		models.User{
			ID: 2, Username: "usersipdua", Password: "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==", Name: "atr 2", Email: "sip2@gmail.com", Phone: "080989991", Pin: "123451", LoginType: "Reguler", UID: "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1", Type: "Sales", Location: "-6.177949, 106.766092", Longitude: "-6.175110", Latitude: "106.865036", SessionToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE", DeviceToken: "-", NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT", RefreshToken: "-", Description: "Description User 2", BusinessID: 1, RoleID: 1,
			UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "username", "password", "name", "email", "phone", "pin", "type_login", "UID", "type", "location", "longitude", "latitude", "session_token", "device_token", "notification_token", "refresh_token", "description", "business_id", "role_id", "updated_at", "created_at"}).
		AddRow(mockDatas[0].ID, mockDatas[0].Username, mockDatas[0].Password, mockDatas[0].Name, mockDatas[0].Email, mockDatas[0].Phone, mockDatas[0].Pin, mockDatas[0].LoginType, mockDatas[0].UID, mockDatas[0].Type, mockDatas[0].Location, mockDatas[0].Longitude, mockDatas[0].Latitude, mockDatas[0].SessionToken, mockDatas[0].DeviceToken, mockDatas[0].NotificationToken, mockDatas[0].RefreshToken, mockDatas[0].Description, mockDatas[0].BusinessID, mockDatas[0].RoleID,
			mockDatas[0].UpdatedAt, mockDatas[0].CreatedAt).
		AddRow(mockDatas[1].ID, mockDatas[1].Username, mockDatas[1].Password, mockDatas[1].Name, mockDatas[1].Email, mockDatas[1].Phone, mockDatas[1].Pin, mockDatas[1].LoginType, mockDatas[1].UID, mockDatas[1].Type, mockDatas[1].Location, mockDatas[1].Longitude, mockDatas[1].Latitude, mockDatas[1].SessionToken, mockDatas[1].DeviceToken, mockDatas[1].NotificationToken, mockDatas[0].RefreshToken, mockDatas[1].Description, mockDatas[1].BusinessID, mockDatas[1].RoleID,
			mockDatas[1].UpdatedAt, mockDatas[1].CreatedAt)

	query := regexp.QuoteMeta("SELECT * FROM `users`")

	suite.mock.ExpectQuery(query).WillReturnRows(rows)

	params := helpers.FindAllParams{Page: "-1", Size: "", Keyword: "", StatusID: ""}

	res, _, err := suite.repository.FindAll(params)

	require.NoError(suite.T(), err)

	assert := assert.New(suite.T())
	assert.Equal(2, len(res), "Total row doesn't match Found: ")
}

func (suite *UserSuite) TestFind() {
	mockData := models.User{
		ID: 1, Username: "usersip", Password: "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==", Name: "atr", Email: "atr@gmail.com", Phone: "080989999", Pin: "123456", LoginType: "Reguler", UID: "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1", Type: "Sales", Location: "-6.177949, 106.766092", Longitude: "-6.175110", Latitude: "106.865036", SessionToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE", DeviceToken: "-", NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT", RefreshToken: "-", Description: "Description User", BusinessID: 1, RoleID: 1,
		UpdatedAt: time.Now(), CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"ID", "username", "password", "name", "email", "phone", "pin", "type_login", "UID", "type", "location", "longitude", "latitude", "session_token", "device_token", "notification_token", "refresh_token", "description", "business_id", "role_id", "updated_at", "created_at"}).AddRow(mockData.ID, mockData.Username, mockData.Password, mockData.Name, mockData.Email, mockData.Phone, mockData.Pin, mockData.LoginType, mockData.UID, mockData.Type, mockData.Location, mockData.Longitude, mockData.Latitude, mockData.SessionToken, mockData.DeviceToken, mockData.NotificationToken, mockData.RefreshToken, mockData.Description, mockData.BusinessID, mockData.RoleID, mockData.CreatedAt, mockData.UpdatedAt)

	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE (id = ?) ORDER BY `users`.`id` ASC LIMIT 1")

	str_id := fmt.Sprint(mockData.ID)

	suite.mock.ExpectQuery(query).WithArgs(str_id).WillReturnRows(rows)

	res, err := suite.repository.Find(str_id)

	require.NoError(suite.T(), err)

	assert := assert.New(suite.T())
	assert.Equal(mockData.Username, res.Username, "Username doesn't match Found: "+res.Username)
}

func (suite *UserSuite) TestCreate() {
	mockData := models.User{
		Code: "USR0000000001", Username: "usersip", Password: "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==", Name: "atr", Email: "atr@gmail.com", Phone: "080989999", Pin: "123456", LoginType: "Reguler", UID: "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1", Type: "Sales", Location: "-6.177949, 106.766092", Longitude: "-6.175110", Latitude: "106.865036", SessionToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE", DeviceToken: "-", NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT", RefreshToken: "-", Description: "Description User", BusinessID: 1, RoleID: 1,
		UpdatedAt: time.Now(), CreatedAt: time.Now(),
	}

	rows := sqlmock.NewResult(1, 1)

	query := regexp.QuoteMeta("INSERT INTO `users` (`code`,`username`,`password`,`name`,`email`,`phone`,`pin`,`login_type`,`uid`,`type`,`location`,`longitude`,`latitude`,`session_token`,`device_token`,`notification_token`,`refresh_token`,`description`,`business_id`,`role_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

	suite.mock.ExpectExec(query).WithArgs(mockData.Code, mockData.Username, mockData.Password, mockData.Name, mockData.Email, mockData.Phone, mockData.Pin, mockData.LoginType, mockData.UID, mockData.Type, mockData.Location, mockData.Longitude, mockData.Latitude, mockData.SessionToken, mockData.DeviceToken, mockData.NotificationToken, mockData.RefreshToken, mockData.Description, mockData.BusinessID, mockData.RoleID, mockData.CreatedAt, mockData.UpdatedAt).
		WillReturnResult(rows)

	res, err := suite.repository.Create(mockData)

	require.NoError(suite.T(), err)

	assert := assert.New(suite.T())
	assert.Equal(mockData.Code, res.Code, "Code doesn't match Found: "+res.Code)
	assert.Equal(mockData.Username, res.Username, "Username doesn't match Found: "+res.Username)
	assert.Equal(mockData.Password, res.Password, "Password doesn't match Found: "+res.Password)
	assert.Equal(mockData.Name, res.Name, "Name doesn't match Found: "+res.Name)
	assert.Equal(mockData.Email, res.Email, "Email doesn't match Found: "+res.Email)
	assert.Equal(mockData.Phone, res.Phone, "Phone doesn't match Found: "+res.Phone)
	assert.Equal(mockData.Pin, res.Pin, "Pin doesn't match Found: "+res.Pin)
	assert.Equal(mockData.LoginType, res.LoginType, "LoginType doesn't match Found: "+res.LoginType)
	assert.Equal(mockData.UID, res.UID, "UID doesn't match Found: "+res.UID)
	assert.Equal(mockData.Type, res.Type, "Type doesn't match Found: "+res.Type)
	assert.Equal(mockData.Location, res.Location, "Location doesn't match Found: "+res.Location)
	assert.Equal(mockData.Longitude, res.Longitude, "Longitude doesn't match Found: "+res.Longitude)
	assert.Equal(mockData.Latitude, res.Latitude, "Latitude doesn't match Found: "+res.Latitude)
	assert.Equal(mockData.SessionToken, res.SessionToken, "SessionToken doesn't match Found: "+res.SessionToken)
	assert.Equal(mockData.DeviceToken, res.DeviceToken, "DeviceToken doesn't match Found: "+res.DeviceToken)
	assert.Equal(mockData.RefreshToken, res.RefreshToken, "RefreshToken doesn't match Found: "+res.RefreshToken)
	assert.Equal(mockData.NotificationToken, res.NotificationToken, "NotificationToken doesn't match Found: "+res.NotificationToken)
	assert.Equal(mockData.Description, res.Description, "Description doesn't match Found: "+res.Description)
	assert.Equal(mockData.BusinessID, int(res.BusinessID), "OutletID doesn't match Found: "+string(res.BusinessID))
	assert.Equal(mockData.RoleID, int(res.RoleID), "RoleID doesn't match Found: "+string(res.RoleID))

}

func (suite *UserSuite) TestUpdate() {
	mockData := models.User{
		ID: 1, Username: "usersip", Password: "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==", Name: "atr", Email: "atr@gmail.com", Phone: "080989999", Pin: "123456", LoginType: "Reguler", UID: "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1", Type: "Sales", Location: "-6.177949, 106.766092", Longitude: "-6.175110", Latitude: "106.865036", SessionToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE", DeviceToken: "-", NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT", RefreshToken: "-", Description: "Description User", BusinessID: 1, RoleID: 1,
		UpdatedAt: time.Now(), CreatedAt: time.Now(),
	}

	str_id := fmt.Sprint(mockData.ID)

	//UPDATE
	mockDataNew := models.User{
		Username: "usersip", Password: "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==", Name: "atr", Email: "atr@gmail.com", Phone: "080989999", Pin: "123456", LoginType: "Reguler", UID: "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1", Type: "Sales", Location: "-6.177949, 106.766092", Longitude: "-6.175110", Latitude: "106.865036", SessionToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE", DeviceToken: "-", NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT", RefreshToken: "-", Description: "Description User", BusinessID: 1, RoleID: 1,
		UpdatedAt: time.Now(),
	}

	query_u := regexp.QuoteMeta("UPDATE `users`")
	rows_u := sqlmock.NewResult(1, 1)
	suite.mock.ExpectExec(query_u).
		WillReturnResult(rows_u)

	_, err := suite.repository.Update(str_id, mockDataNew)

	require.NoError(suite.T(), err)
}

func (suite *UserSuite) TestLogin() {
	mockData := models.User{
		ID: 1, Username: "usersip", Password: "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==", Name: "atr", Email: "atr@gmail.com", Phone: "080989999", Pin: "123456", LoginType: "Reguler", UID: "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1", Type: "Sales", Location: "-6.177949, 106.766092", Longitude: "-6.175110", Latitude: "106.865036", SessionToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE", DeviceToken: "-", NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT", RefreshToken: "-", Description: "Description User", BusinessID: 1, RoleID: 1,
		UpdatedAt: time.Now(), CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"ID", "username", "password", "name", "email", "phone", "pin", "type_login", "UID", "type", "location", "longitude", "latitude", "session_token", "device_token", "notification_token", "refresh_token", "description", "business_id", "role_id", "updated_at", "created_at"}).AddRow(mockData.ID, mockData.Username, mockData.Password, mockData.Name, mockData.Email, mockData.Phone, mockData.Pin, mockData.LoginType, mockData.UID, mockData.Type, mockData.Location, mockData.Longitude, mockData.Latitude, mockData.SessionToken, mockData.DeviceToken, mockData.NotificationToken, mockData.RefreshToken, mockData.Description, mockData.BusinessID, mockData.RoleID, mockData.CreatedAt, mockData.UpdatedAt)

	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE ((email=?) OR (phone=?) OR (username=?)) AND (password=?)")

	suite.mock.ExpectQuery(query).WithArgs(mockData.Username, "password", mockData.Username, mockData.Username).WillReturnRows(rows)

	_, _, err := suite.repository.Login(mockData.Username, "password", mockData.DeviceToken, mockData.NotificationToken)

	require.NoError(suite.T(), err)
}

func (suite *UserSuite) TestDelete() {
	mockData := models.User{
		ID: 1, Username: "usersip", Password: "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==", Name: "atr", Email: "atr@gmail.com", Phone: "080989999", Pin: "123456", LoginType: "Reguler", UID: "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1", Type: "Sales", Location: "-6.177949, 106.766092", Longitude: "-6.175110", Latitude: "106.865036", SessionToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE", DeviceToken: "-", NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT", RefreshToken: "-", Description: "Description User", BusinessID: 1, RoleID: 1,
		UpdatedAt: time.Now(), CreatedAt: time.Now(),
	}

	str_id := fmt.Sprint(mockData.ID)

	query_d := regexp.QuoteMeta("UPDATE `users`")
	rows_d := sqlmock.NewResult(1, 1)
	suite.mock.ExpectExec(query_d). //WithArgs(time.Now(), str_id).
					WillReturnResult(rows_d)

	_, err := suite.repository.Delete(str_id)
	fmt.Println(err)
	// require.NoError(suite.T(), err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRunAll(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
