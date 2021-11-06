// Package models contains url struct records for database
package models

import (
	"database/sql"
	"errors"
	"time"
)


var (
	ErrRecordNotFound = errors.New("record not found")
)

type Url struct {
	ID uint64 `json:"id"`
	CreatedAt time.Time	`json:"-"`
	URL string `json:"url"`
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

func (u UrlModel) Get(id uint64) (*Url, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, url
		FROM urls
		WHERE id = $1`
	
	var url Url

	err := u.DB.QueryRow(query, id).Scan(
		&url.ID,
		&url.CreatedAt,
		&url.URL,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &url, nil
}

func (u UrlModel) Delete(id uint64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
	DELETE FROM urls
	WHERE id = $1`

	result, err := u.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}


