package apiv1

import (
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/api/v1.0/Category"
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/api/v1.0/Images"
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/api/v1.0/Tags"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		Images.ApplyRoutes(v1)
		category.ApplyRoutes(v1)
		tags.ApplyRoutes(v1)
	}
}