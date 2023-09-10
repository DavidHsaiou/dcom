package util

import (
	"testing"
)

func TestNewZapLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TestNewZapLogger",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewZapLogger()
			logger.Info("Info")
			logger.Infof("Infof")
			logger.Debug("Debug")
			logger.Debugf("Debugf")
			logger.Warn("Warn")
			logger.Warnf("Warnf")
			logger.Error("Error")
			logger.Errorf("Errorf")
		})
	}
}
