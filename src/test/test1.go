package test
// https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
// 面试题57 - II. 和为s的连续正数序列

func FindContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	for num := 2; num < target; num++{
		aRes := make([]int, 0)
		split := 0
		tempNum := num
		for tempNum > 0 {
			tempNum--
			split += tempNum
		}
		totalLess := target - split
		if totalLess % num != 0 {
			continue
		}
		begin := (target - split) / num
		if begin <= 0 {
			break
		}
		for ;tempNum < num; tempNum++ {
			aRes = append(aRes, begin + tempNum)
		}
		res = append(res, aRes)
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
