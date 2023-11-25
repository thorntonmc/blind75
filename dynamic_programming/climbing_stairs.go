package dynamic_programming

//You are climbing a staircase. It takes n steps to reach the top.
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

func climb(n int, cache map[int]int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if s, ok := cache[n]; ok {
		return s
	}

	cache[n] = climb(n-1, cache) + climb(n-2, cache)
	return cache[n]
}

func climbStairs(n int) int {
	cache := make(map[int]int, n)
	return climb(n, cache)
}

/*

n = 10 || + 1

n = 9

*/
