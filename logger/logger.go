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
	Log(level Level, message string)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

/*
Log imprime uma mensagem no console com o nível de log informado.

- INFO: Mensagem informativa

- SUCCESS: Mensagem de sucesso

- WARNING: Mensagem de alerta

- ERROR: Mensagem de erro

- PANIC: Mensagem de erro seguida de um panic

- FATAL: Mensagem de erro seguida de um os.Exit(1)

@param level

@param message
*/
func (l *logger) Log(level Level, message string) {
	_, fn, line, _ := runtime.Caller(1)
	threadID := fmt.Sprintf("%d", time.Now().UnixNano())
	logDetails := fmt.Sprintf("%s | %s | %s:%d", time.Now().Format("2006/01/02 15:04:05"), threadID, fn, line)

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
