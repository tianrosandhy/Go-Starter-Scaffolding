package bootstrap

import (
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tianrosandhy/goconfigloader"
	"gopkg.in/natefinch/lumberjack.v2"
)

var LOGRUS_LEVEL_MAP = map[string]logrus.Level{
	"panic":   logrus.PanicLevel,
	"fatal":   logrus.FatalLevel,
	"error":   logrus.ErrorLevel,
	"warning": logrus.WarnLevel,
	"info":    logrus.InfoLevel,
	"debug":   logrus.DebugLevel,
	"trace":   logrus.TraceLevel,
}

var FALLBACK_LOGRUS_LEVEL = "info"

func NewLogger(cfg *goconfigloader.Config) *logrus.Logger {
	log := logrus.New()

	level := cfg.GetString("LOG_LEVEL")
	if _, ok := LOGRUS_LEVEL_MAP[level]; !ok {
		level = FALLBACK_LOGRUS_LEVEL
	}

	log.SetLevel(logrus.Level(LOGRUS_LEVEL_MAP[level]))
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)

	logFilename := cfg.GetString("LOG_PATH")

	if strings.Trim(logFilename, " \t\\/") != "" {
		lumberjackLogger := lumberjack.Logger{
			Filename:   logFilename,
			MaxSize:    cfg.GetInt("LOG_MAX_SIZE"),
			MaxBackups: 8,
			MaxAge:     60,
			Compress:   true,
		}
		log.SetOutput(io.MultiWriter(os.Stdout, &lumberjackLogger))
	}

	return log
}
