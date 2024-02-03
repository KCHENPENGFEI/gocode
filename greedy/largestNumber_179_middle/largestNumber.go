package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/8/19 下午4:52
 * @Desc:
 */

func main() {
	//nums := []int{5, 4, 3, 781, 782, 74, 69, 333, 45, 37, 639}
	nums := []int{2, 10}
	fmt.Println(largestNumber1(nums))
	//a := 1001
	//fmt.Println(getHighNum(a, 4))
}

func largestNumber1(nums []int) string {
	// 排序
	// 对于最高位数字相同的的情况，会有不同的场景
	// [4,42]可以组成最大数442
	// [4,45]则可以组成最大数454
	// 因此排序函数中需要计算出不同顺序情况下的值大小
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return x*sy+y > y*sx+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

func largestNumber(nums []int) string {
	return getLargestNum(nums, 0)
}

func getLargestNum(nums []int, bit int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if ok, result := same(nums); ok {
		return result
	}
	rec := make([][]int, 10)
	for _, num := range nums {
		highNum := getHighNum(num, bit)
		rec[highNum] = append(rec[highNum], num)
	}
	r := ""
	for i := 9; i >= 0; i-- {
		if rec[i] != nil {
			s := getLargestNum(rec[i], bit+1)
			r += s
		}
	}
	return r
}

func same(nums []int) (bool, string) {
	m := make(map[rune]int)
	var result = make([]rune, 0)
	for _, num := range nums {
		strN := strconv.Itoa(num)
		for _, r := range strN {
			m[r] = 1
			result = append(result, r)
		}
	}
	return len(m) == 1, string(result)
}

func getHighNum(n int, bit int) int {
	strN := strconv.Itoa(n)
	if bit >= len(strN) {
		bit = len(strN) - 1
	}
	i := strN[bit] - '0'
	return int(i)
}
