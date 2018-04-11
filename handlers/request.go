package handlers

type (
	requestGetUser struct {
		NickName  string `json:"nickName,omitempty"`
		AvatarURL string `json:"avatarUrl,omitempty"`
		Gender    int    `json:"gender,omitempty"`
		City      string `json:"city,omitempty"`
		Province  string `json:"province,omitempty"`
		Country   string `json:"country,omitempty"`
		Language  string `json:"language,omitempty"`
	}
)
