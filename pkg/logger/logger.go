package logger

import (
	"github.com/rs/zerolog"
)

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
