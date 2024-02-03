package linklist

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/2/27 下午11:28
 * @Desc:
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeLinkList(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := new(ListNode)
	p := head
	for _, l := range list {
		node := &ListNode{
			Val: l,
		}
		p.Next = node
		p = node
	}
	return head.Next
}

func PrintLinkList(list *ListNode) {
	var slice []int
	for {
		if list == nil {
			break
		}
		slice = append(slice, list.Val)
		list = list.Next
	}
	fmt.Println(slice)
}
