package http

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/getmiranda/meli-challenge-api/logger"
	dbRepo "github.com/getmiranda/meli-challenge-api/repository/db"
	"github.com/getmiranda/meli-challenge-api/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	ctx  context.Context
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	pingHandler  PingHandler
	humanHandler HumanHandler
}

func (s *Suite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(s.T(), err)

	s.ctx = context.Background()

	s.pingHandler = MakePingHandler()

	dbRepo := dbRepo.MakeDBRepository(s.DB)
	hService := services.MakeHumansService(dbRepo)
	s.humanHandler = MakeHumanHandler(hService)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
