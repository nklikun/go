package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/
func Run59_2() {
	q := Constructor59_2()
	fmt.Println(q.Max_value())
	fmt.Println(q.Pop_front())
	q.Push_back(10)
	q.Push_back(9)
	q.Push_back(3)
	q.Push_back(2)
	q.Push_back(1)
	q.Push_back(8)
	fmt.Println(q.Max_value())
	fmt.Println(q.Pop_front())
	fmt.Println(q.Pop_front())
	fmt.Println(q.Pop_front())
	fmt.Println(q.Max_value())
	fmt.Println(q.Pop_front())
	fmt.Println(q.Max_value())
}

type MaxQueue struct {
	nums    []int
	ordered []int
}

func Constructor59_2() MaxQueue {
	queue := MaxQueue{nums: make([]int, 0), ordered: make([]int, 0)}
	return queue
}

func (this *MaxQueue) Max_value() int {
	if len(this.ordered) == 0 {
		return -1
	}
	return this.ordered[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.nums = append(this.nums, value)
	for len(this.ordered) > 0 && this.ordered[len(this.ordered)-1] < value {
		this.ordered = this.ordered[:len(this.ordered)-1]
	}
	this.ordered = append(this.ordered, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.nums) == 0 {
		return -1
	}
	num0 := this.nums[0]
	if num0 == this.ordered[0] {
		this.ordered = this.ordered[1:]
	}
	this.nums = this.nums[1:]
	return num0
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
