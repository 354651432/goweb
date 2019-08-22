# goweb

## routes demo
```go
func registerRoutes() {

	// access /
	router.Redirect("/", "/home")

	// access /home
	router.Get("/home", func(res http.ResponseWriter, req *http.Request) {
		router.WriteContent(res, "<h1>it works</h1?")
	})

	// access /about
	router.Get("/about", func(res http.ResponseWriter, req *http.Request) {
		router.WriteContent(res, "<h1>go router code by dual@dr.h</h1?")
	})

	router.Get("/users", func(res http.ResponseWriter, req *http.Request) {
		router.WriteJson(res, []interface{}{
			map[string]interface{}{
				"name": "dual",
				"age":  20,
				"id":   "1",
			},
			map[string]interface{}{
				"name": "ddd",
				"age":  20,
				"id":   "2",
			},
		})
	})

	// access /users/2
	router.Get("/users/@id:[0-9]+", func(res http.ResponseWriter, req *http.Request) {
		id := router.GetParams("id")
		router.WriteJson(res, map[string]interface{}{
			"code":    200,
			"user_id": id,
			"message": "user detail page",
		})
	})

	// access /users/create
	router.Get("/users/create", func(res http.ResponseWriter, req *http.Request) {
		router.WriteContent(res, "<h1>create user page</h1?")
	})

	// access /users/2/scores/3
	router.Get("/users/@user_id:[0-9]+/scores/@score_id", func(res http.ResponseWriter, req *http.Request) {
		userId := router.GetParams("user_id")
		scoreId := router.GetParams("score_id")

		_, _ = fmt.Fprintf(res, "<h2>score of user page user id is %v and score id is %v</h2>", userId, scoreId)
	})
}
```


## functions have not been achieved
0. read config file - supported
0. mysql support - supported
0. static file proc - supported
0. view
0. router like nginx
0. artisan support - supported
