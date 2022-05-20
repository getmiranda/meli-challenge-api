package mutants

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getmiranda/meli-challenge-api/domain"
	_ "github.com/getmiranda/meli-challenge-api/logger"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	mutantsHost     = "DB_MUTANTS_HOST"
	mutantsUser     = "DB_MUTANTS_USER"
	mutantsPassword = "DB_MUTANTS_PASSWORD"
	mutantsDbName   = "DB_MUTANTS_DBNAME"
	mutantsPort     = "DB_MUTANTS_PORT"
)

var (
	client *gorm.DB

	host     = os.Getenv(mutantsHost)
	user     = os.Getenv(mutantsUser)
	password = os.Getenv(mutantsPassword)
	dbName   = os.Getenv(mutantsDbName)
	port     = os.Getenv(mutantsPort)
)

func init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbName, port,
	)

	logger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             499 * time.Millisecond,
		LogLevel:                  logger.Silent,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})

	var err error
	client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		panic(err)
	}

	log := zerolog.Ctx(context.Background())
	log.Info().Msg("Database configured successfully")
}

// Migrate all tables in the database.
func Migrate() error {
	return client.AutoMigrate(domain.Migrate()...)
}

// GetClient returns the database client.
func GetClient() *gorm.DB {
	return client
}
