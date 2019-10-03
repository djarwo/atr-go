package helpers

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/atomic/atr/configs"
	"github.com/atomic/atr/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/leekchan/accounting"
	"gopkg.in/go-playground/validator.v8"
)

type Result struct {
	Status     string
	Message    string
	StatusCode int
	Data       interface{}
}

type BasicTokenEnv struct {
	Type      string
	Key       string
	SecretKey string
}

type SlackRequestBody struct {
	Text    string `json:"text"`
	Status  uint16 `json:"Status"`
	Type    string `json:"Type"`
	Package string `json:"Package"`
	Service string `json:"Service"`
}

type Status struct {
	Name string
	ID   uint
}

type StatusClient struct {
	Err        error
	Status     int
	Permission interface{} //models.Permission
	Jwt        string
	IP         [][]string
}

type LastID struct {
	ID uint64
}

type buffer struct {
	r         []byte
	runeBytes [utf8.UTFMax]byte
}

// result all
type ResultAll struct {
	Status     string
	Message    string
	StatusCode int
	TotalData  int
	Page       string
	Size       string
	Data       interface{}
}

// return in trip
type ReturnInTrip struct {
	Data       interface{}
	Message    string
	StatusCode int
	Error      error
}

// month
type MonthInv struct {
	StartDay time.Time
	EndDay   time.Time
	Month    time.Month
	Week     int
}

// day in week
type DayWeekInv struct {
	DayInt    int
	DayString time.Weekday
	Week      int
	WeekMod   int
	Year      int
	Date      time.Time
}

// table migrate
type TableMigrate struct {
	ModelName string
	TableName string
	Table     interface{}
}

const (
	serverKeyPushNotification = "AAAAZvT7Vs0:APA91bFs6wlz6vyM5GksKZ9Jdd00qrw4QrLVApsI9vdvaUoAFKwHR6Xszc_z1XQIabeZFPK5Ic0MUnttd2Ht3i0VPDRgK3IJmhl38762Cg7oFDbd1F659XYAukLqHE6BFOW4fF1nofSK"
	passwordSalt              = "a99VVoWzmd1C9ujcitK0fIVNE0I5I61AC47C852RoLTsHDyLCltvP+ZHEkIl/2hkzTOW90c3ZEjtYRkdfTWJ1Q=="
)

func DefaultRepo() ReturnRepo {
	returnRepo := ReturnRepo{200, "Suksess", nil, nil}
	return returnRepo
}

func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	return BeginningOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

func monthInterval(y int, m time.Month) (firstDay, lastDay time.Time) {
	firstDay = time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	lastDay = time.Date(y, m+1, 1, 0, 0, 0, -1, time.UTC)
	return firstDay, lastDay
}

func GetWeek() interface{} {
	var arr []interface{}
	month := 12
	y, m, _ := time.Now().Date()
	for i := 1; i <= month; i++ {
		yArr, mArr := y, time.Month(i)
		first, last := monthInterval(yArr, mArr)
		_, week := time.Now().ISOWeek()
		MonthInv := MonthInv{EndDay: last, StartDay: first, Month: m, Week: week}
		arr = append(arr, MonthInv)
	}

	return arr
}

func MessageErr(err error) string {
	if err != nil {
		return err.Error()
	} else {
		return "Record Not Found"
	}
}

func GetDay(date time.Time) (int, time.Weekday) {
	day := date.Weekday()
	return int(day), day
}

func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func DayWeekArr(datas []DayWeekInv) ([]int, []int) {
	var days []int
	var weeks []int
	for _, v := range datas {
		days = append(days, v.DayInt)
		weeks = append(days, v.WeekMod)
	}

	return days, weeks
}

func ConvertRupiah(value int) string {
	ac := accounting.Accounting{Symbol: "Rp. ", Precision: 2}

	Strings := ac.FormatMoney(value)

	return Strings
}

func ParseInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		logrus.Error(err)
	}
	return i
}

