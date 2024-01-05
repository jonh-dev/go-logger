package logger

import (
	"testing"
)

func TestLog(t *testing.T) {
	l := NewLogger()

	levels := []Level{INFO, SUCCESS, WARNING, ERROR, PANIC, FATAL}

	for _, level := range levels {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("A função Log causou um pânico ao ser chamada com o nível de log %v", level)
				}
			}()

			l.Log(level, "Mensagem de teste")
		}()
	}
}
