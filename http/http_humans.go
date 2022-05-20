package http

import (
	"net/http"

	"github.com/getmiranda/meli-challenge-api/domain/humans"
	"github.com/getmiranda/meli-challenge-api/services"
	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type HumanHandler interface {
	IsMutant(c *gin.Context)
	Stats(c *gin.Context)
}

type humanHandler struct {
	humanService services.HumanService
}

// IsMutant checks if the human is a mutant.
func (h *humanHandler) IsMutant(c *gin.Context) {
	ctx := c.Request.Context()
	log := zerolog.Ctx(ctx)

	log.Info().Msg("Checking if human is mutant")

	var input humans.HumanRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		restErr := errors_utils.MakeBadRequestError("error binding JSON")
		c.JSON(restErr.Status(), restErr)
		return
	}

	isMutant, err := h.humanService.IsMutant(ctx, &input)
	if err != nil {
		log.Error().Err(err).Msg("Error from service")
		c.JSON(err.Status(), err)
		return
	}

	if !isMutant {
		c.Status(http.StatusForbidden)
		return
	}

	c.Status(http.StatusOK)
}

func (h *humanHandler) Stats(c *gin.Context) {
	ctx := c.Request.Context()
	log := zerolog.Ctx(ctx)

	log.Info().Msg("Getting stats")

	stats, err := h.humanService.Stats(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error from service")
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, stats)
}

// MakeHumanHandler returns a new HumanHandler.
func MakeHumanHandler(service services.HumanService) HumanHandler {
	return &humanHandler{
		humanService: service,
	}
}
