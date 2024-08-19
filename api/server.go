package api

import (
	"os"
	"template/api/api/config"
	"template/api/api/controller"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Start() {
	log.Info().Msg("initialize the server")
	scope := os.Getenv("SCOPE")
	log.Info().
		Msgf("set scope: %s", scope)

	basepathFile := os.Getenv("BASEPATH")

	appConfig, err := config.Get(basepathFile, scope)
	if err != nil {
		exitSystem("error initializing the server", err)
	}

	/// -----------------   refactorizar este codigo  -----------------------------------
	// Configure Logs
	logLevel, err := zerolog.ParseLevel(appConfig.Log.Lvl)
	if err != nil {
		exitSystem("fail log level configuration", err)
	}

	log.Debug().Msgf("log level %s", logLevel)

	zerolog.SetGlobalLevel(logLevel)

	file, err := os.OpenFile(
		os.Getenv("LOGFILE"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		exitSystem("fail to get configuration", err)
	}
	defer file.Close()
	multi := zerolog.MultiLevelWriter(file, os.Stdout)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
	log.Debug().Msg("starting server")

	// create server
	r := gin.Default()

	// Set router
	setRouter(r)

	// run server
	r.Run(":8080")
}

func setRouter(r *gin.Engine) {
	log.Print("set router")
	// ping
	r.GET("/ping", controller.Ping)
}

func exitSystem(msg string, err error) {
	log.Error().
		Err(err).
		Msg(msg)

	os.Exit(0)
}
