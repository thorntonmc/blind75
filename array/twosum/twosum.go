package main

func twosum(nums []int, target int) []int {
	var result []int

	if len(nums) == 2 {
		return []int{0, 1}
	}

	answerMap := make(map[int]int)

	// naive first`
	for i := 0; i < len(nums); i++ {
		v, ok := answerMap[target-nums[i]]

		if ok {
			return []int{v, i}
		}
		answerMap[nums[i]] = i
	}
	return result
}
