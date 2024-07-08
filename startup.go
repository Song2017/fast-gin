package startup

import (
	"log"

	initial "github.com/Song2017/startup/initial"
	mw "github.com/Song2017/startup/middleware"
	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	initial.InitServerConfig()
	router := initial.InitBaseRouter(nil)

	initial.InitSwagger(router, "./resources")

	return router
}

func RunServer(router *gin.Engine) {
	// init server
	s := initial.InitServer("0.0.0.0:9000", router)
	log.Fatal(s.ListenAndServe().Error())
}

// middleware
func LoggerMiddleware(platform string) gin.HandlerFunc {
	return mw.LoggerMiddleware(platform)
}

func AuthorizationMiddleware(securityKey, securityVal string) gin.HandlerFunc {
	return mw.AuthorizationMiddleware(securityKey, securityVal)
}
