package main

import (
	"fmt"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/11/4 上午10:35
 * @Desc:
 */

func main() {
	m := [][]int{{1}}
	t := 1
	fmt.Println(searchMatrix2(m, t))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	line1, found := findMaxLine(matrix, target)
	if found {
		return true
	}
	line2, found := findMinLine(matrix, target)
	if found {
		return true
	}
	if line1 == -1 || line2 == -1 {
		return false
	}
	for i := line2; i <= line1; i++ {
		if search(matrix[i], target) {
			return true
		}
	}
	return false
}

// findLine
//
//	@Description: 找到target可能存在的行数
//	@param matrix
//	@param target
//	@return int
func findLine(matrix [][]int, target int) int {
	if target < matrix[0][0] {
		return -1
	}
	if target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return -1
	}
	l, r := 0, len(matrix)-1
	for l <= r {
		mid := l + (r-l)/2
		num := matrix[mid][0]
		if num == target {
			return mid
		}
		if num > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func findMaxLine(matrix [][]int, target int) (int, bool) {
	if target < matrix[0][0] {
		return -1, false
	}
	l, r := 0, len(matrix)-1
	for l < r {
		mid := l + (r-l+1)/2
		if matrix[mid][0] == target {
			return mid, true
		}
		if matrix[mid][0] > target {
			r = mid - 1
		} else {
			l = mid
		}
	}

	return l, matrix[l][0] == target
}

func findMinLine(matrix [][]int, target int) (int, bool) {
	m := len(matrix)
	n := len(matrix[0])
	if target > matrix[m-1][n-1] {
		return -1, false
	}
	l, r := 0, m-1
	for l < r {
		mid := l + (r-l)/2
		if matrix[mid][n-1] == target {
			return mid, true
		}
		if matrix[mid][n-1] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r, matrix[r][n-1] == target
}

func search(line []int, target int) bool {
	l, r := 0, len(line)-1
	for l <= r {
		mid := l + (r-l)/2
		if target == line[mid] {
			return true
		}
		if target > line[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

// searchMatrix2
//
//	@Description: 二刷
//	@param matrix
//	@param target
//	@return bool
func searchMatrix2(matrix [][]int, target int) bool {
	// 二分找到第一列小于等于target的最后一行
	row := findRow(matrix, target)
	fmt.Println(row)
	fmt.Println(matrix[row][len(matrix[0])-1])
	if row == -1 || matrix[row][len(matrix[0])-1] < target {
		return false
	}
	// 继续二分查找target
	return findTarget(matrix[row], target)
}

func findRow(matrix [][]int, target int) int {
	l, r := 0, len(matrix)-1
	for l < r {
		mid := l + (r-l+1)/2
		if matrix[mid][0] > target {
			r = mid - 1
		} else {
			l = mid
		}
	}
	if matrix[l][0] > target {
		return -1
	}
	return l
}

func findTarget(row []int, target int) bool {
	l, r := 0, len(row)-1
	for l < r {
		mid := l + (r-l)/2
		if row[mid] < target {
			l = mid + 1
		} else if row[mid] == target {
			return true
		} else {
			r = mid
		}
	}
	return row[l] == target
}
