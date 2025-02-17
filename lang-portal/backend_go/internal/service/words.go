package service

import (
	"database/sql"
	"lang-portal/internal/models"
)

type WordWithStats struct {
	models.Word
	CorrectCount int `json:"correct_count"`
	WrongCount   int `json:"wrong_count"`
}

func GetWords(page int) ([]WordWithStats, int, error) {
	offset := (page - 1) * 100
	query := `
		SELECT 
			w.id,
			w.spanish,
			w.english,
			w.parts,
			COUNT(CASE WHEN wri.correct THEN 1 END) as correct_count,
			COUNT(CASE WHEN NOT wri.correct THEN 1 END) as wrong_count
		FROM words w
		LEFT JOIN word_review_items wri ON w.id = wri.word_id
		GROUP BY w.id
		ORDER BY w.id
		LIMIT 100 OFFSET ?
	`

	rows, err := db.Query(query, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var words []WordWithStats
	for rows.Next() {
		var w WordWithStats
		err := rows.Scan(&w.ID, &w.Spanish, &w.English, &w.Parts, &w.CorrectCount, &w.WrongCount)
		if err != nil {
			return nil, 0, err
		}
		words = append(words, w)
	}

	// Get total count
	var total int
	err = db.QueryRow("SELECT COUNT(*) FROM words").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return words, total, nil
}

func GetWord(id int) (*WordWithStats, error) {
	var word WordWithStats

	err := db.QueryRow(`
		SELECT 
			w.id,
			w.spanish,
			w.english,
			w.parts,
			COUNT(CASE WHEN wri.correct THEN 1 END) as correct_count,
			COUNT(CASE WHEN NOT wri.correct THEN 1 END) as wrong_count
		FROM words w
		LEFT JOIN word_review_items wri ON w.id = wri.word_id
		WHERE w.id = ?
		GROUP BY w.id
	`, id).Scan(&word.ID, &word.Spanish, &word.English, &word.Parts, &word.CorrectCount, &word.WrongCount)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &word, nil
}

func ReviewWord(wordID int, correct bool) error {
	// Get the latest study session
	var sessionID int
	err := db.QueryRow("SELECT id FROM study_sessions ORDER BY created_at DESC LIMIT 1").Scan(&sessionID)
	if err != nil {
		return err
	}

	// Insert the review
	_, err = db.Exec(`
		INSERT INTO word_review_items (word_id, study_session_id, correct) 
		VALUES (?, ?, ?)
	`, wordID, sessionID, correct)

	return err
}
