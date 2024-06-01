package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/internal/pkg"
)

type UpdateUserRequest struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
}

func (u UpdateUserRequest) EmptyUpdate() UpdateUserRequest {
	return UpdateUserRequest{}
}

func (u UpdateUserRequest) SetId(id int) UpdateUserRequest {
	u.Id = id
	return u
}

func (u UpdateUserRequest) HasId() bool {
	return u.Id != 0
}

func (u UpdateUserRequest) SetEmail(email string) UpdateUserRequest {
	u.Email = email
	return u
}

func (u UpdateUserRequest) HasEmail() bool {
	return u.Email != ""
}

func (u UpdateUserRequest) SetPassword(psw string) UpdateUserRequest {
	u.Password, _ = pkg.HashPassword(psw)
	return u
}

func (u UpdateUserRequest) HasPassword() bool {
	return u.Password != ""
}

func (u UpdateUserRequest) SetFname(str string) UpdateUserRequest {
	u.Fname = str
	return u
}

func (u UpdateUserRequest) HasFname() bool {
	return u.Fname != ""
}

func (u UpdateUserRequest) SetLname(str string) UpdateUserRequest {
	u.Lname = str
	return u
}

func (u UpdateUserRequest) HasLname() bool {
	return u.Lname != ""
}

type GetCourseLimitedResponse struct {
	Id               primitive.ObjectID `json:"id"`
	Name             string             `json:"name"`
	Cost             float64            `json:"cost"`
	ModuleCount      string             `json:"module_count"`
	ShortDescription string             `json:"short_description"`
	ImageUrl         string             `json:"image_url"`
}

type MakeAppointmentRequest struct {
	UserEmail   string `json:"userEmail"`
	ExpertEmail string `json:"expertEmail"`
	TimeStart   string `json:"timeStart"`
	TimeEnd     string `json:"timeEnd"`
	RoomId      string `json:"roomId"`
}
