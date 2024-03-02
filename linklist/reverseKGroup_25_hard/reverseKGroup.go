package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/11/12 下午5:08
 * @Desc: 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

func main() {
	l := MakeLinkList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintLinkList(reverseKGroup(l, 4))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := new(ListNode)
	pre.Next = head
	cur := head
	for i := 0; i < k; i++ {
		if cur != nil {
			cur = cur.Next
			pre = pre.Next
		} else {
			return head
		}
	}
	// 对链表进行截断
	pre.Next = nil
	// 先翻转被截断的部分
	newHead, tail := reverse(head)
	// 拼接
	tail.Next = reverseKGroup(cur, k)
	return newHead
}

// reverse
//
//	@Description: 翻转链表，这里还需要返回链表的tail
//	@param head
//	@return *ListNode
//	@return *ListNode
func reverse(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	newHead, _ := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead, head
}

// reverseKGroup2
//
//	@Description: 二刷
//	@param head
//	@param k
//	@return *ListNode
func reverseKGroup2(head *ListNode, k int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	cur := head
	for i := 0; i < k; i++ {
		if cur != nil {
			cur = cur.Next
			pre = pre.Next
		} else {
			return head
		}
	}
	// 先断开pre.Next指针
	pre.Next = nil
	// 翻转链表
	newHead, tail := reverseList(head)
	// 拼接
	tail.Next = reverseKGroup(cur, k)
	return newHead
}

func reverseList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	newHead, _ := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead, head
}
