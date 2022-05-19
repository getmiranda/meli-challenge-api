package db

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getmiranda/meli-challenge-api/domain/humans"
	_ "github.com/getmiranda/meli-challenge-api/logger"
	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	ctx  context.Context

	repository DBRepository
}

func (s *Suite) SetupSuite() {
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
	s.repository = MakeDBRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestSaveHuman() {
	query := `INSERT INTO "humans" ("created_at","updated_at","deleted_at","dna","is_mutant") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`
	human := &humans.Human{
		Dna:      "AAAA-GGGG-CCCC-TTTT",
		IsMutant: true,
	}

	s.T().Run("UnexpectedError", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), human.Dna, human.IsMutant).
			WillReturnError(gorm.ErrInvalidDB)
		s.mock.ExpectRollback()

		err := s.repository.SaveHuman(s.ctx, human)

		require.NotNil(s.T(), err)
		require.EqualValues(s.T(), errors_utils.ErrDatabase.Error(), err.Error())
	})

	s.T().Run("NoError", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), human.Dna, human.IsMutant).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(1))
		s.mock.ExpectCommit()

		err := s.repository.SaveHuman(s.ctx, human)

		require.Nil(s.T(), err)
	})
}
