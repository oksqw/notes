package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"notes"
	"notes/pkg/xerror"
)

type SQLiteNoteRepository struct {
	db *sql.DB
}

func (r *SQLiteNoteRepository) Create(input notes.Note) (*notes.Note, error) {
	var output notes.Note
	query := `INSERT INTO notes (title, text) VALUES (?, ?) RETURNING id, title, text, created_at, updated_at`
	err := r.db.QueryRow(query, input.Title, input.Text).Scan(
		&output.Id,
		&output.Title,
		&output.Text,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert and retrieve note: %w", err)
	}

	return &output, nil
}

func (r *SQLiteNoteRepository) Get(id int) (*notes.Note, error) {
	var output notes.Note
	query := "SELECT id, title, text, created_at, updated_at FROM notes WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(
		&output.Id,
		&output.Title,
		&output.Text,
		&output.CreatedAt,
		&output.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, &xerror.NotFoundError{Message: "note with id %d not found", ID: id}
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update and retrieve note: %w", err)
	}

	return &output, nil
}

func (r *SQLiteNoteRepository) Update(input notes.Note) (*notes.Note, error) {
	var output notes.Note
	query := "UPDATE notes SET title = ?, text = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ? RETURNING id, title, text, created_at, updated_at"
	err := r.db.QueryRow(query, input.Title, input.Text, input.Id).Scan(
		&output.Id,
		&output.Title,
		&output.Text,
		&output.CreatedAt,
		&output.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, &xerror.NotFoundError{Message: "note with id %d not found", ID: input.Id}
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update and retrieve note: %w", err)
	}

	return &output, nil
}

func (r *SQLiteNoteRepository) Delete(id int) (*notes.Note, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer handleRollbackTx(tx, err)

	var output notes.Note
	query := "SELECT id, title, text, created_at, updated_at FROM notes WHERE id = ?"
	err = tx.QueryRow(query, id).Scan(
		&output.Id,
		&output.Title,
		&output.Text,
		&output.CreatedAt,
		&output.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, &xerror.NotFoundError{Message: "note with id %d not found", ID: id}
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query note: %w", err)
	}

	query = "DELETE FROM notes WHERE id = ?"
	_, err = tx.Exec(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete note: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &output, nil
}
func (r *SQLiteNoteRepository) All() ([]notes.Note, error) {
	query := fmt.Sprintf(`SELECT id, title, text, created_at, updated_at FROM notes`)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query notes: %w", err)
	}
	defer handleCloseRows(rows)

	var output []notes.Note
	for rows.Next() {
		var note notes.Note
		if err := rows.Scan(&note.Id, &note.Title, &note.Text, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan note row: %w", err)
		}
		output = append(output, note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over rows: %w", err)
	}

	return output, nil
}

func NewSQLiteNoteRepository(db *sql.DB) *SQLiteNoteRepository {
	return &SQLiteNoteRepository{
		db: db,
	}
}

func handleRollbackTx(tx *sql.Tx, initiator error) {
	if initiator == nil {
		return
	}
	if err := tx.Rollback(); err != nil {
		logrus.Fatalf("update failed: %v, unable to back: %v", initiator, err)
	}
}

func handleCloseRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		logrus.Fatalf("failed to close rows: %w", err)
	}
}
