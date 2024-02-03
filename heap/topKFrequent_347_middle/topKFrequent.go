package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/1/27 下午3:05
 * @Desc:
 */

func main() {
	n := []int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}
	k := 3
	r := topKFrequent(n, k)
	fmt.Println(r)
}

func topKFrequent(nums []int, k int) []int {
	cntMap := make(map[int]int)
	_cntMap := make(map[int][]int)
	var cntSl []int
	for _, num := range nums {
		if _, ok := cntMap[num]; ok {
			cntMap[num]++
		} else {
			cntMap[num] = 1
		}
	}
	for k, v := range cntMap {
		if _, ok := _cntMap[v]; ok {
			_cntMap[v] = append(_cntMap[v], k)
		} else {
			_cntMap[v] = []int{k}
		}
		cntSl = append(cntSl, v)
	}

	var result []int
	buildMaxHeap(cntSl)
	for i := 0; i < k; i++ {
		cntSl[0], cntSl[len(cntSl)-1] = cntSl[len(cntSl)-1], cntSl[0]
		result = append(result, _cntMap[cntSl[len(cntSl)-1]][0])
		_cntMap[cntSl[len(cntSl)-1]] = _cntMap[cntSl[len(cntSl)-1]][1:]
		cntSl = cntSl[:len(cntSl)-1]
		adjust(cntSl, 0)
	}
	return result
}

func buildMaxHeap(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		adjust(nums, i)
	}
}

func adjust(num []int, parent int) {
	leftChild, rightChild := 2*parent+1, 2*parent+2
	idx := parent
	if leftChild < len(num) && num[leftChild] > num[idx] {
		// 调整位置
		idx = leftChild
	}
	if rightChild < len(num) && num[rightChild] > num[idx] {
		idx = rightChild
	}
	// 如果位置有调整
	if idx != parent {
		num[parent], num[idx] = num[idx], num[parent]
		adjust(num, idx)
	}
}
