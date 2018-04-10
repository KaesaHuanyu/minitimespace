package models

import (
	"encoding/json"
	"time"
)

type Session struct {
	OpenId     string `json:"open_id,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
}

func (s *Session) Set(key string) (err error) {
	body, err := json.Marshal(s)
	if err != nil {
		return
	}
	err = client.Set(key, string(body), 30*24*time.Hour).Err()
	return
}

func GetSession(key string) (s *Session, err error) {
	s = new(Session)
	ret, err := client.Get(key).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(ret), s)
	return
}
