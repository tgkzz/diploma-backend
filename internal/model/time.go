package model

import (
	"encoding/json"
	"time"
)

type UnixTime struct {
	time.Time
}

func (t *UnixTime) UnmarshalJSON(b []byte) error {
	var unixTime int64
	if err := json.Unmarshal(b, &unixTime); err != nil {
		return err
	}
	t.Time = time.Unix(unixTime, 0)
	return nil
}

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

type QuoteMsg struct {
	Id   string `json:"id"`
	Body string `json:"body"`
}
