package models

import (
	"github.com/jinzhu/gorm"
)

type Album struct {
	gorm.Model
	Name string

	UserID      uint `gorm:"not null"`
	TimespaceID uint `gorm:"not null"`
	Pictures    []Picture
	Favours     []Favour `gorm:"ForeignKey:TargetID"`
}
