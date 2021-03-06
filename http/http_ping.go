package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type PingHandler interface {
	Ping(*gin.Context)
}

type pingHandler struct{}

// Ping checks if the service is available.
func (h *pingHandler) Ping(c *gin.Context) {
	ctx := c.Request.Context()
	logger := zerolog.Ctx(ctx)

	logger.Info().Msg("Pong")

	c.String(http.StatusOK, "pong")
}

// MakePingHandler returns a new PingHandler.
func MakePingHandler() PingHandler {
	return &pingHandler{}
}
