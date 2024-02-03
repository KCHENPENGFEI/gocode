package main

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午6:39
 * @Desc: 多数元素：给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
摩尔投票法思路
候选人(cand_num)初始化为 nums[0]，票数 count 初始化为 1。
当遇到与 cand_num 相同的数，则票数 count = count + 1，否则票数 count = count - 1。
当票数 count 为 0 时，更换候选人，并将票数 count 重置为 1。
遍历完数组后，cand_num 即为最终答案。

为何这行得通呢？
投票法是遇到相同的则 票数 + 1，遇到不同的则 票数 - 1。
且“多数元素”的个数 > ⌊ n/2 ⌋，其余元素的个数总和 <= ⌊ n/2 ⌋。
因此“多数元素”的个数 - 其余元素的个数总和 的结果 肯定 >= 1。
这就相当于每个 “多数元素” 和其他元素 两两相互抵消，抵消到最后肯定还剩余 至少1个 “多数元素”。
*/

// majorityElement
//
//	@Description: 空间复杂度是O(1)
//	@param nums
//	@return int
func majorityElement(nums []int) int {
	candidate := nums[0]
	cnt := 0
	for _, num := range nums {
		if candidate == num {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			candidate = num
			cnt = 1
		}
	}
	return candidate
}
