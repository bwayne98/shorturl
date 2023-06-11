package short_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"shorturl/m/api"
	"shorturl/m/db/store"
	"shorturl/m/util"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type MakeMacher struct {
}

func (m MakeMacher) Matches(x interface{}) bool {
	_, ok := x.(store.CreateShorturlParams)
	return ok
}
func (m MakeMacher) String() string {
	return "CreateShourturlParams not match"
}

func MakeParamsEq() gomock.Matcher {
	return MakeMacher{}
}

func TestMake(t *testing.T) {

	testCase := []struct {
		name      string
		buildFunc func(mockQuery *store.MockQuerier) []byte
		checkFunc func(t *testing.T, record *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildFunc: func(mockQuery *store.MockQuerier) []byte {

				shorturl := randShorturl()

				mockQuery.EXPECT().
					CountMatchShorturl(gomock.Any(), gomock.Any()).
					AnyTimes().
					Return(int64(0), nil)

				mockQuery.EXPECT().
					CreateShorturl(gomock.Any(), MakeParamsEq()).
					Times(1).
					Return(shorturl, nil)

				value := map[string]interface{}{"origin": shorturl.Origin}
				json, err := json.Marshal(value)
				require.NoError(t, err)

				return json
			},
			checkFunc: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, record.Code)
			},
		},
		{
			name: "Binding fail",
			buildFunc: func(mockQuery *store.MockQuerier) []byte {

				value := map[string]interface{}{}
				json, err := json.Marshal(value)
				require.NoError(t, err)

				return json
			},
			checkFunc: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, record.Code)
			},
		},
		{
			name: "Origin fail",
			buildFunc: func(mockQuery *store.MockQuerier) []byte {

				value := map[string]interface{}{"origin": "https://example.com/test"}
				json, err := json.Marshal(value)
				require.NoError(t, err)

				return json
			},
			checkFunc: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, record.Code)
			},
		},
		{
			name: "Unique Match Sever Error",
			buildFunc: func(mockQuery *store.MockQuerier) []byte {
				shorturl := randShorturl()

				mockQuery.EXPECT().
					CountMatchShorturl(gomock.Any(), gomock.Any()).
					AnyTimes().
					Return(int64(0), sql.ErrConnDone)

				value := map[string]interface{}{"origin": shorturl.Origin}
				json, err := json.Marshal(value)
				require.NoError(t, err)

				return json
			},
			checkFunc: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, record.Code)
			},
		},
		{
			name: "CreateShorturl Server Error",
			buildFunc: func(mockQuery *store.MockQuerier) []byte {

				shorturl := randShorturl()

				mockQuery.EXPECT().
					CountMatchShorturl(gomock.Any(), gomock.Any()).
					AnyTimes().
					Return(int64(0), nil)

				mockQuery.EXPECT().
					CreateShorturl(gomock.Any(), MakeParamsEq()).
					Times(1).
					Return(store.Shorturl{}, sql.ErrConnDone)

				value := map[string]interface{}{"origin": shorturl.Origin}
				json, err := json.Marshal(value)
				require.NoError(t, err)

				return json
			},
			checkFunc: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, record.Code)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockQuery := store.NewMockQuerier(ctrl)

			json := tc.buildFunc(mockQuery)

			route := gin.Default()
			server := api.NewServer(util.Config{}, mockQuery, route)
			server.SetupRoute()
			record := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodPost, "/api/short/make", bytes.NewBuffer(json))
			require.NoError(t, err)

			request.Header.Set("Content-Type", "application/json")
			request.Host = "example.com"
			server.Router.ServeHTTP(record, request)

			tc.checkFunc(t, record)
		})
	}
}

func TestMatch(t *testing.T) {

	testCase := []struct {
		name      string
		buildFunc func(mockQuery *store.MockQuerier) []byte
		checkFunc func(t *testing.T, record *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildFunc: func(mockQuery *store.MockQuerier) []byte {
				shorturl := randShorturl()

				mockQuery.EXPECT().
					GetMatchShorturl(gomock.Any(), gomock.Eq(shorturl.Match)).
					Times(1).
					Return(shorturl.Origin, nil)

				value := map[string]interface{}{"match": shorturl.Match}
				json, err := json.Marshal(value)
				require.NoError(t, err)

				return json
			},
			checkFunc: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, record.Code)

			},
		},
		{
			name: "Binding fail",
			buildFunc: func(mockQuery *store.MockQuerier) []byte {

				value := map[string]interface{}{}
				json, err := json.Marshal(value)
				require.NoError(t, err)

				return json
			},
			checkFunc: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, record.Code)
			},
		},
		{
			name: "No Row Error",
			buildFunc: func(mockQuery *store.MockQuerier) []byte {

				match := "example890"

				mockQuery.EXPECT().
					GetMatchShorturl(gomock.Any(), gomock.Eq(match)).
					Times(1).
					Return("", sql.ErrNoRows)

				value := map[string]interface{}{"match": match}
				json, err := json.Marshal(value)
				require.NoError(t, err)

				return json
			},
			checkFunc: func(t *testing.T, record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, record.Code)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockQuery := store.NewMockQuerier(ctrl)

			json := tc.buildFunc(mockQuery)

			router := gin.Default()
			server := api.NewServer(util.Config{}, mockQuery, router)
			server.SetupRoute()
			record := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodPost, "/api/short/match", bytes.NewBuffer(json))
			require.NoError(t, err)

			server.Router.ServeHTTP(record, request)

			tc.checkFunc(t, record)
		})
	}

}

func randShorturl() store.Shorturl {
	return store.Shorturl{
		ID:        int32(util.RandInt(1, 100)),
		UserID:    int32(util.RandInt(1, 100)),
		Origin:    util.RandUrl(),
		Match:     util.RandString(10),
		ExpiredAt: time.Now().Add(time.Hour * 2),
		CreatedAt: time.Now(),
	}
}
