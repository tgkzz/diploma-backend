package meeting

import (
	"database/sql"
	"fmt"
	"server/internal/model"
	"strings"
)

type MeetingRepo struct {
	pgDb *sql.DB
}

type IMeetingRepo interface {
	CreateMeeting(meeting model.Meeting) error
	GetMeetingByRoomId(roomId string) (model.Meeting, error)
	UpdateMeeting(meeting model.Meeting, meetingId int) error
	GetMeetingsByUserId(userId int) ([]model.Meeting, error)
	GetMeetingsByExpertId(expertId int) ([]model.Meeting, error)
	GetExpertAvailableMeets(expertId int) ([]model.Meeting, error)
}

func NewMeetingRepo(pg *sql.DB) *MeetingRepo {
	return &MeetingRepo{pgDb: pg}
}

func (m *MeetingRepo) CreateMeeting(meeting model.Meeting) error {

	query := `INSERT INTO meeting_transactions(user_id, expert_id, time_start, time_end, total_cost, meeting_link, meeting_id, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	if _, err := m.pgDb.Exec(query,
		meeting.UserId,
		meeting.ExpertId,
		meeting.TimeStart.Time,
		meeting.TimeEnd.Time,
		meeting.TotalCost,
		meeting.MeetingLink,
		meeting.RoomId,
		meeting.Status,
	); err != nil {
		return err
	}

	return nil
}

func (m *MeetingRepo) GetMeetingByRoomId(roomId string) (model.Meeting, error) {
	query := `SELECT id, user_id, expert_id, time_start, time_end, total_cost, meeting_link, meeting_id, status FROM meeting_transactions WHERE meeting_id=$1`

	var res model.Meeting
	if err := m.pgDb.QueryRow(query, roomId).Scan(&res.Id, &res.UserId, &res.ExpertId,
		&res.TimeStart.Time, &res.TimeEnd.Time, &res.TotalCost, &res.MeetingLink, &res.RoomId, &res.Status); err != nil {
		return model.Meeting{}, err
	}

	return res, nil
}

func (m *MeetingRepo) UpdateMeeting(meeting model.Meeting, meetingId int) error {
	setParts := []string{}
	args := []interface{}{}
	argId := 1

	if meeting.UserId != 0 {
		setParts = append(setParts, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, meeting.UserId)
		argId++
	}

	if meeting.ExpertId != 0 {
		setParts = append(setParts, fmt.Sprintf("expert_id=$%d", argId))
		args = append(args, meeting.ExpertId)
		argId++
	}

	if !meeting.TimeStart.Time.IsZero() {
		setParts = append(setParts, fmt.Sprintf("time_start=$%d", argId))
		args = append(args, meeting.TimeStart.Time)
		argId++
	}

	if !meeting.TimeEnd.Time.IsZero() {
		setParts = append(setParts, fmt.Sprintf("time_end=$%d", argId))
		args = append(args, meeting.TimeEnd.Time)
		argId++
	}

	if meeting.TotalCost != 0 {
		setParts = append(setParts, fmt.Sprintf("total_cost=$%d", argId))
		args = append(args, meeting.TotalCost)
		argId++
	}

	if meeting.MeetingLink != "" {
		setParts = append(setParts, fmt.Sprintf("meeting_link=$%d", argId))
		args = append(args, meeting.MeetingLink)
		argId++
	}

	if meeting.Status != "" {
		setParts = append(setParts, fmt.Sprintf("status=$%d", argId))
		args = append(args, meeting.Status)
		argId++
	}

	if len(setParts) == 0 {
		return fmt.Errorf("no updates specified")
	}

	query := fmt.Sprintf("UPDATE meeting_transactions SET %s WHERE id=$%d", strings.Join(setParts, ", "), argId)
	args = append(args, meeting.Id)

	if _, err := m.pgDb.Exec(query, args...); err != nil {
		return err
	}

	return nil
}

func (m *MeetingRepo) GetMeetingsByExpertId(expertId int) ([]model.Meeting, error) {
	query := `SELECT id, user_id, expert_id, time_start, time_end, total_cost, meeting_link, meeting_id, status FROM meeting_transactions WHERE expert_id=$1`

	rows, err := m.pgDb.Query(query, expertId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meetings []model.Meeting
	for rows.Next() {
		var meeting model.Meeting
		err := rows.Scan(&meeting.Id, &meeting.UserId, &meeting.ExpertId, &meeting.TimeStart.Time, &meeting.TimeEnd.Time, &meeting.TotalCost, &meeting.MeetingLink, &meeting.RoomId, &meeting.Status)
		if err != nil {
			return nil, err
		}
		meetings = append(meetings, meeting)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return meetings, nil
}

func (m *MeetingRepo) GetMeetingsByUserId(userId int) ([]model.Meeting, error) {
	query := `SELECT id, user_id, expert_id, time_start, time_end, total_cost, meeting_link, meeting_id, status FROM meeting_transactions WHERE user_id=$1`

	rows, err := m.pgDb.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meetings []model.Meeting
	for rows.Next() {
		var meeting model.Meeting
		err := rows.Scan(&meeting.Id, &meeting.UserId, &meeting.ExpertId, &meeting.TimeStart.Time, &meeting.TimeEnd.Time, &meeting.TotalCost, &meeting.MeetingLink, &meeting.RoomId, &meeting.Status)
		if err != nil {
			return nil, err
		}
		meetings = append(meetings, meeting)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return meetings, nil
}

func (m *MeetingRepo) GetExpertAvailableMeets(expertId int) ([]model.Meeting, error) {
	query := `SELECT id, user_id, expert_id, time_start, time_end, total_cost, meeting_link, meeting_id, status FROM meeting_transactions WHERE expert_id=$1 AND status="available"`

	rows, err := m.pgDb.Query(query, expertId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var meetings []model.Meeting
	for rows.Next() {
		var meeting model.Meeting
		err := rows.Scan(&meeting.Id, &meeting.UserId, &meeting.ExpertId, &meeting.TimeStart.Time, &meeting.TimeEnd.Time, &meeting.TotalCost, &meeting.MeetingLink, &meeting.RoomId, &meeting.Status)
		if err != nil {
			return nil, err
		}
		meetings = append(meetings, meeting)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return meetings, nil
}
