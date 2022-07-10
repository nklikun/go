package jianzhioffer

// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
func Run09() {
	queue := Constructor()
	queue.AppendTail(3)
	queue.DeleteHead()
	queue.DeleteHead()
}

type CQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() CQueue {
	queue := CQueue{}
	queue.inStack = make([]int, 0)
	queue.outStack = make([]int, 0)
	return queue
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		if len(this.inStack) == 0 {
			return -1
		}
		this.in2out()
	}
	res := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return res
}

func (this *CQueue) in2out() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
