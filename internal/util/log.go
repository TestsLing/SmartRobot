package util

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var SugarLogger *zap.SugaredLogger

func SetupLog() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
	//return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {

	hook, err := rotatelogs.New(
		"./runtime/test.log"+".%Y%m%d",
		rotatelogs.WithMaxAge(time.Duration(int64(24*time.Hour)*7)),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		Panic("构造 rotatelogs 对象失败:", err)
	}
	//return hook

	//lumberJackLogger := &lumberjack.Logger{
	//	Filename:   "./runtime/test.log",
	//	MaxSize:    10,
	//	MaxBackups: 5,
	//	MaxAge:     30,
	//	Compress:   false,
	//}
	return zapcore.AddSync(hook)
}

func Info(args ...interface{}) {
	SugarLogger.Info(args...)
}

func InfoF(template string, args ...interface{}) {
	SugarLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	SugarLogger.Warn(args...)
}

func Panic(args ...interface{}) {
	SugarLogger.Panic(args...)
}

func PanicF(template string, args ...interface{}) {
	SugarLogger.Panic(args...)
}

func WarnF(template string, args ...interface{}) {
	SugarLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	SugarLogger.Error(args...)
}

func ErrorF(template string, args ...interface{}) {
	SugarLogger.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	SugarLogger.Fatal(args...)
}
