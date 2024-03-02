package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/21 下午5:50
 * @Desc:
 */

func main() {
	l := MakeLinkList([]int{1, 3, 2, 8, 5, 10, 4})
	PrintLinkList(sortList(l))
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1, p2 := list1, list2
	pre := &ListNode{}
	cur := pre
	for p1 != nil || p2 != nil {
		if p1 == nil {
			cur.Next = p2
			return pre.Next
		}
		if p2 == nil {
			cur.Next = p1
			return pre.Next
		}
		if p1.Val < p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}
		cur = cur.Next
	}
	return pre.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return mergeTwoLists(sort(head, mid), sort(mid, tail))
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
