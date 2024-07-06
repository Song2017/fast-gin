package startup

import (
	_init "startup/initial"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	_init.InitServerConfig()
	router := _init.InitRouter(nil)
	_init.InitSwagger(router, "./resources")

	return router
}
