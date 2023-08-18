package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/quantcdn/backend-init/internal/commands"
	"github.com/rs/zerolog/log"
)

var cli struct {
	Verify commands.Verify `cmd:"" help:"Verify a backend HTTP connection to the given URL."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	if err != nil {
		log.Error().AnErr("error", err)
		os.Exit(1)
	}
	os.Exit(0)
}
