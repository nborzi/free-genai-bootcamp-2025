package service

import (
	"lang-portal/internal/models"
	"time"
)

type StudyActivity struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnail_url"`
	Description  string `json:"description"`
}

type StudySession struct {
	ID              int    `json:"id"`
	ActivityName    string `json:"activity_name"`
	GroupName       string `json:"group_name"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	ReviewItemCount int    `json:"review_items_count"`
}

func GetStudyActivities() ([]models.StudyActivity, error) {
	var activities []models.StudyActivity
	rows, err := db.Query("SELECT id, study_session_id, group_id, created_at FROM study_activities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var activity models.StudyActivity
		if err := rows.Scan(&activity.ID, &activity.StudySessionID, &activity.GroupID, &activity.CreatedAt); err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}

func GetStudyActivity(id int) (*models.StudyActivity, error) {
	var activity models.StudyActivity
	err := db.QueryRow("SELECT id, study_session_id, group_id, created_at FROM study_activities WHERE id = ?", id).
		Scan(&activity.ID, &activity.StudySessionID, &activity.GroupID, &activity.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func StartStudyActivity(id int) (*models.StudySession, error) {
	// Create a new study session for the activity
	result, err := db.Exec("INSERT INTO study_sessions (group_id, study_activity_id, created_at) VALUES (?, ?, ?)",
		1, id, time.Now())
	if err != nil {
		return nil, err
	}

	sessionID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var session models.StudySession
	err = db.QueryRow("SELECT id, group_id, study_activity_id, created_at FROM study_sessions WHERE id = ?", sessionID).
		Scan(&session.ID, &session.GroupID, &session.StudyActivityID, &session.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func GetStudySessions() ([]models.StudySession, error) {
	var sessions []models.StudySession
	rows, err := db.Query("SELECT id, group_id, study_activity_id, created_at FROM study_sessions ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var session models.StudySession
		if err := rows.Scan(&session.ID, &session.GroupID, &session.StudyActivityID, &session.CreatedAt); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func GetStudySession(id int) (*models.StudySession, error) {
	var session models.StudySession
	err := db.QueryRow("SELECT id, group_id, study_activity_id, created_at FROM study_sessions WHERE id = ?", id).
		Scan(&session.ID, &session.GroupID, &session.StudyActivityID, &session.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func GetStudyActivitySessions(activityID int, page int) ([]StudySession, int, error) {
	offset := (page - 1) * 100
	query := `
		SELECT 
			s.id,
			'Vocabulary Quiz' as activity_name,
			g.name as group_name,
			s.created_at as start_time,
			datetime(s.created_at, '+10 minutes') as end_time,
			COUNT(wri.word_id) as review_items_count
		FROM study_sessions s
		JOIN groups g ON s.group_id = g.id
		LEFT JOIN word_review_items wri ON s.id = wri.study_session_id
		WHERE s.study_activity_id = ?
		GROUP BY s.id
		ORDER BY s.created_at DESC
		LIMIT 100 OFFSET ?
	`

	rows, err := db.Query(query, activityID, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var sessions []StudySession
	for rows.Next() {
		var s StudySession
		err := rows.Scan(&s.ID, &s.ActivityName, &s.GroupName, &s.StartTime, &s.EndTime, &s.ReviewItemCount)
		if err != nil {
			return nil, 0, err
		}
		sessions = append(sessions, s)
	}

	// Get total count
	var total int
	err = db.QueryRow("SELECT COUNT(*) FROM study_sessions WHERE study_activity_id = ?", activityID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return sessions, total, nil
}

func CreateStudySession(groupID, activityID int) (*models.StudySession, error) {
	result, err := db.Exec(`
		INSERT INTO study_sessions (group_id, study_activity_id, created_at)
		VALUES (?, ?, CURRENT_TIMESTAMP)
	`, groupID, activityID)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.StudySession{
		ID:              int(id),
		GroupID:         groupID,
		StudyActivityID: activityID,
	}, nil
}
