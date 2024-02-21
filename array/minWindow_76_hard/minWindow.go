package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/17 下午2:47
 * @Desc:
 */

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	r := minWindow(s, t)
	fmt.Println(r)
}

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]] += 1
	}
	needCnt := len(t)
	i := 0
	index := []int{0, len(s) + 1}
	for j := 0; j < len(s); j++ {
		// 先向右扩展j指针
		if _, ok := need[s[j]]; ok {
			// 找到一个
			if need[s[j]] > 0 {
				needCnt--
			}
			need[s[j]]--
		}
		if needCnt == 0 {
			// i,j区间已经包含了所有元素
			// 开始增大i
			for {
				if v, ok := need[s[i]]; ok && v == 0 {
					break
				} else if ok {
					need[s[i]]++
				}
				i++
			}
			//for v, ok := need[s[i]]; !ok || v != 0; i++ {
			//	fmt.Println(s[i])
			//	fmt.Println(need[s[i]])
			//	fmt.Println(v)
			//	// 说明有多余字符
			//	if ok {
			//		need[s[i]]++
			//	}
			//}
			// 记录i,j的结果
			if j-i < index[1]-index[0] {
				index[0], index[1] = i, j
			}
			// i向后扩展一个位置，重新寻找结果
			need[s[i]]++
			needCnt++
			i++
		}
	}
	if index[1]-index[0] > len(s) {
		return ""
	}
	return s[index[0] : index[1]+1]
}
