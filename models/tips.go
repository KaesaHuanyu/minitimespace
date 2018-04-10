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

func (t *Tips) Create() (err error) {
	err = db.Create(t).Error
	return
}

func GetTipsById(tid uint) (t *Tips, err error) {
	t = new(Tips)
	err = db.First(t, tid).Error
	return
}

func (t *Tips) GetFavours() (favs []Favour, err error) {
	favs = make([]Favour, 0)
	err = db.Model(t).Where("type = ?", TipsFavour).Related(&favs, "Favours").Error
	return
}
