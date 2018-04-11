package models

import (
	"github.com/jinzhu/gorm"
)

type Timespace struct {
	gorm.Model
	Topic     string  `gorm:"not null;size:32"`
	Labels    []Label `gorm:"many2many:timespace_labels"`
	Desc      string
	StartTime string `gorm:"not null"`
	EndTime   string `gorm:"not null"`

	UserID uint   `gorm:"not null"`
	Users  []User `gorm:"many2many:timespace_users"`
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

func (t *Timespace) GetUsers() (users []User, err error) {
	users = make([]User, 0)
	err = db.Model(t).Related(&users, "Users").Error
	return
}

func (t *Timespace) GetTips() (tips []Tips, err error) {
	tips = make([]Tips, 0)
	err = db.Model(t).Order("id desc").Related(&tips, "Tips").Error
	return
}

func (t *Timespace) GetChats() (chats []Chat, err error) {
	chats = make([]Chat, 0)
	err = db.Model(t).Order("id desc").Related(&chats, "Chats").Error
	return
}

func (t *Timespace) GetLabels() (labels []Label, err error) {
	labels = make([]Label, 0)
	err = db.Model(t).Related(&labels, "Labels").Error
	return
}

func (t *Timespace) GetAlbums() (albums []Album, err error) {
	albums = make([]Album, 0)
	err = db.Model(t).Order("id desc").Related(&albums, "Albums").Error
	return
}

func (t *Timespace) AddUser(u User) (err error) {
	err = db.Model(t).Association("Users").Append(u).Error
	return
}

func (t *Timespace) AddLabel(labels []Label) (err error) {
	err = db.Model(t).Association("Labels").Append(labels).Error
	return
}
