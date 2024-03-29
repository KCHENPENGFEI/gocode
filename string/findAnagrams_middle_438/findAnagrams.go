package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/9/8 上午11:26
 * @Desc:
 */

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	var ans = make([]int, 0)
	lenP := len(p)
	if lenP > len(s) {
		return ans
	}
	mp := buildMap(0, len(p), p)
	ms := make(map[byte]int)
	for i := 0; i <= len(s)-lenP; i++ {
		if i == 0 {
			// 装载数据
			ms = buildMap(i, len(p), s)
		} else {
			ms[s[i-1]]--
			if ms[s[i-1]] <= 0 {
				delete(ms, s[i-1])
			}
			ms[s[i+lenP-1]]++
		}

		if reflect.DeepEqual(mp, ms) {
			ans = append(ans, i)
		}
	}
	return ans
}

func buildMap(start, end int, s string) map[byte]int {
	var m = make(map[byte]int)
	for i := start; i < end; i++ {
		m[s[i]] += 1
	}
	return m
}

func findAnagrams2(s string, p string) []int {
	var result []int
	lenS, lenP := len(s), len(p)
	if lenS < lenP {
		return nil
	}
	var sCount, pCount [26]int
	// 初始化sCount和pCount
	for i := 0; i < lenP; i++ {
		sCount[s[i]-'a']++
		pCount[p[i]-'a']++
	}
	if sCount == pCount {
		result = append(result, 0)
	}

	for i := 0; i < lenS-lenP; i++ {
		sCount[s[i]-'a']--
		sCount[s[i+lenP]-'a']++
		if sCount == pCount {
			result = append(result, i+1)
		}
	}
	return result
}
