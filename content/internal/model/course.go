package model

type Course struct {
	Id          int
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	Cost        float64 `json:"cost"`
}
