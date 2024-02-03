package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/8/19 下午12:21
 * @Desc:
 */

func main() {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	minus := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		minus[i] = gas[i] - cost[i]
	}

	for i, m := range minus {
		if m < 0 {
			continue
		}
		sum := 0
		success := true
		for j := i; j < len(minus)+i; j++ {
			var index int
			if j >= len(minus) {
				index = j - len(minus)
			} else {
				index = j
			}
			sum += minus[index]
			if sum < 0 {
				success = false
				break
			}
		}
		if success == true {
			return i
		}
	}
	return -1
}

func canCompleteCircuit2(gas []int, cost []int) int {
	totalNum := 0
	curNum := 0
	idx := 0
	for i := 0; i < len(gas); i++ {
		curNum += gas[i] - cost[i]
		totalNum += gas[i] - cost[i]
		if curNum < 0 {
			curNum = 0
			idx = (i + 1) % len(gas)
		}
	}
	if totalNum < 0 {
		return -1
	}
	return idx
}
