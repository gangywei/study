package test

import (
	"fmt"
	"math"
	"strconv"
)

//https://leetcode-cn.com/problems/valid-parentheses/
//有效的括号
func isValid(s string) bool {
	//使用切片维护栈 stack 队列 queue
	var stack []string
	// 括号映射表
	frontBracket := map[string]string{ ")":"(", "]":"[", "}":"{" }

	for _, x := range s {
		if x=='(' || x=='[' || x=='{' {     // 遇到前括号，入栈
			stack = append(stack,string(x))
		} else if x==')' || x==']' || x=='}' {    // 遇到后括号，判断
			if len(stack)!=0 && stack[len(stack)-1] == frontBracket[string(x)] { // 栈非空，和栈顶元素匹配，匹配成功，出栈
				stack = stack[0:len(stack)-1]
			} else {    // 栈空或者匹配失败，返回错误
				return false
			}
		}
	}
	if len(stack)==0 {
		return true
	} else {
		return false
	}
}

//逆波兰表达式求值
//https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	numbers := []int{}
	numLen := -1
	for _,value := range tokens {
		switch value {
		case "+":
			numbers[numLen-1] = numbers[numLen-1] + numbers[numLen]
			numLen--
		case "-":
			numbers[numLen-1] = numbers[numLen-1] - numbers[numLen]
			numLen--
		case "*":
			numbers[numLen-1] = numbers[numLen-1] * numbers[numLen]
			numLen--
		case "/":
			numbers[numLen-1] = numbers[numLen-1] / numbers[numLen]
			numLen--
		default:
			intValue,_ := strconv.Atoi(value)
			if len(numbers) > numLen {
				numbers[numLen+1] = intValue
			} else {
				numbers = append(numbers[:numLen], intValue)
			}
			numLen++
		}
		fmt.Print(numLen)
		fmt.Println(numbers)
	}
	return numbers[0]
}

func evalRPN2(tokens []string) int {
	number := []int{}
	for _, val := range tokens{
		l := len(number)
		switch val {
		case "+":
			number  = append(number[:l -2], number[l-2] + number[l-1])
		case "-":
			number  = append(number[:l -2], number[l-2] - number[l-1])
		case "*":
			number  = append(number[:l -2], number[l-2] * number[l-1])
		case "/":
			number  = append(number[:l -2], number[l-2] / number[l-1])
		default:
			num, _ := strconv.Atoi(val)
			number  = append(number, num)
		}
	}
	return number[0]
}

//https://leetcode-cn.com/problems/design-circular-deque/submissions/
//设计一个双端队列
type MyCircularDeque struct {
	head  *Node
	tail  *Node
	len   int
	count int
}

type Node struct {
	Next *Node
	Pre  *Node
	Val  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor2(k int) MyCircularDeque {
	head := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	tail := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	head.Next = &tail
	tail.Pre = &head
	deque := MyCircularDeque{
		head:  &head,
		tail:  &tail,
		len:   k,
		count: 0,
	}

	return deque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	temp := this.head.Next
	tempNode := Node{
		Next: temp,
		Pre:  this.head,
		Val:  value,
	}
	this.head.Next = &tempNode
	temp.Pre = &tempNode
	this.count++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	temp := this.tail.Pre
	tempNode := Node{
		Next: this.tail,
		Pre:  temp,
		Val:  value,
	}
	this.tail.Pre = &tempNode
	temp.Next = &tempNode
	this.count++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	deleteTemp := this.head.Next
	this.head.Next = deleteTemp.Next
	deleteTemp.Next.Pre = this.head
	deleteTemp.Next, deleteTemp.Pre = nil, nil
	this.count--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	deleteTemp := this.tail.Pre
	deleteTemp.Pre.Next = this.tail
	this.tail.Pre = deleteTemp.Pre
	deleteTemp.Next, deleteTemp.Pre = nil, nil
	this.count--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	return this.head.Next.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	return this.tail.Pre.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.count == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.count
}

//滑动窗口最大值
//https://leetcode-cn.com/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil{
		return []int{}
	}

	numLen := len(nums)
	var maxValue int
	var nowValue int
	var maxIndex int = math.MinInt64 //将上一个最大值的下标缓存起来
	var res []int = make([]int, 0, numLen-k+1)

	for i := k-1; i < numLen; i++ {
		//先使用缓存的 index 进行一步判断
		if maxIndex <= i && maxIndex >= i-k+1 {
			if nums[maxIndex] <= nums[i] {
				res = append(res, nums[i])
				maxIndex = i
			} else {
				res = append(res, nums[maxIndex])
			}
			continue
		}
		maxValue = math.MinInt64
		for j := k-1; j >= 0; j-- {
			nowValue = nums[i-j]
			if nowValue >= maxValue {
				maxIndex = i-j
				maxValue = nowValue
			}
		}
		res = append(res, maxValue)
	}

	return res
}
