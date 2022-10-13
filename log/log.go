// log - 日志包

package log

type Logger interface {
	InfoW(msg string, keyVals ...interface{})
}
