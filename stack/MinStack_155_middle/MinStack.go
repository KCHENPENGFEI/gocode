package main

/**
 * @Author: chenpengfei
 * @Date: 2024/1/13 下午2:59
 * @Desc:
 */

func main() {

}

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		if val < this.minStack[len(this.minStack)-1] {
			this.minStack = append(this.minStack, val)
		} else {
			this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
	if len(this.minStack) > 0 {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) > 0 {
		return this.minStack[len(this.minStack)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
