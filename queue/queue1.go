package main

import (
	"fmt"
)

/*
设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。

循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。

你的实现应该支持如下操作：

MyCircularQueue(k): 构造器，设置队列长度为 k 。
Front: 从队首获取元素。如果队列为空，返回 -1 。
Rear: 获取队尾元素。如果队列为空，返回 -1 。
enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
isEmpty(): 检查循环队列是否为空。
isFull(): 检查循环队列是否已满。

示例：
MyCircularQueue circularQueue = new MycircularQueue(3); // 设置长度为 3
circularQueue.enQueue(1);  // 返回 true
circularQueue.enQueue(2);  // 返回 true
circularQueue.enQueue(3);  // 返回 true
circularQueue.enQueue(4);  // 返回 false，队列已满
circularQueue.Rear();  // 返回 3
circularQueue.isFull();  // 返回 true
circularQueue.deQueue();  // 返回 true
circularQueue.enQueue(4);  // 返回 true
circularQueue.Rear();  // 返回 4
*/

type MyCircularQueue struct {
	currLength  int
	totalLength int
	frontIndex  int
	data        []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		currLength:  0,
		totalLength: k,
		frontIndex:  0,
		data:        make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.currLength < this.totalLength {
		this.currLength += 1
		this.frontIndex = (this.frontIndex + 1) % this.totalLength
		this.data[this.frontIndex] = value
		return true
	}
	return false
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.currLength > 0 {
		this.currLength -= 1
		return true
	}
	return false
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.currLength == 0 {
		return -1
	}
	_rearIndex := this.frontIndex - (this.currLength - 1)
	fmt.Println(this.frontIndex, "  ", this.currLength, "  ", _rearIndex)
	if _rearIndex < 0 {
		_rearIndex = _rearIndex + this.totalLength
	}
	return this.data[_rearIndex]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.currLength == 0 {
		return -1
	}
	return this.data[this.frontIndex]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.currLength == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.currLength == this.totalLength
}

func main() {

	obj := Constructor(3)
	for i := 0; i < 3; i++ {
		fmt.Println(obj.EnQueue(i))
		fmt.Println(obj.Front())
		fmt.Println(obj.Rear())
		fmt.Println(obj.IsFull())
		fmt.Println("=======")
	}

}
