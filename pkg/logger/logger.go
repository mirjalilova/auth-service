package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	INFO  *log.Logger
	WARN  *log.Logger
	ERROR *log.Logger
	DEBUG *log.Logger
	TRACE *log.Logger
}

func NewLogger(basepath, path string) *Logger {
    l := &Logger{}

    // Use the full path including the basepath
    fullpath := basepath + "/" + path
    dir := fullpath[:len(fullpath)-len("/"+path)] // Remove the file name to get the directory

    // Ensure the directory exists
    if err := os.MkdirAll(dir, os.ModePerm); err != nil {
        fmt.Fprintf(os.Stderr, "logger - NewLogger - os.MkdirAll: %v\n", err)
        return l
    }

    file, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Fprintf(os.Stderr, "logger - NewLogger - os.OpenFile: %v\n", err)
        return l
    }

    multiWriter := io.MultiWriter(file, os.Stdout)

    l.INFO = log.New(multiWriter, "[INFO]   ", log.Lshortfile|log.LstdFlags)
    l.WARN = log.New(multiWriter, "[WARN]   ", log.Lshortfile|log.LstdFlags)
    l.ERROR = log.New(multiWriter, "[ERROR]  ", log.Lshortfile|log.LstdFlags)
    l.DEBUG = log.New(multiWriter, "[DEBUG]  ", log.Lshortfile|log.LstdFlags)
    l.TRACE = log.New(multiWriter, "[TRACE]  ", log.Lshortfile|log.LstdFlags)

    l.INFO.Println("Logger initialized successfully")

    return l
}