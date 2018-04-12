package handlers

type (
	getAccessTokenResponse struct {
		AccessToken string `json:"access_token,omitempty"`
		ExpiresIn   int    `json:"expires_in,omitempty"`
		Errcode     int    `json:"errcode,omitempty"`
		Errmsg      string `json:"errmsg,omitempty"`
	}

	loginCredentialsCheckResponse struct {
		Openid     string `json:"openid,omitempty"`
		SessionKey string `json:"session_key,omitempty"`
		Errcode    int    `json:"errcode,omitempty"`
		Errmsg     string `json:"errmsg,omitempty"`
	}

	getTimespaceResponse struct {
		Timespace []timespaceDesc `json:"timespace,omitempty"`
	}

	timespaceDesc struct {
		ID        uint        `json:"id,omitempty"`
		Topic     string      `json:"topic,omitempty"`
		Desc      string      `json:"desc,omitempty"`
		StartTime string      `json:"start_time,omitempty"`
		EndTime   string      `json:"end_time,omitempty"`
		Labels    []labelDesc `json:"label_names,omitempty"`
		Users     []userDesc  `json:"user_avatars,omitempty"`
	}

	labelDesc struct {
		ID   uint   `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}

	userDesc struct {
		ID     uint   `json:"id,omitempty"`
		Name   string `json:"name,omitempty"`
		Avatar string `json:"avatar,omitempty"`
	}
)
