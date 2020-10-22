package testdata

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	// load PostgreSQL driver
	// _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/hunyxv/go-dao"
	"github.com/hunyxv/go-dao/mock"
)

var db *sql.DB

func init() {
	var err error

	host := "hongning"
	database := "test"
	port := 3306
	user := "test"
	passwd := "123456"

	dbinfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8",
		user, passwd, host, port, database)

	db, err = sql.Open("mysql", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Printf("connection to %v was established", dbinfo)
}

type DatabaseTestSuite struct {
	suite.Suite
	manager dao.Manager
	factory dao.Factory
}

func (s *DatabaseTestSuite) SetupSuite() {

	ds := &dao.DataSource{
		DB:   db,
		Name: "travis_ci_test",
	}

	s.manager = dao.NewBaseManager()
	s.manager.RegisterDataSource(ds)
	log.Printf("registered data source %v", ds)

	s.factory = mock.NewFactory(ds)
	s.manager.RegisterFactory(s.factory)
	log.Printf("registered factory %v", s.factory)
}

// TestPackage test suite for the dao package.
func TestPackage(t *testing.T) {
	suite.Run(t, new(DAOTestSuite))
	suite.Run(t, new(HandlerTestSuite))
	suite.Run(t, new(ManagerTestSuite))
}

// helpers

func assertNoError(t *testing.T, err error) {
	if err != nil {
		assert.Fail(t, err.Error())
	}
}
