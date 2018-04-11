package handlers

import (
	"minitimespace/models"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

func (h *Handler) CreateUser(c echo.Context) (err error) {
	r := responses()
	request := new(requestGetUser)
	err = c.Bind(request)
	if err != nil {
		r.Code = BindErr
		r.Error = err.Error()
		h.danger("CreateUser", "c.Bind, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}
	openid := c.Get("openid").(string)
	u, err := models.GetUserByOpenId(openid)
	if err == nil {
		if err == gorm.ErrRecordNotFound {
			//新建用户
			u = &models.User{
				OpenId:   openid,
				Name:     request.NickName,
				Avatar:   request.AvatarURL,
				Gender:   request.Gender,
				Country:  request.Country,
				Province: request.Province,
				City:     request.City,
				Language: request.Language,
			}
			err = u.Create()
			if err != nil {
				r.Code = DatabaseErr
				r.Error = err.Error()
				h.danger("CreateUser", "u.Create, err=[%+v]", err)
				return c.JSON(http.StatusInternalServerError, r)
			}
			r.Info = "CREATED"
			h.info("CreateUser", "新增用户, user=[%+v]", u)
			return c.JSON(http.StatusCreated, r)
		}
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("CreateUser", "models.GetUserByOpenId, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}
	//更新user
	err = u.Update(map[string]interface{}{
		"name":     request.NickName,
		"avatar":   request.AvatarURL,
		"gender":   request.Gender,
		"country":  request.Country,
		"province": request.Province,
		"city":     request.City,
		"language": request.Language,
	})
	if err != nil {
		r.Code = DatabaseErr
		r.Error = err.Error()
		h.danger("CreateUser", "u.Update, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	r.Info = "CREATED"
	h.info("CreateUser", "更新用户基本信息, user=[%+v]", u)
	return c.JSON(http.StatusCreated, r)
}
