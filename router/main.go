package router

import (
	"io"
	"io/ioutil"
	mime2 "mime"
	"net/http"
	"path/filepath"
)

var (
	routes           = make(map[string]http.HandlerFunc)
	params           = make(map[string]string)
	notFoundCallback http.HandlerFunc
	notFoundPage     string
)

func Dispatch(res http.ResponseWriter, req *http.Request) {

	for router, callback := range routes {
		if match(req.Method, req.URL.Path, router) {
			callback(res, req)
			return
		}
	}

	if notFoundCallback == nil {
		notFoundCallback = defaultNotFound
	}
	notFoundCallback(res, req)
}

func defaultNotFound(res http.ResponseWriter, req *http.Request) {
	content := "page not found"
	mime := "text/html"
	if notFoundPage != "" {
		content1, err := ioutil.ReadFile(notFoundPage)
		if err == nil {
			content = string(content1)
		}

		// todo 判断没查找到的情况
		mime = mime2.TypeByExtension(filepath.Ext(notFoundPage))
	}

	res.Header().Set("Content-Type", mime)
	_, err := io.WriteString(res, content)
	if err != nil {

	}
}

func RegisterNotFound(callback http.HandlerFunc) {
	notFoundCallback = callback
}

func RegisterNotFoundPage(path string) {
	notFoundPage = path
}

func GetParams(key string) string {
	return params[key]
}

func HasParams(key string) bool {
	if _, ok := params[key]; ok {
		return true
	}
	return false
}
