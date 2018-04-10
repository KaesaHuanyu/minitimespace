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

func (a *Album) Create() (err error) {
	err = db.Create(a).Error
	return
}

func GetAlbumById(aid uint) (a *Album, err error) {
	a = new(Album)
	err = db.First(a, aid).Error
	return
}

func (a *Album) GetPictures() (pics []Picture, err error) {
	pics = make([]Picture, 0)
	err = db.Model(a).Related(&pics, "Pictures").Error
	return
}

func (a *Album) GetFavours() (favs []Favour, err error) {
	favs = make([]Favour, 0)
	err = db.Model(a).Where("type = ?", AlbumFavour).Related(&favs, "Favours").Error
	return
}
