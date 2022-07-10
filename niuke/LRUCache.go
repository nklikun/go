package niuke

import "fmt"

// https://www.nowcoder.com/practice/5dfded165916435d9defb053c63f1e84?tpId=295&tqId=2427094&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func Run_BM100() {
	c := LRUCacheConstructor(3)
	c.set(1, 1)
	print_BM100(&c)
	c.set(2, 2)
	print_BM100(&c)
	c.set(3, 3)
	print_BM100(&c)
	fmt.Println(c.get(2))
	print_BM100(&c)
	c.set(4, 4)
	print_BM100(&c)
	c.set(5, 5)
	print_BM100(&c)
	fmt.Println(c.get(2))
	print_BM100(&c)
}

func print_BM100(c *LRUCache) {
	// fmt.Println()
	// fmt.Printf("mHead: ")
	// for n := c.mHead; n != nil; n = n.Next {
	// 	fmt.Printf("%d-%d,", n.Key, n.Val)
	// }
	// fmt.Println()
	// fmt.Printf("mTail: ")
	// for n := c.mTail; n != nil; n = n.Prev {
	// 	fmt.Printf("%d-%d,", n.Key, n.Val)
	// }
	// fmt.Println()
	// fmt.Println("member lenght: ", len(c.members))
}

type DListNode struct {
	Key  int
	Val  int
	Next *DListNode
	Prev *DListNode
}
type LRUCache struct {
	// write code here
	mHead   *DListNode
	mTail   *DListNode
	members map[int]*DListNode
	size    int
	cap     int
}

func LRUCacheConstructor(capacity int) LRUCache {
	// write code here
	s := LRUCache{}
	s.cap = capacity
	s.members = make(map[int]*DListNode)
	return s
}

func (this *LRUCache) get(key int) int {
	// write code here
	if node, ok := this.members[key]; ok {
		this.removeNode(node)
		this.addToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) set(key int, value int) {
	// write code here
	var node *DListNode
	var ok bool
	if node, ok = this.members[key]; ok {
		node.Val = value
		this.removeNode(node)
		this.addToHead(node)
	} else {
		node = &DListNode{Key: key, Val: value}
		if this.cap == this.size {
			this.removeNode(this.mTail)
		}
		this.addToHead(node)
	}

}

func (this *LRUCache) removeNode(node *DListNode) {
	delete(this.members, node.Key)
	this.size--
	if node.Prev != nil && node.Next != nil {
		node.Prev.Next, node.Next.Prev = node.Next, node.Prev
	}
	if node.Prev == nil {
		// head
		if node.Next == nil {
			this.mHead = nil
		} else {
			this.mHead = node.Next
			this.mHead.Prev = nil
		}
	}
	if node.Next == nil {
		// tail
		if this.mTail.Prev == nil {
			this.mTail = nil
		} else {
			this.mTail.Prev.Next = nil
			this.mTail = this.mTail.Prev
		}
	}
}

func (this *LRUCache) addToHead(node *DListNode) {
	node.Prev = nil
	node.Next = this.mHead
	if this.mHead != nil {
		this.mHead.Prev = node
	}
	this.mHead = node
	if this.mTail == nil {
		this.mTail = node
	}
	this.members[node.Key] = node
	this.size++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache := Constructor(capacity);
 * output := LRUCache.get(key);
 * LRUCache.set(key,value);
 */
