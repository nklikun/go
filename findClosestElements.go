package main

import "fmt"

/*
658. 找到 K 个最接近的元素
给定一个 排序好 的数组 arr ，两个整数 k 和 x ,从数组中找到最靠
近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
整数 a 比整数 b 更接近 x 需要满足：
|a - x| < |b - x| 或者
|a - x| == |b - x| 且 a < b

示例 1：
输入：arr = [1,2,3,4,5], k = 4, x = 3
输出：[1,2,3,4]

示例 2：
输入：arr = [1,2,3,4,5], k = 4, x = -1
输出：[1,2,3,4]

提示：
1 <= k <= arr.length
1 <= arr.length <= 104
arr 按 升序 排列
-104 <= arr[i], x <= 104

https://leetcode.cn/problems/find-k-closest-elements/
*/

func run658() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))  // 1, 2, 3, 4
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1)) // 1, 2, 3, 4
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 6))  // 2, 3, 4, 5

	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 5))             // 2, 3, 4, 5
	fmt.Println(findClosestElements([]int{0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5)) // 3, 3, 4
	fmt.Println(findClosestElements([]int{1, 1, 2, 2, 2, 2, 2, 3, 3}, 3, 3)) // 2,3,3
}

// O(logn+k)
func findClosestElements(arr []int, k int, x int) []int {
	abs := func(a int, b int) int {
		if a >= b {
			return a - b
		}
		return b - a
	}
	// 此时l，r代表的返回结果的不是左右边界
	// 指的是以[i, i+k]为结果，i的边界位于l，r之间
	l, r := 0, len(arr)-k
	// 找到 l == r，则 l == r == 返回结果的start index
	for l < r {
		mid := (l + r) / 2
		s, e := mid, mid+k
		// arr[e] < x, 收缩左边范围
		// arr[s] > x, 收缩右边范围
		// 如果x在arr[s]到arr[e]之间，则根据abs差值收缩范围
		if arr[e] <= x {
			l = mid + 1
		} else if arr[s] >= x {
			r = mid
		} else if abs(arr[s], x) > abs(arr[e], x) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return arr[l : l+k]
}

/*
// O(n)
func findClosestElements(arr []int, k int, x int) []int {
	abs := func(a int, b int) int {
		if a >= b {
			return a - b
		}
		return b - a
	}

	l, r := 0, len(arr)-1
	for r+1-l > k {
		if abs(arr[l], x) > abs(arr[r], x) {
			l++
		} else {
			r--
		}
	}

	return arr[l : r+1]
}
*/
