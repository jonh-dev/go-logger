package logger

import (
	"testing"
)

func TestLog(t *testing.T) {
	tests := []struct {
		name string
		f    func(string)
	}{
		{"Info", Info},
		{"Success", Success},
		{"Warning", Warning},
		{"Error", Error},
		{"Panic", Panic},
		{"Fatal", Fatal},
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
