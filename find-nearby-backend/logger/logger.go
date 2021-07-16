package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Fields represents key-value pairs and can be used to
// provide additional context in logs
type Fields map[string]interface{}

// Logger represents a generic logging component
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	ErrorWithTag(err error, fields Fields)
}

type logger struct {
	logrus *logrus.Logger
}

type Error struct {
	Error error
}

// Setup the logger with appropriate log-level and format.
// Valid log-levels are debug, info, warn, error, fatal, panic.
// Valid log-formats are plain or json (default: json)
func New(logLevel, logFormat string) Logger {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.WarnLevel
	}

	formatter := &logrus.JSONFormatter{
		DataKey: "xcontext",
	}

	log := &logger{logrus: &logrus.Logger{
		Out:       os.Stderr,
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
		Formatter: formatter,
	}}

	if logFormat != "json" {
		log.logrus.Formatter = &logrus.TextFormatter{}
	}

	return log
}

func (log *logger) ErrorWithTag(err error, fields Fields) {
	if err != nil {
		log.logrus.WithFields(logrus.Fields(fields)).Error(err.Error())
	}
}

func (log *logger) Debugf(format string, args ...interface{}) {
	log.logrus.Debugf(format, args...)
}

func (log *logger) Infof(format string, args ...interface{}) {
	log.logrus.Infof(format, args...)
}

func (log *logger) Warnf(format string, args ...interface{}) {
	log.logrus.Warnf(format, args...)
}

func (log *logger) Errorf(format string, args ...interface{}) {
	log.logrus.Errorf(format, args...)
}

func (log *logger) Fatalf(format string, args ...interface{}) {
	log.logrus.Fatalf(format, args...)
}

func (log *logger) Panicf(format string, args ...interface{}) {
	log.logrus.Panicf(format, args...)
}
