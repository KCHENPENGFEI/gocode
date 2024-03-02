package main

import (
	"fmt"
	"math"
)

/**
 * @Author: chenpengfei
 * @Date: 2024/2/26 上午1:09
 * @Desc:
 */

func main() {
	x := -12
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		// 关键在于不能直接使用int32最大最小进行比较
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		d := x % 10
		x = x / 10
		res = res*10 + d
	}
	return res
}
