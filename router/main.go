package router

import (
	"io"
	"io/ioutil"
	"log"
	mime2 "mime"
	"net/http"
	"os"
	path2 "path"
	"path/filepath"
)

var (
	routes           = make(map[string]http.HandlerFunc)
	params           = make(map[string]string)
	notFoundCallback http.HandlerFunc
	notFoundPage     string
	publicPath       = "public"
	index            = "index.html"
)

func Dispatch(res http.ResponseWriter, req *http.Request) {
	if staticProc(res, req) {
		return
	}
	res.Header().Set("Content-Type", "text/html;charset=utf8")

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

func staticProc(res http.ResponseWriter, req *http.Request) bool {
	var path = path2.Join(publicPath, req.URL.Path)
	stat, _ := os.Stat(path)

	if stat == nil {
		return false
	}

	if stat.IsDir() {
		path = path2.Join(path, index)
		stat, _ = os.Stat(path)
		if stat == nil {
			return false
		}
	}
	var mime = mime2.TypeByExtension(filepath.Ext(path))
	res.Header().Set("Content-Type", mime)

	var content, _ = ioutil.ReadFile(path)
	_, err := res.Write(content)
	if err != nil {
		log.Println(err)
	}
	return true
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

func SetPublicPath(path string) {
	publicPath = path
}

func SetIndex(indexStr string) {
	index = indexStr
}
