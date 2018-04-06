package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

//Activity 是开黑活动的抽象
type Activity struct {
	gorm.Model
	StartTime time.Time `gorm:"not null" json:"start_time,omitempty"`
	Status    int       `gorm:"not null" json:"is_end,omitempty"`
	Topic     string    `gorm:"not null" json:"topic,omitempty"`
	Content   int       `gorm:"not null" json:"content,omitempty"`
	Info      string    `gorm:"not null" json:"info,omitempty"` //most 156 bytes, 52 字
	Addr      string    `gorm:"not null" json:"addr,omitempty"`

	TeamID       uint         `gorm:"not null" json:"team_id,omitempty"`
	InitiatorID  uint         `gorm:"not null" json:"initiator_id,omitempty"`
	Participants []User       `gorm:"many2many:users_acticities" json:"participants,omitempty"`
	Invitations  []Invitation `gorm:"ForeignKey:ActivityID" json:"invitation,omitempty"`
	DoveVote     DoveVote     `gorm:"ForeignKey:ActivityID" json:"dove_vote,omitempty"`
	MvpVote      MvpVote      `gorm:"ForeignKey:ActivityID" json:"mvp_vote,omitempty"`
}

//Create 将当前活动插入到数据库
func (a *Activity) Create() (err error) {
	if a.ID != 0 {
		err = errors.New("Activity is not NIL")
		return
	}
	err = db.Create(a).Error
	return
}

//Delete 将当前活动从数据库中删除
func (a *Activity) Delete() (err error) {
	if a.ID == 0 {
		err = errors.New("Activity not Found")
		return
	}
	err = db.Delete(a).Error
	return
}

//AddParticipant 当前活动调用该函数增加参与者
func (a *Activity) AddParticipant(user *User) (err error) {
	if a.ID == 0 {
		err = errors.New("Activity not Found")
		return
	}
	if user == nil {
		err = errors.New("User not Found")
		return
	}
	if user.ID == 0 {
		err = errors.New("User not Found")
		return
	}

	//Check Content
	if len(a.Participants) <= a.Content {
		err = db.Model(a).Association("Participants").Append(user).Error
		return
	}
	err = errors.New("Reach the upper limit")
	return
}

func (a *Activity) Cancel() (err error) {
	if a.Status != ActivityCreated {
		err = errors.New("status wrong")
		return
	}
	c := db.New()
	err = c.Model(a).UpdateColumn("status", ActivityCanceled).Error
	return
}

func (a *Activity) Start() (err error) {
	if a.Status != ActivityCreated {
		err = errors.New("status wrong")
		return
	}
	c := db.New()
	err = c.Model(a).UpdateColumn("status", ActivityStarted).Error
	return
}

//EndActivity 结束当前活动
func (a *Activity) Success() (err error) {
	if a.Status == ActivityStarted {
		err = errors.New("status wrong")
		return
	}
	c := db.New()
	err = c.Model(a).Update("status", ActivitySuccessed).Error
	return
}

//GetActivityByID 通过id获取活动
func GetActivityByID(id uint) (a *Activity, err error) {
	a = new(Activity)
	err = db.First(a, id).Error
	return
}

//GetActivitiesByTeamID 通过teamID获取活动列表
func GetActivitiesByTeamID(id, limit, page uint, status bool) (activities []Activity, err error) {
	activities = []Activity{}
	if limit == 0 || page == 0 {
		err = db.Find(&activities, "team_id = ?", id).Where("is_end = ?", status).Error
		return
	}
	err = db.Find(&activities, "team_id = ?", id).Where("is_end = ?", status).
		Offset((page - 1) * limit).Limit(limit).Error
	return
}

//GetActivitiesByInitiatorID 通过用户id获取活动列表
func GetActivitiesByInitiatorID(id, limit, page uint, status bool) (activities []Activity, err error) {
	activities = []Activity{}
	if limit == 0 || page == 0 {
		err = db.Find(&activities, "initiator_id = ?", id).Where("is_end = ?", status).Error
		return
	}
	err = db.Find(&activities, "initiator_id = ?", id).Where("is_end = ?", status).
		Offset((page - 1) * limit).Limit(limit).Error
	return
}

//GetInitiator 获取活动的组织用户
func (a *Activity) GetInitiator() (initiator *User, err error) {
	initiator = new(User)
	err = db.Find(initiator, a.InitiatorID).Error
	return
}

//GetParticipants 获取所有的参与者，按id排列
func (a *Activity) GetParticipants() (participants []User, err error) {
	participants = []User{}
	err = db.Model(a).Association("Participants").Find(&participants).Error
	return
}

//GetInvitations 获取所有邀请函，按id排列
func (a *Activity) GetInvitations() (invitations []Invitation, err error) {
	invitations = []Invitation{}
	err = db.Model(a).Related(&invitations, "Invitations").Error
	return
}

//GetDoveVote 获取鸽子投票
func (a *Activity) GetDoveVote() (dv *DoveVote, err error) {
	dv = new(DoveVote)
	err = db.Model(a).Related(&dv, "DoveVote").Error
	return
}

//GetMvpVote 获取MVP投票
func (a *Activity) GetMvpVote() (mv *MvpVote, err error) {
	mv = new(MvpVote)
	err = db.Model(a).Related(mv, "MvpVote").Error
	return
}
