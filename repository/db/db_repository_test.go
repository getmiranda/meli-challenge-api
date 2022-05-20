package db

import (
	"context"
	"database/sql"
	"net/http"
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

		s.NotNil(err)
		s.EqualValues(errors_utils.ErrDatabase.Error(), err.Error())
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

		s.Nil(err)
	})
}

func (s *Suite) TestGetHumanByDna() {
	query := `SELECT * FROM "humans" WHERE dna = $1 AND "humans"."deleted_at" IS NULL ORDER BY "humans"."id" LIMIT 1`
	human := &humans.Human{
		Dna:      "AAAA-GGGG-CCCC-TTTT",
		IsMutant: true,
	}

	s.T().Run("HumanNotFound", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(human.Dna).
			WillReturnError(gorm.ErrRecordNotFound)

		human, err := s.repository.GetHumanByDna(s.ctx, human.Dna)

		s.Nil(human)
		s.Nil(err)
	})

	s.T().Run("UnexpectedError", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(human.Dna).
			WillReturnError(gorm.ErrInvalidDB)

		_, err := s.repository.GetHumanByDna(s.ctx, human.Dna)

		s.NotNil(err)
		s.EqualValues(errors_utils.ErrDatabase.Error(), err.Error())
	})

	s.T().Run("NoError", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(human.Dna).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "dna", "is_mutant"}).
				AddRow(1, human.CreatedAt, human.UpdatedAt, human.DeletedAt, human.Dna, human.IsMutant))

		h, err := s.repository.GetHumanByDna(s.ctx, human.Dna)

		s.Nil(err)
		s.NotNil(human)
		s.EqualValues(human.Dna, h.Dna)
		s.EqualValues(human.IsMutant, h.IsMutant)

	})
}

func (s *Suite) TestCountMutants() {
	query := `SELECT count(*) FROM "humans" WHERE is_mutant = $1 AND "humans"."deleted_at" IS NULL`

	s.T().Run("UnexpectedError", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(true).
			WillReturnError(gorm.ErrInvalidDB)

		count, err := s.repository.CountMutants(s.ctx)

		s.NotNil(err)
		s.EqualValues(errors_utils.ErrDatabase.Error(), err.Error())
		s.EqualValues(http.StatusInternalServerError, err.Status())
		s.EqualValues(0, count)
	})

	s.T().Run("NoError", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(true).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).
				AddRow(5))

		count, err := s.repository.CountMutants(s.ctx)

		s.Nil(err)
		s.EqualValues(5, count)
	})
}

func (s *Suite) TestCountHumans() {
	query := `SELECT count(*) FROM "humans" WHERE is_mutant = $1 AND "humans"."deleted_at" IS NULL`

	s.T().Run("UnexpectedError", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(false).
			WillReturnError(gorm.ErrInvalidDB)

		count, err := s.repository.CountHumans(s.ctx)

		s.NotNil(err)
		s.EqualValues(errors_utils.ErrDatabase.Error(), err.Error())
		s.EqualValues(http.StatusInternalServerError, err.Status())
		s.EqualValues(0, count)
	})

	s.T().Run("NoError", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(false).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).
				AddRow(5))

		count, err := s.repository.CountHumans(s.ctx)

		s.Nil(err)
		s.EqualValues(5, count)
	})
}
