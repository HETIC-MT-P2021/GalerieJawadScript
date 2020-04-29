package models

import (
	"fmt"
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/lib/common"
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	Name        string
	Description string `sql:"type:text;"`
	Category    Category
	Tags        []Tags
	UuidFile    string
}

func (i Image) Serialize() common.JSON {
	return common.JSON{
		"name": i.Name,
		"description": i.Description,
		"category": i.Category.Serialize(),
		"uuid_file": fmt.Sprintf("/images/%s", i.UuidFile),
		"date_created": i.CreatedAt,
		"date_updated": i.UpdatedAt,
		"tags": i.Tags,
	}
}