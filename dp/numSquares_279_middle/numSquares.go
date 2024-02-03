package main

import (
	"fmt"
	"math"
)

/**
 * @Author: chenpengfei
 * @Date: 2024/1/28 下午2:21
 * @Desc:
 */

func main() {
	fmt.Println(numSquares(9))
}

func numSquares(n int) int {
	var dp = make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		if i == 1 {
			dp[i] = 1
		} else {
			maxInt := int(math.Sqrt(float64(i)))
			for j := 1; j <= maxInt; j++ {
				dp[i] = minTwo(dp[i], dp[i-j*j]+1)
			}
		}
	}
	return dp[n]
}

func minTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares1(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = minTwo(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}
