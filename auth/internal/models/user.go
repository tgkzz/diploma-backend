package models

type User struct {
	Id        int
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}
