package tests

import (
	"testing"

	"github.com/ahuigo/glogger"
)

func TestLogger(t *testing.T) {
	glogger.Glogger.Info("https://httpbin.org/json")
}
