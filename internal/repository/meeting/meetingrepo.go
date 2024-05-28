package meeting

import (
	"database/sql"
	"errors"
	"server/internal/model"
)

type MeetingRepo struct {
	pgDb *sql.DB
}

type IMeetingRepo interface {
	CreateMeeting(meeting model.Meeting) error
}

func NewMeetingRepo(pg *sql.DB) *MeetingRepo {
	return &MeetingRepo{pgDb: pg}
}

// TODO: fix me, fix tables
func (m *MeetingRepo) CreateMeeting(meeting model.Meeting) error {
	query := `INSERT INTO meeting_transactions(user_id, expert_id, time_start, time_end, total_cost, meeting_link) VALUES ($1, $2, $3, $4, $5, $6)`

	if _, err := m.pgDb.Exec(query,
		meeting.UserId,
		meeting.ExpertId,
		meeting.TimeStart,
		meeting.TimeEnd,
		meeting.TotalCost,
		meeting.MeetingLink,
	); err != nil {
		return err
	}

	return nil
}

func (m *MeetingRepo) GetMeetingByRoomId(roomId string) (model.Meeting, error) {
	query := `SELECT id, user_id, expert_id, time_start, time_end, total_cost, meeting_link FROM meeting_transactions WHERE `

	_ = query

	return model.Meeting{}, errors.New("method not implemented")
}
