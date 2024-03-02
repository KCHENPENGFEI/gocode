package main

/**
 * @Author: chenpengfei
 * @Date: 2024/2/22 下午10:58
 * @Desc:
 */

func main() {

}

func longestPalindrome(s string) int {
	cnt := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	// 要标记一下是否存在奇数个数的字母
	flag := false
	for _, value := range m {
		if value%2 == 0 {
			cnt += value
		} else {
			flag = true
			// 将奇数拆分成偶数+1就好了
			cnt += value - 1
		}
	}
	// 如果有奇数的字母，最后再选一个字母就好
	if flag {
		cnt++
	}
	return cnt
}
