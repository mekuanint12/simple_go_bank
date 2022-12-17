package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	dbDriver = "postgres"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to the database: ", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
