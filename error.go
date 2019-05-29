package utils

import (
	"github.com/rs/zerolog/log"
)
// 带退出的Log方法
func CheckError(err error) {
	if err != nil {
		log.Fatal().Caller().Msgf("%s", err)
	}
}


