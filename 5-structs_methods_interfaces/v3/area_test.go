package main

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{
			name: "triangle",
			shape: Triangle{base: 12, height: 6}, 
			want: 36.0,
		},
		{
			name: "rectangle",
			shape: Rectangle{width: 12, height: 6}, 
			want: 72.0,
		},
		{
			name: "circle",
			shape: Circle{radius: 10}, 
			want: 314.1592653589793,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want{
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}