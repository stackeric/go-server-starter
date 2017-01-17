// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()

	// TODO: Handle it in a better way
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": c.Errors,
		})
	}
}
