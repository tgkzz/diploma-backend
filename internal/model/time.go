package model

import (
	"encoding/json"
	"time"
)

type UnixTime struct {
	time.Time
}

func (t *UnixTime) UnmarshalJSON(b []byte) error {
	var isoTime string
	if err := json.Unmarshal(b, &isoTime); err != nil {
		return err
	}
	parsedTime, err := time.Parse(time.RFC3339, isoTime)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format(time.RFC3339))
}

type QuoteMsg struct {
	Id   string `json:"id"`
	Body string `json:"body"`
}
