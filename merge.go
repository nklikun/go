package main

// https://leetcode.cn/problems/merge-sorted-array/
func run88() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i1, i2, e1 := m-1, n-1, m+n-1; i1 >= 0 || i2 >= 0; e1-- {
		if i1 >= 0 && i2 >= 0 {
			if nums1[i1] > nums2[i2] {
				nums1[e1] = nums1[i1]
				i1--
			} else {
				nums1[e1] = nums2[i2]
				i2--
			}
		} else if i1 >= 0 {
			nums1[e1] = nums1[i1]
			i1--
		} else {
			nums1[e1] = nums2[i2]
			i2--
		}
	}
}
