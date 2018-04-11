package handlers

import (
	"fmt"
	"minitimespace/models"
	"net/http"
	"time"

	"github.com/levigross/grequests"

	"github.com/labstack/echo"
)

//Login 接收小程序传来的code，获取openid与session_key
func (h *Handler) Login(c echo.Context) (err error) {
	r := responses()
	code := c.QueryParam("code")
	if code == "" {
		err = fmt.Errorf("code参数为空")
		r.Code = RequestErr
		r.Error = err.Error()
		h.danger("Login", "c.QueryParam, err=[%+v]", err)
		return c.JSON(http.StatusBadRequest, r)
	}
	resp, err := grequests.Get(fmt.Sprintf(WeChatLoginCredentialsCheck, code), nil)
	if err != nil {
		r.Code = HTTPGetErr
		r.Error = err.Error()
		h.danger("Login", "grequests.Get, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}
	wxresp := new(loginCredentialsCheckResponse)
	err = resp.JSON(wxresp)
	if err != nil {
		r.Code = JSONErr
		r.Error = err.Error()
		h.danger("Login", "resp.JSON, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}
	if wxresp.Errcode != 0 {
		err = fmt.Errorf("%+v", wxresp)
		r.Code = WxApiErr
		r.Error = err.Error()
		h.danger("Login", "wxresp.Errcode != 0, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	//对openid和session_key哈希加密
	// hash := sha1.New()
	// hash.Write([]byte(wxresp.Openid + wxresp.SessionKey))
	// session := string(hash.Sum(nil))

	//将session存入redis
	s := &models.Session{
		OpenId:     wxresp.Openid,
		SessionKey: wxresp.SessionKey,
	}
	err = s.Set()
	if err != nil {
		r.Code = SessionErr
		r.Error = err.Error()
		h.danger("Login", "s.Set, err=[%+v]", err)
		return c.JSON(http.StatusInternalServerError, r)
	}

	r.Info = "LOGIN SUCCESS"
	r.Data["session_key"] = s.SessionKey
	r.Data["expiration"] = time.Now().Add(30 * 24 * time.Hour).Unix()
	return c.JSON(http.StatusOK, r)
}

//Protect is used to protect data
func (h *Handler) Protect(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		//check if user is logged in
		//TODO
		r := responses()
		sessionKey := c.QueryParam("session_key")
		if sessionKey == "" {
			err = fmt.Errorf("session_key参数为空")
			r.Code = RequestErr
			r.Error = err.Error()
			h.danger("Protect", "c.QueryParam, err=[%+v]", err)
			return c.JSON(http.StatusBadRequest, r)
		}
		openid, err := models.GetOpenIdBySession(sessionKey)
		if err != nil {
			r.Code = SessionErr
			r.Error = err.Error()
			h.danger("Protect", "models.GetSession, err=[%+v]", err)
			return c.JSON(http.StatusInternalServerError, r)
		}
		c.Set("openid", openid)
		return handlerFunc(c)
	}
}
