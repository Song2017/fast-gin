package initial

import (
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server interface {
	ListenAndServe() error
}

func InitRouter(router *gin.Engine) *gin.Engine {
	if router == nil {
		router = gin.New()
	}
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "root page")
	})
	router.Static("/static", "./_resources")
	// swagger UI
	// http://localhost:9000/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler, ginSwagger.URL("/static/openapi.yaml")))
	return router
}

func InitServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20

	LoadServerConfig()
	if O_SERVER_CONFIG.PgConn.DSN != "" {
		O_DB_PG = GormPgSqlByConfig(O_SERVER_CONFIG.PgConn)
	}
	if O_SERVER_CONFIG.Redis.Addr != "" {
		O_REDIS = InitRedisByConfig(O_SERVER_CONFIG.Redis)
	}

	return s
}
