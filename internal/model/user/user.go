package user

type User struct {
	Id        int
	Login     string `json:"login"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}
