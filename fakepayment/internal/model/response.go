package model

type CourseResponse struct {
	Course struct {
		Id          int     `json:"Id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Link        string  `json:"link"`
		Cost        float64 `json:"cost"`
	} `json:"course"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type UserResponse struct {
	Email struct {
		Id       int    `json:"Id"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Fname    string `json:"fname"`
		Lname    string `json:"lname"`
	} `json:"email"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type EmailRequest struct {
	Email string `json:"email"`
}
