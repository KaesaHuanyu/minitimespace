package handlers

import (
	"minitimespace/models"
	"net/http"

	"github.com/labstack/echo"
)

//创建小时空
func (h *Handler) CreateTimespace(c echo.Context) (err error) {
	r := responses()
	//获取当前用户
	openid := c.Get("openid").(string)
	u, err := models.GetUserByOpenId(openid)
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("CreateTimespace", "models.GetUserByOpenId, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	request := new(requestCreateTimespace)
	err - c.Bind(request)
	if err != nil {
		r.Code = JSONErr
		r.Error = err.Error()
		h.danger("CreateTimespace", "c.Bind, err=[%+v]", err)
	}
	
	return c.JSON(http.StatusOK, r)
}

//返回当前用户的小时空列表
func (h *Handler) GetTimespace(c echo.Context) (err error) {
	r := responses()
	//获取当前用户
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
		labels, err := timespace[i].GetLabels(u.ID)
		if err != nil {
			r.Code = DatabaseErr
			r.Error = err.Error()
			h.danger("GetTimespace", "timespace[i].GetLabels, err=[%+v]", err)
			return c.JSON(http.StatusInternalServerError, r)
		}
		labelDescs := make([]labelDesc, len(labels))
		for j := range labels {
			labelDescs[j].ID = labels[j].ID
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
			userDescs[k].ID = users[k].ID
			userDescs[k].Name = users[k].Name
			userDescs[k].Avatar = users[k].Avatar
		}
		timespaceDescs[i] = timespaceDesc{
			ID:        timespace[i].ID,
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

//返回小时空详情
func (h *Handler) GetTimespaceDetail(c echo.Context) (err error) {
	r := responses()
	return c.JSON(http.StatusOK, r)
}
