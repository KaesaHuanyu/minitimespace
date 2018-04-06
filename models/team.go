package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//Team for opening black
type Team struct {
	gorm.Model
	TeamName string `gorm:"not null;unique" json:"team_name,omitempty"`
	Info     string `json:"info,omitempty"`
	Avatar   string `json:"avatar,omitempty"`

	CaptainID  uint       `gorm:"not null" json:"captain_id,omitempty"`
	Teammates  []User     `gorm:"many2many:users_teams" json:"teammates,omitempty"`
	Activities []Activity `gorm:"ForeignKey:TeamID" json:"activities,omitempty"`
}

//Create 将当前队伍插入到数据库中
func (t *Team) Create() (err error) {
	if t.ID != 0 {
		err = errors.New("ID is not NIL")
		return
	}
	if t.CaptainID == 0 {
		err = errors.New("CaptainID not Found")
		return
	}
	err = db.Create(t).Error
	return
}

//Delete 将当前队伍从数据库中删除
func (t *Team) Delete() (err error) {
	if t.ID == 0 {
		err = errors.New("Team not Found")
		return
	}
	err = db.Delete(t).Error
	return
}

//AddTeammate 队伍通过调用该函数添加队员
func (t *Team) AddTeammate(user *User) (err error) {
	if t.ID == 0 {
		err = errors.New("Team not Found")
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

	err = db.Model(t).Association("Teammates").Append(user).Error
	return
}

//GetTeamByID 接收一个id，返回该队伍
func GetTeamByID(id uint) (team *Team, err error) {
	team = new(Team)
	err = db.First(team, id).Error
	return
}

//GetTeamByTeamName 接收一个队名，返回该队伍
func GetTeamByTeamName(name string) (team *Team, err error) {
	team = new(Team)
	err = db.First(team, "team_name = ?", name).Error
	return
}

//GetAllTeams 返回所有的队伍
func GetAllTeams() (teams []Team, err error) {
	teams = []Team{}
	err = db.Find(&teams).Error
	return
}

//GetTeamsPage 接收limit和offset，返回队伍列表
func GetTeamsPage(limit, page uint, order string) (teams []Team, err error) {
	teams = []Team{}
	if limit == 0 {
		err = db.Find(&teams).Order(order).Error
	} else {
		err = db.Offset((page - 1) * limit).Limit(limit).Find(&teams).Order(order).Error
	}
	return
}

//GetCaptain 返回当前队伍的队长
func (t *Team) GetCaptain() (user *User, err error) {
	user = new(User)
	err = db.First(user, t.CaptainID).Error
	return
}

//GetTeammates 返回当前队伍的所有队员
func (t *Team) GetTeammates() (users []User, err error) {
	users = []User{}
	err = db.Model(t).Association("Teammates").Find(&users).Error
	return
}

//GetActivities 返回当前队伍的所有已结束/未结束的活动
func (t *Team) GetActivities(isEnd bool) (activities []Activity, err error) {
	activities = []Activity{}
	err = db.Model(t).Where("is_end = ?", isEnd).Order(ORDER_BY_CREATE_DESC).
		Related(&activities, "Activities").Error
	return
}
