package Q39CombinationSumPkg

import "sort"

func Solve(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := [][]int{}
	deep(candidates, target, []int{}, &ans, 0)
	return ans
}

func deep(candidates []int, target int, carry []int, ans *[][]int, min int) {
	for _, v := range candidates {
		// backtracking的要點是 先不要想剪枝
		if v < min {
			continue
		}
		switch {
		case v > target:
			return
		case v == target:
			t := append([]int{}, carry...)
			t = append(t, v)
			*ans = append(*ans, t)
		case 2*v <= target:
			t := append([]int{}, carry...)
			t = append(t, v)
			deep(candidates, target-v, t, ans, v)
		case v < target:
			t := append([]int{}, carry...)
			t = append(t, v)
			deep(candidates[1:], target-v, t, ans, v)
		}
	}
}
