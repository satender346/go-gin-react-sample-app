package routes

import (
	L "app/controllers"
	"github.com/gin-gonic/gin"
)



// InitializeRoutes : Initializes routes
func InitializeRoutes(router *gin.Engine) {
	r := router.Group("")
	// No Authentication Endpoints
	{
		// Handle the Index Page route
		r.GET("/", L.Index)
	}
}
