package log

import "learnyscape-backend-mono/pkg/logger"

var Logger logger.Logger

func SetLogger(logger logger.Logger) {
	Logger = logger
}
