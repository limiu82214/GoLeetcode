package Q1091ShortestPathinBinaryMatrixPkg

// BFS
// If I use Step to record the step instead edit the grid, it will be OOM
func Solve(grid [][]int) int {
	if len(grid) <= 0 {
		return -1
	}
	l := len(grid)
	if grid[0][0] == 1 || grid[l-1][l-1] == 1 {
		return -1
	}

	direction := [][]int{
		[]int{1, 0},   // i+1, j+0 right
		[]int{1, 1},   // i+1, j+1 right bot
		[]int{0, 1},   // i+0, j+1 bot
		[]int{-1, 1},  // i-1, j+1 left bot
		[]int{-1, 0},  // i-1, j+0 left
		[]int{-1, -1}, // i-1, j-1 left top
		[]int{0, -1},  // i+0, j-1 top
		[]int{1, -1},  // i+1, j-1 right top
	}
	ti, tj := 0, 0
	queue := []Node{Node{0, 0}}
	for len(queue) != 0 {
		// pop now
		n := queue[0]
		queue = queue[1:]
		i, j := n.I, n.J

		if i == l-1 && j == l-1 {
			return grid[i][j] + 1
		}

		// push next
		for k := 0; k < len(direction); k++ {
			ti, tj = i+direction[k][0], j+direction[k][1]
			if !isOutOfBound(ti, tj, l) && grid[ti][tj] == 0 {
				grid[ti][tj] = grid[i][j] + 1 // set visited
				queue = append(queue, Node{ti, tj})
			}
		}
	}

	return -1
}

type Node struct {
	I, J int
}

func isOutOfBound(i, j, n int) bool {
	return i < 0 || j < 0 || i > n-1 || j > n-1
}

/** DFS TLE
func shortestPathBinaryMatrix(grid [][]int) int {
    if len(grid) <= 0 {
        return -1
    }
    n := len(grid)
    if grid[0][0] == 1 || grid[n-1][n-1] == 1{
        return -1
    }
    min := math.MaxInt32
    deep(grid, 0, 0, n, 1,[][]int{}, &min)
    if min == math.MaxInt32 {
        return -1
    }
    return min
}

func deep(grid [][]int, i, j, n, step int, passed [][]int, min *int) {
    if isOutOfBound(i, j, n) || isWall(i,j,grid) || isPassed(i, j, passed) {
        return
    }
    if step >= *min {
        return
    }
    if isGoal(i, j, n) {
        if step < *min {
            *min = step
        }
        return
    }
    t := append([][]int{}, passed...)
    t = append(t, []int{i, j})

    // i+1, j+0 right
    ti := i+1
    tj := j+0
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1, t, min)
    }

    // i+1, j+1 right bot
    ti = i+1
    tj = j+1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t, min)
    }

    // i+0, j+1 bot
    ti = i+0
    tj = j+1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t, min)
    }

    // i-1, j+1 left bot
    ti = i-1
    tj = j+1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t, min)
    }

    // i-1, j+0 left
    ti = i-1
    tj = j+0
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t, min)
    }

    // i-1, j-1 left top
    ti = i-1
    tj = j-1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t, min)
    }

    // i+0, j-1 top
    ti = i+0
    tj = j-1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t, min)
    }

    // i+1, j-1 right top
    ti = i+1
    tj = j-1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t, min)
    }

    return
}

func isOutOfBound(i, j, n int) bool {
    return i<0 || j <0 || i > n-1 || j > n-1
}

func isWall(i, j int, grid [][]int) bool {
    return grid[i][j] == 1
}

func isGoal(i, j, n int) bool {
    return i == n-1 && j == n-1
}

func isPassed(i, j int, passed [][]int) bool {
    for _, s := range passed {
        if s[0] ==i && s[1] == j {
            return true
        }
    }

    return false
}
*/

/** BFS use global variable(min) should reset it before each test case, otherwise it will be wrong
func shortestPathBinaryMatrix(grid [][]int) int {
    if len(grid) <= 0 {
        return -1
    }

    deep(grid, 0, 0, len(grid), 1,[][]int{})
    return min
}
var min int = math.MaxInt32

func deep(grid [][]int, i, j, n, step int, passed [][]int) {
    if isOutOfBound(i, j, n) || isWall(i,j,grid) || isPassed(i, j, passed) {
        return
    }
    if isGoal(i, j, n) {
            fmt.Println(min)
        if step < min {
            min = step
        }
        return
    }
    t := append([][]int{}, passed...)
    t = append(t, []int{i, j})

    // i+1, j+0 right
    ti := i+1
    tj := j+0
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1, t)
    }

    // i+1, j+1 right bot
    ti = i+1
    tj = j+1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t)
    }

    // i+0, j+1 bot
    ti = i+0
    tj = j+1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t)
    }

    // i-1, j+1 left bot
    ti = i-1
    tj = j+1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t)
    }

    // i-1, j+0 left
    ti = i-1
    tj = j+0
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t)
    }

    // i-1, j-1 left top
    ti = i-1
    tj = j-1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t)
    }

    // i+0, j-1 top
    ti = i+0
    tj = j-1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t)
    }

    // i+1, j-1 right top
    ti = i+1
    tj = j-1
    if !isPassed(ti, tj, passed) {
        deep(grid, ti, tj, n, step+1,t)
    }

    return
}

func isOutOfBound(i, j, n int) bool {
    return i<0 || j <0 || i > n-1 || j > n-1
}

func isWall(i, j int, grid [][]int) bool {
    return grid[i][j] == 1
}

func isGoal(i, j, n int) bool {
    return i == n-1 && j == n-1
}

func isPassed(i, j int, passed [][]int) bool {
    for _, s := range passed {
        if s[0] ==i && s[1] == j {
            return true
        }
    }

    return false
}

*/
