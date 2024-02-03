package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/8/19 下午2:25
 * @Desc:
 */

func main() {
	bills := []int{5, 5, 5, 10, 20, 20}
	fmt.Println(lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	ownMoney := map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}
	for _, money := range bills {
		if !updateMoney(ownMoney, money) {
			return false
		}
	}
	return true
}

func updateMoney(myMoney map[int]int, bill int) bool {
	myMoney[bill]++
	cash := bill - 5

	moneyType := []int{20, 10, 5}
	if cash > 0 {
		// 贪心优先使用大面额
		for _, money := range moneyType {
			for myMoney[money] > 0 && cash-money >= 0 {
				myMoney[money]--
				cash -= money
			}
		}
	} else {
		return true
	}

	return !(cash > 0)
}
