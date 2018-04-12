package handlers

import (
	"minitimespace/models"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetTimespace(c echo.Context) (err error) {
	r := responses()
	openid := c.Get("openid").(string)
	u, err := models.GetUserByOpenId(openid)
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("GetTimespace", "models.GetUserByOpenId, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}
	timespace, err := u.GetAddedTimespace()
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("GetTimespace", "u.GetAddedTimespace, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}
	timespaceDescs := make([]timespaceDesc, len(timespace))
	for i := range timespace {
		labels, err := timespace[i].GetLabels()
		if err != nil {
			r.Code = DatabaseErr
			r.Error = err.Error()
			h.danger("GetTimespace", "timespace[i].GetLabels, err=[%+v]", err)
			return c.JSON(http.StatusInternalServerError, r)
		}
		labelDescs := make([]labelDesc, len(labels))
		for j := range labels {
			labelDescs[j].Id = labels[j].ID
			labelDescs[j].Name = labels[j].Name
		}
		users, err := timespace[i].GetUsers()
		if err != nil {
			r.Code = DatabaseErr
			r.Error = err.Error()
			h.danger("GetTimespace", "timespace[i].GetUsers, err=[%+v]", err)
			return c.JSON(http.StatusInternalServerError, r)
		}
		userDescs := make([]userDesc, len(users))
		for k := range users {
			userDescs[k].Id = users[k].ID
			userDescs[k].Name = users[k].Name
			userDescs[k].Avatar = users[k].Avatar
		}
		timespaceDescs[i] = timespaceDesc{
			Id:        timespace[i].ID,
			Topic:     timespace[i].Topic,
			Desc:      timespace[i].Desc,
			StartTime: timespace[i].StartTime,
			EndTime:   timespace[i].EndTime,
			Labels:    labelDescs,
			Users:     userDescs,
		}
	}

	r.Info = "GET SUCCESS"
	r.Data["timespace"] = timespaceDescs
	h.info("GetTimespace", "Get timespace=[%+v]", timespaceDescs)
	return c.JSON(http.StatusOK, r)
}
