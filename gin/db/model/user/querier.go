// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package user

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int32) error
	GetUser(ctx context.Context, lineID string) (GetUserRow, error)
}

var _ Querier = (*Queries)(nil)
