# glogger
Simple global logger for golang.

## examples
Refer to tests/logger_test.go

## test

    $ make test
    go test -v ./tests
    === RUN   TestLogger
    2021-10-09T15:04:42.945+0800    INFO    proj    tests/logger_test.go:11 This is a info log
    2021-10-09T15:04:42.945+0800    DEBUG   root    tests/logger_test.go:13 This is a debug log
    --- PASS: TestLogger (0.00s)
