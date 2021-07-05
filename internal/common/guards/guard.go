package guards

import "github.com/kuritka/doge-action/internal/common/log"

var logger = log.Log

func Must(err error, msg string, v ...interface{}) {
	if err != nil {
		logger.Fatal().Err(err).Msgf(msg, v...)
	}
}
