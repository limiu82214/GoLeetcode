package Q347TopKFrequentElementsPkg

func Solve(nums []int, k int) []int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	/*
	   s := make([]int, k, k)
	   for i, _ := range s {
	       maxC := math.MinInt
	       maxN := 0
	       for n, count := range m {
	           if count > maxC {
	               maxC = count
	               maxN = n
	           }
	       }
	       s[i] = maxN
	       m[maxN] = 0
	   }
	*/

	s := make([][]int, len(nums)+1)
	for n, count := range m {
		s[count] = append(s[count], n)
	}

	ans := []int{}
	//    var b []int
	//    ans = append(ans, b...) // if b is [] and append will doing nothing
	//    fmt.Println(ans, len(ans))
	for i := len(s) - 1; i > 0; i-- {
		ans = append(ans, s[i]...)
		if len(ans) == k {
			return ans
		}
	}

	return nil
}
