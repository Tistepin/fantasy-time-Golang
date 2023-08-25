package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_dialogue/docs"
	"go_dialogue/service"
)

/**
* User:徐国纪
* Create_Time:上午 08:57
**/

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 静态资源
	r.Static("/asset", "asset/")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	//r.LoadHTMLGlob("views/**/*")
	r.GET("/GetContactStates", service.GetUserStateInfo)
	r.GET("/LoginWebSocket", service.LoginChat)
	return r
}
