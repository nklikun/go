package main

//16.最接近的三数之和
// https://leetcode.cn/problems/3sum-closest
func run16() {
	nums := [8]int{-18, -15, -16, -17, -17, -17, -19, -20}
	threeSumClosest(nums[:], -52)
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	for i := 0; i < n-1; i++ {
		for j := i; j < n; j++ {
			if nums[i] > nums[j] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}

	ret := nums[0] + nums[1] + nums[2]

	diffRet := ret - target
	if diffRet < 0 {
		diffRet = -1 * diffRet
	}

	for i := 0; i < n-2; i++ {
		l := i + 1
		r := n - 1
		for r > l {
			sum := nums[i] + nums[l] + nums[r]
			diffSum := sum - target
			if diffSum < 0 {
				diffSum = -1 * diffSum
			}

			if diffSum < diffRet {
				ret = sum
				diffRet = ret - target
				if diffRet < 0 {
					diffRet = -1 * diffRet
				}
			}

			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				return sum
			}
		}
	}
	return ret
}

// func threeSumClosest(nums []int, target int) int {
// 	n := len(nums)
// 	if n < 3 {
// 		return 0
// 	}

// 	for i := 0; i < n-1; i++ {
// 		for j := i; j < n; j++ {
// 			if (target > 0 && nums[i] > nums[j]) ||
// 				(target < 0 && nums[i] < nums[j]) {
// 				tmp := nums[i]
// 				nums[i] = nums[j]
// 				nums[j] = tmp
// 			}
// 		}
// 	}

// 	ret := nums[0] + nums[1] + nums[2]

// 	diffRet := ret - target
// 	if diffRet < 0 {
// 		diffRet = -1 * diffRet
// 	}

// 	for i := 0; i < n-2; i++ {
// 		for j := i + 1; j < n-1; j++ {
// 			for k := j + 1; k < n; k++ {
// 				sum := nums[i] + nums[j] + nums[k]
// 				diffSum := sum - target

// 				if diffSum < 0 {
// 					diffSum = -1 * diffSum
// 				}
// 				if diffSum < diffRet {
// 					ret = sum
// 					diffRet = ret - target
// 					if diffRet < 0 {
// 						diffRet = -1 * diffRet
// 					}
// 				}

// 				if (target > 0 && sum-target > 0) ||
// 					(target < 0 && sum-target < 0) {
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return ret
// }

// func threeSumClosest(nums []int, target int) int {
// 	n := len(nums)
// 	if n < 3 {
// 		return 0
// 	}

// 	ret := nums[0] + nums[1] + nums[2]

// 	diffRet := ret - target
// 	if diffRet < 0 {
// 		diffRet = -1 * diffRet
// 	}
// 	for i := 0; i < n-2; i++ {
// 		for j := i + 1; j < n-1; j++ {
// 			for k := j + 1; k < n; k++ {
// 				sum := nums[i] + nums[j] + nums[k]
// 				diffSum := sum - target
// 				if diffSum < 0 {
// 					diffSum = -1 * diffSum
// 				}
// 				if diffSum < diffRet {
// 					ret = sum
//                     diffRet = ret - target
// 					if diffRet < 0 {
// 						diffRet = -1 * diffRet
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return ret
// }
