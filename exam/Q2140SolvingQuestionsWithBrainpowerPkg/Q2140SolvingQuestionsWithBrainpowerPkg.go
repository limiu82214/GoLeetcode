package Q2140SolvingQuestionsWithBrainpowerPkg

func Solve(questions [][]int) int64 {
	greedy := make([]int, 900050)
	return int64(deep(questions, 0, greedy))
}
func deep(q [][]int, idx int, greedy []int) int {
	l := len(q)
	if idx > l-1 {
		return 0
	}

	first := q[idx]
	firstPoint := first[0]
	firstSkip := first[1]

	a := cacheOrDeep(q, idx+1+firstSkip, greedy) + firstPoint
	b := cacheOrDeep(q, idx+1, greedy)

	return getBig(a, b)
}
func cacheOrDeep(q [][]int, idx int, greedy []int) int {
	if greedy[idx] == 0 {
		greedy[idx] = deep(q, idx, greedy)
	}
	return greedy[idx]
}
func getBig(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

// func deep(q [][]int) int {
//     l:= len(q)

//     first := q[0]
//     firstPoint := first[0]
//     firstSkip := first[1]
//     if l == 1 {
//         return firstPoint
//     }

//     a := 0
//     b := 0
//     if (1 + firstSkip > l-1) {
//         a = firstPoint
//         b = deep(q[1:]) // 不要
//     } else {
//         a = deep(q[1 + firstSkip:]) + firstPoint // 要
//         b = deep(q[1:]) // 不要
//     }
//     if a > b {
//         return a
//     } else {
//         return b
//     }
// }

// func deep(q [][]int, length int) int {
//     l := len(q)
//     if l == 0 {
//         return 0
//     }

//     last := q[l-1]
//     lastPoint := last[0]
//     lastSkip := last[1]

//     if length == 0  {
//         a := deep(q[0:l-1], 1) + lastPoint // 要
//         b := deep(q[0:l-1], 0) // 不要
//         if a > b {
//             return a
//         } else {
//             return b
//         }
//     } else if length > lastSkip {
//         a := deep(q[0:l-1], 1) + lastPoint // 要
//         b := deep(q[0:l-1], length+1) // 不要
//         if a > b {
//             return a
//         } else {
//             return b
//         }
//     } else {
//         return deep(q[0:l-1], length+1)
//     }
// }
