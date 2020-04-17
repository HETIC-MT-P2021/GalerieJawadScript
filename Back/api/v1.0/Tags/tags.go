package tags

import (
	"../../../database/models"
	"../../../lib/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Tags = models.Tags
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Name string `json:"Name" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	Tag := Tags{Name: requestBody.Name}
	db.NewRecord(Tag)
	db.Create(&Tag)
	c.JSON(200, Tag.Serialize())
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")

	var tags []Tags

	if cursor == "" {
		if err := db.Find(&tags).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		if err := db.Find(&tags).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(tags)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = tags[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var tags Tags

	if err := db.Where("id = ?", id).First(&tags).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, tags.Serialize())
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")


	var tag Tags
	if err := db.Where("id = ?", id).First(&tag).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}


	db.Delete(&tag)
	c.Status(204)
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")


	type RequestBody struct {
		Name string `json:"Name" binding:"required"`
	}

	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var tag Tags
	if err := db.Where("id = ?", id).First(&tag).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	tag.Name = requestBody.Name
	db.Save(&tag)
	c.JSON(200, tag.Serialize())
}