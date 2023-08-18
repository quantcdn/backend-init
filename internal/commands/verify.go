package commands

import (
	"errors"

	"github.com/quantcdn/backend-init/internal/backend"
	"github.com/rs/zerolog/log"
)

type Verify struct {
	Url     string `arg:"" help:"Backend URL to verify HTTP connection to"`
	Delay   int    `help:"Delay between requests in seconds" env:"BACKEND_DELAY" default:"5"`
	Retries int    `help:"Number of times to retry before failure" env:"BACKEND_RETRIES" default:"10"`
}

func (v *Verify) Run() error {
	log.Info().Str("url", v.Url).Int("delay", v.Delay).Int("retries", v.Retries).Msg("attempting to connect")
	if backend.Connect(v.Url, v.Delay, v.Retries) {
		log.Info().Str("url", v.Url).Msg("Successfully connected.")
		return nil
	}
	return errors.New("Unable to connect.")
}
