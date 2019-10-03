package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/atomic/atr/configs"
	"github.com/atomic/atr/library/faker"
	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/app/api/user/mocks"
	ucase "github.com/atomic/atr/src/app/api/user/usecase"
)

func TestFindAllWithNoParams(t *testing.T) {

	db1, _, _, _ := configs.DBInitTest()

	mockUserRepo := new(mocks.Repository)

	mockListUser := make([]*models.User, 0)

	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockListUser = append(mockListUser, &mockUser)

	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockListUser = append(mockListUser, &mockUser)

	mockUser = models.User{
		ID:                1,
		Username:          "usersip",
		Password:          "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==",
		Name:              "atr",
		Email:             "atr@gmail.com",
		Phone:             "080989999",
		Pin:               "123456",
		LoginType:         "Reguler",
		UID:               "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1",
		Type:              "Sales",
		Location:          "-6.177949, 106.766092",
		Longitude:         "-6.175110",
		Latitude:          "106.865036",
		SessionToken:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE",
		DeviceToken:       "-",
		NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT",
		RefreshToken:      "-",
		Description:       "Description User Terbaru",
		OutletID:          1,
		BusinessID:        1,
		RoleID:            1,
	}

	mockListUser = append(mockListUser, &mockUser)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindAll", mock.Anything).Return(mockListUser, nil).Once()

		u := ucase.NewUserUsecase(db1, mockUserRepo)

		list, err := u.FindAll("-1", "", "", "-1")
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListUser))

		mockUserRepo.AssertExpectations(t)

	})

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindAll", mock.Anything).Return(mockListUser, nil).Once()

		u := ucase.NewUserUsecase(db1, mockUserRepo)

		list, err := u.FindAll("1", "10", "User Terbaru", "-1")
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListUser))

		mockUserRepo.AssertExpectations(t)

	})

	t.Run("error-failed", func(t *testing.T) {

	})
}

func TestFindAllWithParams(t *testing.T) {

	db1, _, _, _ := configs.DBInitTest()

	mockUserRepo := new(mocks.Repository)

	mockListUser := make([]*models.User, 0)

	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockListUser = append(mockListUser, &mockUser)

	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockListUser = append(mockListUser, &mockUser)

	mockUser = models.User{
		ID:                1,
		Username:          "usersip",
		Password:          "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==",
		Name:              "atr",
		Email:             "atr@gmail.com",
		Phone:             "080989999",
		Pin:               "123456",
		LoginType:         "Reguler",
		UID:               "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1",
		Type:              "Sales",
		Location:          "-6.177949, 106.766092",
		Longitude:         "-6.175110",
		Latitude:          "106.865036",
		SessionToken:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE",
		DeviceToken:       "-",
		NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT",
		RefreshToken:      "-",
		Description:       "Description User Terbaru",
		OutletID:          1,
		BusinessID:        1,
		RoleID:            1,
	}

	mockListUser = append(mockListUser, &mockUser)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindAll", mock.Anything).Return(mockListUser, nil).Once()

		u := ucase.NewUserUsecase(db1, mockUserRepo)

		list, err := u.FindAll("1", "10", "User Terbaru", "-1")

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListUser))

		mockUserRepo.AssertExpectations(t)

	})

	t.Run("error-failed", func(t *testing.T) {

	})
}

func TestFind(t *testing.T) {

	db1, _, _, _ := configs.DBInitTest()

	mockUserRepo := new(mocks.Repository)
	mockUser := &models.User{
		ID:                1,
		Username:          "usersip",
		Password:          "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==",
		Name:              "atr",
		Email:             "atr@gmail.com",
		Phone:             "080989999",
		Pin:               "123456",
		LoginType:         "Reguler",
		UID:               "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1",
		Type:              "Sales",
		Location:          "-6.177949, 106.766092",
		Longitude:         "-6.175110",
		Latitude:          "106.865036",
		SessionToken:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE",
		DeviceToken:       "-",
		NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT",
		RefreshToken:      "-",
		Description:       "Description User Terbaru",
		OutletID:          1,
		BusinessID:        1,
		RoleID:            1,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Find", mock.Anything).Return(mockUser, nil).Once()

		u := ucase.NewUserUsecase(db1, mockUserRepo)

		data, err := u.Find("1")
		assert.NoError(t, err)
		assert := assert.New(t)
		assert.Equal(mockUser.Username, data.Username, "Username doesn't match Found: "+data.Username)

		mockUserRepo.AssertExpectations(t)

	})

	t.Run("error-failed", func(t *testing.T) {

	})
}

// func TestLogin(t *testing.T) {

// 	db1, _, _, _ := configs.DBInitTest()

// 	mockUserRepo := new(mocks.Repository)
// 	mockUser := &models.User{
// 		ID:                1,
// 		Username:          "usersip",
// 		Password:          "qUX_wTAvSoZ4xP1mFNT_lWwrQ4DQb5CT3A3NK8QF49AFfcIU1ks2n1zhpIxa4DNY8QDM2pIOhbVc339V7wKiCg==",
// 		Name:              "atr",
// 		Email:             "atr@gmail.com",
// 		Phone:             "080989999",
// 		Pin:               "123456",
// 		LoginType:         "Reguler",
// 		UID:               "EkJ8L8LAE7MPXBsfRQTJqIVEaGy1",
// 		Type:              "Sales",
// 		Location:          "-6.177949, 106.766092",
// 		Longitude:         "-6.175110",
// 		Latitude:          "106.865036",
// 		SessionToken:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlemFAYXRvbWljLmlkIiwiaWQiOjEsInR5cGUiOiJCdXNpbmVzcyJ9.7o7vpO94Q8LzXT8qiqmp_j4CB61LEG0HMdxpZL-WKFE",
// 		DeviceToken:       "-",
// 		NotificationToken: "cghz2S0Yj7w:APA91bEeNQ18hLj4T6YlK8gN--skGKiH4PMdETxoCNyznzIpD9rZ29USY4GbHCTI3Zm6TeR4GsQlS2xQsk06Q16UZZmJGji4VcUJovl1leDdACQLCmhDmWoaPD8_Zhz-TYAIjbEk8daT",
// 		RefreshToken:      "-",
// 		Description:       "Description User Terbaru",
// 		OutletID:          1,
// 		BusinessID:        1,
// 		RoleID:            1,
// 	}

// 	mockFindAllParam := &helpers.FindAllParams{
// 		Page:     "-1",
// 		Size:     "10",
// 		Keyword:  "-",
// 		StatusID: "1",
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("Login", mock.Anything).Return(mockUser, mockFindAllParam, nil).Once()

// 		u := ucase.NewUserUsecase(db1, mockUserRepo)

// 		data, _, err := u.Login("usersip", "password", "awd", "awd")
// 		assert.NoError(t, err)
// 		assert := assert.New(t)
// 		assert.Equal(mockUser.Username, data.Username, "wrong username or password")

// 		mockUserRepo.AssertExpectations(t)

// 	})

// 	t.Run("error-failed", func(t *testing.T) {

// 	})
// }
