package logger

import "testing"

func TestNewJSON(t *testing.T) {
	logger := NewJSON()
	if logger == nil {
		t.Fatal("logger should not be nil")
	}
}
