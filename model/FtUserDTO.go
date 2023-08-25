package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/**
* User:徐国纪
* Create_Time:上午 09:19
**/

type FtUser struct {
	ID           uint   `gorm:"column:id;primarykey"`
	UserName     string `gorm:"column:username;unique;NOT NULL" json:"username"`
	Password     string `gorm:"column:password" json:"passWord"`
	NickName     string `gorm:"column:nickname" json:"nickName"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Email        string `gorm:"column:email" json:"email"`
	Header       string `gorm:"column:header" json:"header"`
	Gender       uint   `gorm:"column:gender" json:"gender"`
	Birth        string `gorm:"column:birth" json:"birth"`
	City         string `gorm:"column:city" json:"city"`
	Job          string `gorm:"column:job" json:"job"`
	Sign         string `gorm:"column:sign" json:"sign"`
	SourceType   uint   `gorm:"column:source_type" json:"sourceType"`
	Status       uint   `gorm:"column:status" json:"status"`
	CreateTime   string `gorm:"column:create_time" json:"createTime"`
	EditTime     string `gorm:"column:edit_time" json:"editTime"`
	DeleteStatus uint   `gorm:"column:delete_status" json:"deleteStatus"`
}

func (this *FtUser) TableName() string {
	return "ft_user"
}
func (this *FtUser) GetId(FantasyTimetoken string) uint {
	// 获取客户端
	client := &http.Client{}
	// 构建请求
	req, _ := http.NewRequest("GET", "http://localhost:8081/works/user/getUserEntity", nil)
	// 添加请求头
	req.Header.Add("FantasyTimetoken", FantasyTimetoken)
	// 开始请求
	resp, _ := client.Do(req)
	// 延迟关闭
	defer resp.Body.Close()
	// io获取请求信息
	reader := io.Reader(resp.Body)
	var p []byte = make([]byte, 1024)
	n, _ := reader.Read(p)
	var m = make(map[string]interface{}, 10)
	err := json.Unmarshal(p[:n], &m)
	if err != nil {
		fmt.Println(err)
	}
	// 获取UserInfo信息
	k := m["data"]
	mMap, ok := k.(map[string]interface{})
	if !ok {
		return 0
	}
	i := mMap["data"]
	marshal, _ := json.Marshal(i)
	//fmt.Println(user)
	err = json.Unmarshal(marshal, this)
	if err != nil {
		fmt.Println(err)
	}
	return this.ID
}
