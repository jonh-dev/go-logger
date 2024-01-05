package logger

import (
	"testing"

	"github.com/jonh-dev/go-logger/logger"
)

func TestLog(t *testing.T) {
	tests := []struct {
		name string
		f    func(string)
	}{
		{"Info", logger.Info},
		{"Success", logger.Success},
		{"Warning", logger.Warning},
		{"Error", logger.Error},
		{"Panic", logger.Panic},
		{"Fatal", logger.Fatal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The function %s caused a panic", tt.name)
				}
			}()

			tt.f("Test message")
		})
	}
}
