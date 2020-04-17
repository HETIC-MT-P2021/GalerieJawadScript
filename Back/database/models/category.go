package models

import (
	"../../lib/common"
	"github.com/jinzhu/gorm"
)

type Category struct {
	Name   string
	gorm.Model

}

func (p Category) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"name":       p.Name,
		"date_created": p.CreatedAt,
	}
}