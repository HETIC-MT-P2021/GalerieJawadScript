package models

import (
	"../../lib/common"
	"github.com/jinzhu/gorm"
)

type Tags struct {
	Name   string
	gorm.Model

}

func (p Tags) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"name":       p.Name,
		"date_created": p.CreatedAt,
	}
}