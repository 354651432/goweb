package router

import (
	"fmt"
	"net/http"
	"strings"
)

func AddRoute(method string, path string, callback http.HandlerFunc) {
	method = strings.ToLower(method)
	path = strings.ToLower(path)
	routes[fmt.Sprintf("%v:%v", method, path)] = callback
}

func Get(path string, callback http.HandlerFunc) {
	AddRoute("get", path, callback)
}

func Post(path string, callback http.HandlerFunc) {
	AddRoute("post", path, callback)
}

func Delete(path string, callback http.HandlerFunc) {
	AddRoute("delete", path, callback)
}

func Put(path string, callback http.HandlerFunc) {
	AddRoute("put", path, callback)
}

func Patch(path string, callback http.HandlerFunc) {
	AddRoute("patch", path, callback)
}

func Any(path string, callback http.HandlerFunc) {
	AddRoute("any", path, callback)
}

func Redirect(path, redPath string) {
	Any(path, func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Location", redPath)
		res.WriteHeader(302)
	})
}
