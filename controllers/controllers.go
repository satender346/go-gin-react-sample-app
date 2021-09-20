package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index : Returns Index page
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",
		gin.H{
			"title":         "Fleet Home",
		},
	)
}
