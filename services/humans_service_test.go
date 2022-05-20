package services

import (
	"context"
	"database/sql"
	"net/http"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getmiranda/meli-challenge-api/domain/humans"
	_ "github.com/getmiranda/meli-challenge-api/logger"
	"github.com/getmiranda/meli-challenge-api/repository/db"
	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
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

	publicService  HumanService
	privateService *humanService
}

func (s *Suite) SetupSuite() {
	var (
		database *sql.DB
		err      error
	)

	database, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}), &gorm.Config{})
	require.NoError(s.T(), err)

	s.ctx = context.Background()
	s.publicService = MakeHumansService(db.MakeDBRepository(s.DB))
	s.privateService = &humanService{
		dbRepo: nil,
	}
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestIsMutant() {
	s.T().Run("ErrorValidatingInput", func(t *testing.T) {
		input := &humans.HumanRequest{}

		result, err := s.publicService.IsMutant(s.ctx, input)

		s.NotNil(err)
		s.EqualValues(http.StatusBadRequest, err.Status())
		s.EqualValues("dna is required", err.Error())
		s.EqualValues(false, result)
	})

	s.T().Run("ErrorFromGetHumanByDna", func(t *testing.T) {
		input := &humans.HumanRequest{
			Dna: []string{"AAAA", "GGGG", "CCCC", "TTTT"},
		}

		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "humans" WHERE dna = $1 AND "humans"."deleted_at" IS NULL ORDER BY "humans"."id" LIMIT 1`)).
			WithArgs("AAAA-GGGG-CCCC-TTTT").
			WillReturnError(gorm.ErrInvalidDB)

		result, err := s.publicService.IsMutant(s.ctx, input)

		s.NotNil(err)
		s.EqualValues(http.StatusInternalServerError, err.Status())
		s.EqualValues(errors_utils.ErrDatabase.Error(), err.Error())
		s.EqualValues(false, result)
	})
	s.T().Run("ErrorFromSaveHuman", func(t *testing.T) {
		input := &humans.HumanRequest{
			Dna: []string{"AAAA", "GGGG", "CCCC", "TTTT"},
		}
		human := &humans.Human{
			Dna:      "AAAA-GGGG-CCCC-TTTT",
			IsMutant: true,
		}

		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "humans" WHERE dna = $1 AND "humans"."deleted_at" IS NULL ORDER BY "humans"."id" LIMIT 1`)).
			WithArgs(human.Dna).
			WillReturnError(gorm.ErrRecordNotFound)

		s.mock.ExpectBegin()
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`INSERT INTO "humans" ("created_at","updated_at","deleted_at","dna","is_mutant") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), human.Dna, human.IsMutant).
			WillReturnError(gorm.ErrInvalidDB)
		s.mock.ExpectRollback()

		result, err := s.publicService.IsMutant(s.ctx, input)

		s.NotNil(err)
		s.EqualValues(http.StatusInternalServerError, err.Status())
		s.EqualValues(errors_utils.ErrDatabase.Error(), err.Error())
		s.EqualValues(false, result)
	})

	s.T().Run("Success", func(t *testing.T) {
		input := &humans.HumanRequest{
			Dna: []string{"AAAA", "GGGG", "CCCC", "TTTT"},
		}
		human := &humans.Human{
			Dna:      "AAAA-GGGG-CCCC-TTTT",
			IsMutant: true,
		}

		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "humans" WHERE dna = $1 AND "humans"."deleted_at" IS NULL ORDER BY "humans"."id" LIMIT 1`)).
			WithArgs(human.Dna).
			WillReturnError(gorm.ErrRecordNotFound)

		s.mock.ExpectBegin()
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`INSERT INTO "humans" ("created_at","updated_at","deleted_at","dna","is_mutant") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), human.Dna, human.IsMutant).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		s.mock.ExpectCommit()

		result, err := s.publicService.IsMutant(s.ctx, input)

		s.NoError(err)
		s.EqualValues(true, result)
	})
}

func (s *Suite) TestStats() {
	query := `SELECT count(*) FROM "humans" WHERE is_mutant = $1 AND "humans"."deleted_at" IS NULL`
	s.T().Run("ErrorFromCountHumans", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WillReturnError(gorm.ErrInvalidDB)

		result, err := s.publicService.Stats(s.ctx)

		s.Nil(result)
		s.NotNil(err)
		s.EqualValues(http.StatusInternalServerError, err.Status())
		s.EqualValues(errors_utils.ErrDatabase.Error(), err.Error())

	})

	s.T().Run("ErrorFromCountMutants", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(false).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).
				AddRow(5))

		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(true).
			WillReturnError(gorm.ErrInvalidDB)

		result, err := s.publicService.Stats(s.ctx)

		s.Nil(result)
		s.NotNil(err)
		s.EqualValues(http.StatusInternalServerError, err.Status())
		s.EqualValues(errors_utils.ErrDatabase.Error(), err.Error())

	})

	s.T().Run("SuccessHumanCounterIsZero", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(false).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).
				AddRow(0))

		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(true).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).
				AddRow(5))

		result, err := s.publicService.Stats(s.ctx)

		s.Nil(err)
		s.EqualValues(0, result.CountHumanDna)
		s.EqualValues(5, result.CountMutantDna)
		s.EqualValues(1.0, result.Ratio)
	})

	s.T().Run("Success", func(t *testing.T) {
		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(false).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).
				AddRow(100))

		s.mock.ExpectQuery(
			regexp.QuoteMeta(query)).
			WithArgs(true).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).
				AddRow(40))

		result, err := s.publicService.Stats(s.ctx)

		s.Nil(err)
		s.EqualValues(100, result.CountHumanDna)
		s.EqualValues(40, result.CountMutantDna)
		s.EqualValues(0.4, result.Ratio)
	})
}
