package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Ping(c *gin.Context) {
	log.Info().Msg("Ping")
	errorHandler(c, ping(), nil)
}

func ping() string {
	return "pong"
}
