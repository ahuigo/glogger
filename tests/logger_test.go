package tests

import (
	"testing"

	"github.com/ahuigo/glogger"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	glogger.GetLogger("proj", zap.InfoLevel).Info("This is a info log")
	glogger.GetLogger("proj", zap.InfoLevel).Debug("This is a debug log")
	glogger.Glogger.Debug("This is a debug log")
}
