package db

import (
	"context"
	"errors"

	"github.com/getmiranda/meli-challenge-api/domain/humans"
	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type DBRepository interface {
	// SaveHuman saves a human in the database.
	SaveHuman(context.Context, *humans.Human) errors_utils.RestErr
	// GetHumanByDna returns a human from the database.
	GetHumanByDna(context.Context, string) (*humans.Human, errors_utils.RestErr)
}

type dbRepository struct {
	DB *gorm.DB
}

func MakeDBRepository(db *gorm.DB) DBRepository {
	return &dbRepository{
		DB: db,
	}
}

// SaveHuman saves a human in the database.
func (r *dbRepository) SaveHuman(ctx context.Context, human *humans.Human) errors_utils.RestErr {
	logger := zerolog.Ctx(ctx)
	log := logger.With().Str("dna", human.Dna).Logger()

	log.Info().Msg("Saving human in database")

	if err := r.DB.Create(human).Error; err != nil {
		log.Error().Err(err).Msg("Error saving human in database")
		return errors_utils.MakeInternalServerError(errors_utils.ErrDatabase)
	}

	log.Info().Msg("Human saved in database")

	return nil
}

func (r *dbRepository) GetHumanByDna(ctx context.Context, dna string) (*humans.Human, errors_utils.RestErr) {
	logger := zerolog.Ctx(ctx)
	log := logger.With().Str("dna", dna).Logger()

	log.Info().Msg("Getting human from database")

	human := &humans.Human{}
	if err := r.DB.Where("dna = ?", dna).First(human).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info().Msg("Human not found in database")
			return nil, nil
		}
		log.Error().Err(err).Msg("Error getting human from database")
		return nil, errors_utils.MakeInternalServerError(errors_utils.ErrDatabase)
	}

	log.Info().Msg("Human found in database")

	return human, nil
}
