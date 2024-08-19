package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func errorHandler(c *gin.Context, data any, err error) {
	l := log.Ctx(c.Request.Context())
	if err != nil {
		l.Error().
			Err(err).
			Msg("error handler")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		l.Info().Msg("finish ok")
		c.JSON(http.StatusOK, data)
	}
}
