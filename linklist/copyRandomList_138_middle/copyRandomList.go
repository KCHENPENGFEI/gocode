package main

/**
 * @Author: chenpengfei
 * @Date: 2023/12/10 下午5:12
 * @Desc:
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func main() {

}

func copyRandomList(head *Node) *Node {
	// 建立哈希表维护老节点到新节点的映射关系
	var hashMap = make(map[*Node]*Node)
	var preNode = &Node{}
	dummy := preNode
	cur := head
	for cur != nil {
		copyNode := &Node{
			Val: cur.Val,
		}
		hashMap[cur] = copyNode
		preNode.Next = copyNode
		preNode = preNode.Next
		cur = cur.Next
	}
	for k, v := range hashMap {
		r := k.Random
		if r != nil {
			v.Random = hashMap[r]
		}
	}
	return dummy.Next
}
