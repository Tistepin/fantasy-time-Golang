package test

import (
	"fmt"
	"github.com/goccy/go-json"
	"go_dialogue/model"
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
	dsn := "root:xu20010502@tcp(127.0.0.1:3306)/fantasytime?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	user := &model.Contact{}
	db.AutoMigrate(user)
}
