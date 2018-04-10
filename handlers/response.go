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
)
