package service

import (
	"github.com/gin-gonic/gin"
	"go_dialogue/model"
)

/**
* User:徐国纪
* Create_Time:下午 01:50
**/

// LoginChat 用户登录聊天系统
func LoginChat(c *gin.Context) {
	model.Chat(c.Writer, c.Request)
}
