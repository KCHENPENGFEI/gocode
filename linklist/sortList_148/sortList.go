package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/3/7 下午11:59
 * @Desc:
 */

func main() {
	l1 := MakeLinkList([]int{-1, 5, 3, 4, 0})
	PrintLinkList(sortList(l1))
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tmp := slow
	slow = slow.Next
	// 一定要断开指针
	tmp.Next = nil
	left := sortList(head)
	right := sortList(slow)
	return merge(left, right)
}

func merge(left, right *ListNode) *ListNode {
	var dummy = new(ListNode)
	pointer := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			pointer.Next = left
			left = left.Next
		} else {
			pointer.Next = right
			right = right.Next
		}
		pointer = pointer.Next
	}
	if left != nil {
		pointer.Next = left
	} else {
		pointer.Next = right
	}
	return dummy.Next
}

// sortList2
//
//	@Description: 二刷，自底向上归并
//	@param head
//	@return *ListNode
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针找中点后断开链表
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	back := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(back)
	return merge2(left, right)
}

func merge2(left, right *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	p1, p2 := left, right
	for {
		if p1 == nil {
			p.Next = p2
			return dummy.Next
		}
		if p2 == nil {
			p.Next = p1
			return dummy.Next
		}
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
}
