package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type (
	DoveRecord struct {
		gorm.Model

		DoveVoteID uint `gorm:"not null" json:"dove_vote_id,omitempty"`
		VoterID    uint `gorm:"not null" json:"voter_id,omitempty"`
		AdvocateID uint `gorm:"not null" json:"advocate_id,omitempty"`
	}

	MvpRecord struct {
		gorm.Model

		MvpVoteID  uint `gorm:"not null" json:"mvp_vote_id,omitempty"`
		VoterID    uint `gorm:"not null" json:"voter_id,omitempty"`
		AdvocateID uint `gorm:"not null" json:"advocate_id,omitempty"`
	}
)

func (r *DoveRecord) Create() (err error) {
	if r.ID != 0 {
		err = errors.New("DoveRecord is not NIL")
		return
	}
	err = db.Create(r).Error
	return
}

func (r *DoveRecord) Delete() (err error) {
	if r.ID == 0 {
		err = errors.New("DoveRecord not Found")
		return
	}
	err = db.Delete(r).Error
	return
}

func GetDoveRecordByID(id uint) (r *DoveRecord, err error) {
	r = new(DoveRecord)
	err = db.First(r, id).Error
	return
}

func (r *DoveRecord) GetDoveVote() (v *DoveVote, err error) {
	v = new(DoveVote)
	err = db.First(v, r.DoveVoteID).Error
	return
}

func (r *DoveRecord) GetVoter() (u *User, err error) {
	u = new(User)
	err = db.First(u, r.VoterID).Error
	return
}

func (r *DoveRecord) GetAdvocate() (u *User, err error) {
	u = new(User)
	err = db.First(u, r.AdvocateID).Error
	return
}

func (r *MvpRecord) Create() (err error) {
	if r.ID != 0 {
		err = errors.New("MvpRecord is not NIL")
		return
	}
	err = db.Create(r).Error
	return
}

func (r *MvpRecord) Delete() (err error) {
	if r.ID == 0 {
		err = errors.New("MvpRecord not Found")
		return
	}
	err = db.Delete(r).Error
	return
}

func GetMvpRecordByID(id uint) (r *MvpRecord, err error) {
	r = new(MvpRecord)
	err = db.First(r, id).Error
	return
}

func (r *MvpRecord) GetMvpVote() (v *MvpVote, err error) {
	v = new(MvpVote)
	err = db.First(v, r.MvpVoteID).Error
	return
}

func (r *MvpRecord) GetVoter() (u *User, err error) {
	u = new(User)
	err = db.First(u, r.VoterID).Error
	return
}

func (r *MvpRecord) GetAdvocate() (u *User, err error) {
	u = new(User)
	err = db.First(u, r.AdvocateID).Error
	return
}
