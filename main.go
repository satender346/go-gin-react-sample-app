// main.go

package main

import (
	L "app/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	L.InitializeRoutes(router)
	// Start serving the application
	router.Run(":7070")
}
