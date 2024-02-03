package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/2/27 下午11:29
 * @Desc:
 */

//
// main
//  @Description: 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
func main() {
	l1 := MakeLinkList([]int{1, 2, 4, 5, 8})
	l2 := MakeLinkList([]int{1, 3, 4, 6, 7})
	result := mergeTwoLists(l1, l2)
	PrintLinkList(result)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := new(ListNode)
	p := head
	p1 := list1
	p2 := list2
	for p1 != nil || p2 != nil {
		if p1 == nil {
			p.Next = p2
			return head.Next
		}
		if p2 == nil {
			p.Next = p1
			return head.Next
		}
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	return head.Next
}
