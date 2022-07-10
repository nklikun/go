package main

// https://leetcode.cn/problems/plus-one/
func run66() {

}

func plusOne(digits []int) []int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1] += 1
		return digits
	}

	shift := true
	for i := len(digits) - 1; shift && i >= 0; i-- {
		shift = false
		if digits[i] == 9 {
			digits[i] = 0
			shift = true
		} else {
			digits[i] += 1
		}
	}

	if shift {
		digits = append([]int{1}, digits...)
	}
	return digits
}
