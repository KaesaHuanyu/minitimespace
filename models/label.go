package models

import (
	"github.com/jinzhu/gorm"
)

type Label struct {
	gorm.Model
	Name      string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	Timespace []Timespace
}

func (l *Label) Create() (err error) {
	err = db.Create(l).Error
	return
}

func GetLabelById(lid uint) (l *Label, err error) {
	l = new(Label)
	err = db.First(l, lid).Error
	return
}

func (l *Label) GetTimespace() (ts []Timespace, err error) {
	ts = make([]Timespace, 0)
	err = db.Model(l).Related(&ts, "Timespace").Error
	return
}
