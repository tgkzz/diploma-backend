package meeting

import (
	"github.com/google/uuid"
	"server/internal/model"
	"server/internal/repository/auth"
	"server/internal/repository/authexpert"
	"server/internal/repository/meeting"
	"time"
)

type MeetingService struct {
	meetingRepo meeting.IMeetingRepo
	userRepo    auth.IAuthRepo
	expertRepo  authexpert.IExpertRepo
}

type IMeetingService interface {
	CreateMeeting(request model.MakeAppointmentRequest, userEmail string) (string, error)
	GetMeetingByRoomId(roomId string) (model.Meeting, error)
}

func NewMeetingService(repo meeting.IMeetingRepo, authRepo auth.IAuthRepo, expertRepo authexpert.IExpertRepo) *MeetingService {
	return &MeetingService{meetingRepo: repo, userRepo: authRepo, expertRepo: expertRepo}
}

func (m *MeetingService) CreateMeeting(request model.MakeAppointmentRequest, userEmail string) (string, error) {
	expert, err := m.expertRepo.GetExpertByEmail(request.ExpertEmail)
	if err != nil {
		return "", err
	}

	user, err := m.userRepo.GetUserByEmail(userEmail)
	if err != nil {
		return "", err
	}

	req := model.Meeting{
		UserId:      user.Id,
		ExpertId:    expert.Id,
		TotalCost:   expert.Cost,
		MeetingLink: "",
		RoomId:      uuid.New().String(),
	}

	req.TimeStart.Time, err = time.Parse(time.RFC3339, request.TimeStart)
	if err != nil {
		return "", err
	}

	req.TimeEnd.Time, err = time.Parse(time.RFC3339, request.TimeEnd)
	if err != nil {
		return "", err
	}

	if err := m.meetingRepo.CreateMeeting(req); err != nil {
		return "", err
	}

	return req.RoomId, nil
}

func (m *MeetingService) GetMeetingByRoomId(roomId string) (model.Meeting, error) {
	return m.meetingRepo.GetMeetingByRoomId(roomId)
}
