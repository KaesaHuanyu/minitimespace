package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//Invitation is the open black Invitation
type Invitation struct {
	gorm.Model
	Accepted bool `gorm:"not null" json:"accepted,omitempty"`

	ActivityID uint `gorm:"not null" json:"activity_id,omitempty"`
	SenderID   uint `gorm:"not null" json:"sender_id,omitempty"`
	ReceiverID uint `gorm:"not null" json:"receiver_id,omitempty"`
}

func (i *Invitation) Create() (err error) {
	if i.ID != 0 {
		err = errors.New("Invitation is not NIL")
		return
	}
	err = db.Create(i).Error
	return
}

func (i *Invitation) Delete() (err error) {
	if i.ID == 0 {
		err = errors.New("Invitation not Found")
		return
	}
	err = db.Delete(i).Error
	return
}

func (i *Invitation) Accept(status bool) (err error) {
	if i.ID == 0 {
		err = errors.New("Invitation not Found")
		return
	}
	err = db.Model(i).Update("accepted", status).Error
	return
}

func GetInvitationByID(id uint) (i *Invitation, err error) {
	i = new(Invitation)
	err = db.First(i, id).Error
	return
}

func (i *Invitation) GetActivity() (a *Activity, err error) {
	a = new(Activity)
	err = db.First(a, i.ActivityID).Error
	return
}

func (i *Invitation) GetReceiver() (u *User, err error) {
	u = new(User)
	err = db.First(u, i.ReceiverID).Error
	return
}

func (i *Invitation) GetSender() (u *User, err error) {
	u = new(User)
	err = db.First(u, i.SenderID).Error
	return
}
