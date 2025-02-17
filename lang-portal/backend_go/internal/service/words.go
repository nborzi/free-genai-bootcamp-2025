package service

func GetWords(page int) ([]WordWithStats, int, error) {
	offset := (page - 1) * 100
	query := `
		SELECT 
			w.spanish,
			w.english,
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
		err := rows.Scan(&w.Spanish, &w.English, &w.CorrectCount, &w.WrongCount)
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

func GetWord(id int) (*struct {
	WordWithStats
	Groups []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
}, error) {
	word := &struct {
		WordWithStats
		Groups []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"groups"`
	}{}

	// Get word details with stats
	err := db.QueryRow(`
		SELECT 
			w.spanish,
			w.english,
			COUNT(CASE WHEN wri.correct THEN 1 END) as correct_count,
			COUNT(CASE WHEN NOT wri.correct THEN 1 END) as wrong_count
		FROM words w
		LEFT JOIN word_review_items wri ON w.id = wri.word_id
		WHERE w.id = ?
		GROUP BY w.id
	`, id).Scan(&word.Spanish, &word.English, &word.CorrectCount, &word.WrongCount)
	if err != nil {
		return nil, err
	}

	// Get groups for the word
	rows, err := db.Query(`
		SELECT g.id, g.name
		FROM groups g
		JOIN words_groups wg ON g.id = wg.group_id
		WHERE wg.word_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var group struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		err := rows.Scan(&group.ID, &group.Name)
		if err != nil {
			return nil, err
		}
		word.Groups = append(word.Groups, group)
	}

	return word, nil
}

type WordWithStats struct {
	Spanish      string `json:"spanish"`
	English      string `json:"english"`
	CorrectCount int    `json:"correct_count"`
	WrongCount   int    `json:"wrong_count"`
}
