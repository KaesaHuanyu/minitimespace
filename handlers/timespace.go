package handlers

import (
	"minitimespace/models"
	"net/http"
	"strconv"

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
	err = c.Bind(request)
	if err != nil {
		r.Code = JSONErr
		r.Error = err.Error()
		h.danger("CreateTimespace", "c.Bind, err=[%+v]", err)
	}
	labels := make([]models.Label, 0)
	if len(request.Labels) > 0 {
		labels, err = models.GetLabelsByIds(request.Labels)
		if err != nil {
			r.Code = DatabaseErr
			r.Error = err.Error()
			h.danger("CreateTimespace", "models.GetLabelsByIds, err=[%+v]", err)
			return c.JSON(http.StatusInternalServerError, r)
		}
	}
	t := &models.Timespace{
		UserID:    u.ID,
		Topic:     request.Topic,
		Desc:      request.Desc,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		Labels:    labels,
	}
	err = t.Create()
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("CreateTimespace", "t.Create, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	r.Info = "CREATE SUCCESS"
	r.Data["tid"] = t.ID
	h.info("CreateTimespace", "新增小时空, timespace=[%+v]", t)
	return c.JSON(http.StatusOK, r)
}

//更新小时空
func (h *Handler) UpdateTimespace(c echo.Context) (err error) {
	r := responses()
	openid := c.Get("openid").(string)
	u, err := models.GetUserByOpenId(openid)
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("UpdateTimespace", "models.GetUserByOpenId, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
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

//删除小时空
func (h *Handler) DeleteTimespace(c echo.Context) (err error) {
	r := responses()
	openid := c.Get("openid").(string)
	u, err := models.GetUserByOpenId(openid)
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("DeleteTimespace", "models.GetUserByOpenId, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	return c.JSON(http.StatusOK, r)
}

//返回小时空详情
func (h *Handler) GetTimespaceDetail(c echo.Context) (err error) {
	r := responses()
	openid := c.Get("openid").(string)
	u, err := models.GetUserByOpenId(openid)
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("GetTimespaceDetail", "models.GetUserByOpenId, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	return c.JSON(http.StatusOK, r)
}

func (h *Handler) JoinTimespace(c echo.Context) (err error) {
	r := responses()
	openid := c.Get("openid").(string)
	u, err := models.GetUserByOpenId(openid)
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("JoinTimespace", "models.GetUserByOpenId, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	tid, err := strconv.Atoi(c.Param("tid"))
	if err != nil {
		r.Code = RequestErr
		r.Error = err.Error()
		h.danger("JoinTimespace", "strconv.Atoi(c.Param, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}
	t, err := models.GetTimespaceById(uint(tid))
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("JoinTimespace", "用户成功加入小时空，models.GetTimespaceById, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	err = t.AddUser(*u)
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("JoinTimespace", "models.AddUser, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	r.Info = "JOINED SUCCESS"
	h.info("JoinTimespace", "timespace=[%+v], user=[%+v]", t, u)
	return c.JSON(http.StatusOK, r)
}

func (h *Handler) WhetherCurrentUserJoined(c echo.Context) (err error) {
	r := responses()
	openid := c.Get("openid").(string)
	u, err := models.GetUserByOpenId(openid)
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("WhetherCurrentUserJoined", "models.GetUserByOpenId, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	tid, err := strconv.Atoi(c.Param("tid"))
	if err != nil {
		r.Code = RequestErr
		r.Error = err.Error()
		h.danger("WhetherCurrentUserJoined", "strconv.Atoi(c.Param, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}
	t, err := models.GetTimespaceById(uint(tid))
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("WhetherCurrentUserJoined", "models.GetTimespaceById, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	if t.WhetherTheUserJoined(*u) {
		r.Data["status"] = true
		h.info("WhetherCurrentUserJoined", "当前用户已加入该小时空, timespace=[%+v], user=[%+v]", t, u)
	} else {
		r.Data["status"] = false
		h.info("WhetherCurrentUserJoined", "当前用户未加入该小时空, timespace=[%+v], user=[%+v]", t, u)
	}
	return c.JSON(http.StatusOK, r)
}
