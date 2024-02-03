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
