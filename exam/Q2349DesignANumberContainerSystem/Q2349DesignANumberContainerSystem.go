package q2349designanumbercontainersystem

type NumberContainers struct {
	idxValMap      map[int]int              // 紀錄index對應value
	valIdxMapQuery map[int]map[int]struct{} // 紀錄某個value的所有index
	minIdxMap      map[int]int              // 記錄某個value最小的index
}

func Constructor() NumberContainers {
	return NumberContainers{
		make(map[int]int),
		make(map[int]map[int]struct{}),
		make(map[int]int),
	}
}

func (n *NumberContainers) Change(index int, number int) {
	if _, ok := n.idxValMap[index]; ok {
		// 修改
		oldNumber := n.idxValMap[index]
		if oldNumber != number {
			// idxValMap
			n.idxValMap[index] = number

			// valIdxMapQuery
			delete(n.valIdxMapQuery[oldNumber], index)
			q, ok := n.valIdxMapQuery[number]
			if !ok {
				q = make(map[int]struct{})
			}
			q[index] = struct{}{}
			n.valIdxMapQuery[number] = q

			// minIdxMap
			if _, ok := n.minIdxMap[oldNumber]; ok {
				delete(n.minIdxMap, oldNumber)
			}
			min := 1000000001
			for idx := range n.valIdxMapQuery[oldNumber] {
				if min > idx {
					min = idx
				}
			}
			if min != 1000000001 {
				n.minIdxMap[oldNumber] = min
			}
			if v, ok := n.minIdxMap[number]; !ok || index < v {
				n.minIdxMap[number] = index
			}
		}

	} else {
		// 新增
		// idxValMap
		n.idxValMap[index] = number

		// valIdxMapQuery
		q, ok := n.valIdxMapQuery[number]
		if !ok {
			q = make(map[int]struct{})
		}
		q[index] = struct{}{}
		n.valIdxMapQuery[number] = q

		// minIdxMap
		if v, ok := n.minIdxMap[number]; !ok || index < v {
			n.minIdxMap[number] = index
		}
	}
}

func (n *NumberContainers) Find(number int) int {
	v, ok := n.minIdxMap[number]
	if ok {
		return v
	}
	return -1
}
