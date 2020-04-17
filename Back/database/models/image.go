package models

import (
	"../../lib/common"
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	Name     string
	Category string
	UuidFile string
}

func (u *Image) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"username":     u.Name,
		"display_name": u.Category,
	}
}

func (u *Image) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Name = m["name"].(string)
	u.Category = m["Category"].(string)
	u.UuidFile = m["UuidFile"].(string)
}