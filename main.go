package main

import (
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go_dialogue/router"
	"go_dialogue/utils"
)

/**
* User:徐国纪
* Create_Time:下午 04:55
**/
func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitNacos()
	//utils.InitRedis()
	e := router.Router()
	e.Run(":8883")
	defer func() {
		utils.Client.DeregisterInstance(vo.DeregisterInstanceParam{
			Ip:          viper.GetString("port.IP"),
			Port:        8883,
			ServiceName: "go-dialogue",
			Ephemeral:   true,
		})
	}()
}
