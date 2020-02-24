package main

// 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
// push(x) -- 将元素 x 推入栈中。
// pop() -- 删除栈顶的元素。
// top() -- 获取栈顶元素。
// getMin() -- 检索栈中的最小元素。
// 示例:
//
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.

type MinStack struct {
	data   []int
	length int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:   make([]int, 0),
		length: 0,
	}
}

func (this *MinStack) Push(x int) {
	this.length++
	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}
	this.length--
	if this.length == 0 {
		this.data = make([]int, 0)
	} else {
		this.data = this.data[:this.length]
	}
	return
}

func (this *MinStack) Top() int {
	if this.length == 0 {
		return 0
	}
	return this.data[this.length-1]
}

func (this *MinStack) GetMin() int {
	if this.length == 0 {
		return 0
	}
	min := 0
	first := true
	for _, _min := range this.data {
		if first {
			min = _min
			first = false
			continue
		}
		if min > _min {
			min = _min
		}
	}
	return min
}

func main() {
	x := Constructor()
}
