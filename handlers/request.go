package handlers

type (
	requestGetUser struct {
		NickName  string `json:"nick_name,omitempty"`
		AvatarURL string `json:"avatar_url,omitempty"`
		Gender    string `json:"gender,omitempty"`
		City      string `json:"city,omitempty"`
		Province  string `json:"province,omitempty"`
		Country   string `json:"country,omitempty"`
		Language  string `json:"language,omitempty"`
	}
)
