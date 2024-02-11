package model

type ClientInput struct {
	JwtToken   string `json:"jwtToken"`
	CourseName string `json:"courseName"`
}
