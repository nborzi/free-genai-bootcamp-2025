package models

import "time"

type Word struct {
	ID      int    `json:"id"`
	Spanish string `json:"spanish"`
	English string `json:"english"`
	Parts   string `json:"parts"` // JSON string
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type StudySession struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"group_id"`
	StudyActivityID int       `json:"study_activity_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type StudyActivity struct {
	ID             int       `json:"id"`
	StudySessionID int       `json:"study_session_id"`
	GroupID        int       `json:"group_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type WordReviewItem struct {
	WordID         int       `json:"word_id"`
	StudySessionID int       `json:"study_session_id"`
	Correct        bool      `json:"correct"`
	CreatedAt      time.Time `json:"created_at"`
}

type DashboardStats struct {
	LastSession    *StudySession `json:"last_session"`
	Progress       Progress      `json:"progress"`
	QuickStats     QuickStats    `json:"stats"`
}

type Progress struct {
	TotalWords     int     `json:"total_words"`
	WordsStudied   int     `json:"words_studied"`
	CorrectAnswers int     `json:"correct_answers"`
	Accuracy       float64 `json:"accuracy"`
}

type QuickStats struct {
	StudySessions  int `json:"study_sessions"`
	WordsReviewed  int `json:"words_reviewed"`
	GroupsStudied  int `json:"groups_studied"`
	DaysStreak     int `json:"days_streak"`
}
