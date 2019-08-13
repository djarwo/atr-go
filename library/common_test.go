package library

import "github.com/gin-gonic/gin"

func GetRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
		r.Use(setUserStatus()) // new line
	}
	return r
}
