package test

import "fmt"

func Run_findNums() {
	fmt.Println(findNums([]int{3, 3, 3, 3, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}))
}

// 一个整形数组中有两个元素的值的数量分别都在数组中超过了总长度的三分之一，把他们找出来
func findNums(nums []int) []int {
	num1, num2 := 0, 0
	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if cnt1 > 0 && num1 == num {
			cnt1++
		} else if cnt2 > 0 && num2 == num {
			cnt2++
		} else {
			if cnt1 > cnt2 {
				if cnt2 == 0 {
					cnt2++
					num2 = num
				} else {
					cnt2--
				}
			} else {
				if cnt1 == 0 {
					cnt1++
					num1 = num
				} else {
					cnt1--
				}
			}
		}
	}
	return []int{num1, num2}
}
