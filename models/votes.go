package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type (
	DoveVote struct {
		gorm.Model

		ActivityID  uint         `gorm:"not null" json:"activity_id,omitempty"`
		DoveRecords []DoveRecord `gorm:"ForeignKey:DoveVoteID" json:"dove_records,omitempty"`
		WinnerID    uint         `json:"winner_id,omitempty"`
	}

	MvpVote struct {
		gorm.Model

		ActivityID uint        `gorm:"not null" json:"activity_id,omitempty"`
		MvpRecords []MvpRecord `gorm:"ForeignKey:MvpVoteID" json:"mvp_records,omitempty"`
		WinnerID   uint        `json:"winner_id,omitempty"`
	}
)

func (dv *DoveVote) Create() (err error) {
	if dv.ID != 0 {
		err = errors.New("DoveVote is not NIL")
		return
	}
	err = db.Create(dv).Error
	return
}

func (mv *MvpVote) Create() (err error) {
	if mv.ID != 0 {
		err = errors.New("MvpVote is not NIL")
		return
	}
	err = db.Create(mv).Error
	return
}

func (dv *DoveVote) Delete() (err error) {
	if dv.ID == 0 {
		err = errors.New("DoveVote not Found")
		return
	}
	err = db.Delete(dv).Error
	return
}

func (mv *MvpVote) Delete() (err error) {
	if mv.ID == 0 {
		err = errors.New("MvpVote not Found")
		return
	}
	err = db.Delete(mv).Error
	return
}

func GetDoveVoteByID(id uint) (dv *DoveVote, err error) {
	dv = new(DoveVote)
	err = db.First(dv, id).Error
	return
}

func GetMvpVoteByID(id uint) (mv *MvpVote, err error) {
	mv = new(MvpVote)
	err = db.First(mv, id).Error
	return
}

func (dv *DoveVote) GetDoveRecords() (records []DoveRecord, err error) {
	records = []DoveRecord{}
	err = db.Model(dv).Related(&records).Error
	return
}

func (mv *MvpVote) GetRecords() (records []MvpRecord, err error) {
	records = []MvpRecord{}
	err = db.Model(mv).Related(&records).Error
	return
}

func (dv *DoveVote) GetActivity() (a *Activity, err error) {
	a = new(Activity)
	err = db.First(a, dv.ActivityID).Error
	return
}

func (mv *MvpVote) GetActivity() (a *Activity, err error) {
	a = new(Activity)
	err = db.First(a, mv.ActivityID).Error
	return
}

func (dv *DoveVote) GetWinner() (u *User, err error) {
	u = new(User)
	err = db.First(u, dv.WinnerID).Error
	return
}

func (mv *MvpVote) GetWinner() (u *User, err error) {
	u = new(User)
	err = db.First(u, mv.WinnerID).Error
	return
}
