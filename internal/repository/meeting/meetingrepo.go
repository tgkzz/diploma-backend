package meeting

import (
	"database/sql"
	"server/internal/model"
)

type MeetingRepo struct {
	pgDb *sql.DB
}

type IMeetingRepo interface {
	CreateMeeting(meeting model.Meeting) error
	GetMeetingByRoomId(roomId string) (model.Meeting, error)
}

func NewMeetingRepo(pg *sql.DB) *MeetingRepo {
	return &MeetingRepo{pgDb: pg}
}

// TODO: fix me, fix tables
func (m *MeetingRepo) CreateMeeting(meeting model.Meeting) error {

	query := `INSERT INTO meeting_transactions(user_id, expert_id, time_start, time_end, total_cost, meeting_link, meeting_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	if _, err := m.pgDb.Exec(query,
		meeting.UserId,
		meeting.ExpertId,
		meeting.TimeStart.Time,
		meeting.TimeEnd.Time,
		meeting.TotalCost,
		meeting.MeetingLink,
		meeting.RoomId,
	); err != nil {
		return err
	}

	return nil
}

func (m *MeetingRepo) GetMeetingByRoomId(roomId string) (model.Meeting, error) {
	query := `SELECT id, user_id, expert_id, time_start, time_end, total_cost, meeting_link, meeting_id FROM meeting_transactions WHERE meeting_id=$1`

	var res model.Meeting
	if err := m.pgDb.QueryRow(query, roomId).Scan(&res.Id, &res.UserId, &res.ExpertId,
		&res.TimeStart.Time, &res.TimeEnd.Time, &res.TotalCost, &res.MeetingLink, &res.RoomId); err != nil {
		return model.Meeting{}, err
	}

	return res, nil
}
