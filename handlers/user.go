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
	if err != nil {
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
	updates := make(map[string]interface{})
	if u.Name != request.NickName {
		updates["name"] = request.NickName
	}
	if u.Avatar != request.AvatarURL {
		updates["avatar"] = request.AvatarURL
	}
	if u.Gender != request.Gender {
		updates["gender"] = request.Gender
	}
	if u.Country != request.Country {
		updates["country"] = request.Country
	}
	if u.Province != request.Province {
		updates["province"] = request.Province
	}
	if u.City != request.City {
		updates["city"] = request.City
	}
	if u.Language != request.Language {
		updates["language"] = request.Language
	}
	if len(updates) > 0 {
		err = u.Update(updates)
		if err != nil {
			r.Code = DatabaseErr
			r.Error = err.Error()
			h.danger("CreateUser", "u.Update, err=[%+v]", err)
			return c.JSON(http.StatusInternalServerError, r)
		}
	}

	r.Info = "CREATED"
	h.info("CreateUser", "更新用户基本信息, updates=[%+v]", updates)
	return c.JSON(http.StatusCreated, r)
}
