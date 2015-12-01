package glog

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var FLAG_LOG_TRACE = flag.String(
	"log.trace", "", "Regex of loggers to enable trace for.")
var FLAG_LOG_DEBUG = flag.String(
	"log.debug", "", "Regex of loggers to enable debug for.")
var FLAG_LOG_INFO = flag.String(
	"log.info", ".*", "Regex of loggers to enable info for.")
var FLAG_LOG_WARN = flag.String(
	"log.warn", ".*", "Regex of loggers to enable warn for.")
var FLAG_LOG_ERROR = flag.String(
	"log.error", ".*", "Regex of loggers to enable error for.")
var FLAG_LOG_FATAL = flag.String(
	"log.fatal", "", "Regex of loggers to enable fatal for.")

// levels
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

const (
	PrintLevelTrace = "[ TRACE ] "
	PrintLevelDebug = "[ DEBUG ] "
	PrintLevelInfo  = "[ INFO  ] "
	PrintLevelWarn  = "[ WARN  ] "
	PrintLevelError = "[ ERROR ] "
	PrintLevelFatal = "[ FATAL ] "
)

type Logger struct {
	level      int
	baseLogger *log.Logger
	baseFile   *os.File
}

func New(strLevel string, pathname string) (*Logger, error) {
	// level
	var level int
	switch strings.ToLower(strLevel) {
	case "trace":
		level = LevelTrace
	case "debug":
		level = LevelDebug
	case "info":
		level = LevelInfo
	case "warn":
		level = LevelWarn
	case "error":
		level = LevelError
	case "fatal":
		level = LevelFatal
	default:
		return nil, errors.New("unknown level: " + strLevel)
	}

	// logger
	var baseLogger *log.Logger
	var baseFile *os.File
	if pathname != "" {
		now := time.Now()

		filename := fmt.Sprintf("%d%02d%02d_%02d_%02d_%02d.log",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())

		file, err := os.Create(path.Join(pathname, filename))
		if err != nil {
			return nil, err
		}

		baseLogger = log.New(file, "", log.LstdFlags)
		baseFile = file
	} else {
		baseLogger = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	}

	// new
	logger := new(Logger)
	logger.level = level
	logger.baseLogger = baseLogger
	logger.baseFile = baseFile

	return logger, nil
}

// It's dangerous to call the method on logging
func (logger *Logger) Close() {
	if logger.baseFile != nil {
		logger.baseFile.Close()
	}

	logger.baseLogger = nil
	logger.baseFile = nil
}

func (logger *Logger) doPrintf(level int, printLevel string, format string, a ...interface{}) {
	if level < logger.level {
		return
	}
	if logger.baseLogger == nil {
		panic("logger closed")
	}

	format = printLevel + format
	logger.baseLogger.Printf(format, a...)

	if level == LevelFatal {
		os.Exit(1)
	}
}

func (logger *Logger) Trace(format string, a ...interface{}) {
	logger.doPrintf(LevelTrace, PrintLevelTrace, format, a...)
}

func (logger *Logger) Debug(format string, a ...interface{}) {
	logger.doPrintf(LevelDebug, PrintLevelDebug, format, a...)
}

func (logger *Logger) Info(format string, a ...interface{}) {
	logger.doPrintf(LevelInfo, PrintLevelInfo, format, a...)
}

func (logger *Logger) Warn(format string, a ...interface{}) {
	logger.doPrintf(LevelWarn, PrintLevelWarn, format, a...)
}

func (logger *Logger) Error(format string, a ...interface{}) {
	logger.doPrintf(LevelError, PrintLevelError, format, a...)
}

func (logger *Logger) Fatal(format string, a ...interface{}) {
	logger.doPrintf(LevelFatal, PrintLevelFatal, format, a...)
}

var gLogger, _ = New("trace", "")

// It's dangerous to call the method on logging
func Export(logger *Logger) {
	if logger != nil {
		gLogger = logger
	}
}

func Trace(format string, a ...interface{}) {
	gLogger.Trace(format, a...)
}

func Debug(format string, a ...interface{}) {
	gLogger.Debug(format, a...)
}

func Info(format string, a ...interface{}) {
	gLogger.Info(format, a...)
}

func Warn(format string, a ...interface{}) {
	gLogger.Warn(format, a...)
}

func Error(format string, a ...interface{}) {
	gLogger.Error(format, a...)
}

func Fatal(format string, a ...interface{}) {
	gLogger.Fatal(format, a...)
}

func Close() {
	gLogger.Close()
}
