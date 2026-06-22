package main

import (
	"math"
	"testing"
)

// Tolerance for floating-point comparisons
const delta = 1e-9

func TestShapes(t *testing.T) {
	// Table-driven test cases capturing unique shapes and edge cases
	tests := []struct {
		name          string
		shape         Shape
		wantArea      float64
		wantPerimeter float64
	}{
		{
			name:          "Standard Circle",
			shape:         Circle{radius: 5},
			wantArea:      math.Pi * 25,
			wantPerimeter: math.Pi * 10,
		},
		{
			name:          "Zero Radius Circle",
			shape:         Circle{radius: 0},
			wantArea:      0,
			wantPerimeter: 0,
		},
		{
			name:          "Standard Rectangle",
			shape:         Rectangle{width: 10, height: 5},
			wantArea:      50,
			wantPerimeter: 30,
		},
		{
			name:          "Square Rectangle",
			shape:         Rectangle{width: 4, height: 4},
			wantArea:      16,
			wantPerimeter: 16,
		},
		{
			name:          "Zero Dimension Rectangle",
			shape:         Rectangle{width: 0, height: 0},
			wantArea:      0,
			wantPerimeter: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Area method
			gotArea := tt.shape.Area()
			if math.Abs(gotArea-tt.wantArea) > delta {
				t.Errorf("%s Area() = %v, want %v", tt.name, gotArea, tt.wantArea)
			}

			// Test Perimeter method
			gotPerimeter := tt.shape.Perimeter()
			if math.Abs(gotPerimeter-tt.wantPerimeter) > delta {
				t.Errorf("%s Perimeter() = %v, want %v", tt.name, gotPerimeter, tt.wantPerimeter)
			}
		})
	}
}
