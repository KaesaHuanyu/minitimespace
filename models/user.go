package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//User is
type User struct {
	gorm.Model
	OpenID   string `gorm:"not null;unique_index"`
	Name     string `gorm:"not null;"`
	Avatar   string `gorm:"not null"`
	Gender   int    `gorm:"not null"`
	City     string `gorm:"not null"`
	Province string `gorm:"not null"`

	Teams      []Team     `gorm:"many2many:users_teams" json:"joined_teams,omitempty"`
	Activities []Activity `gorm:"many2many:users_acticities" json:"joined_activities,omitempty"`
}

//User's CURD

//Create 将当前用户插入到数据库
func (u *User) Create() (err error) {
	if u.ID != 0 {
		err = errors.New("UserID is not NIL")
		return
	}
	err = db.Create(u).Error
	return
}

//Delete 将当前用户从数据库中软删除
func (u *User) Delete() (err error) {
	if u.ID == 0 {
		err = errors.New("User not Found")
		return
	}
	err = db.Delete(u).Error
	return
}

//GetUsersPage 接收一个limit和一个offset，返回用户列表
func GetUsersPage(limit, page uint) (users []User, err error) {
	users = []User{}
	if limit == 0 || page == 0 {
		err = db.Find(&users).Error
	} else {
		err = db.Offset((page - 1) * limit).Limit(limit).Find(&users).Error
	}
	return
}

//GetAllUsers 返回所有的用户
func GetAllUsers() (users []User, err error) {
	users = []User{}
	err = db.Find(&users).Error
	return
}

//GetUserByID 接收一个ID返回该用户
func GetUserByID(id uint) (user *User, err error) {
	user = new(User)
	err = db.First(user, id).Error
	return
}

//GetUserByOpenID 接收一个OpenID返回该用户
func GetUserByOpenID(OpenID string) (user *User, err error) {
	user = new(User)
	c := db.New()
	err = c.First(user, "open_id = ?", OpenID).Error
	return
}

//GetJoinedTeams 返回当前用户加入的群列表
func (u *User) GetJoinedTeams(limit, page uint, orderBy string) (teams []Team, err error) {
	teams = []Team{}
	if limit == 0 || page == 0 {
		err = db.Model(u).Order(orderBy).Association("JoinedTeams").Find(&teams).Error
		return
	}
	err = db.Model(u).Order(orderBy).Offset((page - 1) * limit).
		Limit(limit).Association("JoinedTeams").Find(&teams).Error
	return
}

//GetCreatedTeams 返回当前用户创建的群列表
func (u *User) GetCreatedTeams(limit, page uint, orderBy string) (teams []Team, err error) {
	teams = []Team{}
	if limit == 0 || page == 0 {
		err = db.Model(u).Order(orderBy).Association("CreatedTeams").Find(&teams).Error
		return
	}
	err = db.Model(u).Order(orderBy).Offset((page-1)*limit).
		Limit(limit).Related(&teams, "CreatedTeams").Error
	return
}

//GetJoinedActivities 返回当前用户参加的活动列表，可选择已结束/未结束
func (u *User) GetJoinedActivities(limit, page uint, isEnd bool, order string) (activities []Activity, err error) {
	activities = []Activity{}
	if limit == 0 || page == 0 {
		err = db.Model(u).Order(order).Where("is_end = ?", isEnd).
			Association("JoinedActivities").Find(&activities).Error
		return
	}
	err = db.Model(u).Order(order).Offset((page-1)*limit).Limit(limit).
		Where("is_end = ?", isEnd).Association("JoinedActivities").Find(&activities).Error
	return
}

//GetCreatedActivities 返回当前用户创建的活动列表
func (u *User) GetCreatedActivities(limit, page uint, order string) (activities []Activity, err error) {
	activities = []Activity{}
	if limit == 0 || page == 0 {
		err = db.Model(u).Order(order).Association("CreatedActivities").Find(&activities).Error
		return
	}
	err = db.Model(u).Order(order).Offset((page-1)*limit).
		Limit(limit).Related(&activities, "CreatedActivities").Error
	return
}

//GetDovePoints 返回当前用户获得的鸽子积分列表
func (u *User) GetDovePoints() (dovePoints []DovePoint, err error) {
	dovePoints = []DovePoint{}
	err = db.Model(u).Order(ORDER_BY_CREATE_DESC).Related(&dovePoints, "DovePoints").Error
	return
}

//GetMvpPoints 返回当前用户获得的MVP积分列表
func (u *User) GetMvpPoints() (mvpPoints []MvpPoint, err error) {
	mvpPoints = []MvpPoint{}
	err = db.Model(u).Order(ORDER_BY_CREATE_DESC).Related(&mvpPoints, "MvpPoints").Error
	return
}

//JoinTeam 用户通过调用该函数加入某队伍
func (u *User) JoinTeam(team *Team) (err error) {
	if u.ID == 0 {
		err = errors.New("User not Found")
		return
	}
	if team == nil {
		err = errors.New("Team not Found")
		return
	}
	if team.ID == 0 {
		err = errors.New("Team not Found")
		return
	}

	err = db.Model(u).Association("JoinedTeams").Append(team).Error
	return
}

func (u *User) ExitTeam(team *Team) (err error) {
	if u.ID == 0 {
		err = errors.New("User not Found")
		return
	}
	if team == nil {
		err = errors.New("Team not Found")
		return
	}
	if team.ID == 0 {
		err = errors.New("Team not Found")
		return
	}

	err = db.Model(u).Association("JoinedTeams").Delete(team).Error
	return
}

//JoinActivity 用户调用这个函数加入某队伍
func (u *User) JoinActivity(activity *Activity) (err error) {
	if u.ID == 0 {
		err = errors.New("User not Found")
		return
	}
	if activity == nil {
		err = errors.New("Activity not Found")
		return
	}
	if activity.ID == 0 {
		err = errors.New("Activity not Found")
		return
	}

	activity.Participants, err = activity.GetParticipants()
	if err != nil {
		return
	}
	if len(activity.Participants) < activity.Content {
		err = db.Model(u).Association("JoinedActivities").Append(activity).Error
	} else {
		err = errors.New("Reach the upper limit")
	}
	return
}

func (u *User) ExitActivity(activity *Activity) (err error) {
	if u.ID == 0 {
		err = errors.New("User not Found")
		return
	}
	if activity == nil {
		err = errors.New("Activity not Found")
		return
	}
	if activity.ID == 0 {
		err = errors.New("Activity not Found")
		return
	}

	err = db.Model(u).Association("JoinedActivities").Delete(activity).Error
	return
}

func (u *User) GetReceivedInvitations(limit, page uint, order string) (invitations []Invitation, err error) {
	invitations = []Invitation{}
	if limit == 0 || page == 0 {
		err = db.Model(u).Order(order).Related(&invitations, "ReceivedInvitations").Error
		return
	}
	err = db.Model(u).Order(order).Offset((page-1)*limit).
		Limit(limit).Related(&invitations, "ReceivedInvitations").Error
	return
}
