# glogger
Simple global logger for golang.

## examples
Refer to tests/logger_test.go


    package tests

    import (
        "testing"

        "github.com/ahuigo/glogger"
        "go.uber.org/zap"
        "fmt"
    )

    type User struct{
        Name string 
        Age int
    }

    func TestLogger(t *testing.T) {
        // this log is named root
        glogger.Glogger.Debug("This is a debug log")

        // get logger named with "proj"
        logger:=glogger.GetLogger("proj", zap.InfoLevel)
        logger.Info("This is a info log")
        logger.Debug("This is a debug log") //not printed

        user:=User{
            Name:"name",
            Age:1,
        }
        // print data log
        fmt.Println("print data log")
        glogger.Glogger.Debug(user)
        glogger.Glogger.Debug(glogger.JsonEncode(user))
    }

## test
    $ make test
    go test -v ./tests
    === RUN   TestLogger
    2021-11-03T19:49:35.754+0800	DEBUG	root	tests/logger_test.go:18	This is a debug log
    2021-11-03T19:49:35.754+0800	INFO	proj	tests/logger_test.go:22	This is a info log
    print data log
    2021-11-03T19:49:35.754+0800	DEBUG	root	tests/logger_test.go:31	{name 1}
    2021-11-03T19:49:35.754+0800	DEBUG	root	tests/logger_test.go:32	{"Name":"name","Age":1}
    --- PASS: TestLogger (0.00s)

