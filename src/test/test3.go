package test

func canThreePartsEqualSum(A []int) bool {
	var sum, avgSum int
	for _, value := range A{
		sum += value
	}
	if sum % 3 != 0 {
		return false
	}
	avgSum = sum / 3

	aryCount := len(A) - 1
	var curSum, curCount int
	for key, value := range A{
		curSum += value
		if curSum == avgSum{
			curCount++
			curSum = 0
		}
		if curCount == 2 && key < aryCount{
			return true
		}
	}

	return false
}
