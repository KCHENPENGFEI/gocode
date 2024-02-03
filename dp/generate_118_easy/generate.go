package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/1/28 下午1:25
 * @Desc:
 */

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	var result = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			result[i] = []int{1}
		} else if i == 1 {
			result[i] = []int{1, 1}
		} else {
			var temp = make([]int, i+1)
			temp[0], temp[len(temp)-1] = 1, 1
			for j := 1; j < len(temp)-1; j++ {
				temp[j] = result[i-1][j-1] + result[i-1][j]
			}
			result[i] = temp
		}
	}
	return result
}
