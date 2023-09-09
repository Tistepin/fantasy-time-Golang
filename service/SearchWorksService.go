package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_dialogue/model"
	"go_dialogue/utils"
	"io"
	"os"
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

func GetWorksImages(r *gin.Context) {
	getImage(r)
}
func getImage(r *gin.Context) {
	ImageDefaultStatus, _ := r.GetQuery("ImageDefaultStatus")
	WorksId, b := r.GetQuery("WorksId")
	worksChapterId, worksChapterIdCheck := r.GetQuery("WorksChapterId")
	imageId, imageIdCheck := r.GetQuery("ImageId")
	// 0 不是 1 是
	Url := ""
	if a, _ := strconv.Atoi(ImageDefaultStatus); a == 1 {
		//  查询封面
		if b {
			entity := model.WorksDefaultImageEntity{}
			utils.DB.Where("works_id=?", WorksId).First(&entity)
			Url = entity.WorksDefaultImage
		} else {
			utils.RFail(r.Writer, "没有作品ID")
		}
	} else {
		if worksChapterIdCheck {
			if imageIdCheck {
				entity := model.WorksChapterDetailedViewingContentEntity{}
				utils.DB.Where("works_id=? and works_chapter_id = ? and image_id=? and delete_status=1 and review_status=1", WorksId, worksChapterId, imageId).First(&entity)
				Url = entity.WorksChapterLocation
			} else {
				utils.RFail(r.Writer, "没有作品章节图片ID")
			}
		} else {
			utils.RFail(r.Writer, "没有作品章节ID")
		}
	}
	if Url != "" {
		file, err := os.Open(Url)
		if err != nil {
			fmt.Println("open file err=", err)
		}
		defer file.Close()

		// 读取图片文件内容
		imageData, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// 从数据库中获取图片数据
		if err != nil {
			//http.Error(r, "Failed to retrieve image", http.StatusInternalServerError)
			return
		}

		// 将图片数据写入响应
		r.Header("Content-Type", "image/jpg")
		r.Writer.Write(imageData)
	} else {
		utils.RFail(r.Writer, "没有图片")
	}

}
