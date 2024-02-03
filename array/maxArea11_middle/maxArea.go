package main

/**
 * @Author: chenpengfei
 * @Date: 2022/9/11 4:11 PM
 * @Desc: 给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。

找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。
*/

func main() {

}

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxVolumn := (j - i) * min(height[i], height[j])
	for i < j {
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		maxVolumn = max(maxVolumn, (j-i)*min(height[i], height[j]))
	}
	return maxVolumn
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
