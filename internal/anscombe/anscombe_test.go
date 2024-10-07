package anscombe_test

import (
	"math"
	"testing"

	"github.com/Saracomethstein/go_day_00/internal/anscombe"
)

func TestMeanCalc(t *testing.T) {
	tests := []struct {
		input    []int
		expected float32
		hasError bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0, false},
		{[]int{10, 20, 30, 40, 50}, 30.0, false},
		{[]int{}, 0.0, true},
	}

	for _, test := range tests {
		result, err := anscombe.MeanCalc(test.input)
		if test.hasError && err == nil {
			t.Errorf("Expected error but got nil")
		}
		if !test.hasError && result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestMedianCalc(t *testing.T) {
	tests := []struct {
		input    []int
		expected float32
		hasError bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0, false},
		{[]int{10, 20, 30, 40, 50}, 30.0, false},
		{[]int{1, 2, 3, 4}, 2.5, false},
		{[]int{}, 0.0, true},
	}

	for _, test := range tests {
		result, err := anscombe.MedianCalc(test.input)
		if test.hasError && err == nil {
			t.Errorf("Expected error but got nil")
		}
		if !test.hasError && result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestModeCalc(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		hasError bool
	}{
		{[]int{1, 2, 2, 3, 3, 3}, 3, false},
		{[]int{10, 10, 20, 20, 20}, 20, false},
		{[]int{}, 0, true},
	}

	for _, test := range tests {
		result, err := anscombe.ModeCalc(test.input)
		if test.hasError && err == nil {
			t.Errorf("Expected error but got nil")
		}
		if !test.hasError && result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestSdCalc(t *testing.T) {
	tests := []struct {
		input    []int
		expected float32
		hasError bool
	}{
		{[]int{1, 2, 3, 4, 5}, 1.4142135, false},
		{[]int{10, 20, 30, 40, 50}, 14.142136, false},
		{[]int{}, 0.0, true},
	}

	for _, test := range tests {
		result, err := anscombe.SDCalc(test.input)
		if test.hasError && err == nil {
			t.Errorf("Expected error but got nil")
		}
		if !test.hasError && math.Abs(float64(result-test.expected)) > 1e-6 {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestMakeCalc(t *testing.T) {
	tests := []struct {
		input    []int
		expected anscombe.Metrics
		hasError bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			anscombe.Metrics{MEAN: 3.0, MEDIAN: 3.0, MODE: 1, SD: 1.4142135},
			false,
		},
		{
			[]int{},
			anscombe.Metrics{},
			true,
		},
	}

	for _, test := range tests {
		calc := anscombe.Metrics{}
		result, err := anscombe.MakeCalc(calc, test.input)
		if test.hasError && err != nil {
			t.Errorf("Expected error but got nil")
		}
		if !test.hasError && result != test.expected {
			t.Errorf("Expected %+v, got %+v", test.expected, result)
		}
	}
}
