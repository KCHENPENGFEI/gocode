package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/3/3 下午11:17
 * @Desc: 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
 * 你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
 */

func main() {
	l := MakeLinkList([]int{1, 2, 3, 4})
	result := swapPairs(l)
	PrintLinkList(result)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 先保存第二个和第三个节点
	second := head.Next
	third := head.Next.Next
	head.Next.Next = head
	head.Next = swapPairs(third)
	return second
}

func swapPairsLoop(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	cur := head
	pre := dummy
	for cur != nil && cur.Next != nil {
		// 这里一定需要保存一下next和nextNext
		next := cur.Next
		nextNext := next.Next

		pre.Next = next
		cur.Next = nextNext
		next.Next = cur
		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}

// swapPairs2
//
//	@Description: 二刷递归
//	@param head
//	@return *ListNode
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := head.Next.Next
	newHead := head.Next
	head.Next.Next = head
	head.Next = swapPairs2(nextHead)
	return newHead
}
