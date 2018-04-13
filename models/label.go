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

func GetLabelsByIds(lids []uint) (labels []Label, err error) {
	labels = make([]Label, 0)
	err = db.Where(lids).Find(labels).Error
	return
}

func (l *Label) GetTimespace() (ts []Timespace, err error) {
	ts = make([]Timespace, 0)
	err = db.Model(l).Where("user_id = ?", l.UserID).Related(&ts, "Timespace").Error
	return
}

func (l *Label) CountTimespace() (count int) {
	return db.Model(l).Where("user_id = ?", l.UserID).Association("Timespace").Count()
}
