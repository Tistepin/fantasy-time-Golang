package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

/**
* User:徐国纪
* Create_Time:上午 08:56
**/

var (
	// DB 数据库连接
	DB *gorm.DB
	//Red *redis.Client
)

func InitConfig() {
	// 写入读取配置的文件名
	viper.SetConfigName("app")
	// 文件读取地址路径
	viper.AddConfigPath("config")
	// 读
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

// InitMysql 初始化mysql
func InitMysql() {
	//自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)
	// 用gorm框架连接mysql  数据库配置获取yml文件的
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println("数据库报错", err)
	}
	DB = db
}

func main() {

}
