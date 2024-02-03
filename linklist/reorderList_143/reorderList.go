package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/3/4 下午11:51
 * @Desc: 给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

方法二：找到链表的中点位置，然后翻转后半段链表，最后插空插入前半段即可

*/

func main() {
	l := MakeLinkList([]int{1, 2})
	PrintLinkList(l)
	reorderList(l)
	PrintLinkList(l)
}

//
// reorderList
//  @Description:使用递归
//  @param head
//
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	second := head.Next
	p := second
	for p.Next.Next != nil {
		p = p.Next
	}
	head.Next = p.Next
	p.Next = nil
	reorderList(second)
	head.Next.Next = second
}
