package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func ZeroConsoleLog() {
	zerolog.TimeFieldFormat = DataTimeMilli
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: DataTimeMilli})

}
