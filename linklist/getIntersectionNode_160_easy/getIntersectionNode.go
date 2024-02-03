package main

import (
	"fmt"
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/21 下午12:37
 * @Desc: 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
 */

func main() {
	l1 := MakeLinkList([]int{1, 2, 3, 4, 5})
	l2 := MakeLinkList([]int{1, 333, 4, 5})
	l2.Next.Next = l1.Next.Next.Next
	result := getIntersectionNode1(l1, l2)
	if result == nil {
		fmt.Println(nil)
	} else {
		fmt.Println(result.Val)
	}
}

// 利用set存储去重
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	pA, pB := headA, headB
	for pA != nil {
		m[pA] = struct{}{}
		pA = pA.Next
	}

	for pB != nil {
		if _, ok := m[pB]; ok {
			return pB
		}
		pB = pB.Next
	}
	return nil
}

// 双指针，当pA为空时，另其指向headB，pB同理
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
