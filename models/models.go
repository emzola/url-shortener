// Package models contains url struct records for database
package models

import (
	"database/sql"
	"time"
)

type Url struct {
	ID uint64 `json:"id"`
	CreatedAt time.Time	`json:"-"`
	URL string `json:"url"`
	Visits int `json:"-"`
}

type UrlModel struct {
	DB *sql.DB
}

func (u UrlModel) Insert(webUrl *Url) error {
	query := `
	INSERT INTO urls (url)
	VALUES ($1)
	RETURNING id, created_at`

	return u.DB.QueryRow(query, webUrl.URL).Scan(&webUrl.ID, &webUrl.CreatedAt)
}



