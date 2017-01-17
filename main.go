// Package main is the CLI.
// You can use the CLI via Terminal.
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stackeric/gostarter/db"
	"github.com/stackeric/gostarter/middlewares"
)

const (
	// Port at which the server starts listening
	Port = "7000"
)

func init() {
	db.Connect()
}

var router *gin.Engine

func main() {

	// Configure
	gin.SetMode(gin.DebugMode)
	router = gin.Default()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	initializeRoutes()

	// Middlewares
	router.Use(middlewares.ErrorHandler)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
