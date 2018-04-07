package models

import (
	"github.com/jinzhu/gorm"
)

type Label struct {
	gorm.Model
	Name   string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}
