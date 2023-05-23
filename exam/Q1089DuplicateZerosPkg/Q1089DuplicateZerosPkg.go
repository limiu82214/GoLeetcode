package Q1089DuplicateZerosPkg

func Solve(arr []int) {
	i := 0
	j := len(arr)
	for ; i < len(arr); i++ {
		if i >= j {
			break
		}
		isZero := arr[i] == 0
		if j < len(arr) {
			arr[i], arr[j] = arr[j], arr[i]
			for k := j; k < len(arr)-1; k++ {
				arr[k], arr[k+1] = arr[k+1], arr[k]
			}
		}

		if isZero {
			j--
			if i >= j {
				break
			}
			arr[j] = 0
			for k := j; k < len(arr)-1; k++ {
				arr[k], arr[k+1] = arr[k+1], arr[k]
			}
		}
	}
	// fmt.Println(i,j,arr)
}
