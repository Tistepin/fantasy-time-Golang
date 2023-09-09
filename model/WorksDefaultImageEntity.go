package model

/**
* User:徐国纪
* Create_Time:上午 11:16
**/

type WorksDefaultImageEntity struct {

	/**
	 * id
	 */
	id int `json:"id,omitempty"`
	/**
	 * 作品id
	 */
	WorksId int `json:"works_id,omitempty"`
	/**
	 * 图片服务请求数据位置
	 */
	WorksDefaultImage string `json:"works_default_image,omitempty"`
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
}

// TableName 获取表名
func (table *WorksDefaultImageEntity) TableName() string {
	return "ft_works_default_image"
}
