package main

import (
	_ "Goshop/bootstrap"
	"Goshop/router"
	"Goshop/utils/yml_config"
	"log"
	"net/http"

	"github.com/arl/statsviz"
)

func main() {
	var (
		StatsvizPort  = yml_config.CreateYamlFactory().GetString("HttpServer.Statsviz.Port")
		WebServerPort = yml_config.CreateYamlFactory().GetString("HttpServer.Web.Port")
	)
	go func() {
		// Real-time monitoring system
		// http://127.0.0.1:8080/debug/statsviz/
		statsviz.RegisterDefault()
		if err := http.ListenAndServe(StatsvizPort, nil); err != nil {
			log.Fatal(err)
		}
	}()

	if err := router.InitRouter().Run(WebServerPort); err != nil {
		log.Fatal(err)
	}
}
