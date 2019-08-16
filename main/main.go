package main

import (
	"app1/router"
	"log"
	"net/http"
)

func main() {

	registerRoutes()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("service begin")
	err := http.ListenAndServe("0.0.0.0:2000", http.HandlerFunc(router.Dispatch))
	if err != nil {

	}
}
