package services

import (
	"context"

	"github.com/getmiranda/meli-challenge-api/domain/humans"
	"github.com/getmiranda/meli-challenge-api/repository/db"
	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
)

type HumanService interface {
}

type humanService struct {
	dbRepo db.DBRepository
}

func (s *humanService) IsMutant(ctx context.Context, input *humans.HumanRequest) errors_utils.RestErr {
	if err := input.Validate(); err != nil {
		return err
	}
	return nil
}

func MakeHumansService(db db.DBRepository) HumanService {
	return &humanService{
		dbRepo: db,
	}
}
