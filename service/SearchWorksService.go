package service

import (
	"github.com/gin-gonic/gin"
	"go_dialogue/model"
	"go_dialogue/utils"
	"strconv"
)

/**
* User:徐国纪
* Create_Time:下午 01:18
**/

func GetWorks(r *gin.Context) {

	db := utils.DB
	db = db.Where("review_status=1")
	// 页码
	pageNum, _ := strconv.ParseInt(r.Query("pageNum"), 0, 0)
	// 每页数量
	pageSize, _ := strconv.ParseInt(r.Query("pageSize"), 0, 0)
	// 评分最高
	highestRated, _ := strconv.ParseBool(r.Query("highestRated"))
	if highestRated {
		db = db.Order("works_score")
	}
	//// 最新发布
	latestRelease, _ := strconv.ParseBool(r.Query("latestRelease"))
	if latestRelease {
		db = db.Order("create_time desc")
	}
	// 最近更新
	recentlyUpdated, _ := strconv.ParseBool(r.Query("recentlyUpdated"))
	if recentlyUpdated {
		db = db.Order("edit_time desc")
	}
	// 人气最旺
	theMostPopular, _ := strconv.ParseBool(r.Query("theMostPopular"))
	if theMostPopular {
		db = db.Order("works_popularity desc")
	}
	// 地区
	if a := r.Query("worksArea"); a != "" {
		worksArea, _ := strconv.ParseInt(a, 0, 0)
		if worksArea != 1 {
			db = db.Where("works_area=?", worksArea)
		}
	}
	// 作品类别标签
	if worksCategory := r.Query("worksCategory"); worksCategory != "" {
		//
		if a, _ := strconv.ParseInt(worksCategory, 0, 0); a != 1 {
			db = db.Where("works_category like ?", "%"+worksCategory+"%")
		}
	}
	// 作品名称 模糊查询
	if worksName := r.Query("worksName"); worksName != "" {
		db = db.Where("works_name like ?", "%"+worksName+"%")
	}
	// 作品状态是否完结 1连载 2 完结
	if a := r.Query("worksStatus"); a != "" {
		worksStatus, _ := strconv.ParseInt(a, 0, 0)
		if worksStatus != 0 {
			db = db.Where("works_status=?", worksStatus)
		}
	}
	//// 作品类型
	worksType, _ := strconv.ParseInt(r.Query("worksType"), 0, 0)
	db = db.Where("works_type=?", worksType)
	offset := (pageNum - 1) * pageSize

	var works []model.Works
	//utils.DB.Find(&works)

	//分页查询
	if err := db.Offset(int(offset)).Limit(int(pageSize)).Find(&works).Error; err != nil {
		utils.RFail(r.Writer, "查询数据异常")
		return
	}
	var count int64 = 0
	db.Count(&count)
	//marshal, _ := json.Marshal(works)
	utils.RespOKList(r.Writer, works, count)
}
