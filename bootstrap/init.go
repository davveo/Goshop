package bootstrap

import (
	_ "Goshop/core/destroy"
	"Goshop/global/my_errors"
	"Goshop/global/variable"
	"Goshop/service/sys_log_hook"
	"Goshop/utils/yml_config"
	"Goshop/utils/zap_factory"
	"log"
	"os"
)

// 检查项目必须的非编译目录是否存在，避免编译后调用的时候缺失相关目录
func checkRequiredFolders() {
	//1.检查配置文件是否存在
	if _, err := os.Stat(variable.BasePath + "/config/config.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
	//2.检查public目录是否存在
	//if _, err := os.Stat(variable.BasePath + "/public/"); err != nil {
	//	log.Fatal(my_errors.ErrorsPublicNotExists + err.Error())
	//}

	//3.检查Storage/logs 目录是否存在
	if _, err := os.Stat(variable.BasePath + "/logs/"); err != nil {
		if err := os.Mkdir(variable.BasePath+"/logs/", 0666); err != nil {
			log.Fatal(my_errors.ErrorsStorageLogsNotExists + err.Error())
		}
	}
}

func init() {
	checkRequiredFolders()

	// 初始化全局日志具柄, 载入钩子处理函数
	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)

	ymlConf := yml_config.CreateYamlFactory()
	ymlConf.ConfigFileChangeListen()

	if ymlConf.GetInt("Websocket.Start") == 1 {
		// websocket 管理中心hub全局初始化一份
		//variable.WebsocketHub = core.CreateHubFactory()
		//if Wh, ok := variable.WebsocketHub.(*core.Hub); ok {
		//	go Wh.Run()
		//}
	}
}
