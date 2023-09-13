package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) // テスト時のログ出力を抑制する
	os.Exit(m.Run())
}
