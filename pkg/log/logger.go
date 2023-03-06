package log

import (
	"os"
	"sync"

	"go.uber.org/zap"
)

var once sync.Once
var logger *zap.SugaredLogger

func generateLogger() {
	level := os.Getenv("LOG_LEVEL")
	var l *zap.Logger

	if level == "dev" {
		l, _ = zap.NewDevelopment()
	} else {
		l, _ = zap.NewProduction()
	}
	logger = l.Sugar()
}

func Logger() *zap.SugaredLogger {
	once.Do(generateLogger)
	return logger
}
