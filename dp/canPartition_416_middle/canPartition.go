package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午3:32
 * @Desc:
 */

func main() {
	nums := []int{9, 1, 2, 4, 10}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	half := sum / 2
	// 找到一个子序列使得和等于half，0-1背包问题
	var dp = make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, half+1)
		for j := 0; j < len(dp[0]); j++ {
			if i == 0 {
				dp[i][j] = false
				continue
			}
			if j == 0 {
				dp[i][j] = true
				continue
			}
			if j-nums[i-1] < 0 {
				dp[i][j] = false
			} else {
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			}
		}

	}
	return dp[len(nums)][half]
}
