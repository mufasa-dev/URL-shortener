package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "urls.db")
	if err != nil {
		return nil, err
	}

	query := `CREATE TABLE IF NOT EXISTS urls (
		id TEXT PRIMARY KEY,
		original_url TEXT NOT NULL
	);`
	_, err = db.Exec(query)
	return db, err
}

func StoreURL(db *sql.DB, shortURL, originalURL string) error {
	_, err := db.Exec("INSERT INTO urls (id, original_url) VALUES (?, ?)", shortURL, originalURL)
	return err
}

func GetOriginalURL(db *sql.DB, shortURL string) (string, error) {
	var originalURL string
	err := db.QueryRow("SELECT original_url FROM urls WHERE id = ?", shortURL).Scan(&originalURL)
	return originalURL, err
}
