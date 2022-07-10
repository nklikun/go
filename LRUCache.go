package main

import "fmt"

// https://leetcode.cn/problems/lru-cache
func run146() {
	// cache := Constructor(2)
	// cache.Put(1, 1)
	// cache.Put(2, 2)
	// fmt.Println(cache.Get(1))
	// cache.Put(3, 3)
	// fmt.Println(cache.Get(2))
	// cache.Put(4, 4)
	// fmt.Println(cache.Get(1))
	// fmt.Println(cache.Get(3))
	// fmt.Println(cache.Get(4))

	// cache := Constructor(1)
	// cache.Put(2, 1)
	// fmt.Println(cache.Get(1))
	// fmt.Println(cache.Get(2))

	/*
		["LRUCache","put","get","put","get","get"]
		[[1],[2,1],[2],[3,2],[2],[3]]
	*/
	// cache := Constructor(1)
	// cache.Put(2, 1)
	// fmt.Println(cache.Get(2))
	// cache.Put(2, 2)
	// fmt.Println(cache.Get(2))

	/*
		["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
		[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
	*/
	cache := Constructor(3)
	fmt.Printf("null,")
	cache.Put(1, 1)
	fmt.Printf("null,")
	cache.Put(2, 2)
	fmt.Printf("null,")
	cache.Put(3, 3)
	fmt.Printf("null,")
	cache.Put(4, 4)
	fmt.Printf("null,")
	fmt.Print(cache.Get(4), ",")
	fmt.Print(cache.Get(3), ",")
	fmt.Print(cache.Get(2), ",")
	fmt.Print(cache.Get(1), ",")
	cache.Put(5, 5)
	fmt.Printf("null,")
	fmt.Print(cache.Get(1), ",")
	fmt.Print(cache.Get(2), ",")
	fmt.Print(cache.Get(3), ",")
	fmt.Print(cache.Get(4), ",")
	fmt.Print(cache.Get(5))
	fmt.Println("\nnull,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5")
}

type LRUListNode struct {
	key  int
	val  int
	prev *LRUListNode
	next *LRUListNode
}

type LRUCache struct {
	members  map[int]*LRUListNode
	head     *LRUListNode
	tail     *LRUListNode
	capacity int
	size     int
}

func Constructor(capacity int) LRUCache {
	var cache LRUCache
	cache.members = make(map[int]*LRUListNode)
	cache.capacity = capacity
	cache.size = 0
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.members[key]
	if !ok {
		return -1
	}
	this.moveToEnd(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.members[key]
	if !ok {
		if this.size == this.capacity {
			delete(this.members, this.head.key)
			// remove old node
			this.head = this.head.next
			this.size--
		}
		// add new node to the end
		new := &LRUListNode{key: key, val: value}
		this.addToEnd(new)
		this.members[key] = new
	} else {
		node.val = value
		this.moveToEnd(node)
	}

}

func (this *LRUCache) moveToEnd(node *LRUListNode) {
	if node.next == nil {
		// node is the tail
		return
	}

	if node.prev == nil {
		// node is head
		this.head = node.next
		this.head.prev = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	this.size--
	// add to the end
	this.addToEnd(node)
}

func (this *LRUCache) addToEnd(node *LRUListNode) {
	if this.head == nil {
		this.head = node
		this.tail = node
		this.head.prev = nil
	} else {
		node.prev = this.tail
		this.tail.next = node
		this.tail = node
	}
	this.tail.next = nil
	this.size++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
