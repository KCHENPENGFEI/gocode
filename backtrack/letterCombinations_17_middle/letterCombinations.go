package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/17 下午5:30
 * @Desc:
 */

func main() {
	d := ""
	r := letterCombinations(d)
	fmt.Println(r)
}

func letterCombinations(digits string) []string {
	lettersMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var (
		res   []string
		track []byte
	)
	if digits == "" {
		return nil
	}
	dfs(&res, 0, len(digits), track, digits, lettersMap)
	return res

}

func dfs(result *[]string, start, n int, track []byte, digits string, lettersMap map[byte]string) {
	if len(track) == n {
		*result = append(*result, string(track))
		return
	}
	// 两次循环，外层循环迭代数字，内层循环迭代字母
	for i := start; i < n; i++ {
		letters := lettersMap[digits[i]]
		for j := 0; j < len(letters); j++ {
			track = append(track, letters[j])
			dfs(result, i+1, n, track, digits, lettersMap)
			track = track[:len(track)-1]
		}
	}
}
