package log

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLevel(t *testing.T) {
	testCases := []struct {
		logLevel      string
		ExpectedLevel Level
	}{
		{
			logLevel:      "log",
			ExpectedLevel: InfoLevel,
		},
		{
			logLevel:      "debug",
			ExpectedLevel: DebugLevel,
		},
		{
			logLevel:      "DeBug",
			ExpectedLevel: DebugLevel,
		},
		{
			logLevel:      "info",
			ExpectedLevel: InfoLevel,
		},
		{
			logLevel:      "warn",
			ExpectedLevel: WarnLevel,
		},
		{
			logLevel:      "error",
			ExpectedLevel: ErrorLevel,
		},
		{
			logLevel:      "fatal",
			ExpectedLevel: FatalLevel,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			l := ParseLevel(tc.logLevel)
			if l != tc.ExpectedLevel {
				t.Fatalf("%d: Expected %v but got %v", i+1, tc.ExpectedLevel, l)
			}
		})
	}
}

func TestLogger_ResetLevel(t *testing.T) {
	buf := CloseBuffer{
		Buffer: &bytes.Buffer{},
	}
	l := NewLogger(buf, WarnLevel)

	assert.Equal(t, l.level, WarnLevel)
	assert.NotEqual(t, l.level, FatalLevel)

	l.ResetLevel(FatalLevel)

	assert.Equal(t, l.level, FatalLevel)
	assert.NotEqual(t, l.level, WarnLevel)
}

func TestNewFileLogger(t *testing.T) {
	logger := NewFileLogger("./log.log", InfoLevel)
	defer os.Remove("./log.log")

	assert.Equal(t, logger.level, InfoLevel)
	assert.NotEqual(t, logger.level, FatalLevel)

	logger.Fatal("aaa")
}

func TestLogger(t *testing.T) {
	buf := CloseBuffer{
		Buffer: &bytes.Buffer{},
	}
	l := NewLogger(buf, DebugLevel)

	l.Info("aaaaa")
	l.Warn("bbbbb")
	l.Error("ccccc")
	l.Debug("ccccc")
	l.Fatal("ddddd")
	l.Println("hehe")
	s := buf.String()
	assert.Contains(t, s, "[INFO]")
	assert.Contains(t, s, "[WARN]")
	assert.Contains(t, s, "[ERROR]")
	assert.Contains(t, s, "[DEBUG]")
	assert.Contains(t, s, "[FATAL]")
	assert.Contains(t, s, "aaaaa")
	assert.Contains(t, s, "bbbbb")
	assert.Contains(t, s, "ccccc")
	assert.Contains(t, s, "ddddd")
	assert.Contains(t, s, "hehe")
	// request ID
	ll := l.NewWithTaskName("request-id")
	ll.Info("haha")
	s = buf.String()
	assert.Contains(t, s, "haha")
	assert.Contains(t, s, "request-id")
	err := l.Close()
	if err != nil {
		t.Fatal("logger close err:", err)
	}
}

func TestLogger_Log(t *testing.T) {
	buf := CloseBuffer{
		Buffer: &bytes.Buffer{},
	}
	l := NewLogger(buf, DebugLevel)

	l.Log(DebugLevel, 3, "aaa")
	l.Log(InfoLevel, 3, "bbb")
	l.Log(WarnLevel, 3, "ccc")
	l.Log(ErrorLevel, 3, "ddd")
	l.Log(FatalLevel, 3, "eee")

	s := buf.String()
	assert.Contains(t, s, "[INFO]")
	assert.Contains(t, s, "[WARN]")
	assert.Contains(t, s, "[ERROR]")
	assert.Contains(t, s, "[DEBUG]")
	assert.Contains(t, s, "[FATAL]")
	assert.Contains(t, s, "aaa")
	assert.Contains(t, s, "bbb")
	assert.Contains(t, s, "ccc")
	assert.Contains(t, s, "eee")

}

func TestLogLevel(t *testing.T) {
	errBuf := CloseBuffer{
		Buffer: &bytes.Buffer{},
	}
	errLogger := NewLogger(errBuf, FatalLevel)
	errLogger.Info("aaa")
	errLogger.Warn("bbb")
	errLogger.Error("ccc")
	errLogger.Fatal("ddd")
	errLogger.Debug("eee")
	errString := errBuf.String()
	assert.NotContains(t, errString, "[INFO]")
	assert.NotContains(t, errString, "aaa")
	assert.NotContains(t, errString, "[WARN]")
	assert.NotContains(t, errString, "bbb")
	assert.NotContains(t, errString, "[DEBUG]")
	assert.NotContains(t, errString, "eee")
	assert.NotContains(t, errString, "[ERROR]")
	assert.NotContains(t, errString, "ccc")
	assert.Contains(t, errString, "[FATAL]")
	assert.Contains(t, errString, "ddd")

	warnBuf := CloseBuffer{
		Buffer: &bytes.Buffer{},
	}
	warnLogger := NewLogger(warnBuf, WarnLevel)
	warnLogger.Info("aaa")
	warnLogger.Warn("bbb")
	warnLogger.Error("ccc")
	warnString := warnBuf.String()
	assert.NotContains(t, warnString, "[INFO]")
	assert.NotContains(t, warnString, "aaa")
	assert.Contains(t, warnString, "[WARN]")
	assert.Contains(t, warnString, "bbb")
	assert.Contains(t, warnString, "[ERROR]")
	assert.Contains(t, warnString, "ccc")
}
