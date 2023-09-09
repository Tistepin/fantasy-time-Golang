package utils

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
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

// 注册nacos 进行服务发现
func InitNacos() {
	// 配置nacos的连接配置
	sc := []constant.ServerConfig{
		constant.ServerConfig{
			ContextPath: "/nacos",
			IpAddr:      viper.GetString("nacos.Ip"),
			Port:        viper.GetUint64("nacos.Port"),
		},
	}
	// 配置客户端注册在哪里
	cc := constant.ClientConfig{
		NamespaceId:         "",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/emp/nacos/log",
		CacheDir:            "/temp/nacos/cache",
		LogLevel:            "debug",
	}
	// 创建服务发现客户端
	client, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err != nil {
		log.Fatal("创建服务发现客户端", err)
	}
	client.RegisterInstance(
		vo.RegisterInstanceParam{
			Ip:       "10.161.139.216", // 配置自己IP  表示谁都可以访问
			Port:     viper.GetUint64("port.server"),
			Weight:   10,
			Enable:   true, // true表示可以访问 其他服务可以根据nacos访问到这个服务
			Healthy:  true, // true表示是否健康
			Metadata: nil,  // 选填
			//ClusterName: "go_dialogue", //集群名称 表示要连接哪个nacos服务
			ServiceName: "go_dialogue",
			GroupName:   "", // 默认不写为default
			//Ephemeral:   false,
		})
	if err != nil {
		log.Fatal("创建服务发现客户端", err)
	}
}
