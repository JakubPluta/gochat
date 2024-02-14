package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

// CreateUser creates a new user in the repository.
//
// ctx context.Context, user *User
// *User, error
func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastInsertId)
	if err != nil {
		return &User{}, err
	}
	user.ID = int64(lastInsertId)
	return user, nil
}

// GetUserByEmail retrieves a user by email from the repository.
//
// ctx: context.Context
// email: string
// *User: user information
// error: any error that occurred
func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// NewRepository creates a new Repository with the given DBTX.
// It returns a Repository.
func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}
