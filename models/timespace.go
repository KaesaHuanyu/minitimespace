package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Timespace struct {
	gorm.Model
	Topic     string    `gorm:"not null;size:32"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`

	UserID uint    `gorm:"not null"`
	Users  []User  `gorm:"many2many:timespace_users"`
	Labels []Label `gorm:"many2many:timespace_labels;"`
	Tips   []Tips
	Chats  []Chat
	Albums []Album
}

func (t *Timespace) Create() (err error) {
	err = db.Create(t).Error
	return
}

func GetTimespaceById(tid uint) (timespace *Timespace, err error) {
	timespace = new(Timespace)
	err = db.First(timespace).Error
	return
}

func GetCreatedTimespaceByUid(uid uint) (timespace []Timespace, err error) {
	timespace = make([]Timespace, 0)
	err = db.Find(&timespace, "user_id = ?", uid).Error
	return
}

func (t *Timespace) GetUsers() (users []User, err error) {
	users = make([]User, 0)
	err = db.Model(t).Related(&users).Error
	return
}

func (t *Timespace) AddUser(uid uint) (err error) {
	return
}
