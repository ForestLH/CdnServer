package src

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func InitRoute(server *server.Hertz) {
	// 使用swag init来更新swag注释
	server.GET("swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/", helpServices)
	g := server.Group("/api")
	g.POST("/video/:name", videoServices)
	g.POST("/images/:name", imagesServices)
}
