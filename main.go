package main

import "go_dialogue/utils"

/**
* User:徐国纪
* Create_Time:下午 04:55
**/
func main() {
	//utils.InitConfig()
	//utils.InitMysql()
	utils.InitNacos()
	////utils.InitRedis()
	//e := router.Router()
	//e.Run(":" + strconv.FormatUint(viper.GetUint64("port.server"), 10))
}