func FormValidation(err error) interface{} {
	var response interface{}
	var message string
	status := "Sukses"
	statusCode := http.StatusOK
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err2 := range castedObject {
			switch err2.Tag {
			case "required":
				message = fmt.Sprintf("%s is required",
					err2.Field)
				status = "Warning"
				statusCode = http.StatusInternalServerError
			case "email":
				message = fmt.Sprintf("%s is not valid email",
					err2.Field)
				status = "Warning"
				statusCode = http.StatusInternalServerError
			case "gte":
				message = fmt.Sprintf("%s value must be greater than %s",
					err2.Field, err2.Param)
				status = "Warning"
				statusCode = http.StatusInternalServerError
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s",
					err2.Field, err2.Param)
				status = "Warning"
				statusCode = http.StatusInternalServerError
			case "max":
				message = fmt.Sprintf("%s value cannot be longer than %s",
					err2.Field, err2.Param)
				status = "Warning"
				statusCode = http.StatusInternalServerError
			case "min":
				message = fmt.Sprintf("%s value must be longer than %s",
					err2.Field, err2.Param)
				status = "Warning"
				statusCode = http.StatusInternalServerError
			case "len":
				message = fmt.Sprintf("%s value must be %s characters",
					err2.Field, err2.Param)
				status = "Warning"
				statusCode = http.StatusInternalServerError
			case "uuid":
				message = fmt.Sprintf("%s value already uuid",
					err2.Field, err2.Param)
				status = "Warning"
				statusCode = http.StatusInternalServerError
			case "unique":
				message = fmt.Sprintf("%s value already uuid",
					err2.Field, err2.Param)
				status = "Warning"
				statusCode = http.StatusInternalServerError
			case "exist":
				message = fmt.Sprintf("%s value already exists",
					err2.Field, err2.Param)
				status = "Warning"
				statusCode = http.StatusInternalServerError
			}
			break
		}
	}

	response = Result{Status: status, StatusCode: statusCode, Message: message}

	result := gin.H{
		"result": response,
	}

	return result
}

func PasswordHasher(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return pwd
}

func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str)+1 > lenght {
			return str[0:lenght]
		}
	}
}

func CodeGenerator(table string, prefix string) string {
	finalcode := prefix

	return finalcode
}

func CodeClientExternal() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 60)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	rand.Seed(time.Now().UnixNano())
	min := 1001
	max := 9999
	random := rand.Intn(max-min) + min
	time := strconv.Itoa(random)
	concat := time + string(b)
	return concat
}

func UniqueCode() string {

	var prefix string
	month := PadNumberWithZero(uint32(time.Now().Month()))
	prefix = strconv.Itoa(time.Now().Year())[2:] + month

	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 6)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	concat := prefix + string(b)

	return concat
}

func PadNumberWithZero(value uint32) string {
	return fmt.Sprintf("%02d", value)
}

func SubStrPermission(handlers string, start string, mid string, end string) (string, string) {
	startIndex := strings.Index(handlers, start) + 2
	centerIndex := strings.Index(handlers, mid)
	endIndex := strings.Index(handlers, end)

	handler := string([]rune(handlers)[startIndex:centerIndex])
	baseHandler := string([]rune(handlers)[centerIndex+2 : endIndex])

	return handler, baseHandler
}

func SubStr(str string, start string, end string) string {

	endIndex := strings.Index(str, end)
	String := string([]rune(str)[1:endIndex])

	return String
}

func Split(src string) (entries []string) {
	if !utf8.ValidString(src) {
		return []string{src}
	}
	entries = []string{}
	var runes [][]rune
	lastClass := 0
	class := 0

	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}

	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}

	for _, s := range runes {
		if len(s) > 0 {
			entries = append(entries, string(s))
		}
	}
	return
}

func FilterFindAll(c *gin.Context) (string, string) {
	page := c.Query("Page")
	size := c.Query("Size")
	if c.Query("Page") == "" {
		page = "-1"
	}
	if c.Query("Size") == "" {
		size = "10"
	}

	return page, size
}

func FindAllParamJoins(table string, params FindAllParams) FindAllParams {
	var status bool
	if params.StatusID != "" {
		status = true
		params.StatusID = table + "." + params.StatusID
	}

	if params.BusinessID != "" {
		businessKey := 5
		if !status {
			businessKey = 0
		}
		businessID := params.BusinessID[businessKey:]
		EqBusinessID := params.BusinessID[:businessKey]
		businessID = table + "." + businessID
		params.BusinessID = EqBusinessID + businessID
	}

	if params.GroupBy != "" {
		params.GroupBy = table + "." + params.GroupBy
	}

	var sort string
	desc := strings.Contains(strings.ToLower(params.SortBy), "desc")
	asc := strings.Contains(strings.ToLower(params.SortBy), "asc")

	if desc {
		sort = "desc"
	}

	if asc {
		sort = "asc"
	}

	if params.SortBy != "" {
		leng := len(params.SortBy)
		sortASC := strings.HasSuffix(strings.ToLower(params.SortBy), "asc")
		sortDESC := strings.HasSuffix(strings.ToLower(params.SortBy), "desc")
		if sortASC {
			leng -= 4
		} else if sortDESC {
			leng -= 5
		}

		ExplodeSortBy := strings.Split(params.SortBy[:leng], ",")
		var newSortBy string
		var no int
		lengExplode := len(ExplodeSortBy)
		for _, v := range ExplodeSortBy {

			if lengExplode-1 != no {
				newSortBy = newSortBy + table + "." + v + " " + sort + ","
			} else {
				newSortBy = newSortBy + table + "." + v + " " + sort
			}
			no++
		}

		params.SortBy = newSortBy
	}

	return params
}

