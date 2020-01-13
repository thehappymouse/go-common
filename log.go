package utils

import (
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"runtime"
)

// 适用于命令行输出观看
func ZeroConsoleLog() {
	zerolog.TimeFieldFormat = DataTimeMilli
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	sysType := runtime.GOOS

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: DataTimeMilli})

	if sysType == "windows" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: colorable.NewColorableStdout(), TimeFormat: DataTimeMilli})
	}
}

// 适用于命令行和日志文件都有打印
func ZeroConsoleAndFileLog(filename string) {
	zerolog.TimeFieldFormat = DataTimeMilli
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	sysType := runtime.GOOS

	var logFile *os.File
	var err error
	if !IsFileExists(filename) {
		logFile, err = os.Create(filename)
	} else {
		logFile, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	}
	CheckError(err)


	var consoleLog zerolog.ConsoleWriter = zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: DataTimeMilli}
	if sysType == "windows" {
		consoleLog = zerolog.ConsoleWriter{Out: colorable.NewColorableStdout(), TimeFormat: DataTimeMilli}
	}

	var writers []io.Writer
	writers = append(writers, logFile)
	writers = append(writers, consoleLog)
	mw := io.MultiWriter(writers...)

	log.Logger = zerolog.New(mw).With().Timestamp().Logger()
}
