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
	m := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	if digits == "" {
		return nil
	}
	var path []byte
	result := new([]string)
	trackback(digits, 0, m, path, result)
	return *result
}

// trackback
//
//	@Description: 通过index来确定每次递归时候的数字位置下标
//	@param digits
//	@param index
//	@param m
//	@param path
//	@param res
func trackback(digits string, index int, m map[byte][]byte, path []byte, res *[]string) {
	if len(path) == len(digits) {
		s := string(path)
		*res = append(*res, s)
		return
	}

	num := digits[index]
	for _, c := range m[num] {
		path = append(path, c)
		trackback(digits, index+1, m, path, res)
		path = path[:len(path)-1]
	}
}
