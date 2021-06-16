package log

import (
	"bytes"
	"testing"
)

func withBenchedLogger(b *testing.B, logLevel Level, f func(Logger)) {
	errBuf := CloseBuffer{
		Buffer: &bytes.Buffer{},
	}
	logger := NewLogger(errBuf, logLevel)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			f(logger)
		}
	})
}

func BenchmarkLogger_Fatal(b *testing.B) {
	withBenchedLogger(b, FatalLevel, func(logger Logger) {
		logger.Info("aaa")
		logger.Warn("bbb")
		logger.Error("ccc")
		logger.Debug("ddd")
		logger.Fatal("eee")
	})
}

func BenchmarkLogger_Error(b *testing.B) {
	withBenchedLogger(b, ErrorLevel, func(logger Logger) {
		logger.Info("aaa")
		logger.Warn("bbb")
		logger.Error("ccc")
		logger.Debug("ddd")
		logger.Fatal("eee")
	})
}

func BenchmarkLogger_Warn(b *testing.B) {
	withBenchedLogger(b, WarnLevel, func(logger Logger) {
		logger.Info("aaa")
		logger.Warn("bbb")
		logger.Error("ccc")
		logger.Debug("ddd")
		logger.Fatal("eee")
	})
}

func BenchmarkLogger_Info(b *testing.B) {
	withBenchedLogger(b, InfoLevel, func(logger Logger) {
		logger.Info("aaa")
		logger.Warn("bbb")
		logger.Error("ccc")
		logger.Debug("ddd")
		logger.Fatal("eee")
	})
}

func BenchmarkLogger_Debug(b *testing.B) {
	withBenchedLogger(b, DebugLevel, func(logger Logger) {
		logger.Info("aaa")
		logger.Warn("bbb")
		logger.Error("ccc")
		logger.Debug("ddd")
		logger.Fatal("eee")
	})
}
