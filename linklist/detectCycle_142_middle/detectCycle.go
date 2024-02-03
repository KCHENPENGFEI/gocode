package main

import (
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/21 下午2:41
 * @Desc: 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
 */

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	// 第一次fast和slow相遇时，fast走了2n个环形距离，slow走了n个环形距离
	for {
		if fast == nil || fast.Next == nil {
			// 不存在环
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			// 第一次相遇
			break
		}
	}
	// 此时另fast回到head处，那么当其和slow再次相遇时候就是在环形的入口处了
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
