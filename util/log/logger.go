package log

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

var logger Logger

func init() {
	logger.log = logrus.New()
	logger.log.SetFormatter(&logrus.JSONFormatter{})
}

func GetLogger() *Logger {
	return &logger
}

func (logger *Logger) Info(ctx, msg string, meta interface{}) {
	logger.log.WithFields(logrus.Fields{
		"context": ctx,
		"meta":    meta,
	}).Info(msg)
}

func (logger *Logger) Warning(ctx, msg string, meta interface{}) {
	logger.log.WithFields(logrus.Fields{
		"context": ctx,
		"meta":    meta,
	}).Warning(msg)
}

func (logger *Logger) Error(ctx, msg string, meta interface{}) {
	logger.log.WithFields(logrus.Fields{
		"context": ctx,
		"meta":    meta,
	}).Error(msg)
}
