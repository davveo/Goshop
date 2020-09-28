package main

import (
	"log"
	_ "Eshop/bootstrap"
	"Eshop/router"
	"Eshop/utils/yml_config"
)

func main() {
	router := router.InitRouter()
	if err := router.Run(yml_config.CreateYamlFactory().GetString("HttpServer.Web.Port")); err != nil {
		log.Fatal(err)
	}
}
