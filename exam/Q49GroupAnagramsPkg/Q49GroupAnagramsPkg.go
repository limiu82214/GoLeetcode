package Q49GroupAnagramsPkg

func Solve(strs []string) [][]string {
	m := make(map[string][]int)
	for idx, str := range strs {
		m[order(str)] = append(m[order(str)], idx)
	}

	ans := [][]string{}
	for _, intlist := range m {
		t := []string{}
		for _, idxs := range intlist {
			t = append(t, strs[idxs])
		}

		ans = append(ans, t)
	}

	return ans
}

func order(s string) string {
	abc := make([]int, 26)
	base := int('a')

	for _, v := range s {
		abc[int(v)-base]++
	}

	ans := ""
	for _, v := range abc {
		ans += string(v + base)
	}

	return ans
}
