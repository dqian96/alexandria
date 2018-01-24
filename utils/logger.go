package utils

import (
	"log"
	"os"
)

var (
	infoPrefix    = "INFO "
	warningPrefix = "WARN "
	errorPrefix   = "ERROR "

	infoLog = log.New(os.Stdout,
		infoPrefix,
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	warningLog = log.New(os.Stdout,
		warningPrefix,
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	errorLog = log.New(os.Stderr,
		errorPrefix,
		log.Ldate|log.Ltime|log.Lshortfile,
	)
)

// Info returns a logger used for logging info messages
func Info(reqID string) *log.Logger {
	if reqID != "" {
		infoLog.SetPrefix(infoPrefix + reqID + " ")
	} else {
		infoLog.SetPrefix(infoPrefix)
	}
	return infoLog
}

// Warning returns a logger used for logging warning messages
func Warning(reqID string) *log.Logger {
	if reqID != "" {
		warningLog.SetPrefix(warningPrefix + reqID + " ")
	} else {
		warningLog.SetPrefix(warningPrefix)
	}
	return warningLog
}

// Error returns a logger used for logging error messages (outputs to stderr)
func Error(reqID string) *log.Logger {
	if reqID != "" {
		errorLog.SetPrefix(errorPrefix + reqID + " ")
	} else {
		errorLog.SetPrefix(errorPrefix)
	}
	return errorLog
}
