package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogger(warnLevel log.Level) *log.Logger {

	var logger = &log.Logger{
		Out:       os.Stderr,
		Formatter: new(log.TextFormatter),
		Hooks:     make(log.LevelHooks),
		Level:     warnLevel,
	}

	return logger
}
