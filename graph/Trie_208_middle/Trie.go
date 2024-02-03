package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/31 下午4:46
 * @Desc:
 */

type Trie struct {
	Value *byte
	Sub   map[byte]*Trie
	IsEnd bool
}

func Constructor() Trie {
	return Trie{
		Value: nil,
		Sub:   make(map[byte]*Trie),
		IsEnd: false,
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		if v, ok := cur.Sub[c]; ok {
			// 已经存在
			cur = v
		} else {
			// 不存在插入
			t := Trie{
				Value: &c,
				Sub:   make(map[byte]*Trie),
			}
			cur.Sub[c] = &t
			cur = &t
		}
	}
	cur.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		if v, ok := cur.Sub[c]; !ok {
			return false
		} else {
			// 存在
			cur = v
		}
	}
	return cur.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if v, ok := cur.Sub[c]; !ok {
			return false
		} else {
			// 存在
			cur = v
		}
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
