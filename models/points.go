package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type (
	DovePoint struct {
		gorm.Model

		UserID uint `gorm:"not null" json:"user_id,omitempty"`
	}
	MvpPoint struct {
		gorm.Model

		UserID uint `gorm:"not null" json:"user_id,omitempty"`
	}
)

func (dp *DovePoint) Create() (err error) {
	if dp.ID != 0 {
		err = errors.New("DovePoint is not NIL")
	}
	err = db.Create(dp).Error
	return
}

func (mp *MvpPoint) Create() (err error) {
	if mp.ID != 0 {
		err = errors.New("MvpPoint is not NIL")
		return
	}
	err = db.Create(mp).Error
	return
}

func (dp *DovePoint) Delete() (err error) {
	if dp.ID == 0 {
		err = errors.New("DovePoint not Found")
		return
	}
	err = db.Delete(dp).Error
	return
}

func (mp *MvpPoint) Delete() (err error) {
	if mp.ID == 0 {
		err = errors.New("MvpPoint not Found")
		return
	}
	err = db.Delete(mp).Error
	return
}

func GetDovePointByID(id uint) (dp *DovePoint, err error) {
	dp = new(DovePoint)
	err = db.First(dp, id).Error
	return
}

func GetMvpPointByID(id uint) (mp *MvpPoint, err error) {
	mp = new(MvpPoint)
	err = db.First(mp, id).Error
	return
}

func (dp *DovePoint) GetUser() (u *User, err error) {
	u = new(User)
	err = db.First(u, dp.UserID).Error
	return
}

func (mp *MvpPoint) GetUser() (u *User, err error) {
	u = new(User)
	err = db.First(u, mp.UserID).Error
	return
}
