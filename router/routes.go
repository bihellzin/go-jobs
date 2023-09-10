package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET opening",
			})
		})
	}
}