// find all multifunction
func FilterFindAllParamExternal(c *gin.Context) FindAllParams {
	var findallparams FindAllParams
	var keywordName string
	var businessID string
	var statusID string
	var outletID string
	var sort string
	var QC string
	var Q string
	var op string

	findallparams = FindAllParams{"-1", "10", "", "code", "", "", "id desc", "", "", "", "", "", ""}
	businessID = *configs.BusinessID
	outletID = c.Query("OutletID")
	keywordName = c.Query("keywordName")
	sortName := Underscore(c.Query("SortName"))
	sortBy := strings.ToLower(c.Query("SortBy"))

	if c.Query("Status ID") == "" {
		statusID = c.Query("StatusID")
	} else {
		statusID = c.Query("Status ID")
	}

	if statusID != "-1" && statusID != "" {
		statusID = "status_id = " + statusID
	} else {
		statusID = ""
	}

	if businessID != "-1" && businessID != "" {
		if statusID != "" {
			op = " AND "
		}
		businessID = op + "business_id = " + businessID
		op = ""
	} else {
		businessID = ""
	}

	if outletID != "-1" && outletID != "" {
		if businessID != "" {
			op = " AND "
		}
		outletID = op + " outlet_id = " + outletID
	} else {
		outletID = ""
	}

	if c.Query("Query") == "" {
		Q = c.Query("Q")
	} else {
		Q = c.Query("Query")
	}
	query := QueryReplaceFindAll(Q)

	if sortName != "" {
		sort = sortName + " " + sortBy
	}

	dataFinder := DataFinder(c.Query("KeywordName"), c.Query("Keyword"))
	page := c.Query("Page")
	size := c.Query("Size")
	keyword := c.Query("Keyword")
	grpupBy := Underscore(c.Query("GroupBy"))
	QC = query + statusID + businessID + outletID + keywordName
	findallparams = FindAllParams{Page: page, Size: size, Keyword: keyword, StatusID: statusID, Query: query, DataFinder: dataFinder, QueryCondition: QC, SortName: sortName, SortBy: sort, GroupBy: grpupBy, BusinessID: businessID, OutletID: outletID}
	return findallparams
}

// find all multifunction
func FilterFindAllParam(c *gin.Context, claims jwt.MapClaims) FindAllParams {
	var findallparams FindAllParams
	var keywordName string
	var businessID string
	var statusID string
	var outletID string
	var sort string
	var QC string
	var Q string
	var op string

	findallparams = FindAllParams{"-1", "10", "", "code", "", "", "id desc", "", "", "", "", "", ""}
	businessID = fmt.Sprintf("%v", claims["BusinessID"])
	outletID = c.Query("OutletID")
	keywordName = c.Query("keywordName")
	sortName := Underscore(c.Query("SortName"))
	sortBy := strings.ToLower(c.Query("SortBy"))

	if c.Query("Status ID") == "" {
		statusID = c.Query("StatusID")
	} else {
		statusID = c.Query("Status ID")
	}

	if statusID != "-1" && statusID != "" {
		statusID = "status_id = " + statusID
	} else {
		statusID = ""
	}

	if businessID != "-1" && businessID != "" {
		if statusID != "" {
			op = " AND "
		}
		businessID = op + "business_id = " + businessID
		op = ""
	} else {
		businessID = ""
	}

	if outletID != "-1" && outletID != "" {
		if businessID != "" {
			op = " AND "
		}
		outletID = op + " outlet_id = " + outletID
	} else {
		outletID = ""
	}

	if c.Query("Query") == "" {
		Q = c.Query("Q")
	} else {
		Q = c.Query("Query")
	}
	query := QueryReplaceFindAll(Q)

	if sortName != "" {
		sort = sortName + " " + sortBy
	}

	dataFinder := DataFinder(c.Query("KeywordName"), c.Query("Keyword"))
	page := c.Query("Page")
	size := c.Query("Size")
	keyword := c.Query("Keyword")
	grpupBy := Underscore(c.Query("GroupBy"))
	QC = query + statusID + businessID + outletID + keywordName
	findallparams = FindAllParams{Page: page, Size: size, Keyword: keyword, StatusID: statusID, Query: query, DataFinder: dataFinder, QueryCondition: QC, SortName: sortName, SortBy: sort, GroupBy: grpupBy, BusinessID: businessID, OutletID: outletID}
	return findallparams
}

