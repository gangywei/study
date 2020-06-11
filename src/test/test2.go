package test
// https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/submissions/
// 面试题59 - II. 队列的最大值
// 当队列后来加入一个大的值之后，前面比较小的值就不是最大值了，将 max 双向队列后小于新加入数据的元素删除

type MaxQueue struct {
	queue []int
	max []int
}

func Constructor() MaxQueue {
	return MaxQueue{make([]int, 0), make([]int, 0)}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int)  {
	this.queue = append(this.queue, value)
	for len(this.max)  != 0 && value > this.max[len(this.max)-1] {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	max := this.queue[0]
	this.queue = this.queue[1:]
	if this.max[0] == max {
		this.max = this.max[1:]
	}

	return max
}
