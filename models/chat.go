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

func (c *Chat) Create() (err error) {
	err = db.Create(c).Error
	return
}

func GetChatById(cid uint) (c *Chat, err error) {
	c = new(Chat)
	err = db.First(c, cid).Error
	return
}

func (c *Chat) GetFather() (f *Chat, err error) {
	f = new(Chat)
	err = db.Model(c).Related(f, "Father").Error
	return
}

func (c *Chat) GetFavours() (favs []Favour, err error) {
	favs = make([]Favour, 0)
	err = db.Model(c).Where("type = ?", ChatFavour).Related(&favs, "Favours").Error
	return
}
