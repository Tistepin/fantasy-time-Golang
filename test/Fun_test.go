package test

import (
	"fmt"
	"github.com/goccy/go-json"
	"go_dialogue/model"
	"go_dialogue/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net/http"
	"testing"
)

/**
* User:徐国纪
* Create_Time:上午 09:03
**/

func TestName(t *testing.T) {
	// 连接数据库
	//dsn := "root:xu20010502@tcp(127.0.0.1:3306)/fantasytime?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//user := &model.Contact{}
	//自动迁移 创建表
	//db.AutoMigrate(user)
	//db.Where("id = 1").First(user)
	//print(user)
	FantasyTimetoken := "eyJhbGciOiJIUzUxMiIsInppcCI6IkdaSVAifQ.H4sIAAAAAAAAAKtWKi5NUrJSMrQwNTexNDY2MDNW0lFKrShQsjI0szSyNDYwMDOtBQC5l_GTJgAAAA.IeasnivUGkx3afIWyujtBHYstlR6wDwa1DMDBPQA7Cquhi3KPpVyShXbG3XacMuvzDh6CycSGsyWvraKC3zjzw"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8081/works/user/getUserEntity", nil)
	req.Header.Add("FantasyTimetoken", FantasyTimetoken)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reader := io.Reader(resp.Body)
	var p []byte = make([]byte, 1024)
	n, _ := reader.Read(p)
	var m = make(map[string]interface{}, 10)
	json.Unmarshal(p[:n], &m)
	k := m["data"]
	mMap, ok := k.(map[string]interface{})
	if !ok {
		return
	}
	i := mMap["data"]
	fmt.Println(i)
	user := &model.FtUser{}
	marshal, _ := json.Marshal(i)
	//fmt.Println(user)
	err := json.Unmarshal(marshal, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func TestItoaji(t *testing.T) {
	//连接数据库
	dsn := "root:root@tcp(192.168.153.134:3306)/fantasytime?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	user := &model.Message{}
	db.AutoMigrate(user)
}

func TestJm(t *testing.T) {

	AddFriend(1, 2)
}

// AddFriend 添加好友   自己的ID  ， 好友的ID
func AddFriend(userId uint, targetId uint) (int, string) {
	if targetId == userId {
		return -1, "不能加自己"
	}
	dsn := "root:xu20010502@tcp(127.0.0.1:3306)/fantasytime?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	contact0 := model.Contact{}
	db.Where("owner_id =?  and target_id =? and type=1", userId, targetId).Find(&contact0)
	if contact0.ID != 0 {
		return -1, "不能重复添加"
	}
	// 开启事务
	tx := db.Begin()
	//事务一旦开始，不论什么异常最终都会 Rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 创建关系
	contact := model.Contact{}
	contact.OwnerId = userId
	contact.TargetId = targetId
	contact.Type = 1
	if err := utils.DB.Create(&contact).Error; err != nil {
		tx.Rollback()
		return -1, "添加好友失败"
	}
	contact1 := model.Contact{}
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

func TestGetworks(t *testing.T) {
	dsn := "root:xu20010502@tcp(127.0.0.1:3306)/fantasytime?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var works []model.Works
	db.Find(&works)
	marshal, _ := json.Marshal(works)
	fmt.Println(string(marshal))
}
