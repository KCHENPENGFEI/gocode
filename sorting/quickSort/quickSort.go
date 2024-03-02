package main

import (
	"fmt"
)

/**
 * @Author: chenpengfei
 * @Date: 2024/2/27 下午10:11
 * @Desc:
 */

func main() {
	arr := []int{4, 1, 6, 5, 7, 2, 3}
	Sort(arr, 0, 6)
	fmt.Println(arr)
}

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	privot := arr[i] //privot就是我们单趟选择的分界值，一般我们是选择左右边界值，不选择中间值
	for i < j {
		//每次找到大于key或者是小于key的值就将 i, j 对应的值进行交换
		for i < j && arr[j] >= privot {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] < privot {
			i++
		}
		arr[j] = arr[i]
	}
	//当for循环退出时，此时i的位置就是key值在排序后应该在的位置
	// 在循环的过程中会在数组中丢失privot的值，因此最后需要重新放回去
	arr[i] = privot
	QuickSort(arr, left, i-1)  //递归将key左边的数组进行排序
	QuickSort(arr, i+1, right) ////递归将key右边的数组进行排序
}

func Sort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	key := nums[i]
	for i < j {
		for i < j && nums[j] >= key {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] < key {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = key
	Sort(nums, left, i-1)
	Sort(nums, i+1, right)
}
