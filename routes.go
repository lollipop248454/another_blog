package main

import "another_blog/handlers"

func initializeRoutes() {

	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)

	router.GET("/article/view/:article_id", handlers.GetArticle)
}
