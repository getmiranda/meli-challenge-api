package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getmiranda/meli-challenge-api/domain/humans"
	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
	"github.com/getmiranda/meli-challenge-api/utils/test_utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Suite) TestIsMutantErrorBindingJson() {
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/mutant/", strings.NewReader(``))
	c := test_utils.GetMockedContext(request, response)

	s.humanHandler.IsMutant(c)

	s.EqualValues(http.StatusBadRequest, response.Code)
	apiErr, err := errors_utils.MakeErrorFromBytes(response.Body.Bytes())
	s.Nil(err)
	s.NotNil(apiErr)
	s.EqualValues(http.StatusBadRequest, apiErr.Status())
	s.EqualValues("error binding JSON", apiErr.Error())
}

func (s *Suite) TestIsMutantErrorFromService() {
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/mutant/", strings.NewReader(`{"dna": ["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTGT"]}`))
	c := test_utils.GetMockedContext(request, response)

	s.humanHandler.IsMutant(c)

	s.EqualValues(http.StatusBadRequest, response.Code)
	apiErr, err := errors_utils.MakeErrorFromBytes(response.Body.Bytes())
	s.Nil(err)
	s.NotNil(apiErr)
	s.EqualValues(http.StatusBadRequest, apiErr.Status())
	s.EqualValues("dna must be a square matrix", apiErr.Error())
}

func (s *Suite) TestIsMutantNoErrorIsMutant() {
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/mutant/", strings.NewReader(`{"dna": ["AAAA", "AAAA", "AAAA", "AAAA"]}`))
	c := test_utils.GetMockedContext(request, response)

	human := &humans.Human{
		Dna:      "AAAA-AAAA-AAAA-AAAA",
		IsMutant: true,
	}

	s.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "humans" WHERE dna = $1 AND "humans"."deleted_at" IS NULL ORDER BY "humans"."id" LIMIT 1`)).
		WithArgs(human.Dna).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "dna", "is_mutant"}).
			AddRow(1, human.CreatedAt, human.UpdatedAt, human.DeletedAt, human.Dna, human.IsMutant))

	s.humanHandler.IsMutant(c)

	s.EqualValues(http.StatusOK, c.Writer.Status())
}

func (s *Suite) TestIsMutantNoErrorIsHuman() {
	gin.SetMode(gin.TestMode)
	req := humans.HumanRequest{
		Dna: []string{"ATTA", "CGGC", "GAGG", "TATT"},
	}
	requestBytes, _ := json.Marshal(req)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/mutant/", bytes.NewReader(requestBytes))
	c := test_utils.GetMockedContext(request, response)

	human := &humans.Human{
		Dna:      "ATTA-CGGC-GAGG-TATT",
		IsMutant: false,
	}

	s.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "humans" WHERE dna = $1 AND "humans"."deleted_at" IS NULL ORDER BY "humans"."id" LIMIT 1`)).
		WithArgs(human.Dna).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "dna", "is_mutant"}).
			AddRow(1, human.CreatedAt, human.UpdatedAt, human.DeletedAt, human.Dna, human.IsMutant))

	s.humanHandler.IsMutant(c)

	log.Println(response)

	s.EqualValues(http.StatusForbidden, c.Writer.Status())
}

func (s *Suite) TestStatsErrorFromService() {
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/stats/", nil)
	c := test_utils.GetMockedContext(request, response)

	s.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT count(*) FROM "humans" WHERE is_mutant = $1 AND "humans"."deleted_at" IS NULL`)).
		WillReturnError(gorm.ErrInvalidDB)

	s.humanHandler.Stats(c)

	s.EqualValues(http.StatusInternalServerError, response.Code)
	apiErr, err := errors_utils.MakeErrorFromBytes(response.Body.Bytes())
	s.Nil(err)
	s.NotNil(apiErr)
	s.EqualValues(http.StatusInternalServerError, apiErr.Status())
	s.EqualValues("database error", apiErr.Error())
}

func (s *Suite) TestStatsNoError() {
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/stats/", nil)
	c := test_utils.GetMockedContext(request, response)

	s.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT count(*) FROM "humans" WHERE is_mutant = $1 AND "humans"."deleted_at" IS NULL`)).
		WithArgs(false).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(100))

	s.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT count(*) FROM "humans" WHERE is_mutant = $1 AND "humans"."deleted_at" IS NULL`)).
		WithArgs(true).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(40))

	s.humanHandler.Stats(c)

	s.EqualValues(http.StatusOK, response.Code)
	stats := &humans.StatsResponse{}
	err := json.Unmarshal(response.Body.Bytes(), stats)
	s.Nil(err)
	s.EqualValues(100, stats.CountHumanDna)
	s.EqualValues(40, stats.CountMutantDna)
	s.EqualValues(0.4, stats.Ratio)
}
