package main

import (
	_ "github.com/getmiranda/meli-challenge-api/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		ctx := c.Request.Context()
		log := zerolog.Ctx(ctx)

		log.Info().Msg("Pong")

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
