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

	requestCreateTimespace struct {
		Topic     string `json:"topic,omitempty"`
		Desc      string `json:"desc,omitempty"`
		StartTime string `json:"start_time,omitempty"`
		EndTime   string `json:"end_time,omitempty"`
		Labels    []uint `json:"labels,omitempty"`
	}

	requestUpdateTimespace struct {
		Topic     string `json:"topic,omitempty"`
		Desc      string `json:"desc,omitempty"`
	}
)
