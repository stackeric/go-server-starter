// routes.go

package main

import (
	"github.com/stackeric/gostarter/handlers"
)

func initializeRoutes() {

	// Handle the index route
	//router.GET("/", showIndexPage)

	testRoutes := router.Group("/v1")
	{
		testRoutes.GET("/user", user.List)
	}

}
