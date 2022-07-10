package jianzhioffer

import (
	"fmt"
)

// https://leetcode.cn/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/
func Run41() {
	mf := Constructor41()
	mf.AddNum(5)
	mf.AddNum(8)
	fmt.Println(mf.FindMedian())
	mf.AddNum(2)
	mf.AddNum(4)
	mf.AddNum(7)
	fmt.Println(mf.FindMedian())
}

type MedianFinder struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor41() MedianFinder {
	mf := MedianFinder{nums: make([]int, 0)}
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	l, r := 0, len(this.nums)-1
	mid := 0
	for l <= r {
		mid = l + (r-l)/2
		if this.nums[mid] > num {
			l = mid + 1
		} else if this.nums[mid] < num {
			r = mid - 1
		} else {
			l = mid
			break
		}
	}
	this.nums = append(this.nums[:l], append([]int{num}, this.nums[l:]...)...)
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.nums) < 0 {
		return 0
	}

	mid := len(this.nums) / 2
	if len(this.nums)%2 == 0 {
		return (float64(this.nums[mid]) + float64(this.nums[mid-1])) / 2
	} else {
		return float64(this.nums[mid])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
