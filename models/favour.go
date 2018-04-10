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

func (f *Favour) Create() (err error) {
	err = db.Create(f).Error
	return
}

func GetFavourById(fid uint) (f *Favour, err error) {
	f = new(Favour)
	err = db.First(f, fid).Error
	return
}
