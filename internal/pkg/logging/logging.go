package logging

// https://www.datadoghq.com/blog/go-logging/

import (
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
  id      int
  message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
  *logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	unexpectedError = Event{1, "Unexpected error: %s"}
)

// UnexpectedError is a standard error message
func(logger *StandardLogger) UnexpectedError(err error) {
	logger.Errorf(unexpectedError.message, err)
}