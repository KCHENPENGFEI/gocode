package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/10 下午5:31
 * @Desc:
 */

func main() {
	lru := Constructor(3)
	value := lru.Get(1)
	fmt.Println(value)
	lru.Put(1, 1)
	value = lru.Get(1)
	fmt.Println(value)
	lru.Put(2, 2)
	value = lru.Get(2)
	fmt.Println(value)
	value = lru.Get(1)
	fmt.Println(value)
	lru.Put(3, 3)
	lru.Put(4, 4)
	value = lru.Get(2)
	fmt.Println(value)
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

type LRUCache struct {
	cap    int
	length int
	data   map[int]int
	index  map[int]*Node
	head   *Node
	tail   *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cap:   capacity,
		data:  make(map[int]int, capacity),
		index: make(map[int]*Node, capacity),
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.data[key]; !ok {
		return -1
	} else {
		newHead := this.index[key]
		if newHead == this.head {
			// do nothing
			return value
		} else if newHead == this.tail {
			newHead.Pre.Next = newHead.Next
			this.tail = newHead.Pre
		} else {
			newHead.Pre.Next = newHead.Next
			newHead.Next.Pre = newHead.Pre
		}
		newHead.Next = this.head
		newHead.Pre = nil
		this.head.Pre = newHead
		this.head = newHead
		return value
	}
}

func (this *LRUCache) Put(key int, value int) {
	oldValue := this.Get(key)
	if oldValue != -1 {
		// 存在则更新
		this.data[key] = value
	} else {
		// 不存在则插入
		newHead := &Node{
			Val:  key,
			Next: this.head,
		}
		if this.head == nil {
			this.head = newHead
			this.tail = newHead
		} else {
			this.head.Pre = newHead
			this.head = newHead
		}
		this.data[key] = value
		this.index[key] = newHead
		if this.length < this.cap {
			// 还存在空余
			this.length++
		} else {
			// 不存在空余则淘汰最久未使用的关键字
			tailKey := this.tail.Val
			this.tail = this.tail.Pre
			this.tail.Next = nil
			delete(this.data, tailKey)
			delete(this.index, tailKey)
		}
	}
}
