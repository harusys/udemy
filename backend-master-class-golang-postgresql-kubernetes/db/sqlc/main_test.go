package db

import (
	"context"
	"log"
	"os"
	"simplebank/util"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	// Setup
	location, _ := time.LoadLocation("JST")
	time.Local = location

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testStore = NewStore(connPool)

	// Do all tests
	code := m.Run()

	// Clean up
	// TODO: テスト実行後にデータ削除する

	// End status
	os.Exit(code)
}
