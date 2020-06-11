package test

/*
emmm 只需要两行代码
sort.Ints(arr)
return arr[:k]
 */

func GetLeastNumbers(arr []int, k int) []int {
	var res = make([]int, k, k) //一般使用 make 初始化切片
	arrLen := len(arr)

	for index := 0; index < arrLen; index++  {
		if index == k {
			break;
		}
		for index2 := index + 1; index2 < arrLen; index2++ {
			if arr[index2] < arr[index] {
				temp := arr[index2]
				arr[index2] = arr[index]
				arr[index] = temp
			}
		}

		res[index] = arr[index]
	}

	return res
}
