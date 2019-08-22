package router

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

func WriteJson(res http.ResponseWriter, content interface{}) {
	str, err := json.Marshal(content)
	if err != nil {

	}
	res.Header().Set("Content-Type", "text/json;charset=utf8")
	_, _ = res.Write(str)
}

func WriteContent(res http.ResponseWriter, content interface{}) {
	_, _ = io.WriteString(res, content.(string))
}

func WriteTemplate(res http.ResponseWriter, tmp string, content interface{}) {
	tpl := template.New("whaterver")
	_, _ = tpl.ParseFiles(tmp)
	_ = tpl.Execute(res, content)
}

func WriteRedirect(res http.ResponseWriter, url string) {
	res.Header().Set("Location", url)
	res.WriteHeader(302)
}

func QueryString(req *http.Request, key string, defValue string) string {
	var querys = req.URL.Query()
	var qs1, _ = querys[key]
	if qs1 != nil {
		if len(qs1) > 0 {
			return qs1[0]
		}
	}

	return defValue
}
