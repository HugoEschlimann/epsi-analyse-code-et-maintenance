package logger

import (
    "sync"

    "github.com/sirupsen/logrus"
)

var (
    once   sync.Once
    logger *logrus.Logger
)

func GetLogger() *logrus.Logger {
    once.Do(func() {
        logger = logrus.New()
        logger.SetFormatter(&logrus.TextFormatter{
            FullTimestamp: true,
        })
        logger.SetLevel(logrus.InfoLevel)
    })
    return logger
}