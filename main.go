package main

import (
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/quantcdn/backend-init/internal/backend"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	app   = kingpin.New("backend-init", "Docker entrypoint to validate containers.")
	url   = kingpin.Flag("url", "The backend url.").Envar("BACKEND_URL").String()
	retry = kingpin.Flag("retry", "Times to retry the backend connection.").Default("10").Envar("BACKEND_RETRY").Int()
	delay = kingpin.Flag("delay", "Delay between backend requests.").Default("5").Envar("BACKEND_DELAY").Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("validating container start")

	if backend.Connect(*url, *delay, *retry) {
		log.Info().Msg("successfully connected to the backend")
		os.Exit(0)
	}
	log.Info().Msg("unable to connect to backend")
	os.Exit(1)
}
