package offer162

func constructArr(a []int) []int {
	len := len(a)
	answer := make([]int, len+1)

	answer[0] = 1
	for i := 1; i < len; i++ {
		answer[i] = answer[i-1] * a[i-1]
	}

	R := 1
	for i := len - 1; i >= 0; i-- {
		answer[i]=answer[i]*R
		R *=a[i]
	}
	return answer[:len]
}
