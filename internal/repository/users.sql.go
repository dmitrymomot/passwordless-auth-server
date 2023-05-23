// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: users.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = $1
`

// Delete a user by ID
func (q *Queries) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteUserByIDStmt, deleteUserByID, id)
	return err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, email, verified, updated_at, created_at FROM users WHERE email = $1
`

// Find a user by email
func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.findUserByEmailStmt, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT id, email, verified, updated_at, created_at FROM users WHERE id = $1
`

// Find a user by ID
func (q *Queries) FindUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.queryRow(ctx, q.findUserByIDStmt, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const storeOrUpdateUser = `-- name: StoreOrUpdateUser :one
INSERT INTO users (email) VALUES ($1) ON CONFLICT (email) DO NOTHING RETURNING id, email, verified, updated_at, created_at
`

// Store or update a user
func (q *Queries) StoreOrUpdateUser(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.storeOrUpdateUserStmt, storeOrUpdateUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserEmailByID = `-- name: UpdateUserEmailByID :exec
UPDATE users SET email = $1 WHERE id = $2
`

type UpdateUserEmailByIDParams struct {
	Email string    `json:"email"`
	ID    uuid.UUID `json:"id"`
}

// Update a user's email by ID
func (q *Queries) UpdateUserEmailByID(ctx context.Context, arg UpdateUserEmailByIDParams) error {
	_, err := q.exec(ctx, q.updateUserEmailByIDStmt, updateUserEmailByID, arg.Email, arg.ID)
	return err
}
