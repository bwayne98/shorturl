// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package shorturl

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateShorturl(ctx context.Context, arg CreateShorturlParams) (Shorturl, error)
	DeleteShorturl(ctx context.Context, arg DeleteShorturlParams) error
	GetMatchShorturl(ctx context.Context, arg GetMatchShorturlParams) (string, error)
	ListUserShorturl(ctx context.Context, userID sql.NullInt32) ([]ListUserShorturlRow, error)
	UpdateExpired(ctx context.Context, arg UpdateExpiredParams) (Shorturl, error)
}

var _ Querier = (*Queries)(nil)
