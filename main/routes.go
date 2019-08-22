package main

import (
	"app1/models"
	"app1/router"
	"fmt"
	"net/http"
)

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
		var users []models.User
		var offset = router.QueryString(req, "offset", "0")
		models.Open().Offset(offset).Limit(20).Find(&users)
		router.WriteJson(res, users)
	})

	// access /users/2
	router.Get("/users/@id:[0-9]+", func(res http.ResponseWriter, req *http.Request) {
		id := router.GetParams("id")
		var user models.User
		models.Open().First(&user, id)
		router.WriteJson(res, user)
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
