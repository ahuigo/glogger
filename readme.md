# glogger
Simple global logger for golang.

## examples
Refer to tests/logger_test.go


    package tests

    import (
        "testing"

        "github.com/ahuigo/glogger"
        "fmt"
    )

    type User struct{
        Name string 
        Age int
    }

    func TestLogger(t *testing.T) {
        // this log is named root
        glogger.Debug("This is a debug log")

        // get logger named with "proj"
        logger:=glogger.GetLogger("proj", glogger.InfoLevel)
        logger.Info("This is a info log")
        logger.Debug("This is a debug log") //not printed

        user:=User{
            Name:"name",
            Age:1,
        }
        // print data log
        fmt.Println("print data log")
        glogger.Debug(user)
    }

## test
    $ make test
    go test -v ./tests
    === RUN   TestLogger
    default log format:
    datetime                	log_level project_name	code_path:line	<message>
    2021-11-03T20:05:58.454+0800	DEBUG	root	tests/logger_test.go:20	This is a debug log
    2021-11-03T20:05:58.454+0800	INFO	proj	tests/logger_test.go:24	This is a info log
    print data log
    2021-11-03T20:05:58.454+0800	DEBUG	root	tests/logger_test.go:33	{name 1}
    2021-11-03T20:05:58.454+0800	DEBUG	root	tests/logger_test.go:34	{"Name":"name","Age":1}

