package logger

import (
	"testing"
)

func TestLog(t *testing.T) {
	log := NewLogger()

	tests := []struct {
		name string
		f    func(string)
	}{
		{"Info", log.Info},
		{"Success", log.Success},
		{"Warning", log.Warning},
		{"Error", log.Error},
		{"Panic", log.Panic},
		{"Fatal", log.Fatal},
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
