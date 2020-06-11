package test

//二分查找法
//https://leetcode-cn.com/problems/sqrtx/submissions/
//https://www.cnblogs.com/kyoner/p/11080078.html
func mySqrt(x int) int {
	left := 0
	right := x/2+1
	var mid int
	for left<=right {
		mid = (left+right+1)>>1
		if mid*mid > x {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	return left
}

//反转字符串
//https://leetcode-cn.com/problems/reverse-string/
func reverseString(s []byte)  {
	n := len(s)
	num := n/2
	n = n-1
	for i := 0; i < num; i++ {
		s[i], s[n-i] = s[n-i], s[i]
	}
}

