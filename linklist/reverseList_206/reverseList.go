package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/3/3 下午11:50
 * @Desc: 翻转链表
 */

func main() {
	l := MakeLinkList([]int{1, 2, 3, 4, 5, 10})
	PrintLinkList(reverseList(l))
}

var res *ListNode

func reverseList(head *ListNode) *ListNode {
	_reverseList(head)
	return res
}

func _reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		res = head
		return head
	}
	p := _reverseList(head.Next)
	p.Next = head
	head.Next = nil
	return head
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
