package sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("sum 5 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		assert.Equal(t, expected, got)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum 2 array", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		assert.Equal(t, expected, got)
	})

	t.Run("sum 1 array", func(t *testing.T) {
		got := SumAll([]int{1, 1, 1})
		expected := []int{3}

		assert.Equal(t, expected, got)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("2 tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4, 5})
		want := []int{5, 12}

		assert.Equal(t, want, got)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 4, 5})
		want := []int{0, 12}

		assert.Equal(t, want, got)
	})
}
