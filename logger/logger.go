package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	//"golang.org/x/text/message"
)

var log *zap.Logger

func init () {
	var err error

	config := zap.NewProductionConfig()

	encondingConfig := zap.NewProductionEncoderConfig()
	encondingConfig.TimeKey = "timestamp"
	encondingConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encondingConfig.StacktraceKey = ""
	config.EncoderConfig = encondingConfig

	log, err = config.Build(zap.AddCallerSkip(1))

	//log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}


func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}


func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}