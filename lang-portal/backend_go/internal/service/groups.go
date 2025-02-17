package service

import "lang-portal/internal/models"

func GetGroups() ([]models.Group, error) {
	var groups []models.Group
	rows, err := db.Query("SELECT id, name FROM groups")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var group models.Group
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func GetGroup(id int) (*models.Group, error) {
	var group models.Group
	err := db.QueryRow("SELECT id, name FROM groups WHERE id = ?", id).Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func GetGroupWords(groupID int) ([]models.Word, error) {
	var words []models.Word
	rows, err := db.Query(`
		SELECT w.id, w.spanish, w.english, w.parts 
		FROM words w 
		JOIN words_groups wg ON w.id = wg.word_id 
		WHERE wg.group_id = ?`, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var word models.Word
		if err := rows.Scan(&word.ID, &word.Spanish, &word.English, &word.Parts); err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	return words, nil
}
