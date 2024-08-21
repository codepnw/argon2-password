package store_test

import (
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type StoreTestSuite struct {
	suite.Suite
	db *sqlx.DB
}

func (s *StoreTestSuite) SetupTest() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	conn := os.Getenv("DB_URL")
	s.db = sqlx.MustConnect("postgres", conn)
	s.db.MustExec("truncate users cascade")
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}
