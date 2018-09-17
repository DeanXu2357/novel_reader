package logs

import (
    "log"
    "os"
)

var (
    Info    *log.Logger
    Error   *log.Logger
)

func init() {
    log_file, err := os.OpenFile("golang.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("file open error : %v", err)
    }

    Info    = log.New(os.Stdout, "INFO  : ", log.Ldate|log.Ltime)
    Error   = log.New(os.Stderr, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)

    Info.SetOutput(log_file)
    Error.SetOutput(log_file)
}
