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
	Labels    []Label
	Tips      []Tips
	Chats     []Chat
	Albums    []Album
	Pictures  []Picture
	Timespace []Timespace
}

func (u *User) Create() (err error) {
	err = db.Create(u).Error
	return
}

func GetUserById(uid uint) (u *User, err error) {
	u = new(User)
	err = db.First(u, uid).Error
	return
}

func GetUserByOpenId(openId string) (u *User, err error) {
	u = new(User)
	err = db.First(u, "open_id = ?", openId).Error
	return
}

func (u *User) GetFavours() (favs []Favour, err error) {
	favs = make([]Favour, 0)
	err = db.Model(u).Related(&favs, "Favours").Error
	return
}

func (u *User) GetLabels() (labels []Label, err error) {
	labels = make([]Label, 0)
	err = db.Model(u).Related(&labels, "Labels").Error
	return
}

func (u *User) GetTips() (tips []Tips, err error) {
	tips = make([]Tips, 0)
	err = db.Model(u).Related(&tips, "Tips").Error
	return
}

func (u *User) GetChats() (chats []Chat, err error) {
	chats = make([]Chat, 0)
	err = db.Model(u).Related(&chats, "Chats").Error
	return
}

func (u *User) GetAlbums() (albums []Album, err error) {
	albums = make([]Album, 0)
	err = db.Model(u).Related(&albums, "Albums").Error
	return
}

func (u *User) GetPictures() (pictures []Picture, err error) {
	pictures = make([]Picture, 0)
	err = db.Model(u).Related(&pictures, "Pictures").Error
	return
}

func (u *User) GetAddedTimespace() (timespace []Timespace, err error) {
	timespace = make([]Timespace, 0)
	err = db.Model(u).Related(&timespace, "Timespace").Error
	return
}