// keyword like full text search
func DataFinder(keywordname string, keyword string) string {
	str := "1=1"
	if keywordname != "" && keyword != "" {
		ExplodeParam := strings.Split(keywordname, ",")
		ExplodeKeyword := strings.Split(keyword, " ")
		for _, vKeyword := range ExplodeKeyword {
			str += " and ( "
			strTmp := ""
			for _, vParam := range ExplodeParam {
				if strTmp != "" {
					strTmp += " or "
				}

				strTmp += " " + Underscore(vParam) + " like '%" + vKeyword + "%' "
			}
			str += strTmp
			str += " )"
		}

	}

	return str
}

func GetFileNameDropbox(code string, folder string) string {
	filename := "/atr/" + folder + "/" + strings.ToUpper(code) + "-" + strconv.Itoa(rand.Intn(10000000)) + ".jpg"
	return filename
}

// migrate db
func AddDbMigrate() []TableMigrate {
	type TableMigrates []TableMigrate
	var tableMigrates = TableMigrates{
		TableMigrate{Table: models.Atr{}},
		TableMigrate{Table: models.User{}},
	}

	return tableMigrates
}

func DBMigrate(db *models.DB) error {
	var errMigrate error
	tx := configs.ActiveDB.Begin()
	tableMigrates := AddDbMigrate()
	for _, v := range tableMigrates {
		errMigrate = tx.AutoMigrate(v.Table).Error
		if errMigrate != nil {
			tx.Rollback()
			fmt.Println(errMigrate)
			return errMigrate
		}
	}
	tx.Commit()

	return errMigrate
}

func DBSeed(db *models.DB) error {
	var errSeed error
	errSeed = models.PackageSeed(db)

	if errSeed != nil {
		return errSeed
	}

	return nil
}

func QueryReplaceFindAll(str string) string {
	strReplace := []byte(str)
	len := len(strReplace)
	for i := 0; i < len; i++ {
		if strReplace[i] == '(' || strReplace[i] == ')' {
			strReplace[i] = ' '
		}
	}
	return string(strReplace)
}

func DistanceLongLat(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	// 'M' is statute miles (default)
	if len(unit) > 0 {
		if unit[0] == "K" { // 'K' is kilometers
			dist = dist * 1.609344
		} else if unit[0] == "N" { // 'N' is nautical miles
			dist = dist * 0.8684
		}
	}

	return dist
}

func (b *buffer) write(r rune) {
	if r < utf8.RuneSelf {
		b.r = append(b.r, byte(r))
		return
	}
	n := utf8.EncodeRune(b.runeBytes[0:], r)
	b.r = append(b.r, b.runeBytes[0:n]...)
}

func (b *buffer) indent() {
	if len(b.r) > 0 {
		b.r = append(b.r, '_')
	}
}

func (b *buffer) indentSpace() {
	if len(b.r) > 0 {
		b.r = append(b.r, ' ')
	}
}

func Underscore(s string) string {
	b := buffer{
		r: make([]byte, 0, len(s)),
	}
	var m rune
	var w bool
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if m != 0 {
				if !w {
					b.indent()
					w = true
				}
				b.write(m)
			}
			m = unicode.ToLower(ch)
		} else if unicode.IsSpace(ch) {
			if m != 0 {
				b.indentSpace()
				m = 0
				w = false
			}
		} else {
			if m != 0 {
				b.indent()
				b.write(m)
				m = 0
				w = false
			}
			b.write(ch)
		}
	}
	if m != 0 {
		if !w {
			b.indent()
		}
		b.write(m)
	}

	// handle ID camel case
	strReplace := []byte(string(b.r))
	countID := strings.Count(string(strReplace), "i_d")
	if countID >= 1 {
		len := len(strReplace)
		for i := 0; i < len; i++ {
			if strReplace[i] == 'i' {
				if strReplace[i+1] == '_' {
					if strReplace[i+2] == 'd' {
						strReplace[i+1] = ' '
					}
				}
			}
		}
	}
	return strings.Replace(string(strReplace), " ", "", -1)
}

func SendSlackNotification(Message string) error {
	envname := "WEBHOOKSLACK"
	webhookslack := os.Getenv(envname)
	webhookUrl := webhookslack // webhooks notif bitu-sip
	slackBody, _ := json.Marshal(SlackRequestBody{Text: Message})

	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}

	return nil
}

func ReturnHandler(status int, err error) StatusClient {
	var StatusClientSave StatusClient
	// Get Filter & Message Here
	// SendSlackNotification(Message)
	return StatusClientSave
}
