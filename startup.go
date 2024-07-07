package startup

import (
	"log"

	initial "github.com/Song2017/startup/initial"

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
