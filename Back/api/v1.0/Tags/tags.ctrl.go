package tags

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	Category := r.Group("/tags")
	{
		Category.POST("/", create)
		Category.GET("/", list)
		Category.GET("/:id", read)
		Category.DELETE("/:id", remove)
		Category.PATCH("/:id", update)
	}
}

