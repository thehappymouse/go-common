package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"runtime"
	"github.com/mattn/go-colorable"
)

// 适用于命令行输出观看
func ZeroConsoleLog() {
	zerolog.TimeFieldFormat = DataTimeMilli
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	sysType := runtime.GOOS

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: DataTimeMilli})

	if sysType == "windows" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: colorable.NewColorableStdout(), TimeFormat:DataTimeMilli})
	}
}
