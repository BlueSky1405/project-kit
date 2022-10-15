// log - 日志包

package log

// DefaultLogFileName 默认日志存储路径
const DefaultLogFileName = "./project.log"

// Logger 通用日志接口
type Logger interface {
	DebugW(msg string, keyVals ...interface{})
	InfoW(msg string, keyVals ...interface{})
	WarnW(msg string, keyVals ...interface{})
	ErrorW(msg string, keyVals ...interface{})
	FatalW(msg string, keyVals ...interface{})
}

type Option interface {
	apply(Logger)
}

type optionFunc func(logger Logger)

func (f optionFunc) apply(logger Logger) {
	f(logger)
}

type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)
