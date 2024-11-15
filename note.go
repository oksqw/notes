package notes

import "time"

type Note struct {
	Id        int       `json:"id"         db:"id"`
	Title     string    `json:"title"      db:"title"`
	Text      string    `json:"text"       db:"text"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
