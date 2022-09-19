package q170twosumiiidatastructuredesign

import "sort"

type TwoSum struct {
	list     []int
	isSorted bool
}

func Constructor() TwoSum {
	return TwoSum{
		list:     []int{},
		isSorted: false,
	}
}

func (t *TwoSum) Add(number int) {
	t.list = append(t.list, number)
	t.isSorted = false
}

func (t *TwoSum) Find(value int) bool {
	if len(t.list) < 1 {
		return false
	}
	if !t.isSorted {
		sort.Ints(t.list)
		t.isSorted = true
	}
	// target = low + high
	low := 0
	high := 0
	high = len(t.list) - 1

	for low != high {
		if t.list[low]+t.list[high] == value {
			return true
		} else if t.list[low]+t.list[high] < value {
			low++
		} else if t.list[low]+t.list[high] > value {
			high--
		}
	}

	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
