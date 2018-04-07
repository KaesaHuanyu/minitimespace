package models

import (
	"github.com/jinzhu/gorm"
)

type Favour struct {
	gorm.Model
	Type     int  `gorm:"not null"`
	UserID   uint `gorm:"not null"`
	TargetID uint `gorm:"not null"`
}
