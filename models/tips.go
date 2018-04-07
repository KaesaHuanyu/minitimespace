package models

import (
	"github.com/jinzhu/gorm"
)

type Tips struct {
	gorm.Model
	Content     string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	TimespaceID uint   `gorm:"not null"`

	Favours []Favour `gorm:"ForeignKey:TargetID"`
}
