package Q238ProductofArrayExceptSelfPkg

func Solve(nums []int) []int {
	l := len(nums)
	// idx:  [0],     [1],   [2],   [3]
	// nums: [1],     [2],   [3],   [4]
	// si:   [0],     [1],   [1*2], [1*2*3]
	// sj:   [2*3*4], [3*4], [4],   [0]
	// i*j:  e:1,     e:2,    e:3,   e:4

	si := make([]int, l)
	iSum := 1
	for i := 1; i < l; i++ {
		iSum *= nums[i-1]
		si[i] = iSum
	}

	sj := make([]int, l)
	jSum := 1
	for j := l - 1; j > 0; j-- {
		jSum *= nums[j]
		sj[j-1] = jSum
	}

	ans := make([]int, l)
	for k := 0; k < l; k++ {
		switch k {
		case 0:
			ans[k] = sj[k]
		case l - 1:
			ans[k] = si[k]
		default:
			ans[k] = si[k] * sj[k]
		}
	}

	return ans
}
