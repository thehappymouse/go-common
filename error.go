package utils

import (
	"github.com/rs/zerolog/log"
	"runtime/debug"
)
// 带退出的Log方法
func CheckError(err error) {
	if err != nil {
		debug.PrintStack()
		log.Fatal().Msgf("%s", err)
	}
}


