package app

import (
	"github.com/getmiranda/meli-challenge-api/logger"
	"github.com/getmiranda/meli-challenge-api/middleware"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func StartApplication() {
	router = gin.Default()

	log := logger.GetLogger()

	log.Info().Msg("Starting application")

	router.Use(
		middleware.WithRequestId(middleware.Config{
			EnabledInRequestContext: true,
			EnabledInRequestHeader:  true,
			EnabledInResponseHeader: true,
			EnabledInZerologContext: true,
		}),
	)

	mapUrls()

	router.Run()
}
