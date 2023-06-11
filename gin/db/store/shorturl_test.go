package store

import (
	"context"
	"database/sql"
	"log"
	"os"
	"shorturl/m/util"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var testQuery Querier

func TestMain(m *testing.M) {
	db, err := sql.Open("postgres", "postgresql://wayne:123456qq@pg:5432/shorturl?sslmode=disable")
	if err != nil {
		log.Fatal("connect db error:", err)
	}
	testQuery = New(db)
	os.Exit(m.Run())
}

func creatRandomShorturl(t *testing.T, userID int) Shorturl {

	params := CreateShorturlParams{
		Origin:    util.RandUrl(),
		Match:     util.RandString(10),
		UserID:    int32(userID),
		ExpiredAt: time.Now().Add(time.Hour * 2),
	}

	shorturl, err := testQuery.CreateShorturl(context.Background(), params)

	require.NoError(t, err)

	require.Equal(t, params.Origin, shorturl.Origin)
	require.Equal(t, params.Match, shorturl.Match)
	require.Equal(t, params.UserID, shorturl.UserID)
	require.WithinDuration(t, params.ExpiredAt, shorturl.ExpiredAt, time.Minute)

	return shorturl
}

func TestCreateShorturl(t *testing.T) {
	creatRandomShorturl(t, 1)
}

func TestGetMatchShorturl(t *testing.T) {

	shorturl := creatRandomShorturl(t, 2)

	origin, err1 := testQuery.GetMatchShorturl(context.Background(), shorturl.Match)

	require.NoError(t, err1)
	require.Equal(t, shorturl.Origin, origin)

	empty, err2 := testQuery.GetMatchShorturl(context.Background(), "")

	require.Error(t, err2)
	require.Empty(t, empty)
}

func TestDeleteShorturl(t *testing.T) {
	shorturl := creatRandomShorturl(t, 3)

	err1 := testQuery.DeleteShorturl(context.Background(), DeleteShorturlParams{
		ID:     shorturl.ID,
		UserID: shorturl.UserID,
	})

	require.NoError(t, err1)

	empty, err2 := testQuery.GetMatchShorturl(context.Background(), shorturl.Match)

	require.Error(t, err2)
	require.Empty(t, empty)
}

func TestListUserShorturl(t *testing.T) {

	userID := 19

	for i := 5; i < 10; i++ {
		creatRandomShorturl(t, userID)
	}

	list, err := testQuery.ListUserShorturl(context.Background(), int32(userID))

	require.NoError(t, err)

	for _, s := range list {
		require.Equal(t, s.UserID, int32(userID))
	}

}

func TestUpdateExpired(t *testing.T) {
	shorturl1 := creatRandomShorturl(t, 11)

	shorturl2, err := testQuery.UpdateExpired(context.Background(), UpdateExpiredParams{
		ID:        shorturl1.ID,
		ExpiredAt: time.Now().Add(time.Hour * 1),
	})

	require.NoError(t, err)
	require.Equal(t, shorturl1.ID, shorturl2.ID)
	require.Equal(t, shorturl1.UserID, shorturl2.UserID)
	require.Equal(t, shorturl1.UserID, shorturl2.UserID)
	require.Equal(t, shorturl1.Origin, shorturl2.Origin)
	require.Equal(t, shorturl1.Match, shorturl2.Match)
	require.Equal(t, shorturl1.CreatedAt, shorturl2.CreatedAt)
	require.NotEqual(t, shorturl1.ExpiredAt, shorturl2.ExpiredAt)

}

func TestCountMatchShorturl(t *testing.T) {
	shorturl := creatRandomShorturl(t, 31)

	count, err := testQuery.CountMatchShorturl(context.Background(), shorturl.Match)

	require.NoError(t, err)
	require.Equal(t, 1, int(count))
}
