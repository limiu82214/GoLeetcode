package Q1557MinimumNumberofVerticestoReachAlNodesPkg

func Solve(n int, edges [][]int) []int {
	m := make([]bool, n)
	for _, p := range edges {
		m[p[1]] = true
	}

	ans := []int{}

	for num, b := range m {
		if !b {
			ans = append(ans, num)
		}
	}

	return ans
}
