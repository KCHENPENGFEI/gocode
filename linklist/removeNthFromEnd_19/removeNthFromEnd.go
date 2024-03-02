package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/3/3 下午9:53
 * @Desc: 删除倒数第n个节点
 * 先正序遍历统计链表的长度
 * 然后根据n和长度的关系删除指定节点
 * 使用了两次遍历
 *
 * 做法2: 只使用一趟遍历
 * 使用两根指针指向dummy，其中一根指针先向前移动n步，然后两根指针同时向后移动
 */

func main() {
	p := MakeLinkList([]int{1, 2, 5})
	ret := removeNthFromEnd(p, 2)
	PrintLinkList(ret)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	p1, p2 := dummy, dummy
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return dummy.Next
}

// removeNthFromEnd2
//
//	@Description: 二刷
//	@param head
//	@param n
//	@return *ListNode
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	p := head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	pp := dummy
	for p != nil {
		pp = pp.Next
		p = p.Next
	}
	// 删除pp的下一个节点
	pp.Next = pp.Next.Next
	return dummy.Next
}
