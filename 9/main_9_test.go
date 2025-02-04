package main

import (
	"9/pkg"
	"testing"
)

func TestCubeConverter(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  float64
	}{
		{"Zero", 0, 0},
		{"One", 1, 1},
		{"Two", 2, 8},
		{"Three", 3, 27},
		{"MaxUint8", 255, float64(255) * 255 * 255},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputChan := make(chan uint8, 1)
			outputChan := pkg.CubeConv(inputChan)

			inputChan <- tt.input
			close(inputChan)

			got := <-outputChan

			if got != tt.want {
				t.Errorf("CubeConv(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
