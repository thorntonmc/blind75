package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
*/

/*
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

/*
Input: n = 4
Output: 4
Explanation: There are three ways to climb to the top.
1. 1 + 1 + 1 + 1
2. 2 + 1 + 1
3. 2 + 2
4. 1 + 2 + 1
5. 1 + 1 + 2
*/

func TestClimbStairs(t *testing.T) {
	result := climbStairs(2)
	assert.Equal(t, 2, result)

	result = climbStairs(3)
	assert.Equal(t, 3, result)

	result = climbStairs(4)
	assert.Equal(t, 5, result)
}
