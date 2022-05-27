package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	assert.Equal(t, expected, got)
}

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {

		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		expected := 72.0

		assert.Equal(t, expected, got)
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{10.0}
		got := c.Area()
		expected := 314.1592653589793

		assert.Equal(t, expected, got)
	})

	t.Run("table driven", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 314.1592653589793},
			{Triangle{12, 6}, 36.0},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		}
	})

}
