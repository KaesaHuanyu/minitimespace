package handlers

import (
	"minitimespace/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

//CURD

//CreateUser create the user accout
func (h *Handler) CreateUser(c echo.Context) (err error) {
	resp := responses()
	//Get data from request
	u := new(models.User)
	err = c.Bind(u)
	if err != nil {
		resp.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, resp, "	")
	}
	//Input Check
	if len(u.OpenID) == 0 || len(u.Name) == 0 {
		resp.Error = "未获取到 open_id 或者 name 参数。 "
		return c.JSONPretty(http.StatusBadRequest, resp, "	")
	}

	//Check whether this openID registered
	user, err := models.GetUserByOpenID(u.OpenID)
	if err != nil {
		//If no record, create it.
		if err == gorm.ErrRecordNotFound {
			err = u.Create()
			if err != nil {
				resp.Error = err.Error()
				return c.JSONPretty(http.StatusInternalServerError, resp, "	")
			}
		} else {
			resp.Error = err.Error()
			return c.JSONPretty(http.StatusInternalServerError, resp, "	")
		}
	} else {
		resp.Warning = "该用户已注册。 "
		resp.Data["user"] = user
		return c.JSONPretty(http.StatusAccepted, resp, "	")
	}

	resp.Info = fmt.Sprintf("创建用户成功，ID: %d, OpenID: %s", int(u.ID), u.OpenID)
	resp.Data["user"] = u
	return c.JSONPretty(http.StatusCreated, resp, "	")
}

func (h *Handler) GetUsers(c echo.Context) (err error) {
	r := responses()
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	if len(limit) == 0 {
		users, err := models.GetAllUsers()
		if err != nil {
			r.Error = err.Error()
			return c.JSONPretty(http.StatusInternalServerError, r, "	")
		}
		r.Info = fmt.Sprintf("总共有%d位用户", len(users))
		r.Data["users"] = users
		r.Data["count"] = len(users)
		return c.JSONPretty(http.StatusOK, r, "	")
	}
	l, err := strconv.Atoi(limit)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	if l < 1 {
		r.Error = "limit参数小于1"
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	if p < 1 {
		r.Error = "page参数小于1"
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	users, err := models.GetUsersPage(uint(l), uint(p))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	r.Info = fmt.Sprintf("第%d页，该页有%d位用户", p, len(users))
	r.Data["users"] = users
	r.Data["count"] = len(users)
	return c.JSONPretty(http.StatusOK, r, "	")
}

func (h *Handler) GetUser(c echo.Context) (err error) {
	r := responses()
	userID := c.Param("userID")
	if len(userID) == 0 {
		r.Error = "未获取到用户id"
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	user, err := models.GetUserByID(uint(id))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	r.Data["user"] = user
	return c.JSONPretty(http.StatusOK, r, "	")
}

func (h *Handler) DeleteUser(c echo.Context) (err error) {
	r := responses()
	userID := c.Param("userID")
	if len(userID) == 0 {
		r.Error = "未获取到用户id"
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	user, err := models.GetUserByID(uint(id))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	err = user.Delete()
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	return c.JSONPretty(http.StatusNoContent, r, "	")
}

// GetUserTeams request{
// 	limit
// 	page
// 	created
// }
func (h *Handler) GetUserTeams(c echo.Context) (err error) {
	r := responses()
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil || userID == 0 {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	request := struct {
		limit   uint
		page    uint
		orderBy string
		created bool
	}{}
	err = c.Bind(request)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	user, err := models.GetUserByID(uint(userID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	teams := []models.Team{}
	if request.created {
		teams, err = user.GetCreatedTeams(uint(request.limit), uint(request.page), request.orderBy)
	} else {
		teams, err = user.GetJoinedTeams(uint(request.limit), uint(request.page), request.orderBy)
	}
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}

	r.Info = fmt.Sprintf("第%d页，共有%d支队伍", request.page, len(teams))
	r.Data["teams"] = teams
	return c.JSONPretty(http.StatusOK, r, "	")
}

// GetUserActivities request{
// 	limit
// 	page
// 	orderBy
// 	created
// }
func (h *Handler) GetUserActivities(c echo.Context) (err error) {
	r := responses()
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil || userID == 0 {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	request := struct {
		limit   uint
		page    uint
		orderBy string
		created bool
		end     bool
	}{}
	err = c.Bind(request)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	user, err := models.GetUserByID(uint(userID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	activities := []models.Activity{}
	if request.created {
		activities, err = user.GetCreatedActivities(uint(request.limit), uint(request.page), request.orderBy)
	} else {
		activities, err = user.GetJoinedActivities(uint(request.limit), uint(request.page),
			request.end, request.orderBy)
	}
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}

	r.Info = fmt.Sprintf("第%d页，共有%d次活动", request.page, len(activities))
	r.Data["activities"] = activities
	return c.JSONPretty(http.StatusOK, r, "	")
}

func (h *Handler) UserJoinsTeam(c echo.Context) (err error) {
	r := responses()
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	teamID, err := strconv.Atoi(c.Param("teamID"))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	if teamID <= 0 || userID <= 0 {
		r.Error = "未获取到合适的 userID 或 teamID"
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}

	user, err := models.GetUserByID(uint(userID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	team, err := models.GetTeamByID(uint(teamID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	err = user.JoinTeam(team)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}

	r.Info = fmt.Sprintf("用户【 %s 】已成功加入队伍【 %s 】", user.Name, team.TeamName)
	r.Data["user"] = user
	r.Data["team"] = team
	return c.JSONPretty(http.StatusOK, r, "	")
}

func (h *Handler) UserExitsTeam(c echo.Context) (err error) {
	r := responses()
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	teamID, err := strconv.Atoi(c.Param("teamID"))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}

	if teamID <= 0 || userID <= 0 {
		r.Error = "未获取到合适的 userID 或 teamID"
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}

	user, err := models.GetUserByID(uint(userID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	team, err := models.GetTeamByID(uint(teamID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	err = user.ExitTeam(team)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}

	r.Info = fmt.Sprintf("用户【 %s 】已成功加入队伍【 %s 】", user.Name, team.TeamName)
	r.Data["user"] = user
	r.Data["team"] = team
	return c.JSONPretty(http.StatusOK, r, "	")
}

func (h *Handler) UserJoinsActivity(c echo.Context) (err error) {
	r := responses()
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	activityID, err := strconv.Atoi(c.Param("activityID"))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	if activityID <= 0 || userID <= 0 {
		r.Error = "未获取到合适的 userID 或 activityID"
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}

	user, err := models.GetUserByID(uint(userID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	activity, err := models.GetActivityByID(uint(activityID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	err = user.JoinActivity(activity)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	r.Info = fmt.Sprintf("用户【 %s 】已成功加入活动【 %d 】", user.Name, activity.ID)
	r.Data["user"] = user
	r.Data["activity"] = activity
	return c.JSONPretty(http.StatusOK, r, "	")
}

func (h *Handler) UserExitsActivity(c echo.Context) (err error) {
	r := responses()
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	activityID, err := strconv.Atoi(c.Param("activityID"))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	if activityID <= 0 || userID <= 0 {
		r.Error = "未获取到合适的 userID 或 activityID"
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}

	user, err := models.GetUserByID(uint(userID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	activity, err := models.GetActivityByID(uint(activityID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	err = user.ExitActivity(activity)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	r.Info = fmt.Sprintf("用户【 %s 】已成功退出活动【 %d 】", user.Name, activity.ID)
	r.Data["user"] = user
	r.Data["activity"] = activity
	return c.JSONPretty(http.StatusOK, r, "	")
}
