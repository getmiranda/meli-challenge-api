package http

import (
	"net/http"
	"net/http/httptest"

	"github.com/getmiranda/meli-challenge-api/utils/test_utils"
)

func (s *Suite) TestPing() {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	c := test_utils.GetMockedContext(request, response)

	s.pingHandler.Ping(c)

	s.EqualValues(http.StatusOK, response.Code)
	s.EqualValues("pong", response.Body.String())
}
