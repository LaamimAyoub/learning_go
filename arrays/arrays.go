package arrays

func Sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumAll(arrs ...[]int) []int {
	var res []int
	for _, arr := range arrs {
		res = append(res, Sum(arr))
	}
	return res
}

func SumAllTail(arrs ...[]int) []int {
	var res []int
	for _, arr := range arrs {
		if len(arr) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(arr)-arr[0])

		}

	}
	return res
}
