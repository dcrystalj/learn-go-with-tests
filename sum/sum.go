package sum

func Sum(arr []int) int {
	c := 0
	for _, num := range arr {
		c += num
	}
	return c
}

func SumAll(arrs ...[]int) []int {
	c := make([]int, 0, len(arrs))
	for _, arr := range arrs {
		c = append(c, Sum(arr))
	}
	return c
}

func SumAllTails(arrs ...[]int) []int {
	c := make([]int, 0, len(arrs))
	for _, arr := range arrs {
		if len(arr) == 0 {
			c = append(c, 0)
		} else {
			c = append(c, Sum(arr[1:]))
		}
	}
	return c
}
