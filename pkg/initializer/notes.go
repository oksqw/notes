package initializer

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitializeNotesSQLiteDb(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, text TEXT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP);")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
