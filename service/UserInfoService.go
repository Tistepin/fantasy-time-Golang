package service

import (
	"github.com/gin-gonic/gin"
	"go_dialogue/model"
	"go_dialogue/utils"
	"strconv"
)

/**
* User:徐国纪
* Create_Time:上午 11:22
**/

// GetUserStateInfo
// @Summary 测试查询用户好友的状态
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags User
// @Accept application/json
// @Produce application/json
// @param ids query string false "ids"
// @Success 200 {string} gin.H
// @Router /GetContactStates [get]
func GetUserStateInfo(r *gin.Context) {
	ids := utils.ParamInt64Slice(r.Query("ids"))
	var rMap = make(map[int64]int64, len(ids))
	for _, id := range ids {
		if model.ClientMap[id] != nil {
			rMap[id] = utils.ON_LINE_STATE
			continue
		}
		rMap[id] = utils.OFF_LINE_STATE
	}
	utils.Rok(r.Writer, rMap, "")
}

// AddFriend
// @Summary 添加好友
// @Description 添加好友
// @Tags User
// @Accept application/json
// @Produce application/json
// @param userId query string false "userId"
// @param targetId query string false "targetId"
// @Success 200 {string} gin.H
// @Router /AddFriend [PUT]
func AddFriend(r *gin.Context) {

	userId, err := strconv.ParseUint(r.Query("userId"), 10, 10)
	if err != nil {
		return
	}
	targetId, err := strconv.ParseUint(r.Query("targetId"), 10, 10)
	if err != nil {
		return
	}
	// 操作类型 1添加关系 2取消关系
	OperationType, err := strconv.ParseUint(r.Query("OperationType"), 10, 10)
	if err != nil {
		return
	}
	var s = ""
	var friend = 0
	if OperationType == 1 {
		friend, s = model.AddFriend(uint(userId), uint(targetId))
		if friend == 0 {
			utils.Rok(r.Writer, friend, s)
			return
		}
	} else if OperationType == 2 {
		friend, s = model.RemoveFriend(uint(userId), uint(targetId))
		if friend == 0 {
			utils.Rok(r.Writer, friend, s)
			return
		}
	}

	utils.RFail(r.Writer, s)
}
