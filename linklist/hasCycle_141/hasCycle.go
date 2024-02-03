package main

import (
	"fmt"
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/3/1 上午12:03
 * @Desc: 给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
*/

func main() {
	l1 := MakeLinkList([]int{3, 2, 0, -4})
	fmt.Println(hasCycle(l1))
	l1.Next.Next.Next.Next = l1.Next.Next
	fmt.Println(hasCycle(l1))
}

//
// hasCycle
//  @Description: 使用快慢指针，开始时均指向第一个节点，然后fast每次跳两步，slow每次跳一步
// 当fast==nil或者fast.Next==nil说明链表没有环；链表有环则fast和slow肯定相遇
//  @param head
//  @return bool
//
func hasCycle(head *ListNode) bool {
	var fast, slow = head, head
	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
}
