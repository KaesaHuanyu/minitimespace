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

func (p *Picture) Create() (err error) {
	err = db.Create(p).Error
	return
}

func GetPictureById(pid uint) (p *Picture, err error) {
	p = new(Picture)
	err = db.First(p, pid).Error
	return
}
