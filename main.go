package main

import (
	_ "Goshop/bootstrap"
	"Goshop/router"
	"Goshop/utils/yml_config"
	"log"
)

func main() {
	router := router.InitRouter()
	if err := router.Run(yml_config.CreateYamlFactory().GetString("HttpServer.Web.Port")); err != nil {
		log.Fatal(err)
	}
}
