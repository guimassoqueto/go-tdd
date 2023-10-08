package shapes

import "testing"

func TestShapes(t *testing.T) {
	areaTests := []struct{
		description string
		shape Shape
		want float64
	} {
		{
			description: "triangle: should return correct area",
			shape: &Triangle{Base: 6, Height: 12},
			want: 36.0,
		},
		{
			description: "circle: should return correct area",
			shape: &Circle{Radious: 10},
			want: 314.1592653589793,
		},
		{
			description: "rectangle: should return correct area",
			shape: &Rectangle{Height: 2, Width: 8},
			want: 16.0,
		},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.description, func(t *testing.T) {
			// CalcArea(testCase.shape) pode ser intercambiado com testCase.shape.Area()
			// pois representam o mesmo dado
			if CalcArea(testCase.shape) != testCase.want {
				t.Errorf("got %v want %v", testCase.shape.Area(), testCase.want)
			}
		})
	}
}