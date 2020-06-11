package test

import (
	"sort"
)

//https://leetcode-cn.com/problems/two-sum/
//两数之和
func twoSum(nums []int, target int) []int {
	var res []int
	var numsLen = len(nums)

	for i, v := range nums{	//这里多了很多复制操作，反而会慢些
		for j:=i+1; j<numsLen; j++ {
			if v + nums[j] == target {
				res = append(res, i, j)
			}
		}
	}

	/*
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	 */
	
	return res
}

//使用 map 存储存在的数字，减少遍历操作
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

//https://leetcode-cn.com/problems/add-two-numbers/
//两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	node1 := l1
	node2 := l2
	l3 := &ListNode{} //比使用 new 函数效率要高些
	node3 := l3
	var a,b,c,sum int

	for node1 != nil || node2 != nil || a>0 {
		node3.Next = &ListNode{}
		node3 = node3.Next
		b,c = 0,0
		if node1 != nil {
			b = node1.Val
		}
		if node2 != nil {
			c = node2.Val
		}
		sum = a+b+c
		node3.Val = sum % 10
		a = sum / 10
		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
	}

	return l3.Next
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	numsLen := len(nums)-1
	for i:=0; i<numsLen; i+=2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[numsLen]
}

func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	arrLen := len(nums)
	for i:=0; i<arrLen; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return 0
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for row, col := len(matrix) - 1, 0; row >= 0 && col < len(matrix[0]);{
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			col++
		} else {
			row--
		}
	}

	return false
}

func reversePrint(head *ListNode) []int {
	var result []int
	//遍历存储val至数组
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	//首尾交换位置，相当于反转后输出
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}

//https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/submissions/
//使用两个栈实现一个队列
type stack []int

func (s *stack) Push(value int) {
	*s = append(*s, value)
}

func (s *stack) Pop() int {
	n := len(*s)
	res := (*s)[n-1]
	*s = (*s)[:n-1]
	return res
}

type CQueue struct {
	in stack
	out stack
}

func Constructor1() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *CQueue) DeleteHead() int  {
	if len(this.out) != 0 {
		return this.out.Pop()
	} else if len(this.in) != 0 {
		for len(this.in) != 0 {
			this.out.Push(this.in.Pop())
		}
		return this.out.Pop()
	}
	return -1
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	pre := 0
	cur, res := 1, 1
	for i:=2; i<=n; i++ {
		res = (cur + pre) % 1000000007
		pre = cur
		cur = res
	}

	return res
}