package service

import (
	"database/sql"
	"lang-portal/internal/models"
)

func GetLastStudySession() (*models.StudySession, error) {
	query := `
		SELECT s.id, s.group_id, s.created_at, s.study_activity_id, g.name
		FROM study_sessions s
		JOIN groups g ON s.group_id = g.id
		ORDER BY s.created_at DESC
		LIMIT 1
	`
	var session models.StudySession
	var groupName string
	err := db.QueryRow(query).Scan(
		&session.ID,
		&session.GroupID,
		&session.CreatedAt,
		&session.StudyActivityID,
		&groupName,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func GetStudyProgress() (map[string]int, error) {
	var totalWords, studiedWords int
	
	// Get total words
	err := db.QueryRow("SELECT COUNT(*) FROM words").Scan(&totalWords)
	if err != nil {
		return nil, err
	}

	// Get studied words
	err = db.QueryRow(`
		SELECT COUNT(DISTINCT word_id) 
		FROM word_review_items
	`).Scan(&studiedWords)
	if err != nil {
		return nil, err
	}

	return map[string]int{
		"total_words_studied":    studiedWords,
		"total_available_words": totalWords,
	}, nil
}

func GetQuickStats() (map[string]interface{}, error) {
	var stats struct {
		SuccessRate        float64
		TotalSessions     int
		ActiveGroups      int
		StudyStreakDays   int
	}

	// Get success rate
	err := db.QueryRow(`
		SELECT COALESCE(
			(SELECT CAST(SUM(CASE WHEN correct THEN 1 ELSE 0 END) AS FLOAT) / COUNT(*) * 100
			FROM word_review_items), 0)
	`).Scan(&stats.SuccessRate)
	if err != nil {
		return nil, err
	}

	// Get total study sessions
	err = db.QueryRow("SELECT COUNT(*) FROM study_sessions").Scan(&stats.TotalSessions)
	if err != nil {
		return nil, err
	}

	// Get active groups (groups with at least one study session)
	err = db.QueryRow(`
		SELECT COUNT(DISTINCT group_id) 
		FROM study_sessions
	`).Scan(&stats.ActiveGroups)
	if err != nil {
		return nil, err
	}

	// Calculate study streak
	err = db.QueryRow(`
		WITH RECURSIVE dates(date) AS (
			SELECT date(MAX(created_at)) FROM study_sessions
			UNION ALL
			SELECT date(date, '-1 day')
			FROM dates
			WHERE date > date((
				SELECT MIN(created_at) FROM study_sessions
			))
		)
		SELECT COUNT(*)
		FROM dates d
		WHERE EXISTS (
			SELECT 1 FROM study_sessions 
			WHERE date(created_at) = d.date
		)
	`).Scan(&stats.StudyStreakDays)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"success_rate":         stats.SuccessRate,
		"total_study_sessions": stats.TotalSessions,
		"total_active_groups":  stats.ActiveGroups,
		"study_streak_days":    stats.StudyStreakDays,
	}, nil
}
