package main

import "fmt"

/*
不使用任何内建的哈希表库设计一个哈希映射（HashMap）。

实现 MyHashMap 类：

MyHashMap() 用空映射初始化对象
void put(int key, int value) 向 HashMap 插入一个键值对 (key, value) 。如果 key 已经存在于映射中，则更新其对应的值 value 。
int get(int key) 返回特定的 key 所映射的 value ；如果映射中不包含 key 的映射，返回 -1 。
void remove(key) 如果映射中存在 key 的映射，则移除 key 和它所对应的 value 。


示例：

输入：
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
输出：
[null, null, null, 1, -1, null, 1, null, -1]

解释：
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // myHashMap 现在为 [[1,1]]
myHashMap.put(2, 2); // myHashMap 现在为 [[1,1], [2,2]]
myHashMap.get(1);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
myHashMap.get(3);    // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
myHashMap.put(2, 1); // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
myHashMap.get(2);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
myHashMap.remove(2); // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
myHashMap.get(2);    // 返回 -1（未找到），myHashMap 现在为 [[1,1]]


提示：

0 <= key, value <= 106
最多调用 104 次 put、get 和 remove 方法

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/design-hashmap
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func run706() {
	m := ConstructorHashMap()
	m.Put(1, 1)
	m.Put(998, 2)
	fmt.Println(m.Get(1))
}

type MyHashMap struct {
	base int
	eles [][]*Pair
}

type Pair struct {
	k int
	v int
}

func ConstructorHashMap() MyHashMap {
	m := MyHashMap{}
	m.base = 997
	m.eles = make([][]*Pair, m.base)
	return m
}

func (this *MyHashMap) Put(key int, value int) {
	hash := key % this.base
	if this.eles[hash] == nil {
		this.eles[hash] = make([]*Pair, 0)
	}
	this.eles[hash] = append(this.eles[hash], &Pair{k: key, v: value})
}

func (this *MyHashMap) Get(key int) int {
	hash := key % this.base
	for _, pair := range this.eles[hash] {
		if pair.k == key {
			return pair.v
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	hash := key % this.base
	for i := 0; i < len(this.eles[hash]); i++ {
		if this.eles[hash][i].k == key {
			this.eles[hash] = append(this.eles[hash][0:i], this.eles[hash][i+1:]...)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
