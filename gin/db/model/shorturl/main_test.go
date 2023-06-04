package shorturl

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

var testQuery *Queries

func TestMain(m *testing.M){
	conn, err := sql.Open("postgres", "postgresql://wayne:123456qq@pg:5432/shorturl?sslmode=disable")
	if err != nil{
		log.Fatal("Cant Connect DB:", err)
	}

	testQuery = New(conn)
	os.Exit(m.Run())
}

func creatRandomShorturl(t *testing.T) Shorturl {
	params := CreateShorturlParams{
		Origin: util.RandUrl(),
		Match: util.RandString(10),
		ExpiredAt: time.Now().Add(time.Hour *2),
	}

	shorturl, err := testQuery.CreateShorturl(context.Background(), params)

	require.NoError(t, err)

	require.Equal(t, params.Origin, shorturl.Origin)
	require.Equal(t, params.Match, shorturl.Match)
	require.WithinDuration(t, params.ExpiredAt, shorturl.ExpiredAt, time.Minute)

	return shorturl
}

func TestCreateShorturl(t *testing.T){
	creatRandomShorturl(t)
}

func TestGetMatchShorturl(t *testing.T){
	shorturl := creatRandomShorturl(t)

	origin, err := testQuery.GetMatchShorturl(context.Background(), shorturl.Match)

	require.NoError(t, err)
	require.Equal(t, shorturl.Origin, origin)
}