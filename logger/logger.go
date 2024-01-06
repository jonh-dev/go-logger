package logger

import (
	"fmt"
	"log"
	"path/filepath"
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

var (
	Info    = logFunc(INFO)
	Success = logFunc(SUCCESS)
	Warning = logFunc(WARNING)
	Error   = logFunc(ERROR)
	Panic   = logFunc(PANIC)
	Fatal   = logFunc(FATAL)
)

/*
Log imprime uma mensagem no console com o nível de log informado.

- level: nível de log

- message: mensagem a ser impressa

- fn: nome da função que chamou o log

- line: linha da função que chamou o log

- threadID: ID da thread que chamou o log

- logDetails: detalhes do log

@params level Level

@params message string

@return void

@struct logger
*/
func logFunc(level Level) func(message string) {
	return func(message string) {
		_, fn, line, _ := runtime.Caller(1)
		threadID := fmt.Sprintf("%d", time.Now().UnixNano())
		logDetails := fmt.Sprintf("| %s | %s:%d", threadID, filepath.Base(fn), line)

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
}
