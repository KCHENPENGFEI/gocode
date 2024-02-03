package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/2/28 下午11:22
 * @Desc:给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

func main() {
	l1 := MakeLinkList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := MakeLinkList([]int{9, 9, 9, 9})
	result := addTwoNumbers(l1, l2)
	PrintLinkList(result)
}

//
// addTwoNumbers
//  @Description: 常规思路，注意l1和l2为空的时候还要考虑addOne是否为1
//  @param l1
//  @param l2
//  @return *ListNode
//
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	pointer := head

	var addOne int

	for l1 != nil || l2 != nil {
		if l1 != nil && l2 == nil {
			pointer.Next = &ListNode{
				Val: (addOne + l1.Val) % 10,
			}
			addOne = (addOne + l1.Val) / 10
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			pointer.Next = &ListNode{
				Val: (addOne + l2.Val) % 10,
			}
			addOne = (addOne + l2.Val) / 10
			l2 = l2.Next
		} else if l1 != nil && l2 != nil {
			pointer.Next = &ListNode{
				Val: (addOne + l1.Val + l2.Val) % 10,
			}
			addOne = (addOne + l1.Val + l2.Val) / 10
			l1 = l1.Next
			l2 = l2.Next
		}
		pointer = pointer.Next
	}
	if addOne == 1 {
		pointer.Next = &ListNode{
			Val: addOne,
		}
	}
	return head.Next
}
