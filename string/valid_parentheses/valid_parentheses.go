package main

func isValid(s string) bool {
	stack := []rune{}

	closers := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	for _, char := range s {
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if char != closers[pop] {
			return false
		}

	}

	if len(stack) != 0 {
		return false
	}

	return true
}
