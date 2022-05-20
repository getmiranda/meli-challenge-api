package services

import (
	"context"

	"github.com/getmiranda/meli-challenge-api/domain/humans"
	"github.com/getmiranda/meli-challenge-api/repository/db"
	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
	"github.com/getmiranda/meli-challenge-api/utils/mutant_utils"
	"github.com/rs/zerolog"
)

type HumanService interface {
	// IsMutant returns true if the human is mutant.
	IsMutant(ctx context.Context, input *humans.HumanRequest) (bool, errors_utils.RestErr)
}

type humanService struct {
	dbRepo db.DBRepository
}

func (s *humanService) IsMutant(ctx context.Context, input *humans.HumanRequest) (bool, errors_utils.RestErr) {
	log := zerolog.Ctx(ctx)

	log.Info().Msg("Checking if human is mutant")

	if err := input.Validate(); err != nil {
		log.Error().Err(err).Msg("Error validating input")
		return false, err
	}

	dna := input.GenerateDna()

	// Check if the human is already in the database.
	human, err := s.dbRepo.GetHumanByDna(ctx, dna)
	if err != nil {
		log.Error().Err(err).Msg("Error getting human by dna")
		return false, err
	}

	// If the human is not in the database, we need to check if it is mutant.
	if human == nil {
		result := mutant_utils.IsMutant(input.Dna)
		// Save the human in the database.
		human = &humans.Human{
			Dna:      dna,
			IsMutant: result.IsMutant,
		}
		if err := s.dbRepo.SaveHuman(ctx, human); err != nil {
			log.Error().Err(err).Msg("Error saving human")
			return false, err
		}
	}

	log.Info().Bool("is_mutant", human.IsMutant).Msg("Alredy checked")

	return human.IsMutant, nil
}

func MakeHumansService(db db.DBRepository) HumanService {
	return &humanService{
		dbRepo: db,
	}
}
