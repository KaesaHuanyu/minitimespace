package models

import (
	"github.com/jinzhu/gorm"
)

type Picture struct {
	gorm.Model
	URL string `gorm:"not null"`

	UserID  uint `gorm:"not null"`
	AlbumID uint `gorm:"not null"`
}
