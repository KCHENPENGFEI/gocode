package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/3/3 下午10:07
 * @Desc: 给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。
* 思路: 构建一个K长度的数组N，分别存储每个链表当前的指针位置，每次循环遍历N找出最小的节点，依次拼接起来
* 时间复杂度是O(k^2n)
* 采用优先队列的方式的可以将复杂度降低为O(kn*log(k))
*/

func main() {
	l1 := MakeLinkList([]int{})
	l2 := MakeLinkList([]int{})
	l3 := MakeLinkList([]int{})
	lists := []*ListNode{l1, l2, l3}
	result := mergeKLists(lists)
	PrintLinkList(result)
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	head := dummy
	// 初始化指针数值
	var arr []*ListNode
	for _, head := range lists {
		arr = append(arr, head)
	}

	for {
		var minValNode *ListNode
		var minValIndex int
		for i, p := range arr {
			old := minValNode
			minValNode = getMinNode(old, p)
			if old != minValNode {
				minValIndex = i
			}
		}
		if minValNode == nil {
			break
		}
		head.Next = minValNode
		arr[minValIndex] = arr[minValIndex].Next
		head = head.Next
	}

	return dummy.Next
}

func getMinNode(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val < b.Val {
		return a
	}
	return b
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
