package main

import (
	"app1/models"
	"app1/router"
	"github.com/jinzhu/configor"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {

	registerRoutes()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("service begin")

	var config = new(Config)
	err := configor.Load(config, "config.yml")
	if err != nil {
		log.Println("config loaded failed")
		return
	}

	models.SetDsn(config.DB.Dsn)
	log.Println("server listened at " + config.Server.Listen)
	err = http.ListenAndServe(config.Server.Listen, http.HandlerFunc(router.Dispatch))
	if err != nil {
		log.Println(err)
		return
	}
}
