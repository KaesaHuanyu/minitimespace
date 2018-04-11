package handlers

import (
	"minitimespace/models"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetHomeTimespace(c echo.Context) (err error) {
	r := responses()
	openid := c.Get("openid").(string)
	u, err := models.GetUserByOpenId(openid)
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("GetHomeTimespace", "models.GetUserByOpenId, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}
	timespace, err := u.GetAddedTimespace()
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("GetHomeTimespace", "u.GetAddedTimespace, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}
	timespaceDescs := make([]timespaceDesc, len(timespace))
	for i := range timespace {
		labels, err := timespace[i].GetLabels()
		if err != nil {
			r.Code = DatabaseErr
			r.Error = err.Error()
			h.danger("GetHomeTimespace", "timespace[i].GetLabels, err=[%+v]", err)
			return c.JSON(http.StatusInternalServerError, r)
		}
		labelNames := make([]string, len(labels))
		for j := range labels {
			labelNames[j] = labels[j].Name
		}
		users, err := timespace[i].GetUsers()
		if err != nil {
			r.Code = DatabaseErr
			r.Error = err.Error()
			h.danger("GetHomeTimespace", "timespace[i].GetUsers, err=[%+v]", err)
			return c.JSON(http.StatusInternalServerError, r)
		}
		userAvatars := make([]string, len(users))
		for k := range users {
			userAvatars[k] = users[k].Avatar
		}
		timespaceDescs[i] = timespaceDesc{
			Topic:       timespace[i].Topic,
			Desc:        timespace[i].Desc,
			StartTime:   timespace[i].StartTime,
			EndTime:     timespace[i].EndTime,
			LabelNames:  labelNames,
			UserAvatars: userAvatars,
		}
	}

	r.Info = "GET SUCCESS"
	r.Data["timespace"] = timespaceDescs
	h.info("GetHomeTimespace", "Get timespace=[%+v]", timespaceDescs)
	return c.JSON(http.StatusOK, r)
}
