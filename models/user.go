package models

import (
	"github.com/jinzhu/gorm"
)

//User is
type User struct {
	gorm.Model
	OpenID   string `gorm:"not null;unique_index"`
	Name     string `gorm:"not null"`
	Avatar   string `gorm:"not null"`
	Gender   int    `gorm:"not null"`
	Country  string `gorm:"not null"`
	Province string `gorm:"not null"`
	City     string `gorm:"not null"`

	Favours   []Favour
	Timespace []Timespace `gorm:"many2many:timespace_users"`
	Labels    []Label
	Tips      []Tips
	Chats     []Chat
	Albums    []Album
	Pictures  []Picture
}

func (u *User) Create() (err error) {
	err = db.Create(u).Error
	return
}

func (u *User) GetTimespace() (timespace []Timespace, err error) {
	timespace = make([]Timespace, 0)
	err = db.Model(u).Related(&timespace).Error
	return
}

func (u *User) JoinTimespace(tid uint) (err error) {
	t, err := GetTimespaceById(tid)
	if err != nil {
		return err
	}
	err = db.Model(u).Association("Timespace").Append(t).Error
	return
}
