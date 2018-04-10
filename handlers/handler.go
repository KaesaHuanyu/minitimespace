package handlers

import (
	"fmt"

	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
)

type (
	//Handler is a struct to manage handlers
	Handler struct {
		miniProgramAccessToken string
		updateAccessTokenChan  chan bool
	}

	response struct {
		Code    int
		Data    map[string]interface{} `json:"data,omitempty"`
		Error   string                 `json:"error,omitempty"`
		Warning string                 `json:"warning,omitempty"`
		Info    string                 `json:"info,omitempty"`
	}
)

func responses() *response {
	res := new(response)
	res.Data = make(map[string]interface{})
	return res
}

//New create a handler instance
func New() (handler *Handler) {
	handler = new(Handler)
	handler.updateAccessTokenChan = make(chan bool)
	go func() {
		// handler.updateAccessTokenChan <- true
	}()
	return
}

func (h *Handler) info(tag, format string, args ...interface{}) {
	logrus.Infoln(fmt.Sprintf("%s: ", tag) + fmt.Sprintf(format, args...))
}

func (h *Handler) warning(tag, format string, args ...interface{}) {
	logrus.Warningln(fmt.Sprintf("%s: ", tag) + fmt.Sprintf(format, args...))
}

func (h *Handler) danger(tag, format string, args ...interface{}) {
	logrus.Errorln(fmt.Sprintf("%s: ", tag) + fmt.Sprintf(format, args...))
}

func (h *Handler) updateAccessToken() (err error) {
	req, err := grequests.Get(fmt.Sprintf(WeChatGetAccessToken, MiniProgramAppID, MiniProgramAppSecret), nil)
	if err != nil {
		return
	}
	data := new(getAccessTokenResponse)
	if err = req.JSON(data); err != nil {
		return
	}
	if data.Errcode != 0 {
		err = fmt.Errorf("%+v", data)
		return
	}
	h.miniProgramAccessToken = data.AccessToken
	return
}

func (h *Handler) AccessTokenUpdater() {
	for {
		select {
		case <-h.updateAccessTokenChan:
			if err := h.updateAccessToken(); err != nil {
				h.danger("AccessTokenUpdater", "更新access_token错误: [%+v]", err)
				h.updateAccessTokenChan <- true
				continue
			}
			h.info("AccessTokenUpdater", "更新access_token成功，access_token: [%s]", h.miniProgramAccessToken)
		}
	}
}
