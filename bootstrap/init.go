package bootstrap

import (
	"log"
	_ "Eshop/core/destroy"
	"Eshop/global/my_errors"
	"Eshop/global/variable"
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
			log.Fatal(my_errors.ErrorsStorageLogsMake + err.Error())
		}
	}
}

func init() {
	checkRequiredFolders()
}
