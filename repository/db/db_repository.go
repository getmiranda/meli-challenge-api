package db

import (
	"context"

	"github.com/getmiranda/meli-challenge-api/domain/humans"
	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type DBRepository interface {
	// SaveHuman saves a human in the database.
	SaveHuman(context.Context, *humans.Human) errors_utils.RestErr
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
