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

type WordGroup struct {
	ID      int `json:"id"`
	WordID  int `json:"word_id"`
	GroupID int `json:"group_id"`
}

type StudySession struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
	StudyActivityID int       `json:"study_activity_id"`
}

type StudyActivity struct {
	ID              int       `json:"id"`
	StudySessionID  int       `json:"study_session_id"`
	GroupID         int       `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type WordReviewItem struct {
	WordID         int       `json:"word_id"`
	StudySessionID int       `json:"study_session_id"`
	Correct        bool      `json:"correct"`
	CreatedAt      time.Time `json:"created_at"`
}
