package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/8/6 下午2:25
 * @Desc:
 */

func main() {
	nums := []int{10, 60, 1, 765, 36, 78, 13, 9, 567, 79}
	fmt.Println(RadixSort(nums))
}

func howManyBit(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func pow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
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

func RadixSort(nums []int) []int {
	numberBit := howManyBit(maximum(nums))
	for i := 0; i < numberBit; i++ {
		rec := make([][]int, 10)
		for _, num := range nums {
			rec[(num/pow10(i))%10] = append(rec[(num/pow10(i))%10], num)
		}
		numsCopy := make([]int, 0)
		for j := 0; j < 10; j++ {
			numsCopy = append(numsCopy, rec[j]...)
		}
		nums = numsCopy
	}
	return nums
}
