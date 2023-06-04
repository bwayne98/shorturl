// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package user

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
  name, line_id
 ) VALUES (
  $1, $2
 )
 RETURNING id, name, line_id, created_at
`

type CreateUserParams struct {
	Name   string `json:"name"`
	LineID string `json:"line_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.LineID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LineID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
 DELETE FROM users
 WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
 SELECT id, name
 FROM users
 WHERE line_id = $1
`

type GetUserRow struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) GetUser(ctx context.Context, lineID string) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, lineID)
	var i GetUserRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
