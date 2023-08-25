package model

import (
	"go_dialogue/utils"
	"gorm.io/gorm"
)

/**
* User:徐国纪
* Create_Time:上午 09:48
**/

// Contact 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁 /群 ID
	Type     int  //对应的类型  1好友  2群  3xx
	DescInfo string
}

// TableName 获取表名
func (table *Contact) TableName() string {
	return "ft_contact"
}

// SearchFriend 搜索好友列表
func SearchFriend(userId uint) []FtUser {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type=1", userId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]FtUser, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users
}

// AddFriend 添加好友   自己的ID  ， 好友的ID
func AddFriend(userId uint, targetId uint) (int, string) {
	if targetId == userId {
		return -1, "不能加自己"
	}

	contact0 := Contact{}
	utils.DB.Where("owner_id =?  and target_id =? and type=1", userId, targetId).Find(&contact0)
	if contact0.ID != 0 {
		return -1, "不能重复添加"
	}
	// 开启事务
	tx := utils.DB.Begin()
	//事务一旦开始，不论什么异常最终都会 Rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 创建关系
	contact := Contact{}
	contact.OwnerId = userId
	contact.TargetId = targetId
	contact.Type = 1
	if err := utils.DB.Create(&contact).Error; err != nil {
		tx.Rollback()
		return -1, "添加好友失败"
	}
	contact1 := Contact{}
	contact1.OwnerId = targetId
	contact1.TargetId = userId
	contact1.Type = 1
	if err := utils.DB.Create(&contact1).Error; err != nil {
		tx.Rollback()
		return -1, "添加好友失败"
	}
	tx.Commit()
	return 0, "添加好友成功"
}
