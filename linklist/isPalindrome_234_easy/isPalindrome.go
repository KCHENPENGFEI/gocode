package main

import (
	"fmt"
	. "golang/linklist"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/21 下午1:50
 * @Desc: 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

 */

func main() {
	l := MakeLinkList([]int{1, 2, 3, 2, 1})
	fmt.Println(isPalindrome2(l))
}

func isPalindrome(head *ListNode) bool {
	length := getLen(head)
	i, p := 0, head
	for i < length/2 {
		p = p.Next
		i++
	}
	r := reverseList(p)
	p = head
	for r != nil {
		if r.Val != p.Val {
			return false
		}
		r = r.Next
		p = p.Next
	}
	return true
}

// 这里可以使用快慢指针来获取链表中间的位置
func getLen(head *ListNode) int {
	count := 0
	for head != nil {
		count += 1
		head = head.Next
	}
	return count
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// isPalindrome2
//
//	@Description: 二刷
//	@param head
//	@return bool
func isPalindrome2(head *ListNode) bool {
	// 快慢指针找到中点
	// 然后翻转后半链表
	// 对比前后一半的数据是否相同
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	back := slow.Next

	slow.Next = nil
	newBack := reverseList(back)
	for head != nil && newBack != nil {
		if head.Val != newBack.Val {
			return false
		}
		head = head.Next
		newBack = newBack.Next
	}
	return true
}
