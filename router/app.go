package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_dialogue/docs"
	"go_dialogue/service"
	"net/http"
)

/**
* User:徐国纪
* Create_Time:上午 08:57
**/

// 跨域设置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,fantasytimetoken")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		c.Next()
	}
}

func Router() *gin.Engine {

	r := gin.Default()
	r.Use(Cors())
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 静态资源
	r.Static("/asset", "asset/")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	//r.LoadHTMLGlob("views/**/*")
	r.GET("/GetContactStates", service.GetUserStateInfo)
	r.GET("/LoginWebSocket", service.LoginChat)
	r.PUT("/AddFriend", service.AddFriend)
	r.GET("/test", service.Test)
	r.GET("/getWorks", service.GetWorks)
	r.GET("/getWorksImages", service.GetWorksImages)
	//r.GET("/zc", func(context *gin.Context) {
	//	utils.InitNacos()
	//})
	//r.GET("/zx", func(context *gin.Context) {
	//
	//	success, err := utils.Client.DeregisterInstance(vo.DeregisterInstanceParam{
	//		Ip:          "127.0.0.1",
	//		Port:        8883,
	//		ServiceName: "go-dialogue",
	//		Ephemeral:   true,
	//	})
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(success)
	//
	//})
	return r
}
