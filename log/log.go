package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Trace   = log.New(io.Discard, "", 0)
	Debug   = log.New(io.Discard, "", 0)
	Error   = log.New(io.Discard, "", 0)
	Info    = log.New(io.Discard, "", 0)
	Warning = log.New(io.Discard, "", 0)
)

var (
	LogLevel = "info"
)

const defaultLogFlags = log.LstdFlags | log.Lshortfile | log.Lmsgprefix

// SetLevel sets the log level
func SetLevel(level string) error {
	LogLevel = level

	switch level {
	case "trace":
		Trace = log.New(os.Stdout, "[TRACE] ", defaultLogFlags)
		fallthrough
	case "debug":
		Debug = log.New(os.Stdout, "[DEBUG] ", defaultLogFlags)
		fallthrough
	case "info":
		Info = log.New(os.Stdout, "[INFO] ", defaultLogFlags)
		fallthrough
	case "warning":
		Warning = log.New(os.Stderr, "[WARNING] ", defaultLogFlags)
		fallthrough
	case "error":
		Error = log.New(os.Stderr, "[ERROR] ", defaultLogFlags)
	default:
		return fmt.Errorf("unknown log level: %s", level)
	}
	return nil
}

func IsTrace() bool {
	return LogLevel == "trace" || LogLevel == "debug" || LogLevel == "info" || LogLevel == "warning" || LogLevel == "error"
}

func IsDebug() bool {
	return LogLevel == "debug" || LogLevel == "info" || LogLevel == "warning" || LogLevel == "error"
}

func IsInfo() bool {
	return LogLevel == "info" || LogLevel == "warning" || LogLevel == "error"
}

func IsWarning() bool {
	return LogLevel == "warning" || LogLevel == "error"
}

func IsError() bool {
	return LogLevel == "error"
}
