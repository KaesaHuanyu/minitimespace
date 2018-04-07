package models

import (
	"github.com/jinzhu/gorm"
)

type Chat struct {
	gorm.Model
	Label string
}
