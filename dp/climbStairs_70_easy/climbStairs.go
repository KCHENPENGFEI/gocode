package main

/**
 * @Author: chenpengfei
 * @Date: 2024/1/28 下午1:23
 * @Desc:
 */

func main() {

}

func climbStairs(n int) int {
	var sli = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if i == 1 {
			sli[i] = 1
		} else if i == 2 {
			sli[i] = 2
		} else {
			sli[i] = sli[i-1] + sli[i-2]
		}
	}
	return sli[n]
}
