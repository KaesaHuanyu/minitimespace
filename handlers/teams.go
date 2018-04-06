package handlers

import (
	"minitimespace/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//CreateTeam create a new black team
func (h *Handler) CreateTeam(c echo.Context) (err error) {
	r := responses()
	team := new(models.Team)
	err = c.Bind(team)
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	if len(team.TeamName) == 0 {
		r.Error = fmt.Sprintln("未获取到合适的team_name")
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}
	_, err = models.GetUserByID(uint(team.CaptainID))
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusBadRequest, r, "	")
	}

	err = team.Create()
	if err != nil {
		r.Error = err.Error()
		return c.JSONPretty(http.StatusInternalServerError, r, "	")
	}
	r.Info = fmt.Sprintf("创建小黑队成功，队名：%s", team.TeamName)
	r.Data["team"] = team
	return c.JSONPretty(http.StatusCreated, r, "	")
}

func (h *Handler) UpdateTeam(c echo.Context) (err error) {
	return
}

// /teams/:id equals /teams?id=
func (h *Handler) GetTeams(c echo.Context) (err error) {
	return
}

func (h *Handler) GetTeam(c echo.Context) (err error) {
	return
}

func (h *Handler) DeleteTeam(c echo.Context) (err error) {
	return
}
