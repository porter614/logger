package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	STARTUP = iota
)

type LogConfig struct {
	// TODO
}

func Configure(c *LogConfig) {
	// TODO
}

var LogMessage = map[int]string{
	STARTUP: "Starting application...",
}

// Returns a new customized instance of a logrus.Logger
func Instance() *log.Entry {
	var logger = log.New()
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(log.WarnLevel)

	return logger.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})
}
