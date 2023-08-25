package service

import (
	"github.com/gin-gonic/gin"
	"go_dialogue/model"
	"go_dialogue/utils"
)

/**
* User:徐国纪
* Create_Time:上午 11:22
**/

// GetUserStateInfo
// @Summary 测试查询用户好友的状态
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
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
