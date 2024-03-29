package tests

import (
	"os"
	"testing"

	"fmt"

	"github.com/ahuigo/glogger"
	"go.uber.org/zap"
)

type User struct {
	Name string
	Age  int
}

func TestLogger(t *testing.T) {
	fmt.Println("default log format:")
	fmt.Println("datetime                	log_level project_name	code_path:line	<message>")
	// this log is named root
	glogger.Debug("This is a debug log")

	// get logger named with "proj"
	logger := glogger.GetLogger("proj", zap.InfoLevel)
	logger.Info("This is a info log")
	logger.Info("This is a info log")
	logger.Debug("This is a debug log") //not printed

	user := User{
		Name: "name",
		Age:  1,
	}
	// print data log
	glogger.Debug(user)
	glogger.Error(user)
	glogger.Debug(glogger.JsonEncode(user))

	os.Setenv("APP_ENV", "")
	logger = glogger.GetLogger("proj", zap.InfoLevel)
	logger.Error(glogger.JsonEncode(user))

}
