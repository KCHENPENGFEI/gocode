package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/8/6 下午3:52
 * @Desc:
 */

type pair struct {
	min int
	max int
}

func main() {
	nums := []int{100, 3, 2, 1}
	fmt.Println(maximumGap(nums))
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	max, min := getMaxMin(nums)
	bucketSize := (max - min) / (len(nums) - 1)
	if bucketSize == 0 {
		bucketSize = 1
	}

	bucketNum := (max-min)/bucketSize + 1
	bucket := make([]*pair, bucketNum)

	for _, num := range nums {
		index := (num - min) / bucketSize
		if bucket[index] == nil {
			bucket[index] = &pair{
				min: num,
				max: num,
			}
		} else {
			if num < bucket[index].min {
				bucket[index].min = num
			}
			if num > bucket[index].max {
				bucket[index].max = num
			}
		}
	}

	i, j := 0, 1
	interval := 0
	for i < j && j < bucketNum {
		if bucket[j] == nil {
			j++
			continue
		}
		if bucket[i] == nil {
			i++
			continue
		}
		sub := bucket[j].min - bucket[i].max
		if sub > interval {
			interval = sub
		}
		i++
		j++
	}
	return interval
}

func getMaxMin(nums []int) (int, int) {
	max, min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max, min
}
