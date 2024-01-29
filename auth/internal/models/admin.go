package models

type Admin struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
}
