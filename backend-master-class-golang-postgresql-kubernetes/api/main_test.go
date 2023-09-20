package api

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	db "simplebank/db/sqlc"
	"simplebank/test"
	"simplebank/util"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   test.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) // テスト時のログ出力を抑制する
	os.Exit(m.Run())
}

// HTTP レスポンスボディ検証
func requireBodyMatch[T any](t *testing.T, body *bytes.Buffer, expected T) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotValue T
	err = json.Unmarshal(data, &gotValue)
	require.NoError(t, err)

	require.Equal(t, expected, gotValue)
}
