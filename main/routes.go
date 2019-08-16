package main

import (
	"app1/router"
	"io"
	"net/http"
)

func registerRoutes() {
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(res, "<h1>it works</h1?")
	})
}
