package Images

import (
	"fmt"
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/database/models"
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/lib/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type Image = models.Image
type JSON = common.JSON


func create(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	type BindImage struct {
		Name  string                `form:"name" binding:"required"`
		Description string                `form:"description" binding:"required"`
		UuidFile *multipart.FileHeader `form:"images" binding:"required"`
	}

	var ImageBindings BindImage

	// Bind file
	if err := c.ShouldBind(&ImageBindings); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	file := ImageBindings.UuidFile

	var extensions = filepath.Ext(file.Filename)
	var uuidImage = uuid.New()

	Image := Image{Name: ImageBindings.Name, Description: ImageBindings.Description, UuidFile: uuidImage.String()+extensions}



	if err := c.SaveUploadedFile(file, "./images/"+uuidImage.String()+extensions); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	db.NewRecord(Image)
	db.Create(&Image)
	c.JSON(200, Image.Serialize())
}
