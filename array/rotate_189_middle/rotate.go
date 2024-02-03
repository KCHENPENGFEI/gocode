package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/9/11 下午9:30
 * @Desc:
 */

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%p\n", &nums)
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	lenNums := len(nums)
	k = k % lenNums
	part1 := nums[:lenNums-k]
	part2 := nums[lenNums-k:]
	nums = append(part2, part1...) // 在函数内部是无法改变切片的指向关系，只能使用copy
	copy(nums, append(part2, part1...))
}

func rotate1(nums []int, k int) {
	lenNums := len(nums)
	k = k % lenNums
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
