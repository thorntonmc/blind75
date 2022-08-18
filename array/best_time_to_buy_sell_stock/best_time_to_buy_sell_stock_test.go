package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{2, 4, 1}, 2},
		{[]int{2, 1, 2, 0, 1}, 1},
	}

	for _, test := range tests {
		got := maxProfit(test.prices)
		assert.Equal(t, test.want, got)
	}
}
