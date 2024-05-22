package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"multifinance-go/internal/consts"
)

func NewLogger(viper *viper.Viper) *logrus.Logger {
	log := logrus.New()

	level := logrus.Level(viper.GetInt32("log.level"))
	log.SetLevel(level)

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,                        // Show full timestamps
		TimestampFormat: consts.LayoutDateTimeFormat, // Custom timestamp format
		ForceColors:     true,                        // Force colored output
		DisableColors:   false,                       // Enable colored output
		PadLevelText:    true,                        // Pad level text for alignment
	}
	log.SetFormatter(formatter)

	log.WithField("level", level.String()).Info("Logger initialized")

	return log
}
