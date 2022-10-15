package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	// zap.logger实例
	logger *zap.Logger
}

func (l *zapLogger) DebugW(msg string, keyVals ...interface{}) {
	l.log(DebugLevel, msg, keyVals...)
}

func (l *zapLogger) InfoW(msg string, keyVals ...interface{}) {
	l.log(InfoLevel, msg, keyVals...)
}

func (l *zapLogger) WarnW(msg string, keyVals ...interface{}) {
	l.log(WarnLevel, msg, keyVals...)
}

func (l *zapLogger) ErrorW(msg string, keyVals ...interface{}) {
	l.log(ErrorLevel, msg, keyVals...)
}

func (l *zapLogger) FatalW(msg string, keyVals ...interface{}) {
	l.log(FatalLevel, msg, keyVals...)
}

func (l *zapLogger) log(level Level, msg string, keyVals ...interface{}) {
	if len(keyVals) == 0 || len(keyVals)%2 != 0 {
		l.logger.Warn("keyVals must appear in pairs")
		return
	}

	var data []zap.Field
	for i := 0; i < len(keyVals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyVals[i]), keyVals[i+1]))
	}

	switch level {
	case DebugLevel:
		l.logger.Debug(msg, data...)
	case InfoLevel:
		l.logger.Info(msg, data...)
	case WarnLevel:
		l.logger.Warn(msg, data...)
	case ErrorLevel:
		// TODO 发送错误告警
		l.logger.Error(msg, data...)
	case FatalLevel:
		// TODO 发送错误告警
		l.logger.Fatal(msg, data...)
	}
}

// NewZapLogger 实例化实现Logger接口的日志对象，path - 日志存储路径，为空则存入DefaultLogFileName内
func NewZapLogger(path string) Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	logPath := DefaultLogFileName
	if path != "" {
		logPath = path
	}

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
	if err != nil {
		panic(any(err))
	}
	ws := zapcore.NewMultiWriteSyncer(file)

	// TODO 根据环境判断哪些日志应该写入 example:生产写入info以上，测试则所有日志都写入
	logger := zap.New(zapcore.NewCore(encoder, ws, zapcore.DebugLevel), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &zapLogger{
		logger: logger,
	}
}
