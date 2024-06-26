package model

type User struct {
	Id        int
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

type Admin struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
}

type Expert struct {
	Id          int
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Email       string  `json:"email"`
	Cost        float64 `json:"cost"`
	Password    string  `json:"password"`
	Description string  `json:"description"`
	ImageLink   string  `json:"imageLink"`
}

type SendEmailCodeRequest struct {
	Email string `json:"email"`
}

type CheckEmailCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
