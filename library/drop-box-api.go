package library

import (
	// "github.com/atomic/sip/library"

	"io"
	"strings"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
)

//PreparingDropboxToken provide access information to Dropbox
func PreparingDropboxToken() dropbox.Config {
	// claims := GetJWTClaims("")
	// BusinessID, _ := strconv.Atoi(fmt.Sprintf("%v", claims["BusinessID"]))
	// var business models.Business
	// configs.ActiveDB.Find(&business, BusinessID)
	data := dropbox.Config{
		Token:    "l8TLMhyrdf8AAAAAAAAAbmCfcD7nC4BLeUvVcPCu18P8a7351mqbhpTtdOKLCbgn",
		LogLevel: dropbox.LogInfo, // if needed, set the desired logging level. Default is off
	}

	return data
}

//InitializeDropboxFile provide initializer for Dropbox API Files
func InitializeDropboxFile() files.Client {

	return files.New(PreparingDropboxToken())
}

//InitializeDropboxSharing provide initializer for Dropbox API Sharing
func InitializeDropboxSharing() sharing.Client {

	return sharing.New(PreparingDropboxToken())
}

//UploadFile provide command for uploading file to Dropbox
func UploadFile(fileget io.Reader, filename string) string {
	fileclient := InitializeDropboxFile()
	arg := files.NewCommitInfo(filename)
	arg.Autorename = true

	resp, err := fileclient.Upload(arg, fileget)
	if err != nil {
		return "Failed Upload File"
	}

	return resp.PathLower
}

//CreateSharingLink provide command for create a sharable link that already modified to able to host the file
func CreateSharingLink(path string) string {
	sharingclient := InitializeDropboxSharing()
	arg := sharing.NewCreateSharedLinkArg(path)

	resp, err := sharingclient.CreateSharedLink(arg)
	if err != nil {
		return "Failed create share link"
	}
	urllink := strings.Replace(resp.Url, "dl=0", "raw=1", -1)

	return urllink
}

func Delete(path string) bool {
	fileclient := InitializeDropboxFile()

	argDelete := files.NewDeleteArg(path)

	_, err := fileclient.Delete(argDelete) // soft delete , jika permanent pakai fileclient.PermanentlyDelete(argDelete)
	if err != nil {
		return false
	}

	return true
}
