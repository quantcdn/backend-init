package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/quantcdn/backend-init/internal/commands"
	"github.com/rs/zerolog/log"
)

var cli struct {
	Verify    commands.Verify `cmd:"" help:"Verify a backend HTTP connection to the given URL."`
	SkipFirst bool            `help:"First run always returns success" default:"false" short:"s"`
	Lock      string          `help:"Lock file location" default:"/data"`
}

func main() {
	ctx := kong.Parse(&cli)
	lock := fmt.Sprintf("%s/init.lock", cli.Lock)
	_, err := os.Stat(lock)

	if cli.SkipFirst && err != nil {
		os.Create(lock)
		log.Info().Msg("Skipping first run,")
		os.Exit(0)
	}

	err = ctx.Run()

	if err != nil {
		log.Error().AnErr("error", err)
		os.Exit(1)
	}

	os.Exit(0)
}
