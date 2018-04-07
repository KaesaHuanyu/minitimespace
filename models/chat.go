package models

import (
	"github.com/jinzhu/gorm"
)

type Chat struct {
	gorm.Model
	Content  string `gorm:"not null"`
	Father   *Chat  `gorm:"ForeignKey:FatherID"`
	FatherID uint

	UserID  uint     `gorm:"not null"`
	Favours []Favour `gorm:"ForeignKey:TargetID"`
}
