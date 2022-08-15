package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
	}

	for _, test := range tests {
		got := twosum(test.nums, test.target)
		assert.Equal(t, test.want, got)
	}
}
