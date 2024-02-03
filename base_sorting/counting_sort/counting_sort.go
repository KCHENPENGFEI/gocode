package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/8/6 下午2:37
 * @Desc: 计数排序：遍历数组，将遍历到的数字累加到中间数组对应位上，然后遍历中间数组将数组还原
 */

func main() {
	nums := []int{10, 60, 1, 765, 36, 78, 13, 9, 567, 79}
	fmt.Println(countingSort(nums))
}

func maximum(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func countingSort(nums []int) []int {
	max := maximum(nums)
	countingArr := make([]int, max+1)
	for _, num := range nums {
		countingArr[num]++
	}

	var result []int
	for index, i := range countingArr {
		for j := 0; j < i; j++ {
			result = append(result, index)
		}
	}
	return result
}
