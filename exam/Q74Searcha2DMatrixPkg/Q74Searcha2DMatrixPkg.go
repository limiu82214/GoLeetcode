package Q74Searcha2DMatrixPkg

func Solve(matrix [][]int, target int) bool {
	i := 0
	j := len(matrix) - 1
	for i <= j {
		midIdx := (i + j) / 2
		lastOfMidArrIdx := len(matrix[midIdx]) - 1
		switch {
		case matrix[midIdx][0] <= target && target <= matrix[midIdx][lastOfMidArrIdx]:
			return miniSearch(matrix[midIdx], target)
		case target < matrix[midIdx][0]:
			j = midIdx - 1
		case target > matrix[midIdx][lastOfMidArrIdx]:
			i = midIdx + 1
		}
	}

	return false
}

func miniSearch(a []int, target int) bool {
	i := 0
	j := len(a) - 1
	for i <= j {
		midIdx := (i + j) / 2
		switch {
		case a[midIdx] == target:
			return true
		case target < a[midIdx]:
			j = midIdx - 1
		case target > a[midIdx]:
			i = midIdx + 1
		}

	}

	return false
}
