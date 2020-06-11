package test

import (
	"sort"
)

func MajorityElement(nums []int) int {
	midIndex := len(nums) / 2
	sort.Ints(nums)
	return nums[midIndex]
}

//https://leetcode-cn.com/problems/first-missing-positive/
//要求时间为O(n)空间为n
//我在做该题目时，因为空间的原因是用了交换，但是开辟个空间更容易一些
func firstMissingPositive(nums []int) int {
	maxVal := len(nums)
	if maxVal == 0 {
		return 1
	}
	if nums[0] == 1 && maxVal == 1 {
		return 2
	}
	countSort := make([]int, maxVal+1)
	for n, _ := range nums {
		if nums[n] <= 0 || nums[n] > maxVal {
			continue
		}
		countSort[nums[n]]++
	}
	var i int
	for i = 1; i <= maxVal; i++ {
		if countSort[i] == 0 {
			return i
		}
	}
	return i
}

//https://leetcode-cn.com/problems/3sum/
//使用双指针解法会更好
//解法链接：https://leetcode-cn.com/problems/3sum/solution/san-shu-zhi-he-cshi-xian-shuang-zhi-zhen-fa-tu-shi/
func threeSum(nums []int) [][]int {
	var res [][]int
	testInt := make(map[int]bool)
	resValue := make(map[[3]int]bool)
	sort.Ints(nums)
	aryLength := len(nums)
	for key,value := range nums {
		if _,ok := testInt[value]; ok == false {
			testInt[value] = true
		} else {
			continue;
		}
		for key1 := key+1; key1 < aryLength; key1++ {
			loop := true
			for key2 := key1+1; key2 < aryLength && loop == true; key2++ {
				addSum := value + nums[key1] + nums[key2]
				if addSum == 0 {
					thrRes := [3]int{value, nums[key1], nums[key2]}
					if _,ok := resValue[thrRes]; ok == false {
						res = append(res, []int{value, nums[key1], nums[key2]})
						resValue[thrRes] = true
					}

				} else if addSum > 0 {
					loop = false
				}
			}
		}
	}

	return res
}