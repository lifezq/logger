package logger

import "testing"

func TestErrorLogger(t *testing.T) {
	ErrorPrintf("ERROR:%s", "something is wrong...")
}

func BenchmarkErrorLogger(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ErrorPrintf("ERROR:%s", "something is wrong...")
	}
}

func TestInfoLogger(t *testing.T) {
	InfoPrintf("INFO:%s", "pending...")
}

func BenchmarkInfoLogger(t *testing.B) {
	for i := 0; i < t.N; i++ {
		InfoPrintf("INFO:%s", "pending...")
	}
}

func TestWarnLogger(t *testing.T) {
	WarnPrintf("Warn:%s", "something is wrong...")
}

func BenchmarkWarnLogger(t *testing.B) {
	for i := 0; i < t.N; i++ {
		WarnPrintf("Warn:%s", "something is wrong...")
	}
}
