package models

import (
	"github.com/jinzhu/gorm"
)

type Timespace struct {
	gorm.Model
	UserID uint
	Topic  string
	Label  string
	News   bool
}
