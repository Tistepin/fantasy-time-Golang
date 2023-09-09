package model

/**
* User:徐国纪
* Create_Time:下午 01:15
**/
type WorksChapterDetailedViewingContentEntity struct {
	/**
	 * id
	 */
	Id int `json:"id,omitempty"`
	/**
	 * 章节ID
	 */
	WorksChapterId int `json:"works_chapter_id,omitempty"`
	/**
	 * 用户ID
	 */
	UserId int `json:"user_id,omitempty"`
	/**
	 * 作品id
	 */
	WorksId int `json:"works_id,omitempty"`
	/**
	 * 该画画作品的该章节的第几个图片
	 */
	ImageId int `json:"image_id,omitempty"`
	/**
	 * 审核状态 0-审核中 1-审核成功 2-审核失败
	 */
	ReviewStatus int `json:"review_status,omitempty"`
	/**
	 * 逻辑删除状态 0-删除 1-存在
	 */
	DeleteStatus int `json:"delete_status,omitempty"`
	/**
	 * 注册时间
	 */
	CreateTime LocalTime `json:"create_time"`
	/**
	 * 修改时间
	 */
	EditTime LocalTime `json:"edit_time"`
	/**
	 * 章节数据存储位置
	 */
	WorksChapterLocation string `json:"works_chapter_Location" gorm:"column:works_chapter_Location"`
}

// TableName 获取表名
func (table *WorksChapterDetailedViewingContentEntity) TableName() string {
	return "ft_works_chapter_detailed_viewing_content"
}
