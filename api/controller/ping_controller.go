package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	message = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus commodo mi non venenatis lacinia. Vivamus interdum ultricies dolor, eu sagittis mi faucibus et. Praesent sit amet sollicitudin velit. Etiam nisi urna, vehicula quis pulvinar vel, blandit ac risus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nam non dui non libero aliquam eleifend at at metus. Nullam pellentesque turpis nec risus egestas gravida. Curabitur at sagittis erat, nec facilisis velit. Fusce in eros non quam cursus commodo ut ut odio. Suspendisse rhoncus ultrices lectus, quis mattis quam vestibulum eget.

Aenean sed orci eget urna aliquet dapibus vitae quis leo. Nulla molestie, justo id suscipit gravida, ipsum lorem efficitur elit, quis iaculis orci nibh id nulla. Vestibulum cursus quam ac nibh placerat, non vulputate libero sodales. Vivamus mollis gravida augue, ac molestie dui tempor vitae. Pellentesque augue massa, dictum quis viverra quis, varius ac dui. Etiam pharetra.`
)

func Ping(c *gin.Context) {
	log.Info().Msg(message)
	errorHandler(c, ping(), nil)
}

func ping() string {
	return "pong"
}
