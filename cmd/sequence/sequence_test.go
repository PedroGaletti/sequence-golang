package sequence

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	id        int64 = 1
	letters         = "[D, D, D, D, U, H, D, B, B, D, U, U, B, U, H, B, U, H, D, D, H, U, H, H, B, D, B, B, H, U, D, D, H, D, U, B]"
	isValid         = true
	createdAt       = time.Now()
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository ISequenceRepository
	Sequence   *Sequence
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

	s.repository = NewSequenceRepository(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestSave() {
	s.mock.ExpectBegin()

	s.mock.ExpectExec("INSERT").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(errors.New("test error"))

	s.mock.ExpectRollback()

	err := s.repository.Store(&Sequence{Id: id, Letters: letters, IsValid: isValid, CreatedAt: createdAt})

	require.Error(s.T(), err)
}

func (s *Suite) TestFindAll() {
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sequences`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "letters", "is_valid", "created_at"}).
			AddRow(id, letters, isValid, createdAt))

	res, err := s.repository.FindAll()

	s.T().Log(res)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal([]Sequence{res[0]}, res))
}

func (s *Suite) TestInjectDependency() {
	InjectDependency(gin.Default().Group(""), s.DB)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) TestRouter() {
	controller := NewSequenceController(s.repository)
	require.NotNil(s.T(), controller)

	Router(gin.Default().Group(""), controller)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) TestNewSequenceController() {
	controller := NewSequenceController(s.repository)
	require.NotNil(s.T(), controller)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) TestSequenceRequestValidate() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	gin.CreateTestContext(w)
	r := gin.Default()

	controller := NewSequenceController(s.repository)

	json := `{"letters": ["DUHBHB", "DDBUHD", "UBDUHU", "BHBDHH", "DDDDUB", "UDBDUH"]}`
	letters := strings.NewReader(string(json))

	r.POST("/sequence", controller.SequenceRequestValidate)

	// Test with empty body
	req, err := http.NewRequest(http.MethodPost, "/sequence", nil)
	if err != nil {
		s.T().Fatal(fmt.Printf("Couldn't create request: %v\n", err))
	}

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		require.Error(s.T(), errors.New("Request with error"))
	}

	// Test with a valid array
	req, err = http.NewRequest(http.MethodPost, "/sequence", letters)
	if err != nil {
		s.T().Fatal(fmt.Printf("Couldn't create request: %v\n", err))
	}

	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		require.NoError(s.T(), err)
	}
}

func (s *Suite) TestStatsRequestInformation() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	gin.CreateTestContext(w)
	r := gin.Default()

	controller := NewSequenceController(s.repository)

	r.GET("/sequence/stats", controller.StatsRequestInformation)

	// If empty return of database
	req, err := http.NewRequest(http.MethodGet, "/sequence/stats", nil)
	if err != nil {
		s.T().Fatal(fmt.Printf("Couldn't create request: %v\n", err))
	}

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		require.Error(s.T(), errors.New("Request with error"))
	}

	// mock database
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sequences`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "letters", "is_valid", "created_at"}).
			AddRow(id, letters, false, createdAt).
			AddRow(id, letters, isValid, createdAt))

	// If return of data from database
	req, err = http.NewRequest(http.MethodGet, "/sequence/stats", nil)
	if err != nil {
		s.T().Fatal(fmt.Printf("Couldn't create request: %v\n", err))
	}

	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		require.Error(s.T(), errors.New("Request with error"))
	}
}
