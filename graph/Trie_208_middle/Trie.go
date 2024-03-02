package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/31 下午4:46
 * @Desc:
 */

type Trie struct {
	Val   *byte
	Son   map[byte]*Trie
	IsEnd bool
}

func Constructor() Trie {
	return Trie{
		Val: nil,
		Son: make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node, ok := cur.Son[c]; ok {
			cur = node
		} else {
			_node := &Trie{
				Val: &c,
				Son: make(map[byte]*Trie),
			}
			cur.Son[c] = _node
			cur = _node
		}
		if i == len(word)-1 {
			cur.IsEnd = true
		}
	}
}

func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node, ok := this.Son[c]; ok {
			this = node
		} else {
			return false
		}
	}
	return this.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if node, ok := this.Son[c]; ok {
			this = node
		} else {
			return false
		}
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("ab")
	fmt.Println(trie.Search("a"))
	fmt.Println(trie.Search("ab"))
	fmt.Println(trie.StartsWith("a"))
	fmt.Println(trie.StartsWith("ab"))
}
