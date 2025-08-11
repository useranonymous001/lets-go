package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbSource = "postgresql://root:secret@localhost:5432/root?sslmode=disable"
	dbDriver = "postgres"
)

func TestMain(m *testing.M) {

	// create a connection with database;
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Error connecting to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
