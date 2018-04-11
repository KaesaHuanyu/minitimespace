package models

import (
	"time"
)

type Session struct {
	OpenId     string `json:"open_id,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
}

func (s *Session) Set() (err error) {
	err = client.Set(s.SessionKey, s.OpenId, 30*24*time.Hour).Err()
	return
}

func GetOpenIdBySession(key string) (openid string, err error) {
	openid, err = client.Get(key).Result()
	return
}
