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
	// CountMutants returns the number of mutants in the database.
	CountMutants(context.Context) (int64, errors_utils.RestErr)
	// CountHumans returns the number of humans in the database.
	CountHumans(context.Context) (int64, errors_utils.RestErr)
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

// GetHumanByDna returns a human from the database by dna.
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

// CountMutants returns the number of mutants in the database.
func (r *dbRepository) CountMutants(ctx context.Context) (int64, errors_utils.RestErr) {
	log := zerolog.Ctx(ctx)

	log.Info().Msg("Getting mutants count from database")

	var count int64
	if err := r.DB.Model(&humans.Human{}).Where("is_mutant = ?", true).Count(&count).Error; err != nil {
		log.Error().Err(err).Msg("Error getting mutants count from database")
		return 0, errors_utils.MakeInternalServerError(errors_utils.ErrDatabase)
	}

	log.Info().Int64("mutants_count", count).Msg("Mutants count found in database")

	return count, nil
}

// CountHumans returns the number of humans in the database.
func (r *dbRepository) CountHumans(ctx context.Context) (int64, errors_utils.RestErr) {
	log := zerolog.Ctx(ctx)

	log.Info().Msg("Getting humans count from database")

	var count int64
	if err := r.DB.Model(&humans.Human{}).Where("is_mutant = ?", false).Count(&count).Error; err != nil {
		log.Error().Err(err).Msg("Error getting humans count from database")
		return 0, errors_utils.MakeInternalServerError(errors_utils.ErrDatabase)
	}

	log.Info().Int64("humans_count", count).Msg("Humans count found in database")

	return count, nil
}
