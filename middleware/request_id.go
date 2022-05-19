package middleware

import (
	"context"

	"github.com/getmiranda/meli-challenge-api/utils/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Config struct {
	EnabledInRequestContext bool
	EnabledInRequestHeader  bool
	EnabledInResponseHeader bool
	EnabledInZerologContext bool
}

func WithRequestId(config Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		reqId := uuid.New().String()

		if config.EnabledInRequestContext {
			ctx = context.WithValue(ctx, types.XRequestId, reqId)
			c.Request = c.Request.WithContext(ctx)
		}
		if config.EnabledInRequestHeader {
			c.Request.Header.Set(string(types.XRequestId), reqId)
		}
		if config.EnabledInResponseHeader {
			c.Header(string(types.XRequestId), reqId)
		}
		if config.EnabledInZerologContext {
			logger := zerolog.Ctx(ctx)
			newLogger := logger.With().Str("request_id", reqId).Logger()
			ctx = newLogger.WithContext(ctx)
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}
