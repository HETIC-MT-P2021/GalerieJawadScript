package category

import (
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/database/models"
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/lib/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Category = models.Category
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

	Category := Category{Name: requestBody.Name}
	db.NewRecord(Category)
	db.Create(&Category)
	c.JSON(200, Category.Serialize())
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")

	var categories []Category

	if cursor == "" {
		if err := db.Find(&categories).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		if err := db.Find(&categories).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(categories)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = categories[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var category Category

	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, Category{}.Serialize())
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")


	var category Category
	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}


	db.Delete(&category)
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

	var category Category
	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	category.Name = requestBody.Name
	db.Save(&category)
	c.JSON(200, category.Serialize())
}