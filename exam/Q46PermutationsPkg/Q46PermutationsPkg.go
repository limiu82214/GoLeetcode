package Q46PermutationsPkg

func Solve(nums []int) [][]int {
	ans := [][]int{}
	deep(nums, 1, []int{}, &ans)
	return ans
}

func deep(nums []int, dive int, carry []int, ans *[][]int) {
	for _, v := range nums {
		switch {
		case contain(carry, v):
			continue
		case dive == len(nums):
			t := append([]int(nil), carry...)
			t = append(t, v)
			*ans = append(*ans, t)
		default:
			t := append([]int(nil), carry...)
			t = append(t, v)
			deep(nums, dive+1, t, ans)
		}
	}
}

func contain(nums []int, target int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}

	return false
}
