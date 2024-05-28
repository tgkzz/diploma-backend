package model

type Meeting struct {
	Id          int
	UserId      int      `json:"userId"`
	ExpertId    int      `json:"expertId"`
	TimeStart   UnixTime `json:"timeStart"`
	TimeEnd     UnixTime `json:"timeEnd"`
	TotalCost   float64  `json:"totalCost"`
	MeetingLink string   `json:"meetingLink"`
	RoomId      string   `json:"roomId"`
}