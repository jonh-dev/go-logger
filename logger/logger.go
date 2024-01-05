package logger

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

type Level int

const (
	INFO Level = iota
	SUCCESS
	WARNING
	ERROR
	PANIC
	FATAL
)

type Logger interface {
	Info(message string)
	Success(message string)
	Warning(message string)
	Error(message string)
	Panic(message string)
	Fatal(message string)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Info(message string) {
	l.log(INFO, message)
}

func (l *logger) Success(message string) {
	l.log(SUCCESS, message)
}

func (l *logger) Warning(message string) {
	l.log(WARNING, message)
}

func (l *logger) Error(message string) {
	l.log(ERROR, message)
}

func (l *logger) Panic(message string) {
	l.log(PANIC, message)
}

func (l *logger) Fatal(message string) {
	l.log(FATAL, message)
}

func (l *logger) log(level Level, message string) {
	_, fn, line, _ := runtime.Caller(2)
	threadID := fmt.Sprintf("%d", time.Now().UnixNano())
	logDetails := fmt.Sprintf("| %s | %s:%d", threadID, fn, line)

	switch level {
	case INFO:
		log.Printf("\033[34m%s | INFO: %s\033[0m", logDetails, message)
	case SUCCESS:
		log.Printf("\033[32m%s | SUCCESS: %s\033[0m", logDetails, message)
	case WARNING:
		log.Printf("\033[33m%s | WARNING: %s\033[0m", logDetails, message)
	case ERROR:
		log.Printf("\033[31m%s | ERROR: %s\033[0m", logDetails, message)
	case PANIC:
		log.Printf("\033[35m%s | PANIC: %s\033[0m", logDetails, message)
	case FATAL:
		log.Printf("\033[2m%s | FATAL: %s\033[0m", logDetails, message)
	default:
		log.Printf("%s | %s", logDetails, message)
	}
}
