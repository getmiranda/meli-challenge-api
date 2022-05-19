package http

import (
	"context"
	"testing"

	_ "github.com/getmiranda/meli-challenge-api/logger"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	ctx context.Context

	pingHandler PingHandler
}

func (s *Suite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	s.ctx = context.Background()

	s.pingHandler = MakePingHandler()
}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
