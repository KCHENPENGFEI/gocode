package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/8/27 上午9:58
 * @Desc:
 */

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string]*[]string)
	for _, str := range strs {
		tmp := []byte(str)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		t := string(tmp)
		if slice, ok := hashMap[t]; ok {
			*slice = append(*slice, str)
		} else {
			hashMap[t] = &[]string{str}
		}
	}
	var result [][]string
	for _, v := range hashMap {
		result = append(result, *v)
	}
	return result
}

// groupAnagrams2
//
//	@Description: 二刷，使用[26]byte来作为key，数组是可以比较的
//	@param strs
//	@return [][]string
func groupAnagrams2(strs []string) [][]string {
	var result [][]string
	m := make(map[[26]byte]*[]string)
	for _, str := range strs {
		key := count(str)
		if value, ok := m[key]; ok {
			*value = append(*value, str)
		} else {
			m[key] = &[]string{str}
		}
	}
	for _, value := range m {
		result = append(result, *value)
	}
	return result
}

func count(str string) [26]byte {
	var cnt [26]byte
	for i := 0; i < len(str); i++ {
		cnt[str[i]-'a']++
	}
	return cnt
}
