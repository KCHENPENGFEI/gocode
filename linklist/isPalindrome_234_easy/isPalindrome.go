package main

import (
	"fmt"
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/21 下午1:50
 * @Desc: 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

 */

func main() {
	l := MakeLinkList([]int{1, 2, 3, 4, 5, 3, 2, 1})
	fmt.Println(isPalindrome(l))
}

func isPalindrome(head *ListNode) bool {
	length := getLen(head)
	i, p := 0, head
	for i < length/2 {
		p = p.Next
		i++
	}
	r := reverseList(p)
	p = head
	for r != nil {
		if r.Val != p.Val {
			return false
		}
		r = r.Next
		p = p.Next
	}
	return true
}

// 这里可以使用快慢指针来获取链表中间的位置
func getLen(head *ListNode) int {
	count := 0
	for head != nil {
		count += 1
		head = head.Next
	}
	return count
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
