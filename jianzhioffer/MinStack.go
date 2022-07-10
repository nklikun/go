package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
func Run30() {
	m := Constructor1()
	m.Push(2147483646)
	Run30_print(m)
	m.Push(2147483646)
	Run30_print(m)
	m.Push(2147483647)
	Run30_print(m)
	m.Top()
	Run30_print(m)
	m.Pop()
	Run30_print(m)
	m.Min()
	Run30_print(m)
	m.Pop()
	Run30_print(m)
	m.Min()
	Run30_print(m)
	m.Pop()
	Run30_print(m)
	m.Push(2147483647)
	Run30_print(m)
	m.Top()
	Run30_print(m)
	fmt.Println(m.Min())
	Run30_print(m)
	m.Push(-2147483648)
	Run30_print(m)
	m.Top()
	Run30_print(m)
	m.Min()
	Run30_print(m)
	m.Pop()
	Run30_print(m)
	fmt.Println(m.Min())
	Run30_print(m)
}

func Run30_print(m MinStack) {
	fmt.Println("stack: ", m.stack)
	fmt.Println("minStack: ", m.minStack)
	fmt.Println("-------------------")
}

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	s := MinStack{}
	s.stack = make([]int, 0)
	s.minStack = make([]int, 0)
	return s
}

func (this *MinStack) Push(x int) {
	min := int(^uint32(0) >> 1)
	if len(this.minStack) != 0 {
		min = this.minStack[len(this.minStack)-1]
	}
	if x < min {
		min = x
	}
	this.minStack = append(this.minStack, min)
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) < 1 {
		return
	}

	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) < 1 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.minStack) < 1 {
		return 0
	}
	return this.minStack[len(this.stack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
