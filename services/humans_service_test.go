package services

import (
	"context"
	"testing"

	_ "github.com/getmiranda/meli-challenge-api/logger"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	ctx context.Context

	publicService  HumanService
	privateService *humanService
}

func (s *Suite) SetupSuite() {
	s.ctx = context.Background()

	s.publicService = MakeHumansService(nil)

	s.privateService = &humanService{
		dbRepo: nil,
	}
}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
