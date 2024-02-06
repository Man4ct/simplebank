package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	testQueries = New(db)
	os.Exit(m.Run())
}
