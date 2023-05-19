package Q785IsGraphBipartitePkg

func Solve(graph [][]int) bool {
	l := len(graph)
	m := make([]int, l)

	for i, base := range m {
		if base == NONE {
			base = A
		}

		if deep(graph, m, base, i) {
			continue
		}

		return false
	}

	return true
}

func deep(graph [][]int, m []int, base, baseIdx int) bool {
	m[baseIdx] = base
	for _, v := range graph[baseIdx] {
		switch m[v] {
		case NONE:
			deep(graph, m, revert(base), v)
		case base:
			return false
			// default:
			// continue
		}
	}

	return true
}

func revert(i int) int {
	if i == A {
		return B
	}

	return A
}

const (
	NONE int = 0
	A    int = 1
	B    int = 2
)
