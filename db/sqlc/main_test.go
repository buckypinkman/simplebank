package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/buckypinkman/simplebank/config/db"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(db.DBDriver, db.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}