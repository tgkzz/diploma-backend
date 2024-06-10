package meeting

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"server/internal/model"
	"server/internal/pkg/sorting"
	"server/internal/repository/auth"
	"server/internal/repository/authexpert"
	"server/internal/repository/meeting"
	"time"
)

const (
	AVAILABLE = "available"
	BOOKED    = "booked"
	CONDUCTED = "conducted"
)

type MeetingService struct {
	meetingRepo meeting.IMeetingRepo
	userRepo    auth.IAuthRepo
	expertRepo  authexpert.IExpertRepo
}

type IMeetingService interface {
	CreateMeeting(request model.MakeAppointmentRequest, userEmail string) (string, error)
	GetMeetingByRoomId(roomId string) (model.Meeting, error)
	CreateMeetByExpert(ctx context.Context, req model.MakeAppointmentRequest) (string, error)
	PlaceAppointment(req model.MakeAppointmentRequest, userEmail string) error
	GetExpertMeets(email string) ([]model.Meeting, error)
	GetUserMeets(email string) ([]model.Meeting, error)
	GetExpertAvailableMeets(expertId int) ([]model.Meeting, error)
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

func (m *MeetingService) CreateMeetByExpert(ctx context.Context, req model.MakeAppointmentRequest) (string, error) {
	expert, err := m.expertRepo.GetExpertByEmail(req.ExpertEmail)
	if err != nil {
		return "", err
	}

	var timeStart model.UnixTime
	err = timeStart.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", req.TimeStart)))
	if err != nil {
		return "", err
	}

	var timeEnd model.UnixTime
	err = timeEnd.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", req.TimeEnd)))
	if err != nil {
		return "", err
	}

	if timeStart.Time.Before(time.Now()) || timeEnd.Time.Before(time.Now()) || timeEnd.Time.Before(timeStart.Time) {
		return "", model.ErrTimeInPast
	}

	roomId := uuid.New().String()
	if err := m.meetingRepo.CreateMeeting(model.Meeting{
		ExpertId:    expert.Id,
		TotalCost:   expert.Cost,
		MeetingLink: "",
		RoomId:      roomId,
		TimeStart:   timeStart,
		TimeEnd:     timeEnd,
		Status:      AVAILABLE,
	}); err != nil {
		return "", err
	}

	return roomId, nil
}

func (m *MeetingService) PlaceAppointment(req model.MakeAppointmentRequest, userEmail string) error {
	user, err := m.userRepo.GetUserByEmail(userEmail)
	if err != nil {
		return err
	}

	meet, err := m.meetingRepo.GetMeetingByRoomId(req.RoomId)
	if err != nil {
		return err
	}

	if meet.Status != AVAILABLE {
		return model.ErrMeetingAlreadyBooked
	}

	if meet.TimeStart.Time.Before(time.Now()) {
		return model.ErrTimeInPast
	}

	if err := m.meetingRepo.UpdateMeeting(model.Meeting{
		Id:     meet.Id,
		UserId: user.Id,
		Status: BOOKED,
	}, meet.Id); err != nil {
		return err
	}

	return nil
}

func (m *MeetingService) GetExpertMeets(email string) ([]model.Meeting, error) {
	expert, err := m.expertRepo.GetExpertByEmail(email)
	if err != nil {
		return nil, err
	}

	res, err := m.meetingRepo.GetMeetingsByExpertId(expert.Id)
	if err != nil {
		return nil, err
	}

	sorting.SortByTime(res)

	return res, nil
}

func (m *MeetingService) GetUserMeets(email string) ([]model.Meeting, error) {
	user, err := m.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	res, err := m.meetingRepo.GetMeetingsByUserId(user.Id)
	if err != nil {
		return nil, err
	}

	sorting.SortByTime(res)

	return res, nil
}

func (m *MeetingService) GetExpertAvailableMeets(expertId int) ([]model.Meeting, error) {
	res, err := m.meetingRepo.GetExpertAvailableMeets(expertId)
	if err != nil {
		return nil, err
	}

	sorting.SortByTime(res)

	return res, nil
}
