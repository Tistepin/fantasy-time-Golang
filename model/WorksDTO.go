package model

import (
	"fmt"
	"time"
)

/**
* User:徐国纪
* Create_Time:上午 11:14
**/

type Works struct {
	WorksId           int       `json:"worksId"`
	WorksName         string    `json:"worksName"`
	DefaultImage      string    `json:"defaultImage"`
	Creator           string    `json:"creator"`
	WorksCreator      string    `json:"worksCreator"`
	WorksCreateTime   LocalTime `json:"worksCreateTime"`
	WorksArea         string    `json:"worksArea"`
	WorksAreaName     string    `json:"worksAreaName"`
	WorksType         int       `json:"worksType"`
	WorksScore        float64   `json:"worksScore"`
	WorksRenew        int       `json:"worksRenew"`
	WorksPopularity   int       `json:"worksPopularity"`
	WorksDescribe     string    `json:"worksDescribe"`
	WorksCategory     string    `json:"worksCategory"`
	WorksCategoryName string    `json:"worksCategoryName"`
	WorksStatus       int       `json:"worksStatus"`
	CreateTime        LocalTime `json:"createTime"`
	EditTime          LocalTime `json:"editTime"`
}
type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// TableName 获取表名
func (table *Works) TableName() string {
	return "ft_works"
}
