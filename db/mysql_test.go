package db

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestInitMyqlDb() {
	_, err := InitMyqlDb("", "", "", "", "")
	require.Error(s.T(), err)
}

func (s *Suite) TestMigrate() {
	rows := s.mock.NewRows([]string{"id", "letters", "is_valid", "created_at"})

	s.mock.ExpectQuery("SHOW TABLES FROM `` WHERE `Tables_in_` = ?").WithArgs("sequences").WillReturnRows(rows)

	Migrate(s.DB)
}
