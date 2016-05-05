package content

import (
	"database/sql"
	"time"
)

// FindByID function
func FindByID(db *sql.DB, contentID int64) (*Content, error) {
	foundContent := new(Content)
	err := db.QueryRow(`SELECT id, paper_id, title, body, created_at FROM contents WHERE id = ? LIMIT 1`, contentID).Scan(&foundContent.ID, &foundContent.PaperID, &foundContent.Title, &foundContent.Body, &foundContent.CreatedAt)
	return foundContent, err
}

// FindByPaperID function
func FindByPaperID(db *sql.DB, paperID int64) ([]*Content, error) {
	var foundContents []*Content
	rows, err := db.Query(`SELECT id, paper_id, title, body, created_at FROM contents WHERE paper_id = ? LIMIT 1`, paperID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		content := new(Content)
		if err = rows.Scan(&content.ID, &content.PaperID, &content.Title, &content.Body, &content.CreatedAt); err != nil {
			return nil, err
		}
		foundContents = append(foundContents, content)
	}

	return foundContents, err
}

// create function
func create(db *sql.DB, content Content) (int64, error) {
	res, err := db.Exec(`INSERT INTO contents (paper_id, title, body, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`, content.PaperID, content.Title, content.Body, time.Now(), time.Now())
	if err != err {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}
