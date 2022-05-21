package services

import (
	"context"
	"sync"

	"github.com/getmiranda/meli-challenge-api/domain/humans"
	"github.com/getmiranda/meli-challenge-api/repository/db"
	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
	"github.com/getmiranda/meli-challenge-api/utils/mutant_utils"
	"github.com/getmiranda/meli-challenge-api/utils/types"
	"github.com/rs/zerolog"
)

type HumanService interface {
	// IsMutant returns true if the human is mutant.
	IsMutant(ctx context.Context, input *humans.HumanRequest) (bool, errors_utils.RestErr)
	// Stats returns the number of humans, mutants and ratio of mutants.
	Stats(ctx context.Context) (*humans.StatsResponse, errors_utils.RestErr)
}

type humanService struct {
	dbRepo db.DBRepository
	sync   sync.Mutex
}

// IsMutant returns true if the human is mutant.
func (s *humanService) IsMutant(ctx context.Context, input *humans.HumanRequest) (bool, errors_utils.RestErr) {
	s.sync.Lock()
	defer s.sync.Unlock()

	log := zerolog.Ctx(ctx)

	log.Info().Msg("Checking if human is mutant")

	if err := input.Validate(); err != nil {
		log.Error().Err(err).
			Interface("dna", input.Dna).Msg("Error validating input")
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

// Stats returns the number of humans, mutants and ratio of mutants.
func (s *humanService) Stats(ctx context.Context) (*humans.StatsResponse, errors_utils.RestErr) {
	log := zerolog.Ctx(ctx)

	log.Info().Msg("Getting stats")

	humanCounter, err := s.dbRepo.CountHumans(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting humans")
		return nil, err
	}

	mutantCounter, err := s.dbRepo.CountMutants(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error getting mutants")
		return nil, err
	}

	ratio := 1.0
	if humanCounter != 0 {
		ratio = float64(mutantCounter) / float64(humanCounter)
	}

	log.Info().
		Int64("humans", humanCounter).
		Int64("mutants", mutantCounter).
		Float64("ratio", ratio).
		Msg("Stats")

	return &humans.StatsResponse{
		CountMutantDna: mutantCounter,
		CountHumanDna:  humanCounter,
		Ratio:          types.Float(ratio),
	}, nil
}

// MakeHumansService returns a new instance of the HumanService.
func MakeHumansService(db db.DBRepository) HumanService {
	return &humanService{
		dbRepo: db,
		sync:   sync.Mutex{},
	}
}
