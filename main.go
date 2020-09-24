package main

import (
	"log"
	_ "orange/bootstrap"
	"orange/router"
	"orange/utils/yml_config"
)

func main() {
	router := router.InitRouter()
	if err := router.Run(yml_config.CreateYamlFactory().GetString("HttpServer.Web.Port")); err != nil {
		log.Fatal(err)
	}
}
