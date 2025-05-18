package tool

import (
	"math"
	"testing"
)

func TestRoundFloat(t *testing.T) {
	tests := []struct {
		value    float64
		places   int
		expected float64
	}{
		{123.4567, 0, 123},
		{123.4567, 1, 123.5},
		{123.4567, 2, 123.46},
		{123.454, 2, 123.45},
		{-123.4567, 2, -123.46},
		{0.0, 2, 0.0},
		{1.005, 2, 1.01}, // 浮点数四舍五入常见陷阱
		{math.Pi, 3, 3.142},
		{1.9999, 3, 2.0},
		{1.2345, -2, 1}, // places为负
	}

	for _, tt := range tests {
		result := RoundFloat(tt.value, tt.places)
		if result != tt.expected {
			t.Errorf("RoundFloat(%v, %d) = %v, want %v", tt.value, tt.places, result, tt.expected)
		}
	}
}

func TestFormatFloat(t *testing.T) {
	tests := []struct {
		value    float64
		places   int
		expected float64
	}{
		{123.4567, 0, 123},
		{123.4567, 1, 123.5},
		{123.4567, 2, 123.46},
		{123.454, 2, 123.45},
		{-123.4567, 2, -123.46},
		{0.0, 2, 0.0},
		{1.005, 2, 1.01},
		{math.Pi, 3, 3.142},
		{1.9999, 3, 2.0},
		{1.2345, -2, 1}, // places为负
	}

	for _, tt := range tests {
		result, err := FormatFloat(tt.value, tt.places)
		if err != nil {
			t.Errorf("FormatFloat(%v, %d) error: %v", tt.value, tt.places, err)
		}
		if result != tt.expected {
			t.Errorf("FormatFloat(%v, %d) = %v, want %v", tt.value, tt.places, result, tt.expected)
		}
	}
}
