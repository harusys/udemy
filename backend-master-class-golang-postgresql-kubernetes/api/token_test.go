package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	mockdb "simplebank/db/mock"
	db "simplebank/db/sqlc"
	"simplebank/test"
	"simplebank/token"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRenewAccessToken(t *testing.T) {
	username := test.RandomOwner()

	// サーバ側と同じ固定の秘密鍵でトークンを作成
	maker, err := token.NewPasetoMaker(test.FixedTokenSymmetricKey())
	require.NoError(t, err)
	refreshToken, _, err := maker.CreateToken(username, time.Minute)
	require.NoError(t, err)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"refresh_token": refreshToken,
			},
			buildStubs: func(store *mockdb.MockStore) {
				session := db.Session{
					ID:           uuid.New(),
					Username:     username,
					RefreshToken: refreshToken,
					IsBlocked:    false,
					ExpiresAt:    time.Now().Add(time.Minute),
				}
				store.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(session, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest_InvalidRefreshToken",
			body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Unauthorized_IncorrectToken",
			body: gin.H{
				"refresh_token": "incorrect",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "NotFound_GetSession",
			body: gin.H{
				"refresh_token": refreshToken,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Session{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalServerError_GetSession",
			body: gin.H{
				"refresh_token": refreshToken,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Session{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "Unauthorized_BlockedSession",
			body: gin.H{
				"refresh_token": refreshToken,
			},
			buildStubs: func(store *mockdb.MockStore) {
				session := db.Session{
					ID:           uuid.New(),
					Username:     username,
					RefreshToken: refreshToken,
					IsBlocked:    true,
					ExpiresAt:    time.Now().Add(time.Minute),
				}
				store.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(session, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Unauthorized_IncorrectSessionUser",
			body: gin.H{
				"refresh_token": refreshToken,
			},
			buildStubs: func(store *mockdb.MockStore) {
				session := db.Session{
					ID:           uuid.New(),
					Username:     "incorrect",
					RefreshToken: refreshToken,
					IsBlocked:    false,
					ExpiresAt:    time.Now().Add(time.Minute),
				}
				store.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(session, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Unauthorized_MismatchedSessionToken",
			body: gin.H{
				"refresh_token": refreshToken,
			},
			buildStubs: func(store *mockdb.MockStore) {
				session := db.Session{
					ID:           uuid.New(),
					Username:     username,
					RefreshToken: "mismatched",
					IsBlocked:    false,
					ExpiresAt:    time.Now().Add(time.Minute),
				}
				store.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(session, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Unauthorized_ExpiredToken",
			body: gin.H{
				"refresh_token": refreshToken,
			},
			buildStubs: func(store *mockdb.MockStore) {
				session := db.Session{
					ID:           uuid.New(),
					Username:     username,
					RefreshToken: refreshToken,
					IsBlocked:    false,
					ExpiresAt:    time.Now().Add(-time.Minute),
				}
				store.EXPECT().
					GetSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(session, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// Act
			server := newTestFixedKeyServer(t, store) // テスト用の秘密鍵を固定
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/tokens/renew_access"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			// Assert
			tc.checkResponse(recorder)
		})
	}
}
