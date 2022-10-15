package log

import "testing"

func TestZapLogger(t *testing.T) {
	logger := NewZapLogger("")
	logger.DebugW("test", "age", 100, "name", "skyLe")
	logger.ErrorW("testErr", "age", 3, "name", "skyLe")
}