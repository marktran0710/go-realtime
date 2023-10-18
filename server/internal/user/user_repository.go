package user

import (
	"context"
	"database/sql"
)

type repository struct {
	db DBTX
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertID int
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) returing id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertID)
	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInsertID)
	return user, nil
}

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Rows
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}
