package main

import (
	"fmt"
	"sort"
	"testing"

	q170twosumiiidatastructuredesign "github.com/limiu82214/GoLeetcode/exam/Q170TwoSumIIIDataStructureDesign"
	q2349designanumbercontainersystem "github.com/limiu82214/GoLeetcode/exam/Q2349DesignANumberContainerSystem"
	q2353designafoodratingsystem "github.com/limiu82214/GoLeetcode/exam/Q2353DesignAFoodRatingSystem"
	"github.com/stretchr/testify/assert"
)

// TestQ2121IntervalsBetweenIdenticalElements Medium
func TestQ2121IntervalsBetweenIdenticalElements(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected []int64
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[2,1,3,1,2,3,3]`),
			Expected: JsonToSlice[int64](`[4,2,7,2,4,4,5]`),
		},
		{
			Arg1:     JsonToSlice[int](`[10,5,10,10]`),
			Expected: JsonToSlice[int64](`[5,0,3,4]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2121IntervalsBetweenIdenticalElements(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ215KthLargestElementInAnArray Medium
func TestQ215KthLargestElementInAnArray(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[3,2,1,5,6,4]`),
			Arg2:     2,
			Expected: 5,
		},
		{
			Arg1:     JsonToSlice[int](`[3,2,3,1,2,4,5,5,6]`),
			Arg2:     4,
			Expected: 4,
		},
		{
			Arg1:     JsonToSlice[int](`[0]`),
			Arg2:     1,
			Expected: 0,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q215KthLargestElementInAnArray(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ260SingleNumberIII Medium
func TestQ260SingleNumberIII(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected []int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,2,1,3,2,5]`),
			Expected: JsonToSlice[int](`[3,5]`),
		},
		{
			Arg1:     JsonToSlice[int](`[-1,0]`),
			Expected: JsonToSlice[int](`[-1,0]`),
		},
		{
			Arg1:     JsonToSlice[int](`[0,1]`),
			Expected: JsonToSlice[int](`[1,0]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.ElementsMatch(t, d.Expected, Q260SingleNumberIII(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ704BinarySearch Easy
func TestQ704BinarySearch(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,0,3,5,9,12]`),
			Arg2:     9,
			Expected: 4,
		},
		{
			Arg1:     JsonToSlice[int](`[1,0,3,5,9,12]`),
			Arg2:     2,
			Expected: -1,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q704BinarySearch(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ90SubsetsII Medium
func TestQ90SubsetsII(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected [][]int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,2,2]`),
			Expected: JsonToSliceSlice[int](`[[],[1],[1,2],[1,2,2],[2],[2,2]]`),
		},
		{
			Arg1:     JsonToSlice[int](`[0]`),
			Expected: JsonToSliceSlice[int](`[[],[0]]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.ElementsMatch(t, d.Expected, Q90SubsetsII(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ78Subsets Medium
func TestQ78Subsets(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected [][]int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,2,3]`),
			Expected: JsonToSliceSlice[int](`[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]`),
		},
		{
			Arg1:     JsonToSlice[int](`[0]`),
			Expected: JsonToSliceSlice[int](`[[],[0]]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q78Subsets(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ113PathSumII Medium
func TestQ113PathSumII(t *testing.T) {
	type param struct {
		Arg1     *TreeNode
		Arg2     int
		Expected [][]int
	}
	data := []param{
		{
			Arg1:     JSONArrayToTreeNode(`[5,4,8,11,null,13,4,7,2,null,null,5,1]`),
			Arg2:     22,
			Expected: JsonToSliceSlice[int](`[[5,4,11,2],[5,8,4,5]]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[1,2,3]`),
			Arg2:     5,
			Expected: JsonToSliceSlice[int](`[]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[1,2]`),
			Arg2:     0,
			Expected: JsonToSliceSlice[int](`[]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[]`),
			Arg2:     1,
			Expected: JsonToSliceSlice[int](`[]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[1,2]`),
			Arg2:     1,
			Expected: JsonToSliceSlice[int](`[]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[-2,null,-3]`),
			Arg2:     -5,
			Expected: JsonToSliceSlice[int](`[[-2,-3]]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q113PathSumII(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ110BalancedBinaryTree Easy
func TestQ110BalancedBinaryTree(t *testing.T) {
	type param struct {
		Arg1     *TreeNode
		Expected bool
	}
	data := []param{
		{
			Arg1:     JSONArrayToTreeNode(`[3,9,20,null,null,15,7]`),
			Expected: true,
		},
		{
			Arg1:     JSONArrayToTreeNode(`[1,2,2,3,3,null,null,4,4]`),
			Expected: false,
		},
		{
			Arg1:     JSONArrayToTreeNodeV2(`[1,null,2,null,3]`),
			Expected: false,
		},
		{
			Arg1:     JSONArrayToTreeNode(`[]`),
			Expected: true,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q110BalancedBinaryTree(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ111MinimumDepthOfBinaryTree Easy
func TestQ111MinimumDepthOfBinaryTree(t *testing.T) {
	type param struct {
		Arg1     *TreeNode
		Expected int
	}
	data := []param{
		{
			Arg1:     JSONArrayToTreeNode(`[3,9,20,null,null,15,7]`),
			Expected: 2,
		},
		// {
		// 	Arg1:     JSONArrayToTreeNode(`[2,null,3,null,4,null,5,null,6]`),
		// 	Expected: 5, // 此題目的輸入是有左邊節點就會一直往增加左邊
		// },
		{
			Arg1:     JSONArrayToTreeNode(`[]`),
			Expected: 0,
		},
		{
			Arg1:     JSONArrayToTreeNode(`[1]`),
			Expected: 1,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q111MinimumDepthOfBinaryTree(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ107BinaryTreeLevelOrderTraversalII Medium
func TestQ107BinaryTreeLevelOrderTraversalII(t *testing.T) {
	type param struct {
		Arg1     *TreeNode
		Expected [][]int
	}
	data := []param{
		{
			Arg1:     JSONArrayToTreeNode(`[3,9,20,null,null,15,7]`),
			Expected: JsonToSliceSlice[int](`[[15,7],[9,20],[3]]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[1]`),
			Expected: JsonToSliceSlice[int](`[[1]]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[]`),
			Expected: JsonToSliceSlice[int](`[]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q107BinaryTreeLevelOrderTraversalII(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ104MaximumDepthOfBinaryTree Easy
func TestQ104MaximumDepthOfBinaryTree(t *testing.T) {
	type param struct {
		Arg1     *TreeNode
		Expected int
	}
	data := []param{
		{
			Arg1:     JSONArrayToTreeNode(`[3,9,20,null,null,15,7]`),
			Expected: 3,
		},
		{
			Arg1:     JSONArrayToTreeNode(`[1,null,2]`),
			Expected: 2,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q104MaximumDepthOfBinaryTree(d.Arg1), d)
		}
	}, 1, 0)
}

// TestForCopy Easy ============= 複製一個copy並開始你的表演
func TestForCopy(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     []interface{}
		Expected float64
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,3]`),
			Arg2:     JsonStringToSliceAny(`[2]`),
			Expected: 2,
		},
		{
			Arg1:     JsonToSlice[int](`[1,2]`),
			Arg2:     JsonStringToSliceAny(`[3,4]`),
			Expected: 2,
		},
	}

	Decorate(func() {
		for _, d := range data {

			var param2 []int
			Arg2 := ConvertSlice[float64](d.Arg2)
			param2 = make([]int, 0, len(Arg2))
			for _, v := range Arg2 {
				param2 = append(param2, int(v))
			}
			assert.Equal(t, d.Expected, Q4MedianOfTwoSortedArrays(d.Arg1, param2), d)
		}
	}, 1, 0)
}

// TestQ103BinaryTreeZigzagLevelOrderTraversal Medium
func TestQ103BinaryTreeZigzagLevelOrderTraversal(t *testing.T) {
	type param struct {
		Arg1     *TreeNode
		Expected [][]int
	}
	data := []param{

		{
			Arg1:     JSONArrayToTreeNode(`[1,2,3,4,null,null,5]`),
			Expected: JsonToSliceSlice[int](`[[1],[3,2],[4,5]]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[3,9,20,null,null,15,7]`),
			Expected: JsonToSliceSlice[int](`[[3],[20,9],[15,7]]`),
		},

		{
			Arg1:     JSONArrayToTreeNode(`[1]`),
			Expected: JsonToSliceSlice[int](`[[1]]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[]`),
			Expected: JsonToSliceSlice[int](`[]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q103BinaryTreeZigzagLevelOrderTraversal(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ102BinaryTreeLevelOrderTraversal Medium
func TestQ102BinaryTreeLevelOrderTraversal(t *testing.T) {
	type param struct {
		Arg1     *TreeNode
		Expected [][]int
	}
	data := []param{
		{
			Arg1:     JSONArrayToTreeNode(`[3,9,20,null,null,15,7]`),
			Expected: JsonToSliceSlice[int](`[[3],[9,20],[15,7]]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[1]`),
			Expected: JsonToSliceSlice[int](`[[1]]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[]`),
			Expected: JsonToSliceSlice[int](`[]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q102BinaryTreeLevelOrderTraversal(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ41FirstMissingPositive Hard
func TestQ41FirstMissingPositive(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1]`),
			Expected: 2,
		},
		{
			Arg1:     JsonToSlice[int](`[1,2,0]`),
			Expected: 3,
		},
		{
			Arg1:     JsonToSlice[int](`[3,4,-1,1]`),
			Expected: 2,
		},
		{
			Arg1:     JsonToSlice[int](`[7,8,9,11,12]`),
			Expected: 1,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q41FirstMissingPositive(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ645SetMismatch Easy
func TestQ645SetMismatch(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected []int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,2,2,4]`),
			Expected: JsonToSlice[int](`[2,3]`),
		},
		{
			Arg1:     JsonToSlice[int](`[1,1]`),
			Expected: JsonToSlice[int](`[1,2]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q645SetMismatch(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ442FindAllDuplicatesInAnArray Medium
func TestQ442FindAllDuplicatesInAnArray(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected []int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[4,3,2,7,8,2,3,1]`),
			Expected: JsonToSlice[int](`[3,2]`),
		},
		{
			Arg1:     JsonToSlice[int](`[1,1,2]`),
			Expected: JsonToSlice[int](`[1]`),
		},
		{
			Arg1:     JsonToSlice[int](`[1]`),
			Expected: JsonToSlice[int](`[]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q442FindAllDuplicatesInAnArray(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ287FindTheDuplicateNumber Medium *fail*
func TestQ287FindTheDuplicateNumber(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,3,4,2,2]`),
			Expected: 2,
		},
		{
			Arg1:     JsonToSlice[int](`[3,1,3,4,2]`),
			Expected: 3,
		},
		{
			Arg1:     JsonToSlice[int](`[1,1]`),
			Expected: 1,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q287FindTheDuplicateNumber(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ136SingleNumber Easy
func TestQ136SingleNumber(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[2,2,1]`),
			Expected: 1,
		},
		{
			Arg1:     JsonToSlice[int](`[4,1,2,1,2]`),
			Expected: 4,
		},
		{
			Arg1:     JsonToSlice[int](`[1]`),
			Expected: 1,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q136SingleNumber(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ268MissingNumber Easy
func TestQ268MissingNumber(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[3,0,1]`),
			Expected: 2,
		},
		{
			Arg1:     JsonToSlice[int](`[0,1]`),
			Expected: 2,
		},
		{
			Arg1:     JsonToSlice[int](`[9,6,4,2,3,5,7,0,1]`),
			Expected: 8,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q268MissingNumber(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ495TeemoAttacking Easy
func TestQ495TeemoAttacking(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,4]`),
			Arg2:     2,
			Expected: 4,
		},
		{
			Arg1:     JsonToSlice[int](`[1,2]`),
			Arg2:     2,
			Expected: 3,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q495TeemoAttacking(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ253MeetingRoomsII Medium
func TestQ253MeetingRoomsII(t *testing.T) {
	type param struct {
		Arg1     [][]int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSliceSlice[int](`[[0,30],[5,10],[15,20]]`),
			Expected: 2,
		},
		{
			Arg1:     JsonToSliceSlice[int](`[[7,10],[2,4]]`),
			Expected: 1,
		},
		{
			Arg1:     JsonToSliceSlice[int](`[[2,7]]`),
			Expected: 1,
		},
		{
			Arg1:     JsonToSliceSlice[int](`[[5,8],[6,8]]`),
			Expected: 2,
		},
		{
			Arg1:     JsonToSliceSlice[int](`[[6,15],[13,20],[6,17]]`),
			Expected: 3,
		},
		{
			Arg1:     JsonToSliceSlice[int](`[[13,15],[1,13]]`),
			Expected: 1,
		},
		{
			Arg1:     JsonToSliceSlice[int](`[[1293,2986],[848,3846],[4284,5907],[4466,4781],[518,2918],[300,5870]]`),
			Expected: 4,
		},
	}

	Decorate(func() {
		for _, d := range data {
			actual := Q253MeetingRoomsII(d.Arg1)
			assert.Equal(t, d.Expected, actual, d)
		}
	}, 1, 0)

}

// TestQ252MeetingRooms Easy
func TestQ252MeetingRooms(t *testing.T) {
	type param struct {
		Arg1     [][]int
		Expected bool
	}
	data := []param{
		{
			Arg1:     JsonToSliceSlice[int](`[[0,30],[5,10],[15,20]]`),
			Expected: false,
		},
		{
			Arg1:     JsonToSliceSlice[int](`[[7,10],[2,4]]`),
			Expected: true,
		},
	}

	Decorate(func() {
		for _, d := range data {
			actual := Q252MeetingRooms(d.Arg1)
			assert.Equal(t, d.Expected, actual, d)
		}
	}, 1, 0)

}

// TestQ57InsertInterval Medium
func TestQ57InsertInterval(t *testing.T) {
	type param struct {
		Arg1     [][]int
		Arg2     []int
		Expected [][]int
	}
	data := []param{
		{
			Arg1:     JsonToSliceSlice[int](`[[1,3],[6,9]]`),
			Arg2:     JsonToSlice[int](`[2,5]`),
			Expected: JsonToSliceSlice[int](`[[1,5],[6,9]]`),
		},
		{
			Arg1:     JsonToSliceSlice[int](`[[1,2],[3,5],[6,7],[8,10],[12,16]]`),
			Arg2:     JsonToSlice[int](`[4,8]`),
			Expected: JsonToSliceSlice[int](`[[1,2],[3,10],[12,16]]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			for i := range d.Expected {
				sort.Ints(d.Expected[i])
			}
			actual := Q57InsertInterval(d.Arg1, d.Arg2)
			for i := range actual {
				sort.Ints(actual[i])
			}
			assert.Equal(t, d.Expected, actual, d)
		}
	}, 1, 0)

}

// TestQ56MergeIntervals Medium
func TestQ56MergeIntervals(t *testing.T) {
	type param struct {
		Arg1     [][]int
		Expected [][]int
	}
	data := []param{
		{
			Arg1:     JsonToSliceSlice[int](`[[1,3],[2,6],[8,10],[15,18]]`),
			Expected: JsonToSliceSlice[int](`[[1,6],[8,10],[15,18]]`),
		},
		{
			Arg1:     JsonToSliceSlice[int](`[[1,4],[4,5]]`),
			Expected: JsonToSliceSlice[int](`[[1,5]]`),
		},
		{
			Arg1:     JsonToSliceSlice[int](`[[1,4],[0,2],[3,5]]`),
			Expected: JsonToSliceSlice[int](`[[0,5]]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			for i := range d.Expected {
				sort.Ints(d.Expected[i])
			}
			actual := Q56MergeIntervals(d.Arg1)
			for i := range actual {
				sort.Ints(actual[i])
			}
			assert.Equal(t, d.Expected, actual, d)
		}
	}, 1, 0)

}

// TestQ163SumClosest Medium
func TestQ2593SumSmaller(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[-2,0,1,3]`),
			Arg2:     2,
			Expected: 2,
		},
		{
			Arg1:     JsonToSlice[int](`[]`),
			Arg2:     0,
			Expected: 0,
		},
		{
			Arg1:     JsonToSlice[int](`[0]`),
			Arg2:     0,
			Expected: 0,
		},
		{
			Arg1:     JsonToSlice[int](`[3,1,0,-2]`),
			Arg2:     4,
			Expected: 3,
		},
		{
			Arg1:     JsonToSlice[int](`[30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30]`),
			Arg2:     1,
			Expected: 0, // TLE test

		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2593SumSmaller(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ163SumClosest Medium
func TestQ163SumClosest(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[-1,2,1,-4]`),
			Arg2:     1,
			Expected: 2,
		},
		{
			Arg1:     JsonToSlice[int](`[0,0,0]`),
			Arg2:     1,
			Expected: 0,
		},
		{
			Arg1:     JsonToSlice[int](`[0,1,2]`),
			Arg2:     0,
			Expected: 3,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q163SumClosest(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ19RemoveNthNodeFromEndOfList Medium
func TestQ19RemoveNthNodeFromEndOfList(t *testing.T) {
	type param struct {
		Arg1     *ListNode
		Arg2     int
		Expected *ListNode
	}
	data := []param{
		{
			Arg1:     JSONArrayToListNode(`[1,2,3,4,5]`),
			Arg2:     2,
			Expected: JSONArrayToListNode(`[1,2,3,5]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[1]`),
			Arg2:     1,
			Expected: JSONArrayToListNode(`[]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[1,2]`),
			Arg2:     1,
			Expected: JSONArrayToListNode(`[1]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[1,2]`),
			Arg2:     2,
			Expected: JSONArrayToListNode(`[2]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q19RemoveNthNodeFromEndOfList(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ11ContainerWithMostWater Medium
func TestQ11ContainerWithMostWater(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,8,6,2,5,4,8,3,7]`),
			Expected: 49,
		},
		{
			Arg1:     JsonToSlice[int](`[1,1]`),
			Expected: 1,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q11ContainerWithMostWater(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ5LongestPalindromicSubstring Medium
func TestQ5LongestPalindromicSubstring(t *testing.T) {
	type param struct {
		Arg1     string
		Expected string
	}
	data := []param{
		{
			Arg1:     "babad",
			Expected: "bab",
		},
		{
			Arg1:     "cbbd",
			Expected: "bb",
		},
		{
			Arg1:     "aacabdkacaa",
			Expected: "aca",
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q5LongestPalindromicSubstring(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ6ZigzagConversion Medium
func TestQ6ZigzagConversion(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     int
		Expected string
	}
	data := []param{
		{
			Arg1:     "PAYPALISHIRING",
			Arg2:     3,
			Expected: "PAHNAPLSIIGYIR",
		},
		{
			Arg1:     "PAYPALISHIRING",
			Arg2:     4,
			Expected: "PINALSIGYAHRPI",
		},
		{
			Arg1:     "A",
			Arg2:     1,
			Expected: "A",
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q6ZigzagConversion(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ2212MaximumPointsInAnArcheryCompetition Medium *fail*
func TestQ2212MaximumPointsInAnArcheryCompetition(t *testing.T) {
	type param struct {
		Arg1     int
		Arg2     []int
		Expected []int
	}
	data := []param{
		{
			Arg1:     9,
			Arg2:     JsonToSlice[int](`[1,1,0,1,0,0,2,1,0,1,2,0]`),
			Expected: JsonToSlice[int](`[0,0,0,0,1,1,0,0,1,2,3,1]`),
		},
		{
			Arg1:     3,
			Arg2:     JsonToSlice[int](`[0,0,1,0,0,0,0,0,0,0,0,2]`),
			Expected: JsonToSlice[int](`[0,0,0,0,0,0,0,0,1,1,1,0]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2212MaximumPointsInAnArcheryCompetition(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ516LongestPalindromicSubsequence Medium
func TestQ516LongestPalindromicSubsequence(t *testing.T) {
	type param struct {
		Arg1     string
		Expected int
	}
	data := []param{
		{
			Arg1:     "a",
			Expected: 1,
		},
		{
			Arg1:     "bbbab",
			Expected: 4,
		},
		{
			Arg1:     "cbbd",
			Expected: 2,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q516LongestPalindromicSubsequence(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ159LongestSubstringWithAtMostTwoDistinctCharacters Medium
func TestQ159LongestSubstringWithAtMostTwoDistinctCharacters(t *testing.T) {
	type param struct {
		Arg1     string
		Expected int
	}
	data := []param{
		{
			Arg1:     "a",
			Expected: 1,
		},
		{
			Arg1:     "eceba",
			Expected: 3,
		},
		{
			Arg1:     "ccaabbb",
			Expected: 5,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q159LongestSubstringWithAtMostTwoDistinctCharacters(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ2405OptimalPartitionOfString Medium
func TestQ2405OptimalPartitionOfString(t *testing.T) {
	type param struct {
		Arg1     string
		Expected int
	}
	data := []param{
		{
			Arg1:     "abacaba",
			Expected: 4,
		},
		{
			Arg1:     "ssssss",
			Expected: 6,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2405OptimalPartitionOfString(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ2401LongestNiceSubarray Medium
func TestQ2401LongestNiceSubarray(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,3,8,48,10]`),
			Expected: 3,
		},
		{
			Arg1:     JsonToSlice[int](`[3,1,5,11,13]`),
			Expected: 1,
		},
		{
			Arg1:     JsonToSlice[int](`[135745088,609245787,16,2048,2097152]`),
			Expected: 3,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2401LongestNiceSubarray(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ2260MinimumConsecutiveCardsToPickUp Medium
func TestQ2260MinimumConsecutiveCardsToPickUp(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[3,4,2,3,4,7]`),
			Expected: 4,
		},
		{
			Arg1:     JsonToSlice[int](`[1,0,5,3]`),
			Expected: -1,
		},
		{
			Arg1:     JsonToSlice[int](`[95,11,8,65,5,86,30,27,30,73,15,91,30,7,37,26,55,76,60,43,36,85,47,96,6]`),
			Expected: 3,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2260MinimumConsecutiveCardsToPickUp(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ2067NumberOfEqualCountSubstrings Medium *fail*
func TestQ2067NumberOfEqualCountSubstrings(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     "aaabcbbcc",
			Arg2:     3,
			Expected: 3,
		},
		{
			Arg1:     "abcd",
			Arg2:     2,
			Expected: 0,
		},
		{
			Arg1:     "a",
			Arg2:     5,
			Expected: 0,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2067NumberOfEqualCountSubstrings(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ1695MaximumErasureValue Medium *fail*
func TestQ1695MaximumErasureValue(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[4,2,4,5,6]`),
			Expected: 17,
		},
		{
			Arg1:     JsonToSlice[int](`[5,2,1,2,5,2,1,2,5]`),
			Expected: 8,
		},
		{
			Arg1:     JsonToSlice[int](`[449,154,934,526,429,732,784,909,884,805,635,660,742,209,742,272,669,449,766,904,698,434,280,332,876,200,333,464,12,437,269,355,622,903,262,691,768,894,929,628,867,844,208,384,644,511,908,792,56,872,275,598,633,502,894,999,788,394,309,950,159,178,403,110,670,234,119,953,267,634,330,410,137,805,317,470,563,900,545,308,531,428,526,593,638,651,320,874,810,666,180,521,452,131,201,915,502,765,17,577,821,731,925,953,111,305,705,162,994,425,17,140,700,475,772,385,922,159,840,367,276,635,696,70,744,804,63,448,435,242,507,764,373,314,140,825,34,383,151,602,745]`),
			Expected: 30934,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q1695MaximumErasureValue(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ992SubarraysWithKDifferentIntegers Hard
func TestQ992SubarraysWithKDifferentIntegers(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,2,1,2,3]`),
			Arg2:     2,
			Expected: 7,
		},
		{
			Arg1:     JsonToSlice[int](`[1,2,1,3,4]`),
			Arg2:     3,
			Expected: 3,
		},
		{
			Arg1:     JsonToSlice[int](`[2,1,1,1,2]`),
			Arg2:     1,
			Expected: 8,
		},
		// {
		// 	Arg1:     JsonToSlice[int](`[638,1049,105,823,583,1051,1002,1032,333,958,1285,801,474,1342,405,1403,1600,1017,1484,663,932,1022,1331,271,151,787,46,164,1122,1543,90,480,84,856,1159,399,984,1471,1067,1346,573,907,411,152,1143,358,6,1235,420,1218,660,1188,264,831,856,761,52,1104,1099,772,124,199,1041,728,1063,1425,1134,1180,1095,1384,420,251,761,549,860,1229,1188,84,1087,273,326,111,1497,475,403,656,1477,853,328,1413,1211,669,235,450,628,640,690,1006,1590,1144,852,1337,563,266,660,1288,1149,1132,689,1349,1591,1043,1060,845,1239,364,69,809,1005,1483,853,1,1119,1408,448,443,41,73,383,772,961,935,571,164,1266,1344,1465,1496,535,85,25,350,210,1443,738,676,1554,118,685,603,634,872,646,843,1140,968,1061,1524,893,1202,1390,413,387,260,709,1298,1269,463,284,891,969,1140,290,1098,391,519,559,777,745,998,1469,439,34,896,690,542,923,861,85,39,444,1441,131,680,797,818,1134,625,1248,1392,1278,505,1052,1149,1223,796,146,939,249,1588,1598,551,261,320,554,1343,1454,190,536,1410,1512,1082,1489,406,951,844,1444,1401,679,372,105,606,50,856,1509,386,1132,536,836,1411,1115,926,208,58,1296,1111,1492,823,328,710,1343,193,338,1219,1541,999,1462,201,48,329,114,1494,1461,339,392,764,296,510,464,718,135,941,578,769,21,687,1117,35,893,903,642,84,179,1023,1354,79,708,306,209,516,950,778,497,827,338,380,1173,1160,654,1388,307,741,444,145,479,153,1523,1444,488,1376,1371,1495,1309,824,625,535,1133,155,160,1170,711,288,24,1386,1000,957,1240,416,1213,775,1130,178,646,435,204,1038,1098,1236,161,520,612,893,54,464,488,903,134,1428,1305,189,1307,511,781,743,781,364,965,26,518,466,446,380,17,1192,1067,546,9,56,465,1558,634,37,825,923,591,941,1165,239,342,769,1180,1301,1392,67,1529,441,1343,1025,1513,1401,923,1406,188,1497,508,1009,1131,1013,423,171,581,85,257,1119,651,469,638,1252,1181,813,579,61,1504,523,746,542,1002,1074,964,1419,966,259,1267,1183,1211,487,180,327,1558,735,848,1459,1259,618,1277,513,335,1342,175,740,829,321,137,646,382,120,1306,1361,367,224,739,738,1576,727,166,595,1382,619,485,953,440,15,800,1201,835,337,45,417,1416,583,932,823,959,1338,115,1198,1278,317,1299,435,159,1398,568,523,631,153,1416,122,238,1280,1115,1445,98,187,1184,1530,1041,250,959,135,1169,308,915,176,1479,1176,34,426,260,1577,16,1023,377,1293,671,137,154,1528,457,637,299,1046,1091,415,1508,1376,81,1425,209,446,404,134,1260,557,516,587,43,936,299,1487,976,81,1339,1016,1328,172,130,1412,1221,41,1438,963,728,1494,790,1384,1510,1077,1063,1537,815,899,469,1087,1363,815,461,1495,605,1127,714,35,834,851,567,111,1214,196,161,1453,1249,1564,732,1309,963,959,306,1125,1262,1539,445,1513,1181,1561,1535,313,1274,1241,633,1551,543,140,1066,290,1394,1118,136,8,461,945,497,1472,741,1221,315,1344,807,128,691,735,806,218,701,137,1443,1518,273,671,1004,437,1075,1003,1295,1074,785,76,1461,79,262,1578,1587,839,820,1557,18,724,1457,218,353,1470,373,928,1599,1337,362,724,1062,1282,669,1329,849,303,490,778,666,1168,1525,81,1242,132,1188,680,1079,434,840,1422,1532,39,255,626,854,587,622,138,981,1464,877,1220,1227,1303,214,376,1341,698,169,13,1371,152,1575,1264,531,1500,402,85,48,1063,829,31,546,271,519,1137,357,550,1069,1569,1526,977,816,1249,1254,720,1363,1036,214,1085,157,805,971,56,785,1519,1442,1516,447,391,1524,1133,1422,1417,985,1217,290,59,1573,563,268,188,1199,958,44,1062,1111,1256,1563,781,931,1167,508,159,168,1407,774,230,1346,911,745,1480,655,624,1254,426,168,187,495,652,961,680,1273,411,497,1474,327,1467,1458,1034,1183,1509,1105,530,110,294,1580,1161,1584,191,1155,397,514,1234,1323,312,335,1374,277,1382,666,1364,913,489,584,1039,232,1084,513,1312,1121,1363,91,281,1139,701,694,1132,1128,119,426,1232,1490,1137,767,674,1140,121,1185,97,1547,989,1365,761,1398,975,214,316,570,78,875,366,466,278,429,1492,2,1412,1085,287,1575,1519,435,169,960,423,1555,473,720,358,375,35,569,1362,1348,677,836,956,355,1088,1237,91,1124,401,968,1124,473,1042,1232,1428,1488,1598,346,703,920,529,1583,111,910,1374,631,1128,739,1431,161,1573,641,1418,1082,1084,122,646,16,1118,1388,622,945,167,1070,980,947,413,723,333,1024,690,139,1138,255,917,598,1460,1057,727,867,1053,410,319,1397,598,1180,944,1566,872,1539,687,640,1000,1295,757,134,1309,1149,206,864,50,775,804,1281,177,536,1572,902,1178,1589,1089,1123,202,402,586,806,357,1374,1536,212,378,623,1262,828,494,289,926,354,709,956,549,673,914,1597,295,1517,1590,436,220,1531,935,607,721,286,386,952,1467,505,65,512,1142,1130,709,1325,1068,109,1290,1520,1088,386,646,1047,334,44,693,854,668,1088,476,1319,1168,341,317,942,1174,176,1144,874,677,1324,1532,101,136,1468,1333,1062,153,303,33,805,856,146,482,485,365,730,1288,1024,1398,138,84,1478,1296,576,1140,148,1100,1282,628,425,1223,1166,1263,1357,1580,1319,92,798,323,1536,1532,707,883,1006,29,584,413,319,536,1117,1473,101,1332,1518,1516,1568,1280,826,1243,1426,57,197,1265,355,70,130,594,1462,1529,425,354,447,1448,272,970,608,791,537,1228,1145,195,763,1490,357,1590,218,1240,457,1127,1345,1066,1489,542,1435,825,433,543,1333,771,1037,1352,1164,1505,1382,392,244,278,433,60,878,1474,1383,1452,580,694,240,1080,1285,1215,65,1258,1178,595,1374,822,207,1210,696,759,1314,1572,418,1288,693,968,908,1345,1224,733,958,1387,641,296,141,808,1222,923,475,331,556,717,174,202,1190,381,1260,605,340,62,23,762,1334,529,135,723,946,1179,75,812,1307,951,1236,1134,743,265,1110,1076,934,679,931,1078,566,878,1337,597,698,530,1060,1375,114,137,1530,1403,965,1390,1067,217,206,768,1564,455,933,1376,1038,1118,1206,973,60,1187,337,503,1029,1368,1591,489,127,531,1597,537,1134,1161,1055,988,1367,184,558,881,1433,1514,811,871,587,1017,849,984,1083,575,415,1477,1072,1027,24,745,1514,240,440,71,1510,477,1333,930,985,706,646,1091,444,429,947,59,1384,1407,355,1013,1405,572,950,1413,1127,672,98,1523,844,1341,26,900,231,1293,1524,422,568,1353,1419,100,1582,523,148,199,68,826,55,930,987,295,477,672,1096,698,583,1153,224,1052,888,810,660,332,1099,1242,1484,711,1503,1465,608,203,1113,479,331,1269,1181,1347,286,1033,1362,576,1507,351,335,406,299,1209,619,780,671,1382,1009,200,376,1109,1136,1595,424,556,1208,1327,1323,966,1568,504,472,320,1078,14,475,815,1337,404,1267,260,99,267,1231,219,778,1462,566,649,411,822,1241,1521,1420,579,921,796,1485,1292,1462,378,1392,790,1329,1341,814,734,802,777,62,737,721,314,1305,1036,255,631,1340,634,799,1330,1459,41,608,880,272,543,81,1381,112,172,180,808,554,359,1285,765,109,237,696,82,649,726,455,1172,234,719,1496,579,1593,474,703,1311,1468,1199,406,579,369,83,497,1272,29,1102,1135,584,6,179,636,560,471,1450,96,1242,835,1511,362,359,873,1067,525,706,1269,1015,573,16,1599,1324,809,1295,481,269,1046,387,1540,1100,1329,230,313,942,461,627,1402,249,26,764,1393,71,435,1040,316,514,1177,590,1368,1557,834,490,543,1094,1252,176,542,168,1335,54,1048,1203,122,644,524,1235,258,702,105,1040,1314,1031,299,180,978,393,611,370,1332,787,873,1446,1188,57,638,134,1465,694,969,1488,960,321,344,1084,1422,1074,73,1403,932,701,847,1492,267,1544,1403,103,1542,153,34,299,1437,1552,1409,806,1528,1336,1060,885,297,714,190,198,992,628,70,183,1471,1454,1192,477,824,310,27,1546,495,1493,397,440,1520,1251,986,25,574,1296,90,239,1405,665,596,428,1198,1430,614,1440,1197,1502,1055,1448,1,1249,262,445,226,843,115,832,662,893,311,386,126,1566,913,1392,301,178,111,149,1185,133,622,1023,122,1395,574,1189,956,874,12,432,692,1523,247,174,356,930,1449,80,1170,194,1238,361,1599,443,761,324,1302,501,204,1011,1407,1258,522,405,1102,590,1562,649,1230,609,1281,1308,82,1508,279,388,775,533,14,300,397,196,1026,984,297,1416,806,1315,1217,125,104,1567,317,953,967,278,85,239,529,1433,1145,1152,545,1274,319,1226,1352,683,393,1007,589,434,1352,358,1016,749,751,723,205,253,328,597,1266,459,942,1342,1244,930,779,495,1096,602,1092,84,1392,995,1200,1489,1252,476,767,1097,785,937,504,1016,817,1503,88,1164,1121,1179,407,674,495,1330,1337,786,425,50,762,1140,515,1449,856,1481,91,976,1261,1484,380,1177,848,375,714,837,492,200,122,1294,722,343,141,592,1039,1215,291,600,217,274,547,61,1285,1200,615,949,495,1262,56,289,572,318,730,936,575,207,1563,1324,1355,1388,425,709,22,795,1029,670,562,701,1442,1575,471,842,333,662,1228,566,991,1096,1079,1111,533,724,1438,342,329,55,1220,210,1578,711,2,729,1098,296,16,55,1316,963,1029,1309,1176,1070,262,565,353,840,120,415,513,1500,1534,708,617,1596,55,1043,1173,1311,1331,443,1025,1265,957,546,906,1364,70,1334,1595,1186,1209,850,239,1245,820,78,413,873,281,89,399,901,1209,1295,836,1587,1102,1041,1571,970,1378,779,676,116,1443,1484,1003,984,1398,417,940,242,301,1190,1243,920,1135,1097,645,1529,980,1337,458,419,627,1053,371,1045,860,909,1394,672,1279,411,369,1206,721,1330,1512,1149,181,1466,1311,284,263,1345,688,609,464,821,533,955,436,1375,393,1371,777,1216,375,1265,871,135,330,286,1340,165,1093,1003,1524,281,119,307,1238,55,1578,219,698,183,1347,6,23,1348,1262,319,581,1580,6,153,1178,1176,1189,990,1153,99,679,406,1411,921,1279,192,667,1445,991,1345,358,408,1321,120,400,1009,1259,1002,163,900,1471,1161,703,1376,1387,535,614,509,126,1581,1379,67,607,1518,357,348,736,1554,1357,1269,666,1141,1069,606,446,1360,641,1472,224,133,671,1014,1086,1359,1350,713,1127,1249,827,1138,935,989,198,178,1529,871,179,53,1404,763,1137,1384,776,655,1380,984,1128,388,849,709,704,251,193,1382,31,1278,922,569,1283,351,1374,1151,114,421,1062,881,612,967,761,1204,258,468,961,1556,1254,223,786,944,88,1217,744,717,265,1517,285,1596,432,151,775,135,897,614,673,1383,143,453,389,545,193,1510,190,880,909,94,1535,22,791,340,506,916,536,203,980,999,1488,1346,1448,742,74,1004,1372,827,647,1044,1571,1104,212,1391,1191,1529,1073,353,1067,345,360,942,94,374,1544,1507,1511,1231,547,1546,1452,33,1027,153,180,1294,1559,321,287,364,1109,1490,1024,1303,324,1518,1236,1427,887,770,759,446,619,349,411,648,733,29,1123,78,891,880,1241,436,1386,1565,1411,562,1287,753,907,1429,1151,1012,995,734,115,1389,92,779,285,777,80,1352,125,1137,1401,615,591,1193,875,830,241,276,951,633,892,560,696,583,459,166,1055,1327,497,902,667,276,1574,763,204,709,1173,1342,1277,125,77,444,1053,125,520,934,1060,1017,934,207,283,1277,1456,1502,1135,738,1383,865,609,842,1566,200,185,53,285,78,1019,955,1051,1470,91,115,1478,956,312,1211,213,828,782,1450,233,1408,193,1548,1091,1140,1267,298,340,445,330,1091,103,1182,282,1284,1428,604,490,1203,66,915,428,492,823,775,418,1368,259,953,1156,684,427,1282,1445,8,54,1139,946,266,1399,58,1279,1458,1471,143,1287,800,1133,1256,70,474,21,1001,771,375,528,745,98,579,566,629,677,975,1321,1036,525,1135,926,125,1361,208,10,487,218,1541,321,1312,82,657,362,1477,329,988,531,988,27,1292,1101,862,478,474,1122,627,255,795,1130,436,1232,1438,1313,1247,803,1137,1351,948,1362,807,606,1241,431,1435,851,1012,1553,1550,118,678,460,340,644,1285,501,1325,193,1585,1325,45,987,114,1029,1073,387,858,1327,1128,1483,1335,297,315,1362,655,237,1304,390,364,857,1493,813,835,1483,399,202,1443,616,795,170,1011,1029,887,808,661,1346,841,1592,743,965,1339,476,3,750,571,1598,1360,1061,642,1503,349,1366,847,1438,412,1259,869,474,1475,997,792,1475,1444,18,552,438,970,970,822,909,849,300,973,240,307,453,1058,844,191,1235,580,729,1163,1212,268,465,1086,1371,495,456,521,46,1435,365,585,1548,14,1456,443,1099,921,846,253,815,462,1408,304,1329,1347,678,136,1175,935,86,511,548,1345,1423,715,327,815,310,360,894,1278,1465,1069,1429,805,782,696,969,1095,159,562,165,862,441,16,145,1532,1403,974,1588,1331,67,16,382,233,994,594,142,7,1100,462,1000,1418,789,410,233,507,68,960,656,981,838,38,1431,1212,768,414,746,1533,472,400,1169,221,1304,1324,964,372,1522,1132,1272,1255,1065,849,1486,6,959,961,710,16,1396,1362,1552,753,950,478,554,1552,553,825,948,1303,739,474,1555,805,746,25,333,166,1460,830,996,1404,5,1591,1237,353,642,1188,66,549,1201,1507,773,1252,623,530,353,346,79,947,1503,1396,854,1428,322,1593,122,1343,1348,1070,1511,1241,1227,1428,987,1445,981,1073,140,56,1553,1460,1423,1127,1544,288,1089,1,535,476,525,698,365,273,1159,12,1015,1419,192,1354,861,77,66,150,1587,339,624,1496,670,979,859,1101,82,178,977,116,197,175,144,334,317,848,657,575,1400,1158,56,29,572,1063,429,704,989,545,808,1335,904,1176,383,1031,1293,402,1052,153,307,126,1060,1320,604,829,923,219,1099,420,216,1286,939,590,92,1280,39,611,1387,471,1491,283,3,420,1307,868,1100,1014,959,304,241,1599,757,1234,7,94,849,1459,1545,1206,248,1479,1133,705,1265,1594,858,1212,1097,1487,487,1279,593,443,175,8,1158,244,329,1214,235,1337,812,206,1091,1406,877,818,271,26,779,242,573,37,42,965,1420,374,1124,841,902,1086,683,238,1488,1138,1042,993,610,1311,215,866,1592,796,871,739,920,461,626,1581,792,739,153,961,464,1165,140,334,1322,613,393,170,230,28,146,1498,1336,784,1205,600,1287,783,1149,1129,1538,658,144,829,96,70,56,434,640,1189,1118,1516,1162,232,1354,1334,311,1014,1358,1244,668,421,1324,635,1185,419,1153,156,1144,1333,1569,205,304,708,474,264,1204,333,1298,1588,1065,1489,1245,197,638,457,238,914,175,686,620,457,1367,1182,1252,261,188,903,156,1030,366,1440,808,833,630,386,441,84,1283,451,1270,414,1167,1175,1105,537,1249,556,1077,342,330,1151,994,1407,1370,870,1028,1215,556,1105,977,1455,46,1521,602,522,663,558,106,609,1147,142,589,492,802,578,155,702,57,166,876,216,1314,1334,1589,50,99,812,1409,1270,714,1003,902,1466,26,1303,113,1346,69,1234,1492,1369,1477,453,1338,1231,650,558,1444,486,1518,1229,257,1358,476,570,619,1483,858,479,1350,307,347,917,80,1067,1149,335,1347,1107,1512,1474,1308,986,1330,883,1428,785,1123,493,1415,438,17,303,3,873,107,226,1036,1079,938,1203,758,106,102,773,1378,1249,593,1192,958,1121,961,712,309,23,31,1361,653,1201,179,902,1381,178,190,1011,1572,409,1016,355,1421,1294,540,525,664,365,1142,154,278,681,982,1391,1210,232,143,1419,976,148,519,1142,621,1180,1012,1032,286,284,410,229,553,104,885,970,1024,95,993,1410,81,1447,1171,612,129,544,1247,864,226,1588,694,1197,430,1264,212,551,604,1519,584,765,23,792,475,569,1178,988,170,609,1563,920,305,69,1162,9,1419,835,1379,1377,144,442,477,592,420,555,1335,157,1022,963,1211,1335,480,1498,652,769,1248,1227,1155,445,31,1467,844,1226,458,546,1252,1208,1151,277,848,1463,306,1404,645,1020,744,1340,1578,1248,487,84,1545,32,295,1441,160,139,245,1600,109,1339,224,161,106,94,469,313,933,1505,37,1503,911,373,1390,195,282,40,690,333,1230,472,435,1549,168,847,952,1568,881,220,1151,374,1288,820,17,717,1365,277,872,1589,102,1406,903,946,1533,513,1235,892,327,335,17,97,151,348,511,1355,140,1010,914,1487,623,946,1513,463,1056,136,389,473,203,1251,237,1472,328,517,978,848,320,1075,1543,231,1103,1183,187,1599,1588,565,1155,824,90,1237,523,1189,319,1377,298,1168,314,1040,1097,858,1494,1085,129,1152,326,1483,123,812,1224,25,1333,1169,939,432,289,613,884,1591,978,1536,389,461,149,1217,1225,251,1542,1098,684,155,1351,501,99,895,107,7,430,862,1061,1336,985,212,500,1338,1015,603,278,287,1182,583,602,482,1001,971,176,1386,528,383,1019,494,707,1454,249,664,604,1188,301,822,134,499,1014,1382,1305,140,918,536,180,391,312,806,778,1371,303,1391,1230,227,1072,1335,1245,1268,1238,1244,57,109,1307,1285,496,1326,373,605,135,120,652,868,671,1564,947,35,269,1200,1586,1351,1399,507,84,679,318,279,1541,338,915,135,959,156,582,469,262,296,492,1194,164,1098,490,1133,528,299,778,1064,906,857,1082,1539,1249,161,433,41,109,434,753,284,665,450,1479,892,752,612,1440,1458,1538,939,156,307,240,436,1504,1081,200,1561,426,1139,53,383,1447,1208,1584,668,273,582,596,1325,1053,658,1182,381,1384,234,906,1465,1465,442,187,454,97,474,127,165,1598,753,588,1029,729,1188,720,874,992,863,504,887,527,476,1069,510,1497,581,791,648,695,828,787,1264,90,946,1373,862,1104,570,1582,816,1199,1273,1258,174,1043,530,298,178,240,997,472,605,1202,603,1555,848,220,267,1021,173,1259,1344,254,1521,1199,154,589,1545,1327,329,126,811,834,894,514,1596,77,1415,459,196,375,414,1416,1329,498,1306,136,327,965,1023,187,330,946,1374,1315,1186,1499,371,898,213,738,1356,1262,1481,1045,1155,1131,936,1254,516,1121,510,935,576,773,1299,1408,655,1551,567,1339,774,1535,1372,1158,651,298,1092,215,603,775,475,85,1098,1005,1359,1265,796,1570,555,240,660,662,1417,610,1341,619,1040,545,1248,709,362,609,893,359,1580,205,182,1045,1539,1292,1281,1419,863,613,1577,1500,796,1519,1104,366,1513,1025,319,229,566,778,679,51,60,461,1348,209,1192,1404,1006,1265,1236,995,897,734,311,1380,195,399,1544,136,702,1469,383,561,1424,1454,622,1004,939,582,399,856,1580,908,1314,451,580,1186,422,1062,707,100,193,374,600,1008,1067,1043,1232,1357,140,442,1400,608,452,316,1160,1377,234,709,107,1029,611,844,71,544,256,199,1409,846,1303,1395,1575,1068,1461,323,1562,340,931,815,980,1044,511,831,916,1323,981,607,290,1539,1204,832,1021,161,708,819,175,1037,182,44,891,1536,238,230,335,248,97,663,994,72,1512,799,550,1176,1006,1165,770,517,1266,998,218,419,1555,222,389,384,10,790,193,808,1350,1056,1330,906,882,1347,286,339,739,44,106,631,177,1279,569,17,372,1525,847,997,1402,483,775,278,1192,154,1235,516,998,1163,1412,145,1184,70,90,1247,407,445,964,1290,656,1425,521,126,43,564,303,1238,840,941,757,1104,1251,1277,488,234,1246,471,132,116,662,1239,1324,1208,1413,368,612,933,1449,1462,165,639,1259,321,167,340,200,1540,1337,192,959,1285,701,1133,1216,297,1476,820,598,1447,796,202,81,1083,1232,549,1539,1110,689,364,243,293,125,1539,831,1481,1561,1172,924,1295,739,1400,747,1224,326,1509,532,1061,1022,1201,346,561,861,1020,1143,388,1006,258,566,872,1021,1293,1325,1510,389,1511,1175,744,1427,212,1437,1452,1410,1498,574,1545,1520,862,35,463,522,492,867,32,1423,217,1254,1502,1567,757,1574,700,814,1141,1495,1445,1442,33,784,130,287,1477,1178,1227,300,536,125,1017,968,904,136,1232,533,859,1597,786,498,1515,637,624,755,541,1165,1422,825,55,1355,399,1484,1129,670,1292,305,109,567,480,523,227,991,693,601,592,466,974,1374,260,256,456,311,1162,792,545,20,1012,1155,1562,1559,1104,724,1454,120,603,1397,1544,1253,841,267,1557,735,1056,1085,135,795,1111,1144,1002,9,680,313,1233,557,419,1441,1324,706,111,1235,774,319,173,195,1171,1103,1019,1233,1396,529,659,905,1256,109,182,724,1241,654,792,262,349,1518,38,716,1336,1544,857,324,173,914,569,1475,567,1465,1289,1272,728,1563,1133,625,1495,1548,320,227,9,1315,601,1265,1030,590,866,465,277,664,1065,719,1456,798,1569,478,1230,1532,1534,1423,292,119,1094,1422,29,1187,118,139,125,372,152,1216,320,127,1557,946,358,204,197,1037,148,452,70,1530,1043,568,384,1504,167,829,1554,561,1221,212,895,1006,410,294,662,893,951,1215,1301,1217,897,156,1462,964,743,718,1028,284,1054,590,1466,1585,804,397,946,1483,54,1049,1080,1573,488,1149,664,785,648,547,1380,1574,553,1226,981,50,632,118,1591,1509,1131,1499,33,1292,459,92,1396,780,1135,280,773,889,1143,504,356,899,1447,595,1095,890,394,1230,1579,415,616,394,943,327,332,710,1038,357,1052,1369,500,244,1351,1073,514,1570,851,854,1557,361,528,1238,975,1341,1171,276,1302,112,1178,922,1419,1119,870,1296,361,1103,580,51,811,245,1421,1265,1553,358,681,798,1110,577,492,1283,804,200,91,626,1376,1306,1070,1560,164,151,1383,1481,752,381,715,73,889,710,721,817,1363,986,1207,967,1575,1561,999,1598,370,758,1354,1338,79,482,1474,502,1277,540,1441,645,1360,1559,127,611,1177,136,1541,938,909,1251,156,577,1599,1231,1318,387,331,328,1094,294,712,1262,949,1252,455,278,6,764,1360,1041,869,248,819,757,612,293,1315,534,929,294,253,848,982,800,1541,19,1350,1036,573,102,1584,1514,928,1091,493,938,139,188,896,596,301,1437,238,890,141,574,962,525,618,1291,1258,45,722,1123,1512,68,554,510,446,847,1471,1160,1204,24,1576,592,193,844,445,1253,685,1325,697,1559,1534,1175,124,533,972,1379,84,973,1317,1452,42,1336,317,1499,1258,1591,481,1252,440,449,1406,1238,49,755,564,919,1578,969,705,1546,727,124,723,39,623,1034,656,71,1586,1492,1599,1544,1068,461,1008,1596,1077,1434,387,600,714,928,1338,516,1296,624,1222,336,1224,382,969,942,970,1047,1002,237,713,1116,202,1398,1105,1367,1202,1561,531,1443,440,591,1060,1278,810,1028,1366,564,1192,747,259,1510,1344,1597,293,472,1502,1103,503,1507,161,299,762,835,1580,599,706,363,210,10,346,278,716,1457,1244,1355,1134,1073,1074,314,1187,466,519,499,735,1560,1009,144,565,1163,721,784,1510,634,1170,797,95,1128,761,453,1065,1506,542,565,350,982,1160,229,473,695,684,356,1576,46,626,570,1076,814,774,715,318,281,1133,559,783,414,945,260,209,325,806,1106,462,431,1320,1071,1204,1190,251,135,851,1322,585,1292,782,784,734,1033,34,637,375,1027,767,227,916,1458,120,414,1348,638,192,796,693,1071,685,503,1292,1192,117,1085,1388,1024,994,84,1562,872,1136,1504,1124,28,584,1361,1396,1507,612,1456,1525,85,1066,541,1095,1245,1265,142,1413,1231,247,338,122,405,1358,1091,460,1015,854,651,1286,1129,1435,486,1515,43,674,694,861,1357,1201,285,777,526,1355,1295,1099,296,1030,645,550,652,258,377,1423,1041,856,272,1392,780,1518,204,1173,1566,1273,838,441,866,281,469,1401,696,1417,503,1364,526,1404,499,647,1090,1444,716,358,829,512,173,1484,1069,1331,1468,41,1025,292,1441,82,786,669,645,1599,272,748,1168,1247,1093,720,379,438,1578,1077,1433,1146,217,710,1288,351,1088,451,1535,632,59,1511,1107,616,497,53,1277,585,885,1149,540,483,397,357,1548,562,1222,675,1209,629,951,181,1040,629,918,507,813,1485,456,904,1574,613,841,730,1472,705,648,704,247,1129,1021,218,1039,626,1186,877,271,1408,129,472,312,1547,552,767,986,902,1532,569,564,924,1254,186,318,773,730,1129,329,91,1544,38,374,538,1403,1519,1095,1466,964,141,369,467,1033,489,1259,1381,1362,324,852,1485,721,782,1187,1383,719,1398,1543,320,831,1580,1247,1500,588,1520,491,603,1583,951,745,173,1343,2,412,732,264,1078,317,518,668,563,611,333,1264,421,780,447,620,147,203,171,1083,188,1018,765,1118,601,485,331,818,369,396,1498,63,1260,666,449,267,432,637,973,1097,433,855,698,854,598,759,1249,801,316,657,705,132,373,177,447,437,1261,1458,56,1109,93,1538,780,1180,492,482,737,685,1573,929,875,646,755,1201,1416,233,18,437,543,795,591,781,1149,169,875,165,1057,1562,944,1306,952,366,1062,1459,938,1512,384,1147,236,298,898,1280,172,476,685,706,976,1592,1197,69,1038,416,1339,665,1071,721,1055,995,1192,1263,922,1348,1549,209,1261,37,86,973,605,1303,1224,1246,1069,775,1521,716,247,657,692,131,1301,1029,1322,206,1569,519,1033,551,170,1072,12,588,737,1573,1575,1510,88,854,1037,1112,589,838,850,992,1236,1523,1003,1260,743,214,1109,1348,321,220,660,710,814,1078,966,416,987,353,1205,1445,685,1129,812,350,13,622,903,342,931,1085,1300,990,1423,1222,748,176,235,742,869,1369,1532,909,1159,1385,521,161,972,289,1359,1047,123,1269,923,1349,1443,533,823,1058,1512,1395,889,22,6,372,990,726,1383,93,90,248,1309,347,994,422,1094,964,503,10,1123,1263,714,1083,1255,197,1424,1112,1163,1597,1013,1560,317,629,912,794,737,526,711,727,1563,1530,8,1484,1441,155,231,867,1325,1122,1180,168,15,956,238,1020,1477,255,1321,623,1256,364,1045,945,819,1199,404,81,357,390,164,1499,1014,172,158,531,94,1474,1490,838,803,1300,230,334,415,212,1138,60,643,246,39,73,700,243,403,308,937,210,641,1238,1351,1120,759,4,1368,241,931,890,701,1509,34,424,1568,1591,1102,993,1566,145,1175,586,1336,891,1359,906,397,936,94,463,856,636,1062,1408,980,1314,259,394,137,1592,776,1582,650,519,286,366,1235,1495,1154,1080,655,358,1417,836,407,1471,573,630,1042,1006,297,1138,1523,1465,233,115,1351,536,264,718,855,167,1192,1468,860,1235,1485,29,774,183,276,648,1579,1434,929,1316,1056,86,786,637,1409,168,1148,709,445,284,129,1403,1225,1599,975,1251,1028,1479,813,1548,626,348,380,232,845,545,1349,255,928,774,366,73,402,981,318,410,448,533,729,431,1307,555,1032,745,1220,388,290,1482,232,659,712,61,20,1467,120,386,1481,667,1068,563,218,487,1473,626,192,1240,1524,1487,332,891,412,874,815,1308,547,1204,731,942,1531,253,1058,419,419,830,1506,303,509,188,1144,1279,314,551,464,832,792,432,353,988,307,228,64,770,726,773,1229,1376,94,350,510,1502,1307,1084,7,659,134,929,985,837,27,1102,682,937,608,439,333,792,1172,776,1324,107,808,12,617,747,362,154,1235,297,779,1481,1158,1414,170,853,458,311,1163,188,487,1368,1401,1279,1463,671,304,320,673,500,953,672,499,610,323,1587,232,617,1245,77,923,1427,1478,393,1132,1399,1565,1237,1294,567,407,1045,437,604,248,1550,500,308,89,1053,1152,1426,195,108,1473,125,1502,1038,264,753,1257,357,1433,801,1344,311,176,807,1504,770,100,459,259,438,484,559,1336,1554,1301,1400,784,371,620,962,1416,116,413,1294,1490,1157,1076,29,749,378,766,869,660,176,1507,206,63,748,967,368,1440,1121,1140,1408,1117,328,1522,185,963,305,39,660,574,411,748,118,76,397,1371,668,518,551,1156,239,657,309,900,1599,358,1542,1346,283,817,1027,444,1193,1575,343,1277,1386,1097,599,658,1367,838,1064,265,756,137,291,1561,839,1014,1598,1187,837,473,54,875,1547,889,200,803,31,524,606,1542,1477,1163,304,1159,1123,197,1377,1467,1539,1098,37,193,667,606,778,1524,1432,1493,1575,1483,393,223,1229,1037,212,826,483,1286,385,763,393,1520,783,772,1060,768,1579,221,1256,391,154,379,1477,287,1045,351,410,751,181,71,13,1351,434,1570,23,958,89,365,1322,79,476,340,617,485,472,1444,802,518,493,472,1380,999,182,1014,1141,1012,1404,1428,412,290,18,1008,709,55,47,1252,728,613,223,33,932,1318,1217,1137,1030,373,573,410,1103,837,515,591,274,1586,499,919,1386,412,409,838,1335,771,294,369,534,736,1389,327,1499,1234,31,944,317,453,591,915,1077,447,835,28,312,192,60,382,840,63,1178,600,836,983,1239,1284,949,1008,477,78,786,1074,186,1519,51,120,219,368,1506,1209,449,594,1107,1131,1427,1122,910,1418,756,1053,835,976,683,834,1397,1365,1276,1579,1514,751,89,621,437,363,177,1018,1255,549,131,667,851,301,297,1129,1307,770,1514,204,410,1304,1501,692,573,466,749,1562,1057,1288,10,1276,273,1346,905,641,1173,295,2,608,581,93,317,189,775,699,120,1213,1350,1023,107,35,1419,76,671,339,251,850,414,113,1208,319,1424,1125,950,1214,765,294,916,706,842,1402,839,1221,1179,909,571,1371,368,1219,1552,1105,1178,1177,649,428,1309,939,131,24,1026,216,1166,605,1201,1060,1513,649,1254,819,527,1086,470,460,14,540,1593,1521,1136,1406,1417,1504,155,238,1388,1079,719,662,153,1090,764,1115,129,593,526,53,512,740,1285,760,353,799,1095,51,666,601,1187,632,1020,128,448,923,1042,148,816,1295,326,107,343,621,825,1162,909,1390,792,1277,455,124,1213,1126,653,965,904,1061,1385,1178,1083,407,866,1312,1408,1123,378,1177,501,1585,644,370,394,336,607,62,1303,1422,453,1068,1560,1564,671,1391,449,927,931,1504,505,148,830,891,46,731,1538,111,963,419,172,829,1586,1454,1180,557,1579,1320,393,583,698,329,1550,106,493,1000,766,747,1309,861,356,1328,622,211,329,895,1053,610,1546,555,1510,1283,1059,679,254,607,1307,1442,1427,1591,197,1232,951,1068,866,883,15,558,1081,1099,921,24,1234,220,205,448,643,81,1399,1473,984,240,189,1270,766,600,2,952,545,677,970,1365,1111,235,129,1352,687,650,833,998,653,1254,1033,1141,877,774,743,831,1283,166,1186,844,463,1410,181,464,889,560,51,1283,114,1059,708,595,1069,300,661,620,1212,45,275,1519,970,669,1184,1584,408,957,1003,1267,409,33,659,467,670,890,846,1244,821,1577,866,434,589,912,1297,551,618,213,295,455,626,775,459,1028,1249,175,720,1344,997,1210,303,894,678,804,233,480,909,18,626,650,613,409,1434,468,1496,571,450,256,1103,1008,1594,1154,1302,1411,731,1434,110,1031,1053,89,1168,59,920,1022,705,515,586,132,1373,576,1239,115,355,446,1537,1336,168,319,428,910,1208,1222,511,222,1463,755,918,455,269,428,125,1098,1105,91,1498,1120,255,1350,465,1350,471,210,17,186,1021,900,1316,1261,907,346,769,516,750,1474,609,767,1138,1174,548,436,1125,1333,1373,1245,1322,1573,273,755,1395,344,598,961,15,1084,794,192,674,213,1462,991,1503,1418,390,51,395,33,899,535,535,468,298,1568,231,1444,1021,310,1036,362,1376,581,849,1506,1490,712,513,1428,896,994,1165,615,1576,45,1096,395,895,529,1453,1024,1399,1112,836,1075,1017,1429,1170,1192,282,990,963,29,95,452,1446,376,831,124,1033,522,12,611,1011,347,1530,61,333,1432,161,156,555,203,1335,43,1407,1510,254,271,1023,559,1358,1128,579,459,542,1405,826,176,380,601,1423,1507,473,977,377,1306,297,670,707,818,927,430,247,804,460,1115,56,1220,259,544,73,177,422,744,1319,860,249,354,1273,181,365,1573,672,469,706,234,117,104,164,539,1373,384,1045,1532,530,1042,51,438,522,1434,618,102,318,702,1270,608,34,143,1533,1043,19,1266,1350,346,1193,1405,1587,1240,1346,1423,1215,602,1137,642,1087,630,1563,374,1454,363,158,699,1084,1425,793,1187,168,891,200,1461,299,656,1154,1484,510,1566,820,889,918,1471,909,655,838,134,952,123,181,1585,144,635,955,482,502,1360,701,1071,1144,952,727,23,996,189,276,908,579,757,971,1203,1559,766,425,1231,1347,433,1510,1005,719,1264,1355,871,772,643,1409,1405,1505,61,738,65,1147,849,1475,895,320,199,1573,692,367,935,899,90,1324,1033,749,1153,1276,56,1015,1166,680,1432,1222,472,1597,1447,343,225,1000,873,975,1292,330,417,1239,1028,294,1130,392,1572,1461,1481,107,1106,1196,1248,954,671,125,1326,1513,1078,1198,749,261,643,1430,913,922,370,798,354,1470,1405,587,362,1362,337,216,1434,453,813,709,1212,93,966,29,1173,572,657,252,218,64,697,534,780,1102,28,1565,596,1495,505,587,1029,124,66,946,336,1361,561,842,550,668,915,1518,1228,1593,1379,1075,888,1113,195,1413,490,833,604,971,1528,51,550,894,1058,982,1292,396,1550,800,486,593,835,77,1069,538,1422,1046,1199,175,365,1248,1217,1122,1600,782,938,920,221,1093,269,673,775,856,562,196,120,1598,1508,1472,1551,845,1290,987,1164,1107,1152,201,1483,377,329,1384,2,1559,581,289,1204,1309,602,1083,728,1069,228,468,877,454,627,404,1333,1569,1514,427,985,897,474,1522,1404,1179,645,1375,252,15,1137,524,285,304,850,75,857,528,37,279,1147,161,170,983,1103,1139,433,1412,261,1067,1136,476,1043,331,832,322,112,209,411,835,51,1537,1551,1366,1336,1511,384,1165,703,392,530,761,862,777,41,1214,611,365,1567,139,1444,972,1337,1027,757,335,1324,558,780,969,921,1530,605,112,846,1284,1201,1542,430,243,1369,42,1579,313,764,833,87,704,1226,434,84,1543,306,1089,1584,537,854,733,915,1398,107,824,106,248,227,1516,71,899,1370,772,848,1159,962,437,446,1564,558,953,531,288,1111,522,128,824,812,545,305,384,1117,324,705,544,415,775,1383,268,472,1428,98,1475,1323,1045,980,504,64,648,162,1505,528,678,1297,375,931,1103,101,28,765,905,1449,850,1583,1595,663,468,144,1159,851,455,1545,1079,638,426,143,400,167,1068,98,602,340,761,796,313,196,1477,865,1314,973,927,1578,1508,1468,1358,1523,736,1316,1445,338,1446,1175,1218,519,1512,186,939,1566,1108,1306,1084,76,520,980,728,525,1059,753,310,687,1408,468,311,1465,1301,1260,665,39,90,1200,641,180,370,532,461,1577,426,279,1386,345,1431,1535,793,416,1156,1515,686,183,1051,145,666,1567,1533,1419,595,39,913,853,334,619,1060,3,1450,1503,1174,755,1335,377,619,108,942,416,181,1374,718,1591,670,495,1454,1572,1393,825,346,465,1504,1219,268,8,949,5,911,953,690,621,996,1097,1593,1276,858,705,431,108,670,1203,911,677,1094,994,227,57,14,1262,533,1440,383,1188,863,7,904,1176,53,608,28,1337,1428,1247,1147,802,922,601,85,1299,1131,18,75,804,1176,1350,114,618,1408,760,195,189,901,543,476,1382,1000,1085,809,351,521,885,1198,1284,1561,1134,820,265,1285,1360,905,872,66,342,1121,1348,851,137,182,632,1128,77,832,785,973,970,1208,1028,677,1448,234,1521,685,73,1394,1197,582,1292,398,1362,715,1363,451,210,127,646,1302,153,1428,185,494,1327,916,291,1288,666,186,1375,695,1139,743,528,1268,1368,56,1559,358,1527,1508,1025,995,546,1459,1465,1595,975,429,718,1008,483,578,221,12,857,389,1347,61,1373,687,237,76,28,1117,31,143,1510,600,335,853,634,94,704,1102,1487,1305,736,877,579,1234,785,833,1481,369,1426,42,77,1544,898,1573,241,311,1070,182,446,1483,61,1278,397,1211,294,640,511,956,705,1292,993,845,1052,45,1230,138,873,1177,1174,831,171,694,508,1200,842,657,1124,1432,274,1487,1006,977,731,642,1013,967,501,575,949,726,1499,1310,1349,1463,1536,800,255,544,456,1317,1114,877,1086,578,1269,338,711,847,190,830,962,1393,1527,619,884,652,43,691,179,1313,273,400,1273,912,60,1528,1141,210,1201,320,1290,405,1326,652,797,448,965,53,981,1392,852,913,1436,40,291,630,1452,209,1064,1204,1196,826,851,225,1441,929,748,1159,46,1043,200,572,654,1400,1405,1355,1102,1268,926,1036,559,181,1125,1275,763,1379,740,305,1080,1118,577,800,928,553,118,1082,645,195,1070,886,805,1164,1022,493,540,1396,714,981,403,1203,920,706,91,667,1216,282,1344,1298,1142,375,38,756,564,15,1338,389,1303,63,1047,109,922,510,605,1132,1083,1237,186,1011,1516,267,805,999,480,638,37,55,484,1328,198,235,1337,1594,962,1311,1071,1046,202,1405,454,1182,276,746,353,1239,1293,188,1390,798,1186,311,567,731,990,291,409,429,1454,988,536,847,664,301,20,526,1191,712,1464,548,450,489,1455,1293,347,1524,149,1126,935,1264,1491,536,494,613,31,1361,74,1573,538,849,490,136,848,1324,1264,325,937,475,1482,21,505,401,553,914,495,1383,926,707,529,1155,1418,804,354,1231,434,480,1161,1338,1056,174,337,1244,1331,340,843,1221,345,440,162,217,1232,439,183,538,571,878,704,1057,1294,817,944,859,921,757,249,266,456,836,413,343,874,920,119,910,1510,797,1455,1263,1018,787,704,1179,1474,891,1296,1484,1015,428,1102,511,1394,471,1449,1075,1549,153,612,145,459,1379,24,1271,1526,342,1387,1190,497,119,1236,1460,907,240,1149,1368,519,680,1312,1080,210,1450,148,243,187,726,1043,478,661,1197,53,1234,912,581,1412,809,1282,599,1119,688,327,399,5,486,1068,1458,758,242,1117,1335,204,1176,1175,1505,937,877,559,651,669,110,1539,36,691,848,182,733,760,656,1251,203,724,66,22,1246,1262,439,1560,1529,1083,1022,946,1312,583,652,1134,1349,100,765,250,706,108,265,1402,1327,605,1225,1194,1318,190,38,1464,101,360,1496,1212,604,1036,232,1124,1129,518,1182,853,505,671,220,478,1056,932,1131,1456,285,551,361,659,824,494,279,1303,1341,1125,1589,1076,138,1413,1252,169,1416,754,458,203,811,1501,542,438,503,136,1539,74,582,1105,1553,68,1515,789,661,1050,790,1447,450,1076,745,633,1290,847,932,667,409,175,1557,240,489,1013,1065,674,1401,274,945,1204,175,506,857,63,230,753,672,888,169,1281,1074,1141,381,150,1295,242,803,788,679,87,681,600,1119,384,50,930,1003,1017,643,522,1059,297,647,797,274,1247,1583,745,241,835,341,217,1532,523,1568,1354,1433,315,959,376,893,209,824,1049,1439,381,1576,843,912,346,1560,91,1452,171,1075,1421,145,1256,803,1471,445,1275,153,773,795,430,394,1130,614,383,821,1230,1434,631,128,1585,1379,1306,451,1350,1378,943,454,230,981,614,450,577,286,75,1324,338,573,1483,980,1036,1277,1524,1482,756,878,50,1141,1424,1399,508,1421,1063,704,1401,785,815,1448,31,251,920,1166,1533,474,635,672,1019,1069,1580,777,292,357,999,1482,1153,709,253,467,611,1331,26,1514,1025,1371,752,624,1095,1299,487,470,1585,1020,226,184,800,1177,365,469,346,1255,576,151,75,94,1561,1402,70,744,26,781,1046,707,1394,712,79,1065,1167,884,231,659,487,13,130,830,1046,1282,1503,630,63,1146,1084,307,157,1308,63,616,239,1309,1107,131,706,720,968,806,317,397,1101,739,1600,1167,899,210,1532,1051,976,1542,52,1320,871,1552,493,692,1407,1204,1040,795,1346,1418,746,1235,1306,1467,1111,1499,130,112,1435,954,552,1593,1273,698,331,1523,8,1244,389,789,1049,596,325,239,981,1011,1382,1239,593,722,96,564,14,441,407,1426,714,261,686,344,27,757,31,923,409,1518,1118,1201,99,206,1285,341,1056,1431,395,790,116,627,9,965,135,251,392,605,598,115,1181,242,212,1125,669,416,925,402,1081,367,541,627,1588,674,1038,336,1377,1035,891,425,965,1260,486,1067,140,1343,1410,346,752,483,1349,1031,1112,1414,1267,1549,127,340,965,546,504,250,1530,1375,777,783,557,1555,1079,323,344,303,1199,469,207,1398,807,76,1216,1228,930,963,1060,1397,1047,614,1359,1018,226,1364,1406,295,1395,149,824,822,109,1327,1149,1321,610,517,232,1591,1285,742,892,754,1357,940,266,903,1495,1106,361,1057,1432,149,719,868,236,1517,757,217,601,416,185,518,764,263,1012,1571,796,1004,1321,322,505,1369,988,594,1521,712,1023,1572,684,1178,379,1599,759,449,79,391,107,1149,1512,1509,1497,1437,636,1169,815,1352,1546,1128,975,1199,1595,796,801,312,600,1403,252,520,1036,1250,760,1292,1061,1487,391,1089,767,1450,171,1477,618,866,147,1251,1227,1404,1061,1222,583,1205,912,424,121,1594,1198,1545,1340,218,1373,13,1565,998,891,817,1122,1407,1044,1508,1543,1102,1526,1456,1231,1437,979,1055,1176,537,983,11,164,811,1106,1268,1438,470,335,387,1411,520,227,124,674,528,196,793,825,243,793,80,26,405,47,1377,408,98,278,377,1439,1366,728,488,611,529,288,1096,1112,1033,138,153,1398,1369,1430,587,363,191,563,656,161,1208,1003,1498,831,421,622,1152,1231,909,1507,376,355,229,1396,1311,1071,485,1201,4,210,801,952,906,482,280,905,1379,884,1086,475,1034,68,369,227,1384,180,770,1156,647,1047,899,1041,601,1046,745,338,1188,210,323,1581,1041,732,1214,1137,259,874,1586,560,396,1001,786,582,1426,563,230,889,936,1078,936,818,1469,672,1355,1184,118,1066,1161,317,281,449,469,453,100,1348,300,152,409,1308,1081,136,1215,560,875,576,595,435,765,920,1131,917,295,887,1202,1040,1098,1430,918,1368,695,196,87,372,1285,438,478,314,271,590,1055,643,962,1438,963,1587,682,1028,920,1130,817,175,984,1088,808,995,156,1214,117,1164,1007,301,503,1421,610,46,1060,529,1223,785,1589,864,273,840,1205,759,145,249,118,1575,392,285,970,1069,1254,1022,1589,1375,486,1063,1143,901,787,1328,1578,1520,1142,1570,1412,906,625,1408,5,1598,1411,640,1328,90,180,619,202,918,1019,114,1219,155,739,1100,384,1074,856,483,833,895,930,676,449,1338,1358,1345,12,141,1518,1087,1074,1189,923,1234,550,856,1189,113,689,110,1174,380,452,1165,1256,1594,1448,1034,781,680,1094,1146,16,838,402,1099,1311,749,764,1477,671,1368,269,612,15,304,194,808,1077,670,914,532,1278,1186,831,541,1425,679,1525,958,664,277,1481,29,1159,250,341,1366,52,675,756,726,19,525,267,1337,1575,462,1060,675,1276,789,867,891,170,1050,10,312,1410,725,753,600,1001,236,670,1206,1018,1072,317,294,18,783,11,1308,69,72,1315,780,233,1428,1310,1229,1276,1113,1428,1441,41,1435,1360,1115,574,221,277,118,734,1062,467,1596,1308,299,843,1388,980,663,692,443,538,991,1434,1253,1219,581,1571,550,598,1278,612,67,638,1099,1302,711,664,329,102,115,593,1578,1583,1480,1533,1158,171,1063,240,251,379,227,349,1486,554,444,636,1491,313,695,541,338,255,700,425,831,279,1068,907,253,1333,102,1119,960,356,43,975,219,1161,1395,1044,1447,547,1469,1364,86,1104,1150,1441,86,1079,1490,1294,1153,1006,1581,783,542,388,887,1243,331,1447,966,1446,125,1032,1289,927,630,146,485,811,321,1362,1546,1584,117,1175,157,1223,1200,901,1290,667,1079,574,670,1202,894,598,81,26,1006,1181,151,166,820,184,1430,1420,571,73,525,549,1261,1112,530,604,1547,1177,708,23,536,1481,113,1506,450,539,527,1186,117,799,397,578,860,607,932,602,164,22,732,761,25,431,1587,434,665,512,86,1500,1052,675,731,1371,572,1597,364,726,707,827,720,656,33,1599,987,448,1533,1521,506,1152,1397,1392,1583,1438,606,1032,511,821,1434,222,555,185,716,1213,163,917,766,1184,1080,786,1220,670,1249,1373,1071,824,650,1276,842,1280,930,1178,1539,140,296,1295,1542,998,1232,1591,739,32,1387,179,1322,367,872,935,669,978,1056,1344,640,234,1141,1280,1151,1128,736,310,169,403,676,1430,891,853,1219,941,858,1478,583,465,909,1434,1223,311,1419,537,1370,397,1088,678,1459,1259,903,1439,590,781,634,333,966,954,1179,299,746,499,741,222,400,1120,46,1174,1594,1265,507,320,353,9,171,751,356,1383,961,832,156,32,1495,1536,690,229,479,1273,1562,1005,945,1387,1145,932,690,89,361,757,501,186,1043,1385,396,956,1213,815,37,446,397,410,1166,278,1260,161,613,256,126,83,614,946,204,962,646,1384,894,1309,669,1265,536,746,892,1092,707,355,164,370,422,161,172,1565,978,242,611,1279,56,1344,1213,394,1045,798,748,1424,58,1109,651,1157,501,1310,385,1331,1298,67,783,1534,812,1242,200,759,1213,1566,578,388,521,62,1573,1046,601,1486,1474,998,1457,567,602,42,397,1383,811,95,543,1187,1562,1390,960,1456,715,149,1052,1468,1347,918,986,20,995,1120,1173,878,908,798,697,1403,254,1490,161,92,385,454,916,64,1058,163,290,413,424,1119,899,287,1305,953,429,154,758,1120,1149,873,732,907,814,865,310,359,496,835,248,666,27,969,1570,109,725,523,466,571,773,407,1028,1144,1056,647,657,452,455,366,975,142,953,154,111,255,1156,1136,1005,873,315,360,863,1551,831,829,697,752,85,509,990,181,1149,1537,1540,1218,22,1235,501,943,766,521,1564,279,781,2,1181,1023,175,244,669,1062,45,1572,913,71,1138,215,710,1148,70,1308,1020,204,561,179,1079,682,1140,709,84,1497,1562,1142,345,1397,200,1513,105,1117,222,1149,1021,1538,961,734,1352,1379,281,1413,101,1305,1073,822,900,423,1088,43,242,650,477,217,1240,1494,32,1027,1052,908,1294,1000,208,1088,122,1576,980,1335,24,719,785,1202,60,312,1164,809,390,144,418,104,523,999,208,1283,663,1505,1102,1023,1324,32,212,1442,588,308,517,1495,416,1178,969,92,1427,1022,1381,1140,1005,885,444,466,1593,555,771,1517,321,1348,639,187,1305,1452,1535,313,895,167,786,1393,1007,928,357,1162,1185,516,1569,979,1436,992,1453,158,1157,416,1317,185,1328,218,958,539,782,286,1028,134,1211,1255,1526,837,938,658,1399,596,1384,87,666,349,251,1246,568,319,1019,1423,1492,699,1443,164,1098,613,651,704,1333,147,1460,1378,119,124,1057,1121,1003,1108,114,1198,1260,722,1352,62,723,403,125,890,1287,1215,150,1555,1311,387,1181,108,730,465,441,1346,155,406,209,403,775,290,1538,400,1261,1155,509,202,1147,1594,1220,1285,644,1182,762,1017,320,1148,1500,855,1053,836,205,349,686,35,1243,706,32,1222,105,491,351,999,1081,1353,977,1486,294,834,1278,575,1149,1311,1290,1404,382,1439,421,707,1319,166,505,838,1426,978,838,693,442,820,451,110,642,362,827,207,407,790,1324,1289,708,307,1507,1013,979,357,571,1343,1402,630,1201,1082,1129,478,1375,1397,1575,633,1036,737,1124,337,1307,534,1092,1267,334,580,688,739,1425,1454,552,672,514,1599,759,1278,1147,635,254,1000,1541,742,150,512,1399,528,1509,132,1494,432,270,1384,965,1581,525,1161,1002,668,952,1519,660,491,665,677,1536,200,39,438,1135,1228,952,175,1023,1449,431,1258,539,1094,768,1113,1250,247,141,266,1359,388,64,471,276,327,491,1598,638,1201,414,1101,832,51,728,500,118,859,1327,71,796,1351,843,745,810,489,659,298,688,30,1248,1250,695,37,1345,1104,320,774,92,1043,1204,638,1174,379,1127,910,10,836,1399,437,741,1254,951,311,867,1025,1434,1461,1225,592,1567,730,830,159,615,705,45,1009,1495,1421,136,989,686,1110,713,817,1021,290,1143,14,470,1161,345,732,243,649,460,149,964,1596,1353,1472,943,1460,618,342,934,177,1283,772,1285,339,1375,1139,1373,229,751,886,1233,1373,481,1208,1136,252,217,47,360,217,736,1494,66,1195,1489,467,1035,767,1441,285,195,521,264,677,74,1512,249,835,1256,571,1562,858,896,605,1304,116,87,349,1257,917,1207,1225,375,515,426,994,728,573,233,1599,426,1507,385,522,1435,1479,591,1460,465,230,1303,236,1021,1390,867,839,1430,276,1472,1572,1134,487,1471,59,713,1441,1172,1457,1171,910,453,91,593,1173,1554,265,128,870,1184,272,1552,861,1497,532,848,951,600,251,623,1021,1174,1338,1480,241,784,665,192,753,202,56,1440,1532,1255,504,270,43,843,1224,99,570,1047,1561,1531,1407,1288,175,749,699,1182,303,929,1238,1450,836,742,770,1402,1260,1500,1,1245,1093,1251,807,120,225,436,457,1529,859,273,806,726,1177,101,1101,1094,99,734,1419,125,1353,453,1126,1322,28,434,278,1109,325,814,714,1260,34,132,587,1114,936,233,96,719,1461,865,1587,1053,246,5,783,1128,121,133,736,848,803,1221,184,862,994,979,1341,1286,1359,199,911,807,636,350,984,522,1083,915,1436,1243,653,1107,392,1286,530,1211,1447,262,860,363,198,1205,299,1155,389,1303,998,581,145,190,1586,63,1292,1198,1257,284,15,1392,1002,21,64,959,326,401,545,644,353,1535,70,1246,1081,1182,1087,264,1138,1191,211,293,361,306,378,919,1152,626,1113,1331,821,363,134,163,875,1366,1361,1002,169,526,295,170,941,614,1181,414,824,251,1540,1575,1296,115,589,125,1326,822,24,789,1207,849,212,1138,811,512,589,512,431,926,323,36,1139,40,371,340,1062,399,1441,1496,713,1492,102,266,913,703,447,610,520,1291,608,688,516,1225,395,177,294,471,1202,1382,1289,1041,887,223,547,665,189,327,330,506,1524,1447,1044,384,59,574,1521,1419,699,688,581,1124,1496,1537,1432,85,296,59,1497,412,1098,1228,1539,1190,1560,194,154,403,807,1253,240,957,315,1325,642,847,507,1517,363,456,498,46,152,1426,935,888,932,147,793,791,416,1439,1451,146,646,1009,193,1248,530,1524,484,1245,1571,183,1353,1586,900,1590,1540,1054,97,141,1100,1187,910,953,1204,299,779,629,1054,478,185,412,400,156,1261,28,736,451,149,432,1046,87,1565,1477,178,1358,1085,311,1339,212,839,1021,658,1226,776,924,8,1001,1482,1218,725,783,1317,1075,559,965,311,536,878,175,374,5,189,129,1413,1464,76,1130,350,1462,920,1032,490,946,196,70,770,10,866,1147,1022,381,169,20,1226,1593,552,332,1455,686,1034,461,1506,339,22,1211,1502,155,479,1105,815,1214,62,1569,181,669,664,1337,935,1041,1184,612,1429,811,676,13,821,900,1165,904,1079,1542,587,86,1241,1211,876,328,216,1398,1260,1318,653,50,1127,689,1041,81,392,209,500,666,1265,737,363,960,797,52,1211,1578,770,1396,583,831,1358,310,524,1038,1214,477,1047,1590,1532,982,1044,237,1304,664,622,1491,1485,434,1037,458,988,1297,741,1286,1437,1079,773,1581,782,1197,36,518,907,1112,608,236,1580,1454,1014,143,1358,103,528,470,734,562,8,185,1500,1344,354,618,477,820,1437,1504,1600,239,558,1393,1286,1066,679,1536,1364,1459,47,1285,1149,724,1261,582,186,1507,475,806,1347,212,505,1123,1368,1360,176,1496,1362,712,409,736,825,1340,1294,1016,829,1358,1141,1480,454,375,716,136,732,565,949,738,484,1551,641,129,365,991,839,1219,1007,1065,926,730,884,1457,441,666,1127,1130,295,1540,828,1278,43,220,186,776,383,1456,565,1183,944,725,255,242,1118,517,1076,685,795,1062,1413,1264,1002,369,913,1513,831,363,401,365,1126,266,252,611,1575,1584,1238,236,441,1178,1570,617,1020,574,747,1162,1232,478,632,1205,314,1197,1515,924,578,35,1203,1435,339,1218,1129,366,39,1252,1009,1138,257,174,1037,276,501,904,1529,1113,1110,1391,1411,500,714,1357,1571,537,895,989,1578,818,679,474,372,480,818,1544,1115,288,1520,252,1567,1078,610,421,786,275,738,201,44,417,844,994,170,804,967,1010,1166,888,543,88,1258,28,1152,589,588,1310,42,720,1076,1108,262,1175,125,665,1475,572,852,793,771,447,1363,1013,114,1504,1056,106,166,159,505,1291,39,1206,933,753,572,1109,1234,210,678,1386,919,1215,484,938,1147,1412,444,1136,621,1280,674,369,1137,1489,521,1254,973,707,277,1129,1207,240,1415,429,339,1506,29,691,892,1226,654,939,892,560,545,730,692,490,1083,1449,1042,581,1379,1547,54,156,450,1133,1456,459,193,1591,1069,664,1064,937,759,581,398,1274,584,741,349,636,1197,1321,749,768,384,370,221,925,696,987,247,67,573,1015,1474,835,725,702,381,1216,805,1427,736,859,673,160,492,355,1083,102,807,1508,475,718,1048,983,252,780,1508,1495,1157,703,38,1173,195,874,29,1100,177,443,1443,671,1513,744,165,1114,708,863,512,1171,250,968,1125,831,17,226,787,325,1188,616,1521,108,1595,1364,1359,808,118,1313,402,1060,51,734,314,686,611,222,1588,1155,432,1247,1565,806,1561,969,131,129,914,201,1360,217,1482,1003,1264,1333,352,558,270,45,621,1060,923,573,40,573,73,1441,709,878,201,1012,1482,360,1121,1384,98,1202,120,1424,1479,1116,1027,1251,1228,1438,1302,246,381,1282,559,1180,1090,416,718,1304,926,1451,131,878,1342,229,277,1418,1376,199,160,1369,846,1315,464,432,857,1022,1435,455,473,209,1592,2,228,246,879,408,726,1042,106,472,933,405,1310,1008,63,97,375,1001,1100,1503,324,265,618,1529,1361,603,1573,1143,417,1095,899,1034,639,345,580,726,1025,306,1082,829,506,1592,1377,585,1512,362,906,534,254,1367,837,59,545,882,126,1493,1086,1537,690,1173,116,208,950,1438,837,510,665,1518,1415,11,1368,666,544,38,1375,615,100,960,625,508,993,1005,190,95,732,282,210,1286,111,1398,1598,1045,283,346,1067,998,523,465,581,948,729,898,1266,1264,1477,1463,355,1392,1596,149,1106,757,1358,356,442,1564,882,110,722,1124,1221,771,1149,778,1504,943,300,1420,496,1586,634,350,814,121,592,1255,893,664,1228,86,1087,1472,1402,931,338,119,1557,1557,430,1047,628,1436,1555,450,406,1016,1531,964,330,689,459,965,539,522,810,51,1337,1271,954,457,329,231,1553,1328,635,1310,94,82,1110,90,1108,54,251,59,1459,1365,265,1369,15,912,1027,1360,156,1211,1533,922,1373,1372,625,1215,1336,156,1178,870,156,128,1,949,965,408,998,491,775,1099,497,1015,552,780,656,707,390,505,1482,628,64,1493,936,26,72,831,173,1546,372,1084,616,859,668,576,862,974,31,1325,1375,282,1123,1281,365,1371,1260,440,35,1540,31,728,1004,297,111,1131,632,396,177,370,1528,1350,1075,705,1515,119,591,766,694,1223,1134,83,167,782,1104,100,1547,196,220,398,1206,463,1491,1106,98,1366,1048,188,1389,544,665,1551,1263,1562,1166,923,764,1529,176,781,25,870,700,280,917,1269,800,731,1459,725,49,117,199,1061,187,444,411,437,482,1371,884,950,1591,848,1465,583,814,239,514,1021,1493,649,624,1347,713,384,875,587,518,940,500,767,552,724,194,1067,1525,538,1543,1056,164,759,561,153,1295,1002,1575,1306,1182,1209,1013,162,463,1484,428,588,565,577,1552,966,327,1257,1459,823,190,719,473,1171,1595,845,664,1076,537,626,1416,1437,1453,727,793,3,1260,1131,891,1052,987,628,1294,1393,1149,624,183,140,1169,591,1054,467,333,1471,1336,1166,2,108,1152,293,1189,495,1496,1399,865,841,984,791,961,85,653,660,30,1153,357,592,1369,817,980,1263,746,84,1429,7,774,581,1548,1020,832,1463,1546,447,389,719,138,451,925,712,482,1039,276,1287,695,1500,242,1448,778,548,501,53,1267,32,1509,453,315,1546,868,471,1387,1490,1027,1236,464,234,369,426,608,1226,874,81,280,1563,1286,1326,396,1028,1313,992,365,768,734,1514,1046,1599,186,1362,1440,323,998,651,1041,921,1002,314,946,231,398,1209,1377,552,171,1364,1459,1378,517,198,23,1230,808,1133,1081,1232,798,73,99,1343,951,350,712,227,851,1058,209,987,713,1155,621,318,349,1447,1162,839,326,1567,125,325,742,1421,919,928,171,1003,1007,457,550,391,1196,171,417,24,67,610,1469,511,822,1188,1085,1442,1094,40,1028,1447,929,304,26,1408,992,498,1160,36,1372,1078,1592,1131,1016,15,765,1419,210,539,768,803,163,1044,820,108,1194,1367,1182,466,802,580,1168,133,337,554,676,1031,255,549,789,1168,459,1383,921,1035,670,514,30,589,680,1218,452,419,1189,910,428,1210,769,648,378,446,628,473,908,610,427,941,1331,289,579,1311,75,1462,902,1111,1293,496,834,226,1565,663,203,637,99,1353,1586,442,868,102,1553,895,408,589,82,647,1565,1343,473,522,494,672,220,978,1090,256,58,402,16,1346,1162,566,660,210,1138,52,135,585,1189,1388,81,1242,791,50,1434,682,133,1061,1031,717,241,1404,958,1467,183,1254,790,1335,1409,1222,408,701,1272,1453,301,15,616,281,681,1519,380,619,566,454,1111,233,1435,188,1267,1,1568,951,1544,1235,290,1177,1167,667,384,989,1306,1226,1178,65,1196,611,590,44,789,1303,1259,357,1045,1090,1122,1438,1414,788,505,1245,1047,1281,1341,76,1147,1339,787,220,741,937,1422,143,193,593,364,1537,928,1397,1050,234,645,303,825,415,508,439,676,925,623,821,551,1503,1120,334,1409,638,957,1321,913,279,1356,1265,621,1460,1075,371,46,913,1550,971,169,87,1495,1281,1126,869,200,1241,1542,1509,1475,560,1042,1528,1558,67,1428,822,622,1026,65,947,1416,412,874,1116,416,194,86,1181,167,1460,703,1013,1570,714,589,810,1515,1266,1173,216,1528,1385,1075,631,1295,844,896,346,416,204,229,1029,437,47,1395,898,628,965,542,682,440,88,173,410,1532,297,1344,983,1142,612,1585,1321,1388,1383,272,390,310,75,129,96,1533,1430,1380,554,1049,1007,324,1171,998,923,1272,474,931,422,386,1577,926,1252,177,16,126,846,1341,1208,1030,476,1442,833,700,893,649,1527,338,1236,171,1361,1452,894,701,1064,673,1555,901,885,577,37,1507,166,58,883,505,1494,294,791,942,993,12,1302,635,396,373,205,1185,171,900,201,296,470,1369,753,645,408,911,1350,475,593,225,1473,536,1597,548,433,1495,1250,479,922,240,1283,819,253,433,1064,1130,757,1371,134,663,781,824,853,1470,865,1298,433,873,421,425,356,1234,4,835,991,1260,715,457,802,1202,180,940,66,1011,149,1520,1263,800,447,1467,646,457,1020,130,1302,549,675,1497,981,1398,335,1430,448,628,893,849,1255,543,456,781,1562,477,887,310,123,1159,416,548,774,787,1255,1227,235,1452,559,1589,792,641,1064,747,115,1151,762,646,1388,1435,572,1115,671,55,204,51,1569,1261,54,799,1398,402,1305,693,903,990,1116,1006,318,26,900,310,548,319,1086,642,1425,1484,1111,1531,814,1533,776,1420,913,390,305,1579,949,1497,832,1473,821,1251,33,963,652,373,536,1414,418,1539,1400,58,595,891,41,111,302,1373,362,1062,458,1256,1235,1094,535,949,594,1198,484,1123,142,921,242,727,616,198,1066,33,549,303,1384,462,55,866,418,481,1026,734,519,949,364,872,270,1481,1312,306,1272,15,350,1138,561,90,1531,845,265,1530,283,119,1119,265,1518,466,565,182,1469,588,276,1218,1049,1265,1068,113,687,306,1267,62,1156,69,1370,447,1306,1421,1539,1130,1493,854,1506,693,1038,703,1047,430,1595,1081,1212,654,1347,439,1214,685,273,445,1517,887,455,56,783,783,274,452,210,1056,283,918,215,1242,489,710,29,677,1312,1313,1406,290,623,418,957,941,942,1278,439,169,1155,546,430,178,760,1410,1598,262,615,281,950,1195,409,410,382,495,18,1532,759,251,1281,461,290,151,1227,488,1288,1203,116,1250,1196,938,650,1235,852,591,367,682,771,73,1519,1451,111,1438,22,121,475,1192,1174,1079,1021,1591,455,1254,525,1477,293,535,16,763,1238,242,1081,296,797,77,1171,191,173,592,192,1129,701,1263,1177,1468,111,1224,196,443,1282,456,707,1528,1540,96,743,922,219,63,556,486,1401,240,1412,495,839,921,1228,825,1350,821,1438,63,1183,1040,823,1593,1463,1292,314,1138,949,1483,1091,589,709,537,55,1046,100,155,635,529,526,598,1066,1455,574,1331,338,426,1126,1414,1561,83,740,1485,1551,1492,1450,111,1006,725,1416,1497,337,1422,110,892,732,530,755,1444,165,469,621,323,264,908,1570,514,466,1370,1583,197,735,1475,538,802,260,685,887,109,341,976,552,878,386,1150,1166,714,176,39,137,64,1087,1082,105,299,830,1443,1246,378,76,1580,262,1267,777,1275,358,1437,673,961,749,1508,964,1110,437,1031,1220,355,1106,610,223,1412,1473,37,28,1034,1463,643,797,733,594,145,557,1431,170,1501,983,1240,120,1305,1109,445,1587,902,1504,1439,1566,711,905,302,21,786,474,336,1395,182,602,65,1544,272,1333,1460,808,1133,1287,295,1249,1213,784,893,262,525,67,1261,291,1577,366,278,766,519,186,1000,44,463,924,548,1120,764,965,1095,1464,1228,501,181,446,1107,973,1079,569,1,230,979,897,298,1074,1522,90,293,1158,1354,64,1377,917,1208,158,746,96,1299,785,1420,952,1039,1029,629,742,187,277,1042,732,629,796,841,1436,1179,1488,1333,216,353,637,519,144,937,440,799,892,883,262,1565,1402,1221,137,1334,554,1556,204,450,795,1079,648,1536,1482,565,101,1031,1130,1024,566,1384,874,11,1065,661,602,446,1006,1462,868,525,179,1184,290,601,923,319,1575,748,974,982,548,497,153,792,174,83,518,289,1285,67,640,1218,1219,1421,1262,559,1554,583,799,1290,1433,89,1542,245,1387,771,891,994,1283,719,982,792,529,1209,1060,1219,1209,549,1591,25,2,1310,1280,513,722,160,1571,810,688,45,504,1182,1366,126,7,392,459,520,241,536,609,374,261,900,739,123,1320,901,1176,1290,277,666,289,711,1192,1513,727,1161,620,1233,870,1106,510,1498,759,648,551,10,1249,792,777,716,564,208,1202,1395,152,915,1015,723,1407,60,115,1180,38,1392,570,1050,811,1025,207,1040,1200,1165,296,846,482,1574,552,573,259,1395,797,1183,143,1097,43,284,484,156,872,1075,1388,1537,556,1455,722,836,373,1340,68,1157,66,967,314,190,391,1595,411,369,1383,833,1355,1266,240,458,179,1095,1194,753,134,1110,1054,87,302,289,1093,49,272,1273,710,704,1331,51,1301,816,1519,773,1537,595,1531,490,1338,1430,92,815,859,1292,137,1567,262,1452,640,388,1003,1207,418,1210,1560,168,982,1393,1289,1000,429,425,154,1085,792,1425,22,1031,753,739,1496,1075,1416,1452,78,580,96,1441,71,406,1585,181,301,1358,985,628,1107,1124,515,1219,1310,1156,312,365,306,1453,1450,1339,940,860,1059,1486,123,360,1184,1478,453,378,1396,16,1524,1326,1018,593,222,560,352,703,417,718,667,495,90,128,717,1442,290,602,1035,1412,1552,1279,944,670,289,89,16,88,671,239,345,492,739,409,1080,9,347,1518,1340,315,907,1006,465,501,495,583,1317,148,785,813,373,432,435,3,642,353,659,899,165,52,377,1288,762,156,990,1209,1545,852,414,924,387,153,413,952,74,1563,405,1528,224,992,889,1098,1090,492,4,353,617,641,1394,910,807,330,1333,523,469,193,1244,1509,433,943,7,1105,579,978,567,602,167,1286,110,541,807,522,1237,922,593,1514,1156,1143,43,1524,894,67,810,1363,1518,910,609,37,1479,1581,1203,1047,423,1464,112,742,1334,1309,1590,1148,314,1037,352,1207,1375,1500,585,348,1392,281,6,303,1025,1290,721,810,415,412,1108,528,1476,948,577,691,167,443,137,20,800,27,104,16,1279,872,216,1357,868,605,527,779,1152,1423,445,832,1080,1287,500,357,1141,832,1226,1421,1184,688,896,943,33,1207,1298,651,916,1399,380,593,1593,606,832,354,935,478,10,350,463,1237,852,304,1355,1331,1547,265,126,37,901,137,887,483,599,891,854,763,1343,710,25,735,265,396,622,1504,305,630,272,1155,1292,234,1341,1252,659,52,1180,518,671,728,435,585,1319,487,380,1198,1449,154,1187,1543,1554,425,1441,84,836,548,865,431,1516,724,1526,174,475,539,883,1018,1004,701,72,37,242,843,580,927,543,1483,1297,687,321,897,719,939,1398,227,399,934,1210,29,180,1024,532,1258,1463,836,392,1455,996,1349,805,533,1020,1038,1263,523,722,1384,320,558,201,1572,233,594,1413,1108,1057,1149,648,1305,583,1217,1034,1298,664,729,412,140,481,1146,365,28,21,1148,599,635,128,258,443,721,386,1438,1489,1600,1032,1442,1255,45,1161,684,734,1248,1233,812,45,795,640,309,309,1111,1049,1593,1473,1169,328,378,1571,797,1501,716,509,479,1160,1395,550,1189,934,1170,643,1109,184,1401,889,255,1274,1057,874,23,883,494,955,473,1352,26,1025,193,1486,571,944,727,590,942,1336,698,1036,1266,485,1562,91,1307,1386,1025,586,616,225,859,491,32,411,163,1539,205,57,1513,731,1053,235,1077,133,538,705,1529,1084,1328,1550,727,1158,715,831,1056,452,902,131,269,318,7,983,171,859,1112,1121,689,1020,374,1213,1360,877,871,1343,1276,1484,907,288,417,1191,333,8,887,338,269,689,553,314,1370,997,1091,537,1286,281,1194,1066,340,1163,15,571,1285,1120,1245,961,81,1383,76,1152,123,360,566,651,981,939,164,1505,1521,581,1502,728,510,926,943,911,1533,687,210,915,1261,1344,1482,1449,320,998,307,435,95,1339,1587,1300,376,600,197,132,534,803,1529,1100,1594,1452,372,1246,584,51,1285,447,574,3,1107,289,1407,879,526,914,1239,784,385,571,177,908,85,1325,32,657,1320,90,1318,1208,1189,248,230,1264,205,1319,1384,993,605,305,1587,480,1069,226,980,710,501,1083,1286,1415,922,1359,414,162,915,279,591,280,72,857,645,247,1180,1191,1464,575,1257,34,511,591,825,1472,527,463,37,280,746,128,1061,891,1251,339,1136,152,410,254,275,1105,524,1442,70,623,861,635,1291,1137,1418,20,725,934,109,1312,496,453,1303,196,376,690,358,1513,957,798,16,722,1078,251,848,1207,1396,458,1272,1401,646,1308,1047,1108,229,1093,1191,300,45,1241,474,71,1523,1161,972,1331,1239,211,1352,772,1225,798,522,880,1509,58,122,695,192,760,1501,471,733,1248,590,1388,1198,96,1001,1546,850,1266,118,1285,903,451,276,1478,848,500,441,98,1091,1232,1077,354,1332,811,270,596,1127,1104,421,996,111,1572,717,10,770,668,446,112,723,207,1317,716,877,637,873,868,448,868,200,735,657,400,528,104,1550,1257,1593,1579,572,1243,406,520,183,1114,1253,645,426,1121,450,814,1540,1511,955,193,1358,365,227,824,491,1317,180,352,78,1434,1061,1102,860,903,736,346,1418,638,1273,297,243,403,587,477,483,1593,831,297,985,468,418,443,730,814,1120,1317,1166,401,1418,562,336,415,196,626,751,172,900,1225,958,1039,1510,1495,697,1325,501,783,1433,815,771,129,281,1287,124,1449,600,516,1515,21,310,988,530,1263,227,1382,1048,984,519,1295,1207,100,248,513,1534,1395,1457,728,218,1396,411,261,548,1072,1231,390,1111,1018,106,1191,235,1022,416,924,694,1036,916,731,1387,478,1590,526,543,1532,71,1132,1252,1124,210,1226,314,923,595,1185,1523,1266,1452,604,906,1568,1588,1373,906,31,990,2,397,1536,885,1575,788,195,1548,1430,1353,970,308,839,608,1499,444,516,1208,1558,1016,377,655,1534,844,1187,207,1472,660,1446,257,809,833,759,295,1533,1447,1372,609,220,907,731,432,432,970,1495,1223,766,583,476,1031,768,546,1559,635,324,1552,1263,76,1274,1201,633,762,1159,910,943,363,474,1525,1368,1542,375,1307,210,1099,225,1206,1246,1590,1152,904,979,566,812,138,1264,164,262,535,1276,120,977,91,1166,1171,3,1499,805,1572,176,1196,1172,595,434,156,1389,247,1470,660,651,1486,1178,142,950,1529,1497,1595,269,382,184,790,169,763,440,686,1160,974,1085,850,568,631,414,643,1350,214,1515,796,808,311,1588,86,1193,1081,1583,319,1596,1122,1147,1158,1563,591,33,128,149,1505,800,362,192,1243,1587,118,132,562,704,967,758,1026,1178,646,1136,1332,1582,1052,1383,247,588,137,411,1331,1533,857,1048,1580,1123,1084,709,768,510,354,65,1541,207,1211,139,179,1053,550,426,1363,89,254,1326,61,319,1392,723,849,1263,1457,242,59,1598,163,1506,507,1088,254,799,346,251,942,1551,266,1454,1387,50,929,735,371,120,306,97,23,863,1478,1228,585,966,1054,1389,1410,1385,1336,1304,650,669,526,1057,323,1259,768,1260,984,1235,313,1400,682,1199,725,1464,462,1301,1505,6,684,1136,155,102,1216,1023,59,530,473,83,601,1017,932,321,1301,1462,944,469,1275,1181,1031,631,1355,901,387,57,58,183,613,192,1142,1598,1027,849,782,536,443,781,66,786,909,1384,1251,1061,1439,277,1100,1207,628,1318,507,1109,1526,1435,703,1332,171,167,657,911,1237,1554,886,462,597,400,565,764,811,1086,638,603,414,660,653,1416,1075,1318,965,32,172,940,333,643,1268,1210,707,991,1572,282,1470,108,645,299,1537,579,1498,400,1552,1590,981,1294,263,54,1445,291,1456,275,937,442,497,599,1094,1108,167,41,543,1428,1355,185,1169,864,1041,949,1590,228,848,666,550,1506,504,1199,769,693,865,789,1058,410,970,1133,1313,864,512,101,891,1197,606,1531,1149,329,574,1393,89,622,365,905,360,916,1275,1321,1144,521,1324,1242,652,1558,1477,474,116,1598,638,1141,916,18,541,323,1380,346,605,886,763,1101,1190,724,495,1132,555,408,700,594,898,716,1283,68,1468,552,1337,612,837,554,193,6,1062,594,872,642,395,615,533,1268,1343,1373,475,806,1308,762,627,116,786,1585,1575,345,640,169,783,1073,1016,1491,981,913,1093,1028,1205,987,1342,1468,1331,923,470,963,1058,1031,453,1171,1553,1139,163,964,1091,1586,337,353,1524,476,584,1191,466,1372,1082,1163,1354,1267,1537,980,312,1399,730,1137,1197,1421,1004,1384,727,893,532,1561,759,577,948,91,73,187,1119,14,1229,1482,520,242,917,1412,1493,986,205,660,67,34,477,1380,1008,723,1222,1371,699,1468,266,13,136,1397,121,1465,245,722,909,259,1373,303,807,767,263,1362,1040,868,57,429,966,128,269,775,916,143,593,842,1397,1494,672,1008,615,330,1033,1550,558,449,504,1099,261,46,423,769,1216,1456,152,630,549,274,1203,600,405,519,181,674,1112,1221,157,339,1582,603,638,742,1273,1053,1086,1389,578,651,1298,1330,572,271,1070,7,460,1242,243,301,1053,1308,869,1503,237,1510,7,937,999,842,832,1385,37,162,1498,339,757,114,514,912,484,897,498,1451,1449,915,840,1511,1194,1518,385,814,1249,632,670,139,1245,115,379,254,1363,310,958,202,800,927,30,1594,29,785,1146,1319,559,451,735,173,1207,1445,322,1159,1594,92,1045,1096,168,113,357,581,1280,1130,377,537,919,256,158,983,1089,1281,474,639,1381,517,1593,1570,918,364,1302,857,575,1580,48,899,23,253,302,659,1374,1116,1324,938,52,1515,1009,801,558,1254,385,548,736,220,790,1208,1097,752,763,130,1463,320,917,747,124,921,16,288,1111,525,317,1467,1058,1307,1014,1479,87,429,517,24,618,896,331,95,1568,557,1415,1540,1170,1181,324,920,282,305,346,427,779,1378,542,275,1281,1595,91,221,687,1216,194,589,67,1264,518,509,713,344,511,1344,163,681,894,345,353,1296,347,209,536,1284,68,697,1322,1341,399,1575,742,164,494,567,1191,634,609,1536,348,719,522,671,1452,1199,826,105,1348,885,850,214,36,435,39,904,1593,669,316,1384,1469,980,1239,341,1221,1137,946,472,1554,1096,282,1342,1523,275,909,758,1559,833,36,160,449,1133,855,1112,1582,1089,182,1324,1481,271,253,88,973,91,1334,1460,377,1473,1529,496,771,1282,352,260,652,1023,1511,265,945,308,1113,418,31,1529,1211,1430,598,1370,718,1094,1228,1242,535,473,687,1504,733,1456,541,1260,50,418,919,776,999,552,312,694,450,616,715,189,789,909,60,409,764,1336,1062,124,151,491,694,1382,570,1453,1245,515,1571,783,1015,609,1574,665,99,66,278,694,1598,1127,517,541,252,478,736,458,1396,903,1213,148,1207,311,1009,826,542,1346,762,1085,1164,963,1423,336,1328,583,1541,850,1488,1183,1153,987,798,90,414,100,173,941,251,488,1080,980,1373,1165,985,747,1339,1370,235,964,1404,608,1288,841,1117,1534,1565,1454,860,638,113,501,722,1153,953,1032,1450,993,1090,633,556,885,697,1103,441,182,810,1559,1397,784,970,1440,343,613,652,1459,954,1247,41,1119,584,38,188,1159,359,1235,837,1429,103,1518,1007,67,775,412,457,1142,1406,1429,35,1338,1417,522,442,1370,979,253,1140,659,1472,130,1550,82,883,1345,696,756,1263,1388,125,1008,1162,627,1325,1527,1012,1037,767,1311,1297,311,114,635,961,716,1037,771,838,346,1055,702,1042,597,1586,327,1023,987,709,314,278,984,1388,1279,363,854,537,405,896,1088,610,680,540,1131,878,1460,1563,861,1401,461,810,49,1231,875,892,671,454,1390,1565,773,976,420,647,962,198,330,1185,254,753,1413,1224,149,569,26,997,1268,1180,686,1282,641,375,1027,1356,1327,996,824,707,740,95,1057,1566,974,610,1513,867,1503,128,1067,34,977,1256,724,390,312,1007,1409,650,288,947,1328,608,1546,502,429,824,1418,1489,1478,180,871,722,228,841,1205,686,320,345,1333,529,440,737,379,560,1338,615,767,593,690,352,1438,1387,1240,767,75,1041,480,124,1200,219,504,1263,1362,1094,1451,567,1598,1321,613,331,1139,837,933,860,563,405,718,356,615,1555,52,593,1523,462,186,1518,459,1080,121,398,1062,673,1442,694,1429,1595,1553,661,1559,1488,1015,1276,752,186,1132,307,947,1096,1262,1031,100,125,1313,650,449,536,312,50,211,111,521,466,1001,727,1093,517,1341,1025,1555,102,791,836,464,1062,898,1058,1409,629,1227,842,1098,1521,28,639,772,1443,955,697,1505,1154,574,223,1465,406,661,145,1068,649,390,33,807,519,1043,193,1263,1115,1339,235,643,1107,1539,1296,554,652,900,1553,1040,1033,234,1033,1544,1010,1213,1477,114,1280,445,1050,1349,1017,7,1336,1091,987,410,680,1239,814,734,201,1554,694,1188,727,716,1372,1464,125,1150,406,189,1248,158,387,1257,413,166,1474,1310,1419,1309,1306,1542,1090,227,1360,602,411,160,593,835,164,1551,1224,799,70,303,1557,976,620,1342,516,1510,1442,154,295,640,1532,1010,1367,1059,53,1461,1109,825,315,809,119,176,485,1497,366,422,654,513,5,1501,1447,372,896,1230,272,337,294,564,858,964,1309,583,737,1196,909,264,1375,635,1546,1023,839,1544,1300,1071,1050,951,1596,616,1573,199,1435,1000,199,1236,999,789,82,149,1502,1363,744,1575,151,103,80,409,720,820,203,1023,1090,1580,1411,1042,805,366,171,59,1419,389,874,709,432,1238,668,1274,1233,546,629,529,663,1436,1321,1311,1171,1021,903,154,132,1121,1206,1536,368,468,599,461,906,1511,1342,980,1457,1466,1060,935,879,219,1311,306,1249,389,1461,268,1439,326,1379,1187,1567,1587,1526,1391,312,1230,733,1285,1585,1490,559,460,1375,609,255,264,562,93,482,976,350,581,928,1031,1308,1519,14,200,814,1287,1405,14,633,221,1263,1181,473,105,690,1335,1241,450,436,30,906,275,735,1439,385,1318,106,348,1258,802,7,445,937,528,890,1463,329,768,986,25,196,1085,566,283,1182,1593,350,350,964,88,464,1591,1413,156,1079,59,655,1340,136,485,1335,326,1569,153,373,632,526,1279,517,687,112,737,18,1190,517,847,1514,1200,1448,447,248,880,148,429,860,510,975,1508,380,1309,470,913,1365,1306,1253,722,214,491,326,1059,760,628,165,185,183,842,440,1566,1210,1346,741,1547,1339,1432,471,1391,1179,216,1328,1278,1272,1119,1273,285,1056,1338,649,748,295,255,1342,509,48,1132,379,825,1112,190,1110,1314,1532,419,917,509,1120,1111,517,344,592,1016,313,1051,263,1404,1023,1279,646,667,1535,54,17,374,465,826,1238,1000,1376,240,907,1345,138,1197,1088,762,37,309,841,842,609,916,410,552,1600,1159,1109,1295,1495,1309,1521,941,128,1468,1128,1316,191,1,301,1285,1560,742,722,327,259,1415,938,365,809,194,854,743,1211,520,763,404,862,612,332,923,931,714,576,872,1329,401,1280,747,1223,756,1585,1116,281,1094,248,1345,663,121,1132,961,659,19,337,340,229,272,509,998,1584,143,468,1367,362,365,1042,277,1339,824,1496,709,681,468,1218,629,1592,1431,1093,858,5,1266,154,1433,1175,1030,389,413,1309,135,1338,914,1178,1222,448,84,939,1119,1377,88,617,1469,1133,1133,17,341,343,1020,785,1095,1546,1190,625,429,678,687,713,961,638,706,1200,60,1391,1397,1565,1090,1432,1131,1024,520,1501,472,1164,173,884,1141,208,1092,1486,287,970,414,516,463,193,979,680,1514,1271,1464,1541,945,978,768,1316,1223,7,1095,1,556,1219,157,878,141,113,773,188,52,320,1012,1092,1332,910,26,883,1548,399,987,767,739,1030,1292,962,808,558,1380,1185,1119,594,1345,1243,353,1076,1157,1493,747,906,612,1159,150,1542,1133,496,1082,1014,802,1554,402,64,1057,1139,1368,718,835,642,893,1021,460,181,161,1222,975,614,1261,351,499,575,1399,1206,707,60,1342,1588,913,414,964,1530,569,1356,86,193,991,1079,1008,1283,121,1274,339,563,180,1515,1275,1333,1481,253,1109,1315,567,641,185,6,1486,1301,1018,479,113,402,855,58,206,56,22,246,1484,350,102,493,301,1368,782,1445,803,1529,120,804,266,1235,838,61,1186,1336,1562,1310,266,1450,1044,168,1532,707,635,151,523,362,751,34,1438,102,533,1389,910,1596,943,694,583,117,1150,342,1432,858,428,283,598,708,1252,1039,29,209,571,1514,1091,1525,1414,1575,1403,1517,241,262,1217,653,288,832,1600,828,683,898,367,391,1555,286,875,888,186,1465,1155,1256,1058,319,336,661,1057,475,1435,214,1101,409,337,1510,410,1292,445,628,912,812,197,733,625,1151,669,1343,1060,1354,343,905,1402,1409,442,666,732,157,461,1319,943,584,1332,1274,676,423,397,1351,315,1042,753,1443,151,391,62,1219,573,898,584,1060,1016,1068,1500,907,1351,1291,490,188,767,1454,456,1537,914,1114,84,1590,1359,1431,1431,310,1489,1027,787,470,25,691,1404,868,413,900,1430,1143,24,1206,66,345,721,359,192,1567,437,1303,158,1027,221,945,979,286,49,1452,221,1544,1427,1188,196,3,40,919,1362,863,412,701,27,693,1119,915,678,221,1244,257,1032,986,1089,1461,732,242,994,260,746,121,1082,1028,1088,700,56,224,945,439,468,392,1086,867,905,912,1407,1301,988,1293,128,922,306,514,1385,572,477,289,341,675,1030,880,499,197,186,1294,711,252,1083,511,1573,832,1250,36,1019,612,298,259,687,868,941,1320,708,1007,829,1524,917,761,1104,618,1019,775,719,1086,244,459,563,11,1124,537,1522,1466,638,656,1346,322,1091,115,384,311,130,191,133,295,330,807,347,361,622,265,575,667,985,1532,1459,506,368,67,358,1166,63,1091,1514,1298,721,1106,985,557,1136,602,1532,1035,177,1572,433,840,987,1518,821,925,506,1593,1304,1247,1431,893,254,226,301,163,770,1208,1075,151,947,432,1064,605,1264,530,490,344,757,1144,417,880,1278,1418,1224,670,1454,427,889,42,95,803,265,1510,82,438,236,597,412,948,1328,977,1258,807,1242,1535,988,1172,874,360,443,261,703,1323,1442,998,1325,619,325,745,613,946,915,1395,1457,1577,757,1266,1474,222,224,313,70,432,807,674,218,569,177,920,1430,1554,452,537,1085,469,1281,1174,1199,661,1177,84,535,740,83,1486,416,1428,1515,511,671,636,1209,825,834,1413,1460,556,1293,1220,1044,1136,1193,1585,796,671,1010,1520,489,806,342,856,672,451,170,1343,1181,333,1395,1049,162,598,894,558,556,342,384,223,1263,1364,526,176,1548,640,1410,889,1335,1558,702,1269,1482,357,1362,136,1351,1099,550,160,611,1090,1514,192,206,11,536,1060,5,1363,252,1168,940,1241,725,600,849,816,1192,782,703,136,321,1181,1139,1572,1499,76,66,572,691,141,1421,587,1082,1378,1422,4,861,441,1096,571,948,149,132,1597,932,60,1473,1588,1462,270,476,1147,230,698,817,454,1135,259,1056,572,247,703,1360,1170,1509,1016,1375,1320,919,1362,1404,1299,616,261,691,1485,683,1406,654,792,780,949,458,440,1480,949,1057,393,670,1444,1175,1181,570,1511,364,159,273,407,711,188,623,1103,1509,1086,757,876,708,479,1334,967,115,1076,321,1339,840,351,857,1326,1019,472,110,703,104,981,706,969,114,1273,1211,485,973,1566,293,1383,383,506,1148,918,611,288,400,617,296,605,1028,211,962,307,1286,1258,597,807,1106,1425,638,870,540,1353,1437,603,1236,875,679,233,957,845,299,689,1211,603,204,674,1025,1198,1090,1022,616,1588,885,151,604,693,1141,1448,200,824,1191,60,416,47,294,1240,1004,824,447,1454,122,8,509,294,1096,1315,798,175,1151,831,1406,1298,340,222,535,794,617,954,687,759,372,1256,128,1502,575,1588,1333,537,1256,435,845,280,162,1501,507,888,777,705,303,174,1288,1280,755,1323,1264,648,129,1356,767,65,213,295,1391,259,1150,177,983,1435,713,614,971,312,117,1282,460,642,454,1544,894,1185,1250,766,442,622,765,629,311,1518,326,773,697,1589,549,677,1139,813,866,787,987,159,310,565,1219,1290,167,636,314,1357,1194,161,30,355,455,287,866,65,902,607,292,801,332,509,279,1103,148,552,645,341,1300,1409,777,1575,1042,1147,73,1592,540,532,527,308,110,332,850,153,369,556,454,413,261,720,1399,1178,1216,533,15,1058,214,396,1528,1276,1151,894,559,339,1414,52,1294,1437,1236,433,1536,727,229,25,1496,1144,507,651,652,845,1268,1320,1425,88,141,802,256,430,1332,607,158,247,1533,442,1507,909,370,1245,939,1322,591,55,1155,1439,475,556,1349,903,192,221,1204,1189,1308,898,626,138,1415,1550,472,778,1531,1423,1249,341,25,253,810,1102,1179,507,587,302,245,1382,1207,928,178,83,599,1226,1339,251,1150,774,395,1518,121,95,248,1172,390,920,432,932,1187,1335,318,796,397,1077,19,647,1447,570,12,446,310,1185,1509,1013,1449,278,1151,1430,1220,429,306,577,315,176,1013,1329,973,1324,1058,444,479,517,1569,632,1287,914,886,411,828,488,1555,504,1492,746,70,363,1437,76,10,1020,1130,1327,784,750,580,1300,1044,202,310,1126,1169,414,616,1192,12,1014,133,1387,1448,1269,711,979,1197,764,582,860,622,5,894,896,1267,728,1488,10,860,411,293,1014,1456,443,136,1280,62,1189,274,706,457,1191,1087,1334,424,458,183,1340,171,677,53,600,902,541,1421,842,1405,339,20,1456,481,1443,18,1573,976,931,781,539,640,650,395,8,208,432,1061,735,1542,1215,1433,378,893,1320,1286,1412,1567,900,921,1125,1393,975,813,1107,445,503,945,796,1400,1466,54,1469,542,467,56,215,958,645,1295,912,1590,162,643,1504,18,1163,1279,751,1004,462,859,1280,1294,71,1295,414,948,469,709,586,344,98,72,452,4,815,106,1352,890,1522,1163,1553,724,998,71,1260,1284,764,224,1517,67,188,24,1351,1302,425,251,1497,100,284,507,1135,503,12,703,1219,1561,1268,1223,1351,1513,1543,1181,86,708,1048,1471,596,473,1387,1038,507,1576,1270,595,533,1181,720,1084,264,1193,450,702,848,1271,1308,1256,442,6,1486,828,895,275,1404,891,39,800,229,1594,243,1588,1213,813,721,1465,1465,865,1129,542,25,477,1479,1586,710,1391,89,676,450,895,111,924,1573,1220,1033,291,800,1474,482,310,1162,128,1553,1582,1169,564,1252,781,1526,1277,1061,874,1107,141,1316,576,1255,1203,761,1016,1162,928,276,1257,1369,830,212,1501,600,577,564,1185,361,487,950,101,417,1131,1008,1412,939,982,1021,580,967,1259,1063,448,886,273,408,360,1480,1019,361,577,273,27,730,1064,133,209,1539,348,490,650,549,57,683,806,739,482,1010,1090,1577,265,653,1468,705,933,1010,1368,1230,164,537,234,527,474,559,917,355,665,798,59,395,944,1435,1421,1354,1594,832,1117,678,669,972,1480,371,928,403,1257,458,1368,1100,1053,589,461,872,768,363,1329,608,688,351,1589,1263,749,1009,36,687,1161,772,261,749,747,1295,323,289,432,14,155,742,1568,917,664,786,1213,769,308,852,753,1072,246,718,827,859,750,1282,141,1403,981,561,322,315,1325,791,209,1015,712,333,284,939,1182,1261,186,444,638,1527,971,969,5,385,351,1520,240,668,500,182,131,1428,559,1439,881,521,57,651,10,650,344,38,1376,786,138,646,1445,1446,514,22,85,1380,188,875,294,571,35,117,21,787,1095,707,55,1166,218,366,1219,1410,757,413,990,7,52,114,1293,441,716,445,739,887,1174,1121,1205,1167,1269,44,425,704,1546,964,1170,1198,240,1317,304,1445,1517,967,1018,826,559,380,676,912,963,1097,916,180,1573,270,760,1315,620,926,656,578,550,320,1073,218,1101,1357,382,1232,881,1466,924,602,1174,285,730,684,1339,1301,262,957,764,1470,900,983,729,1263,1101,222,1245,800,826,863,689,712,1053,823,62,1208,47,1231,190,1186,1419,466,158,1562,447,283,1265,534,1189,383,452,442,842,459,1180,1332,992,64,1371,850,1160,1085,411,998,181,326,368,341,733,573,1450,801,716,499,171,388,821,984,370,724,24,269,1195,1499,282,83,401,533,706,596,43,1446,1324,1313,651,1282,1132,210,367,1337,387,732,745,226,388,1016,457,352,301,880,1105,1329,1162,131,227,133,494,1485,1456,942,345,325,395,530,816,569,1225,876,1059,275,1589,1160,1369,705,803,1056,64,528,1184,1401,108,268,1314,613,945,979,1141,841,596,1444,475,614,1057,872,692,1236,1554,951,1088,1470,483,163,1321,1265,20,989,1504,287,501,1135,77,1598,1149,321,1308,965,508,493,436,1142,318,595,1550,254,625,681,1076,613,952,602,1497,1121,1080,1176,1590,846,864,1200,887,692,244,1020,1454,1256,237,1099,968,936,536,537,1194,1457,1372,1328,699,1067,1371,1070,320,1238,778,1014,514,374,770,961,910,111,787,792,1180,692,1417,606,1450,694,1037,384,1574,1350,1017,933,318,1001,1349,1318,554,709,884,1058,165,1574,380,1530,697,1244,1440,98,568,1010,868,155,928,793,55,927,699,915,429,282,1056,277,859,829,619,221,1566,885,1331,1417,1387,742,862,759,1476,272,1112,1086,509,500,1343,346,141,965,52,1403,428,586,1322,1412,483,1457,453,1296,1346,674,1558,1334,1201,34,1226,1565,43,134,83,1056,992,1300,131,1308,357,473,327,1137,115,710,1356,491,1136,592,671,226,1134,906,1586,1439,399,803,1425,288,1188,1173,1216,1039,841,1175,887,1558,590,365,216,253,376,133,1379,1053,1099,809,1409,1249,2,418,247,1202,80,729,868,1212,366,1044,1209,163,472,104,984,763,1109,622,211,1570,1543,31,711,561,216,1566,999,828,599,251,1297,1529,446,1163,192,902,887,1071,1372,1510,1263,787,202,1246,597,586,1327,906,669,1383,978,903,1144,1506,1445,506,1175,974,1057,1513,625,1119,269,143,827,854,1487,933,337,1105,834,225,276,269,1581,882,251,1172,670,837,1179,987,1476,387,16,392,460,151,671,1089,803,479,978,226,1512,273,205,466,670,914,435,544,1376,1204,642,1005,241,1319,75,28,1560,1119,448,1169,166,123,1361,1580,79,618,1381,1247,522,412,1513,877,866,929,258,384,868,277,472,1064,1424,362,991,413,1394,1003,1059,34,1504,30,423,712,1015,1346,1071,1213,1229,1242,1104,606,300,38,1137,1256,559,1474,518,1276,146,1366,344,1523,429,1443,1550,848,137,1382,94,610,118,109,145,217,255,145,693,1132,787,388,774,850,382,509,788,774,606,116,1522,432,496,1501,659,1175,224,1076,1302,1538,738,960,1267,1269,568,966,1481,612,1042,540,1583,1357,1501,1425,1120,916,504,575,797,396,740,708,1175,1286,56,65,36,1271,553,1444,1528,565,1296,1489,631,232,1541,688,635,1379,57,1364,525,131,956,995,1109,894,368,588,699,1102,533,1110,1499,1552,131,1191,937,953,280,978,1358,1588,1514,247,1420,428,576,1411,767,59,1142,1549,994,1057,857,72,181,1240,1070,543,404,128,663,972,301,302,970,710,1367,1171,1287,1093,1318,1246,349,110,1370,485,424,1556,671,413,212,1007,1068,1006,1278,1334,1422,362,282,1075,979,454,262,1174,306,599,1004,980,1429,1150,1410,1366,1551,1007,647,241,410,431,701,108,1317,1138,703,1278,661,507,1240,1416,13,852,1527,992,895,1281,912,1054,493,66,195,635,469,524,1508,967,658,1545,709,1584,81,1111,1364,672,294,1529,964,652,240,317,913,1021,1428,410,1129,1334,794,62,1558,577,1506,168,99,1461,459,1216,767,1219,815,1329,95,1186,1016,178,773,1019,602,738,275,209,1238,1358,262,967,1500,1069,260,360,645,856,190,1202,450,1469,1421,1487,1006,716,824,652,1423,1430,1432,698,491,725,584,1507,538,262,726,1297,625,1396,390,1027,1373,1464,1305,1142,821,1487,237,716,697,1435,821,1275,1195,341,648,1142,1361,1088,1233,320,969,1251,935,248,48,981,785,899,1328,675,455,158,772,577,308,611,1000,1249,486,1166,336,1491,1382,242,1512,701,355,914,1572,506,920,308,1103,458,998,1454,1142,1506,1030,1325,453,264,1517,1487,878,1125,1163,1148,67,1315,1310,826,1308,1,1213,1031,165,577,747,1490,1568,804,1444,1538,1525,906,1081,660,701,846,205,479,1205,167,358,576,1226,392,334,694,21,612,977,1222,290,767,1422,112,131,1485,991,1183,715,1112,858,715,947,335,46,1329,520,984,1013,607,642,57,630,45,791,532,1118,175,748,1499,314,498,1409,495,956,1469,865,683,1353,41,230,1105,970,433,1161,1187,988,674,289,411,990,437,1320,1323,620,1438,1315,451,239,619,276,1489,606,1092,243,460,1245,306,361,1212,141,1066,678,1510,1519,1236,1223,127,340,1431,1197,666,1277,1343,530,1334,400,366,681,1202,1220,559,559,613,864,1141,921,543,1296,832,651,629,1188,227,495,672,676,495,1464,1311,1236,1539,1273,270,1317,42,833,1125,226,487,164,294,1327,1196,690,1406,1269,155,583,1502,1100,308,45,205,1002,119,582,595,56,1255,1183,1235,1414,77,235,130,909,720,436,1554,294,722,151,296,508,1302,1363,301,539,1209,1213,1176,971,1291,361,381,895,1523,1564,154,1200,438,1052,1390,494,1477,744,461,1190,230,1459,129,1572,104,906,1408,1227,578,1035,1241,1499,548,951,814,1445,57,84,1057,396,655,673,1412,1277,161,1083,8,684,903,1388,457,926,660,243,1556,71,1010,552,1050,1139,1381,1552,1393,359,1323,333,828,67,1318,1393,1332,1524,1138,1535,1220,14,239,1329,1028,693,204,820,239,571,1013,387,252,1460,1422,87,1086,1226,387,1222,1331,1170,1168,1543,219,687,245,1511,1478,1302,1105,918,194,597,1498,239,67,1548,306,1453,889,396,1064,1004,735,574,622,247,1404,853,366,1571,139,1043,340,760,1490,1131,78,614,1288,843,1287,824,73,1554,897,1121,1566,1104,1118,849,1434,1149,165,1530,637,571,1155,542,1251,1089,761,1485,383,720,1364,793,1365,600,77,129,331,78,810,304,344,441,161,943,1124,1417,376,1365,496,522,1229,881,1571,611,1551,1483,1357,664,1483,1563,1271,1038,564,912,186,1179,1585,39,688,1548,397,612,1398,120,1190,640,981,818,672,526,1418,387,1002,192,1324,1375,867,1526,277,512,566,325,933,557,183,598,347,5,372,384,184,1469,1296,643,1233,924,1104,801,1061,194,395,1387,1437,738,706,512,1473,1351,363,666,295,1415,422,1226,18,1388,1383,1247,1260,23,477,417,719,1565,1193,1195,146,1318,1393,1596,93,1411,106,393,709,581,881,323,1034,807,1211,208,217,165,442,426,603,1203,1502,1148,642,752,9,1581,621,749,707,50,1441,788,1565,8,1576,1339,1315,923,1457,752,218,57,818,638,33,909,1363,988,498,1034,761,857,1098,650,751,102,266,1115,756,577,996,48,1557,858,820,759,357,1376,525,1078,542,333,1200,845,532,843,180,1142,654,1407,136,1212,1245,1265,715,999,1511,1482,1081,986,518,1051,1253,1317,1093,905,690,1138,1012,1432,585,905,1368,519,997,412,211,647,1465,272,1183,1522,1555,635,1448,41,1004,960,180,751,858,907,1180,453,1392,1056,1275,1349,1435,1196,929,266,171,823,995,1400,731,1338,387,273,44,915,1109,55,838,269,976,1488,737,279,864,1170,637,1364,160,1212,224,1397,430,252,140,161,331,1065,969,1572,570,378,101,690,573,843,42,1205,34,543,924,132,872,311,837,861,1268,145,1320,934,1180,1270,67,1332,1498,300,1484,138,1589,1458,1479,880,701,58,1259,496,67,1089,304,1226,145,1418,340,1460,943,1354,956,807,906,1372,455,13,1195,439,1237,673,630,898,701,553,1152,481,476,419,25,202,686,233,1392,1419,1589,797,1185,385,1248,241,1237,236,634,421,890,1292,609,873,1324,452,950,940,1029,1155,7,154,598,1038,1582,448,205,561,337,1398,475,283,1513,1086,1456,661,920,172,347,1291,1405,238,744,495,1272,1452,1567,902,1435,352,1296,840,1492,721,1326,1542,1598,446,1329,1028,864,953,813,1525,1435,507,145,508,157,1521,1301,1299,988,61,866,184,216,1229,731,972,134,793,171,1175,1027,781,110,126,535,216,1219,309,1146,411,94,1097,655,1355,362,522,1071,302,90,692,1376,1206,1578,610,87,985,88,629,266,642,864,315,1257,452,1401,713,1571,1292,216,192,777,1348,985,972,1327,596,142,388,851,1223,1257,263,77,1219,311,852,1160,981,147,1575,1219,1330,679,1533,395,969,409,989,282,1089,434,633,487,1217,1208,668,1471,1493,997,1489,141,613,664,785,219,349,662,1114,575,1025,1201,756,992,412,11,20,1216,95,870,1185,133,693,266,1238,484,1542,1585,705,1426,447,400,1267,1007,857,247,977,927,975,1272,760,826,116,225,125,902,356,446,915,827,1229,375,840,1215,1114,807,1276,1235,327,794,737,1055,270,897,521,876,1552,1148,1211,1534,1062,261,147,217,628,633,1542,804,626,1192,327,498,520,1031,128,916,1142,179,1411,1314,898,659,294,388,890,757,9,1277,1299,1234,529,699,927,1539,1586,15,215,1191,1,1164,1341,87,27,1027,572,695,109,1146,182,1207,1153,332,1440,513,646,588,754,871,1348,951,1268,879,299,669,626,1541,530,285,68,556,432,1041,822,1468,1489,1473,171,436,836,1552,773,1390,1515,1171,663,1386,1505,34,1307,258,1508,1356,911,321,552,955,33,157,177,92,338,5,1497,1490,809,791,652,1264,494,691,4,1470,1577,1323,188,1027,263,427,1032,952,206,1512,423,1238,819,1219,232,618,628,1125,465,1325,1100,1300,1515,1480,1082,1513,58,910,268,255,259,1215,183,323,290,1397,1134,277,522,445,984,1071,899,585,261,309,295,613,1564,877,478,1492,59,22,82,1411,238,1155,1451,637,1217,1596,728,31,1189,1471,3,987,77,1523,993,156,212,335,673,1477,967,1037,985,1316,1071,982,965,952,695,687,1001,832,407,229,1194,1321,538,498,24,1421,208,1424,636,1549,425,219,947,156,395,992,1,542,1546,581,1550,896,1359,221,1228,584,446,1253,1382,1366,215,174,948,462,609,659,1292,701,762,799,900,304,855,873,1527,1393,1391,151,258,1066,704,896,343,1203,1403,1292,388,269,206,1449,260,950,395,447,84,954,167,197,1073,605,1521,951,276,193,1195,910,938,590,470,176,59,1437,661,547,708,280,68,612,946,1055,872,1413,1429,432,1054,849,1107,887,1039,249,60,430,1118,49,1207,484,408,1108,362,204,1058,1166,578,482,394,532,102,301,1189,1540,1453,220,1403,364,530,1319,42,576,86,96,319,1505,1127,1105,1132,1549,556,843,251,536,250,390,358,1205,562,1484,1522,220,351,1098,1189,794,1002,1309,797,156,1092,1204,337,1039,769,999,197,626,576,298,419,892,163,900,514,26,1126,638,1294,1538,72,1445,375,839,881,1000,1139,645,1462,1549,1484,662,293,672,615,1400,1520,894,1170,559,904,681,436,520,389,1132,846,1042,81,147,1321,1354,910,323,606,665,1310,1202,1077,664,305,1093,672,10,570,1316,1572,768,1185,849,1417,954,1358,1042,138,19,669,786,1124,591,1457,832,1514,1234,962,1027,552,800,191,1427,90,1395,787,965,1396,672,171,284,829,811,226,473,960,1464,896,308,1531,890,85,810,376,862,964,527,363,799,1117,96,1037,642,517,1033,555,77,1585,602,350,1206,881,498,1015,304,952,716,588,699,1210,1517,1391,1082,670,944,400,536,408,302,1159,1203,292,1055,97,699,808,362,881,1156,244,584,422,1450,584,450,255,1232,489,805,1596,1404,610,282,1385,1030,381,804,167,33,850,1170,509,519,1424,1155,747,1584,1503,876,907,1342,844,780,357,303,1321,1485,1021,736,640,96,1035,500,1149,312,1482,273,260,1037,977,44,603,1108,1350,375,190,616,926,908,108,735,1318,1054,178,836,268,946,1011,726,1299,773,1571,704,1030,583,374,905,730,911,360,610,418,362,216,1222,399,1481,1595,568,1124,479,1112,178,183,1220,1108,1062,1578,1369,246,261,1189,1308,785,705,369,729,1533,763,1584,47,422,284,555,75,330,1183,580,984,137,1433,1349,154,432,101,1200,1089,167,252,1550,104,1310,1201,251,1538,497,129,461,236,881,610,344,1380,1578,69,1328,971,179,591,1553,371,1586,1052,663,1510,1255,232,213,875,558,747,1182,464,1197,1462,1485,861,283,690,1597,115,1459,1208,1451,563,636,1014,1569,567,723,500,304,539,212,107,1238,1588,545,1582,453,702,39,1188,1017,1339,681,122,1018,337,195,1379,901,1418,1099,54,1385,861,482,1199,1326,1142,87,155,919,524,1418,287,637,168,712,1299,684,89,486,671,319,779,1489,1205,1423,595,550,391,1236,1377,568,568,753,1236,157,158,1480,996,109,635,260,1029,1564,1030,1576,662,592,288,1129,393,1474,399,1442,788,1407,1462,1534,667,762,315,1100,150,177,432,509,114,1045,244,324,1114,423,775,451,843,215,381,584,949,243,480,858,1478,47,819,406,930,920,905,594,120,31,266,930,823,793,620,497,740,634,830,1474,1004,927,1008,236,795,589,497,1244,427,160,1075,1250,1206,1009,228,180,1511,1382,1297,1104,1422,1355,1574,449,1549,1320,699,597,1223,822,1300,285,1049,1387,1285,992,1326,1576,1220,1465,1272,47,926,1482,672,693,1402,796,1150,1219,831,1569,1054,597,415,375,893,337,468,113,1143,1321,1514,588,311,628,100,1577,1311,599,659,855,1411,1436,754,32,106,1493,34,317,1086,1596,1388,109,1589,876,550,802,1277,1191,1220,404,1362,474,149,664,1263,1517,1098,1589,1004,354,1154,674,414,1109,861,1040,1095,98,1369,406,748,601,406,146,538,1002,1364,48,1470,1599,871,913,98,1211,647,943,1094,1145,1532,1525,772,14,643,961,1115,1292,682,123,191,1366,471,1351,269,1504,1230,859,46,1281,1176,748,1080,1046,1204,1131,1202,805,448,299,1303,1318,1364,359,1293,842,696,870,65,1170,1471,703,716,802,24,1396,999,476,512,1591,1239,1468,658,721,1094,339,431,986,129,37,876,521,457,1444,937,243,1377,1349,135,1465,99,1452,677,988,195,951,1251,1087,589,687,1134,274,1399,1106,1119,801,465,262,32,1534,1571,85,815,553,1007,1543,42,539,1122,1498,1524,593,1454,1057,1053,824,1272,650,1380,210,765,428,936,424,594,482,45,506,85,1302,1073,1212,1197,472,43,802,389,25,779,1402,679,1547,1156,59,202,301,1212,1584,607,1294,515,1546,1117,726,879,430,488,62,1244,564,1077,1366,1240,831,1598,255,1069,493,1254,1362,866,1129,1553,1004,1013,1160,316,9,798,1055,1505,1407,537,1261,1365,702,209,1495,654,1335,950,143,407,1326,38,912,938,546,291,802,997,1527,1288,468,752,1474,523,551,235,125,481,1134,421,464,1192,579,16,705,669,478,1334,1301,1503,30,357,952,211,1595,831,284,188,522,1512,978,1379,1473,20,550,83,273,1123,1353,479,1110,668,83,1368,46,246,1311,919,872,848,429,1278,311,801,1340,1335,276,1216,205,671,1518,907,1596,215,1508,349,1089,951,976,335,1049,610,614,1345,378,70,643,1060,290,1468,773,292,502,425,1022,388,507,5,423,23,1499,572,1119,1341,868,704,1324,888,799,9,956,515,134,289,1437,238,72,1091,1257,1076,1022,1076,93,1017,106,1018,592,713,192,1528]`),
		// 	Arg2:     1000,
		// 	Expected: 48932, // time limit test
		// },
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q992SubarraysWithKDifferentIntegers(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ340LongestSubstringWithAtMostKDistinctCharacters Medium
func TestQ340LongestSubstringWithAtMostKDistinctCharacters(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     "eceba",
			Arg2:     2,
			Expected: 3,
		},
		{
			Arg1:     "aa",
			Arg2:     1,
			Expected: 2,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q340LongestSubstringWithAtMostKDistinctCharacters(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ3LongestSubstringWithoutRepeatingCharacters Medium
func TestQ3LongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	type param struct {
		Arg1     string
		Expected int
	}
	data := []param{
		// {
		// 	Arg1:     "brnk",
		// 	Expected: 4,
		// },
		{
			Arg1:     "abcabcbb",
			Expected: 3,
		},
		{
			Arg1:     "bbbbb",
			Expected: 1,
		},
		{
			Arg1:     "pwwkew",
			Expected: 3,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q3LongestSubstringWithoutRepeatingCharacters(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ2002MaximumProductOfTheLengthOfTwoPalindromicSubsequences Medium
func TestQ2002MaximumProductOfTheLengthOfTwoPalindromicSubsequences(t *testing.T) {
	type param struct {
		Arg1     string
		Expected int
	}
	data := []param{
		{
			Arg1:     "leetcodecom",
			Expected: 9,
		},
		{
			Arg1:     "bb",
			Expected: 1,
		},
		{
			Arg1:     "accbcaxxcxx",
			Expected: 25,
		},
		{
			Arg1:     "eabedb",
			Expected: 9,
		},
		{
			Arg1:     "qzq",
			Expected: 2,
		},
		{
			Arg1:     "jee",
			Expected: 2,
		},
		{
			Arg1:     "bbabb",
			Expected: 6,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2002MaximumProductOfTheLengthOfTwoPalindromicSubsequences(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ2108FindFirstPalindromicStringInTheArray Easy
func TestQ2108FindFirstPalindromicStringInTheArray(t *testing.T) {
	type param struct {
		Arg1     []string
		Expected string
	}
	data := []param{
		{
			Arg1:     JsonToSlice[string](`["abc","car","ada","racecar","cool"]`),
			Expected: "ada",
		},
		{
			Arg1:     JsonToSlice[string](`["notapalindrome","racecar"]`),
			Expected: "racecar",
		},
		{
			Arg1:     JsonToSlice[string](`["def","ghi"]`),
			Expected: "",
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2108FindFirstPalindromicStringInTheArray(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ2330ValidPalindromeIV Medium
func TestQ2330ValidPalindromeIV(t *testing.T) {
	type param struct {
		Arg1     string
		Expected bool
	}
	data := []param{
		{
			Arg1:     "abcdba",
			Expected: true,
		},
		{
			Arg1:     "aa",
			Expected: true,
		},
		{
			Arg1:     "abcdef",
			Expected: false,
		},
		{
			Arg1:     "zbcfedcba",
			Expected: true,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2330ValidPalindromeIV(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ1216ValidPalindromeIII Hard
func TestQ1216ValidPalindromeIII(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     int
		Expected bool
	}
	data := []param{
		{
			Arg1:     "abcdeca",
			Arg2:     2,
			Expected: true,
		},
		{
			Arg1:     "abbababa",
			Arg2:     1,
			Expected: true,
		},
		{
			Arg1:     "fcgihcgeadfehgiabegbiahbeadbiafgcfchbcacedbificicihibaeehbffeidiaiighceegbfdggggcfaiibefbgeegbcgeadcfdfegfghebcfceiabiagehhibiheddbcgdebdcfegaiahibcfhheggbheebfdahgcfcahafecfehgcgdabbghddeadecidicchfgicbdbecibddfcgbiadiffcifiggigdeedbiiihfgehhdegcaffaggiidiifgfigfiaiicadceefbhicfhbcachacaeiefdcchegfbifhaeafdehicfgbecahidgdagigbhiffhcccdhfdbd",
			Arg2:     216,
			Expected: true, // 超時
		},
		{
			Arg1:     "bacabaaa",
			Arg2:     2,
			Expected: false,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q1216ValidPalindromeIII(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ2130MaximumTwinSumOfALinkedList Medium
func TestQ2130MaximumTwinSumOfALinkedList(t *testing.T) {
	type param struct {
		Arg1     *ListNode
		Expected int
	}
	data := []param{
		{
			Arg1:     JSONArrayToListNode(`[5,4,2,1]`),
			Expected: 6,
		},
		{
			Arg1:     JSONArrayToListNode(`[4,2,2,3]`),
			Expected: 7,
		},
		{
			Arg1:     JSONArrayToListNode(`[1,100000]`),
			Expected: 100001,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2130MaximumTwinSumOfALinkedList(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ2074ReverseNodesInEvenLengthGroups Medium
func TestQ2074ReverseNodesInEvenLengthGroups(t *testing.T) {
	type param struct {
		Arg1     *ListNode
		Expected *ListNode
	}
	data := []param{
		{
			Arg1:     JSONArrayToListNode(`[5,2,6,3,9,1,7,3,8,4]`),
			Expected: JSONArrayToListNode(`[5,6,2,3,9,1,4,8,3,7]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[1,1,0,6]`),
			Expected: JSONArrayToListNode(`[1,0,1,6]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[1,1,0,6,5]`),
			Expected: JSONArrayToListNode(`[1,0,1,5,6]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2074ReverseNodesInEvenLengthGroups(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ680ValidPalindromeII Easy
func TestQ680ValidPalindromeII(t *testing.T) {
	type param struct {
		Arg1     string
		Expected bool
	}
	data := []param{
		{
			Arg1:     "aba",
			Expected: true,
		},
		{
			Arg1:     "abca",
			Expected: true,
		},
		{
			Arg1:     "abc",
			Expected: false,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q680ValidPalindromeII(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ125ValidPalindrome Easy
func TestQ125ValidPalindrome(t *testing.T) {
	type param struct {
		Arg1     string
		Expected bool
	}
	data := []param{
		{
			Arg1:     "A man, a plan, a canal: Panama",
			Expected: true,
		},
		{
			Arg1:     "race a car",
			Expected: false,
		},
		{
			Arg1:     " ",
			Expected: true,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q125ValidPalindrome(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ234PalindromeLinkedList Easy
func TestQ234PalindromeLinkedList(t *testing.T) {
	type param struct {
		Arg1     *ListNode
		Expected bool
	}
	data := []param{
		{
			Arg1:     JSONArrayToListNode(`[1,2,2,1]`),
			Expected: true,
		},
		{
			Arg1:     JSONArrayToListNode(`[1,2]`),
			Expected: false,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q234PalindromeLinkedList(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ156BinaryTreeUpsideDown Medium
func TestQ156BinaryTreeUpsideDown(t *testing.T) {
	type param struct {
		Arg1     *TreeNode
		Expected *TreeNode
	}
	data := []param{
		{
			Arg1:     JSONArrayToTreeNode(`[1,2,3,4,5]`),
			Expected: JSONArrayToTreeNode(`[4,5,2,null,null,3,1]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[]`),
			Expected: JSONArrayToTreeNode(`[]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[1]`),
			Expected: JSONArrayToTreeNode(`[1]`),
		},
		{
			Arg1:     JSONArrayToTreeNode(`[1,2]`),
			Expected: JSONArrayToTreeNode(`[2, null, 1]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q156BinaryTreeUpsideDown(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ92ReverseLinkedListII Medium
func TestQ92ReverseLinkedListII(t *testing.T) {
	type param struct {
		Arg1     *ListNode
		Arg2     int
		Arg3     int
		Expected *ListNode
	}
	data := []param{
		{
			Arg1:     JSONArrayToListNode(`[1,2,3,4,5]`),
			Arg2:     2,
			Arg3:     4,
			Expected: JSONArrayToListNode(`[1,4,3,2,5]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[5]`),
			Arg2:     1,
			Arg3:     1,
			Expected: JSONArrayToListNode(`[5]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[3,5]`),
			Arg2:     1,
			Arg3:     2,
			Expected: JSONArrayToListNode(`[5,3]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q92ReverseLinkedListII(d.Arg1, d.Arg2, d.Arg3), d)
		}
	}, 1, 0)
}

// TestQ206ReverseLinkedList Easy
func TestQ206ReverseLinkedList(t *testing.T) {
	type param struct {
		Arg1     *ListNode
		Expected *ListNode
	}
	data := []param{
		{
			Arg1:     JSONArrayToListNode(`[1,2]`),
			Expected: JSONArrayToListNode(`[2,1]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[]`),
			Expected: JSONArrayToListNode(`[]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q206ReverseLinkedList(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ445AddTwoNumbersII Medium
func TestQ445AddTwoNumbersII(t *testing.T) {
	type param struct {
		Arg1     *ListNode
		Arg2     *ListNode
		Expected *ListNode
	}
	data := []param{
		{
			Arg1:     JSONArrayToListNode(`[7,2,4,3]`),
			Arg2:     JSONArrayToListNode(`[5,6,4]`),
			Expected: JSONArrayToListNode(`[7,8,0,7]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[2,4,3]`),
			Arg2:     JSONArrayToListNode(`[5,6,4]`),
			Expected: JSONArrayToListNode(`[8,0,7]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[0]`),
			Arg2:     JSONArrayToListNode(`[0]`),
			Expected: JSONArrayToListNode(`[0]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q445AddTwoNumbersII(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ371SumOfTwoIntegers Medium
func TestQ371SumOfTwoIntegers(t *testing.T) {
	type param struct {
		Arg1     int
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     1,
			Arg2:     2,
			Expected: 3,
		},
		{
			Arg1:     2,
			Arg2:     3,
			Expected: 5,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q371SumOfTwoIntegers(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ2AddTwoNumbers Medium
func TestQ2AddTwoNumbers(t *testing.T) {
	type param struct {
		Arg1     *ListNode
		Arg2     *ListNode
		Expected *ListNode
	}
	data := []param{
		{
			Arg1:     JSONArrayToListNode(`[2,4,3]`),
			Arg2:     JSONArrayToListNode(`[5,6,4]`),
			Expected: JSONArrayToListNode(`[7,0,8]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[0]`),
			Arg2:     JSONArrayToListNode(`[0]`),
			Expected: JSONArrayToListNode(`[0]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[9,9,9,9,9,9,9]`),
			Arg2:     JSONArrayToListNode(`[9,9,9,9]`),
			Expected: JSONArrayToListNode(`[8,9,9,9,0,0,0,1]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2AddTwoNumbers(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ415AddStrings Easy
func TestQ415AddStrings(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     string
		Expected string
	}
	data := []param{
		{
			Arg1:     "11",
			Arg2:     "123",
			Expected: "134",
		},
		{
			Arg1:     "456",
			Arg2:     "77",
			Expected: "533",
		},
		{
			Arg1:     "0",
			Arg2:     "0",
			Expected: "0",
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q415AddStrings(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ989AddToArray_FormOfInteger Easy
func TestQ989AddToArray_FormOfInteger(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Expected []int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,2,0,0]`),
			Arg2:     34,
			Expected: JsonToSlice[int](`[1,2,3,4]`),
		},
		{
			Arg1:     JsonToSlice[int](`[2,7,4]`),
			Arg2:     181,
			Expected: JsonToSlice[int](`[4,5,5]`),
		},
		{
			Arg1:     JsonToSlice[int](`[2,1,5]`),
			Arg2:     806,
			Expected: JsonToSlice[int](`[1,0,2,1]`),
		},
		{
			Arg1:     JsonToSlice[int](`[9,9,9,9,9,9,9,9,9,9]`),
			Arg2:     1,
			Expected: JsonToSlice[int](`[1,0,0,0,0,0,0,0,0,0,0]`),
		},
		{
			Arg1:     JsonToSlice[int](`[0]`),
			Arg2:     123,
			Expected: JsonToSlice[int](`[1,2,3]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q989AddToArray_FormOfInteger(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ369PlusOneLinkedList Medium
func TestQ369PlusOneLinkedList(t *testing.T) {
	type param struct {
		Arg1     *ListNode
		Expected *ListNode
	}
	data := []param{
		{
			Arg1:     JSONArrayToListNode(`[1,2,3]`),
			Expected: JSONArrayToListNode(`[1,2,4]`),
		},
		{
			Arg1:     JSONArrayToListNode(`[0]`),
			Expected: JSONArrayToListNode(`[1]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q369PlusOneLinkedList(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ43MultiplyStrings Easy
func TestQ43MultiplyStrings(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     string
		Expected string
	}
	data := []param{
		{
			Arg1:     "2",
			Arg2:     "3",
			Expected: "6",
		},
		{
			Arg1:     "123",
			Arg2:     "456",
			Expected: "56088",
		},
		{
			Arg1:     "9133",
			Arg2:     "0",
			Expected: "0",
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q43MultiplyStrings(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ66PlusOne Easy
func TestQ66PlusOne(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected []int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,2,3]`),
			Expected: JsonToSlice[int](`[1,2,4]`),
		},
		{
			Arg1:     JsonToSlice[int](`[4,3,2,1]`),
			Expected: JsonToSlice[int](`[4,3,2,2]`),
		},
		{
			Arg1:     JsonToSlice[int](`[9]`),
			Expected: JsonToSlice[int](`[1,0]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q66PlusOne(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ67AddBinary Easy
func TestQ67AddBinary(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     string
		Expected string
	}
	data := []param{
		{
			Arg1:     "11",
			Arg2:     "1",
			Expected: "100",
		},
		{
			Arg1:     "1010",
			Arg2:     "1011",
			Expected: "10101",
		},
		{
			Arg1:     "11",
			Arg2:     "111101",
			Expected: "1000000",
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q67AddBinary(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ718MaximumLengthOfRepeatedSubarray Medium *fail*
func TestQ718MaximumLengthOfRepeatedSubarray(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,2,3,2,1]`),
			Arg2:     JsonToSlice[int](`[3,2,1,4,7]`),
			Expected: 3,
		},
		{
			Arg1:     JsonToSlice[int](`[0,0,0,0,0]`),
			Arg2:     JsonToSlice[int](`[0,0,0,0,0]`),
			Expected: 5,
		},
		{
			Arg1:     JsonToSlice[int](`[0,0,0,0,1]`),
			Arg2:     JsonToSlice[int](`[1,0,0,0,0]`),
			Expected: 4,
		},
		{
			Arg1:     JsonToSlice[int](`[0,0,0,0,0,0,1,0,0,0]`),
			Arg2:     JsonToSlice[int](`[0,0,0,0,0,0,0,1,0,0]`),
			Expected: 9,
		},
		{
			Arg1:     JsonToSlice[int](`[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]`),
			Arg2:     JsonToSlice[int](`[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]`),
			Expected: 1000,
		},
	}
	Decorate(func() {
		for _, d := range data {
			fmt.Println(len(d.Arg1))
			assert.Equal(t, d.Expected, Q718MaximumLengthOfRepeatedSubarray(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)
}

// TestQ1995CountSpecialQuadruplets Easy
func TestQ1995CountSpecialQuadruplets(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,2,3,6]`),
			Expected: 1,
		},
		{
			Arg1:     JsonToSlice[int](`[3,3,6,4,5]`),
			Expected: 0,
		},
		{
			Arg1:     JsonToSlice[int](`[1,1,1,3,5]`),
			Expected: 4,
		},
		{
			Arg1:     JsonToSlice[int](`[28,8,49,85,37,90,20,8]`),
			Expected: 1,
		},
		{
			Arg1:     JsonToSlice[int](`[2,16,9,27,3,39]`),
			Expected: 2,
		},
		{
			Arg1:     JsonToSlice[int](`[8,73,11,9,28,92,52]`),
			Expected: 2,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q1995CountSpecialQuadruplets(d.Arg1), d)
		}
	}, 1, 0)
}

// TestQ4544SumII Medium
func TestQ4544SumII(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     []int
		Arg3     []int
		Arg4     []int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,2]`),
			Arg2:     JsonToSlice[int](`[-2,-1]`),
			Arg3:     JsonToSlice[int](`[-1,2]`),
			Arg4:     JsonToSlice[int](`[0,2]`),
			Expected: 2,
		},
		{
			Arg1:     JsonToSlice[int](`[0]`),
			Arg2:     JsonToSlice[int](`[0]`),
			Arg3:     JsonToSlice[int](`[0]`),
			Arg4:     JsonToSlice[int](`[0]`),
			Expected: 1,
		},
		{
			Arg1:     JsonToSlice[int](`[0,1,-1]`),
			Arg2:     JsonToSlice[int](`[-1,1,0]`),
			Arg3:     JsonToSlice[int](`[0,0,1]`),
			Arg4:     JsonToSlice[int](`[-1,1,1]`),
			Expected: 17,
		},
		{
			Arg1:     JsonToSlice[int](`[-1,1,1,1,-1]`),
			Arg2:     JsonToSlice[int](`[0,-1,-1,0,1]`),
			Arg3:     JsonToSlice[int](`[-1,-1,1,-1,-1]`),
			Arg4:     JsonToSlice[int](`[0,1,0,-1,-1]`),
			Expected: 132,
		},
		{
			Arg1:     JsonToSlice[int](`[-39,78,-65,-35,50,-81,-72,83,-92,-26,-21,39,-1,94,2,-2,95,18,-81,85,-68,49,80,48,-62,46,80,-36,-78,-33,-6,25,44,-48,-41,-66,87,-52,-36,65,30,-42,-77,-65,-27,57,-68,-89,66,54,-20,-64,-100,22,-67,-8,1,-100,82,-90,93,66,-82,-74,71,-12,80,73,53,-9,31,-74,81,-44,-39,-55,66,-3,6,-27,56,60,-90,-35,55,66,36,79,-70,-55,38,2,-75,-34,18,95,55,9,-4,-64,-86,-32,-4,-70,-95,-38,9,83,-35,-84,-69,-37,-65,-36,44,27,73,-61,-81,-74,7,4,-1,-50,39,96,4,-21,-52,58,3,58,-20,-3,-76,71,28,-76,12,-83,-97,-19,-77,2,5,30,-13,-66,-8,-82,-35,-3,75,43,77,-50,40,-67,-56,26,-25,29,-55,-5,6,56,62,26,-77,-71,64,42,-15,-8,-80,-23,-1,35,57,-23,-68,0,-65,82,-19,-27,-58]`),
			Arg2:     JsonToSlice[int](`[27,-76,-27,51,2,-23,50,-37,-4,-47,23,-56,-10,90,41,77,48,-53,65,78,-92,64,84,89,36,78,-94,-15,-42,50,97,86,54,50,-21,11,-45,-9,-90,-94,65,79,28,-81,75,92,-64,97,41,-92,-36,82,69,79,1,-100,88,43,1,-55,30,48,5,22,85,88,-83,5,42,-88,88,65,40,73,92,-27,8,-56,-73,-56,-46,-93,53,-92,-72,1,-5,57,-47,67,-18,15,22,30,-70,-24,59,-20,35,-1,-21,-52,-70,-81,-8,-17,21,-58,-70,-29,27,-66,28,19,25,-58,49,61,-74,89,53,37,89,39,-48,-33,-30,-31,-1,-1,83,-56,63,85,96,60,-32,27,-4,48,-19,-30,80,-3,4,93,-65,-91,-26,-34,-6,-28,54,4,24,50,43,-76,-76,10,42,-75,-11,93,-72,41,59,60,6,48,-43,-82,-31,59,85,40,-99,-33,94,36,-68,69,92,56,6,3,58]`),
			Arg3:     JsonToSlice[int](`[-74,27,32,49,-35,91,-41,66,-100,-39,69,-98,-72,42,-8,-58,-36,62,32,-69,53,-13,-49,-79,-32,31,10,-38,-75,-15,1,-58,-100,66,-14,-70,-97,-91,57,44,70,86,49,-2,-46,-96,-71,61,-10,-40,-13,-55,77,-56,8,-69,-88,-91,88,-34,-24,13,47,64,62,-82,-84,-100,28,38,-70,95,-23,51,-9,90,1,38,9,-91,-86,-99,73,63,-98,30,-73,-63,14,40,81,28,26,12,53,93,71,98,-97,-69,-65,84,-69,71,-84,38,-70,72,-30,-89,72,-70,75,30,-21,-86,-47,24,-93,38,33,-9,57,76,35,36,-25,-22,92,-41,-83,-79,35,-43,-64,35,72,49,92,96,98,14,2,-21,-42,88,97,-4,-2,-64,12,9,-59,63,-56,100,-8,91,-34,100,15,-87,-80,89,-57,40,40,-14,75,85,-88,-95,-63,-100,-97,-1,60,23,93,59,61,-54,-44,84,70,17,89]`),
			Arg4:     JsonToSlice[int](`[60,11,22,-35,-57,-80,-78,90,-47,-67,5,78,52,-42,41,19,-47,-35,-79,30,60,-64,-67,-48,-76,-51,-28,79,45,53,81,58,61,52,48,98,68,22,47,-29,86,65,35,76,34,4,-18,-65,32,-91,-53,-9,51,86,-32,42,3,41,78,-64,-94,55,-98,26,-60,17,-98,-3,-67,42,10,52,45,34,53,39,-24,78,95,-74,91,-54,-63,14,40,-12,59,13,71,87,-1,-63,-14,-42,36,63,-52,-94,-80,-88,-66,-30,-97,-87,64,-70,42,70,-78,-14,-35,-96,95,-42,-76,-97,-8,-76,9,46,-98,-13,-27,54,-89,24,-28,-92,-90,-90,74,-86,-42,53,62,0,89,64,7,-48,86,75,91,-49,-20,-62,52,-12,46,55,-24,-68,47,-3,-55,67,-50,48,-8,-32,-14,57,41,76,2,23,-76,-74,-73,58,0,-52,6,12,39,-82,62,-16,-59,-82,-53,-12,38,-57,-72,58,37]`),
			Expected: 4013055, // 超時測試 4013055
		},
		{
			Arg1:     JsonToSlice[int](`59,92,-43,-93,48,89,-4,-19,37,36,35,9,13,64,-100,68,41,72,57,33,-64,-52,43,-72,87,-44,-69,-39,93,74,-47,-98,-24,75,-97,46,-49,-76,12,-34,16,29,-73,57,9,79,-90,-71,-16,-59,57,59,-55,-89,-46,66,27,82,5,62,86,97,47,72,-20,-66,76,2,75,-51,-34,33,43,-78,7,-86,-45,-93,-50,30,-55,-54,-63,-53,-7,27,-70,85,-30,-7,3,-78,-1,-62,-34,5,-51,66,-6,-48,33,63,98,59,-13,-39,-78,-60,-51,43,-30,35,-29,-57,7,9,-21,12,58,34,74,50,33,50,25,-55,67,62,-3,40,-11,-97,57,93,34,28,43,92,18,-100,63,45,13,36,90,63,85]`),
			Arg2:     JsonToSlice[int](`49,1,-9,-7,-68,76,0,-69,17,20,99,-73,-15,-2,-8,36,27,22,84,39,-5,-25,-74,-14,-14,-41,-56,-17,-9,-15,16,57,-82,-32,-10,-97,-68,-99,-65,-72,-75,15,3,20,-21,92,-66,62,-71,70,80,46,-83,60,-87,-86,94,85,-39,-45,65,79,92,97,60,-10,27,-64,-69,-7,-43,27,-48,1,-65,72,-88,93,-97,-5,54,68,62,55,-82,-94,43,-66,-51,-91,19,-11,-4,32,-62,8,91,-89,39,40,-44,26,43,15,47,-93,18,49,-59,-75,83,-32,6,-90,81,46,91,-23,2,-49,33,33,58,-52,-37,-91,-30,-79,-16,93,-62,40,74,-100,35,70,-75,46,23,-60,-47,-80,-81,89,-49,100,47]`),
			Arg3:     JsonToSlice[int](`-38,30,58,35,-46,-22,-13,-52,61,67,-88,29,20,80,-87,58,79,21,-45,-17,31,-64,85,82,-89,-67,-53,21,-30,-96,40,81,-25,22,27,-27,85,-28,-41,-95,94,22,-18,13,-55,43,5,-99,70,-77,96,73,-23,-55,-15,-84,96,-54,-92,100,-64,74,77,90,38,52,-29,67,55,97,-36,83,98,-39,22,17,89,-16,-85,81,-13,61,-90,-69,-76,-83,74,-63,-19,-58,44,-3,-99,-94,-7,92,-73,-19,-83,92,35,-91,72,-17,-78,76,-28,15,75,66,-58,-44,-11,-52,-90,-37,-46,84,60,68,-74,70,-9,100,-47,56,-78,23,-30,26,80,95,92,34,3,-28,-19,12,58,-44,-30,26,16,-79,-23,82,94]`),
			Arg4:     JsonToSlice[int](`-67,77,-65,65,-27,-48,51,32,-10,89,-43,-19,90,-13,-72,-37,-78,97,-41,-2,80,-30,51,88,16,15,-18,37,18,7,-95,18,16,-81,-81,34,-75,-17,-91,-89,-22,57,-39,17,27,-94,-78,-49,-35,12,-27,-52,52,61,87,46,30,36,-61,3,-19,-30,97,-7,46,43,16,30,-91,42,86,68,4,-88,17,32,96,-49,-67,71,-14,17,-37,73,-2,94,-87,-17,-90,-32,-60,86,-65,-99,14,54,-77,-11,-5,93,10,80,-83,-6,18,-3,92,-15,99,1,-40,-56,-27,18,39,-48,57,-63,-32,-1,72,53,-79,-95,-35,40,-10,70,-12,-71,39,-44,8,47,-12,-32,-83,-100,-83,54,54,-2,-33,69,-96,-6,-98]`),
			Expected: 0, // 超時測試
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q4544SumII(d.Arg1, d.Arg2, d.Arg3, d.Arg4), d)
		}
	}, 1, 0)
}

// TestQ153Sum Medium
func TestQ153Sum(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected [][]int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[-1,0,1,2,-1,-4]`),
			Expected: JsonToSliceSlice[int](`[[-1,-1,2],[-1,0,1]]`),
		},
		{
			Arg1:     JsonToSlice[int](`[0,1,1]`),
			Expected: JsonToSliceSlice[int](`[]`),
		},
		{
			Arg1:     JsonToSlice[int](`[0,0,0]`),
			Expected: JsonToSliceSlice[int](`[[0,0,0]]`),
		},
		{
			Arg1:     JsonToSlice[int](`[1,2,-2,-1]`),
			Expected: JsonToSliceSlice[int](`[]`),
		},
		// {
		// 	Arg1:     JsonToSlice[int](`[82597,-9243,62390,83030,-97960,-26521,-61011,83390,-38677,12333,75987,46091,83794,19355,-71037,-6242,-28801,324,1202,-90885,-2989,-95597,-34333,35528,5680,89093,-90606,50360,-29393,-27012,53313,65213,99818,-82405,-41661,-3333,-51952,72135,-1523,26377,74685,96992,92263,15929,5467,-99555,-43348,-41689,-60383,-3990,32165,65265,-72973,-58372,12741,-48568,-46596,72419,-1859,34153,62937,81310,-61823,-96770,-54944,8845,-91184,24208,-29078,31495,65258,14198,85395,70506,-40908,56740,-12228,-40072,32429,93001,68445,-73927,25731,-91859,-24150,10093,-60271,-81683,-18126,51055,48189,-6468,25057,81194,-58628,74042,66158,-14452,-49851,-43667,11092,39189,-17025,-79173,13606,83172,92647,-59741,19343,-26644,-57607,82908,-20655,1637,80060,98994,39331,-31274,-61523,91225,-72953,13211,-75116,-98421,-41571,-69074,99587,39345,42151,-2460,98236,15690,-52507,-95803,-48935,-46492,-45606,-79254,-99851,52533,73486,39948,-7240,71815,-585,-96252,90990,-93815,93340,-71848,58733,-14859,-83082,-75794,-82082,-24871,-15206,91207,-56469,-93618,67131,-8682,75719,87429,-98757,-7535,-24890,-94160,85003,33928,75538,97456,-66424,-60074,-8527,-28697,-22308,2246,-70134,-82319,-10184,87081,-34949,-28645,-47352,-83966,-60418,-15293,-53067,-25921,55172,75064,95859,48049,34311,-86931,-38586,33686,-36714,96922,76713,-22165,-80585,-34503,-44516,39217,-28457,47227,-94036,43457,24626,-87359,26898,-70819,30528,-32397,-69486,84912,-1187,-98986,-32958,4280,-79129,-65604,9344,58964,50584,71128,-55480,24986,15086,-62360,-42977,-49482,-77256,-36895,-74818,20,3063,-49426,28152,-97329,6086,86035,-88743,35241,44249,19927,-10660,89404,24179,-26621,-6511,57745,-28750,96340,-97160,-97822,-49979,52307,79462,94273,-24808,77104,9255,-83057,77655,21361,55956,-9096,48599,-40490,-55107,2689,29608,20497,66834,-34678,23553,-81400,-66630,-96321,-34499,-12957,-20564,25610,-4322,-58462,20801,53700,71527,24669,-54534,57879,-3221,33636,3900,97832,-27688,-98715,5992,24520,-55401,-57613,-69926,57377,-77610,20123,52174,860,60429,-91994,-62403,-6218,-90610,-37263,-15052,62069,-96465,44254,89892,-3406,19121,-41842,-87783,-64125,-56120,73904,-22797,-58118,-4866,5356,75318,46119,21276,-19246,-9241,-97425,57333,-15802,93149,25689,-5532,95716,39209,-87672,-29470,-16324,-15331,27632,-39454,56530,-16000,29853,46475,78242,-46602,83192,-73440,-15816,50964,-36601,89758,38375,-40007,-36675,-94030,67576,46811,-64919,45595,76530,40398,35845,41791,67697,-30439,-82944,63115,33447,-36046,-50122,-34789,43003,-78947,-38763,-89210,32756,-20389,-31358,-90526,-81607,88741,86643,98422,47389,-75189,13091,95993,-15501,94260,-25584,-1483,-67261,-70753,25160,89614,-90620,-48542,83889,-12388,-9642,-37043,-67663,28794,-8801,13621,12241,55379,84290,21692,-95906,-85617,-17341,-63767,80183,-4942,-51478,30997,-13658,8838,17452,-82869,-39897,68449,31964,98158,-49489,62283,-62209,-92792,-59342,55146,-38533,20496,62667,62593,36095,-12470,5453,-50451,74716,-17902,3302,-16760,-71642,-34819,96459,-72860,21638,47342,-69897,-40180,44466,76496,84659,13848,-91600,-90887,-63742,-2156,-84981,-99280,94326,-33854,92029,-50811,98711,-36459,-75555,79110,-88164,-97397,-84217,97457,64387,30513,-53190,-83215,252,2344,-27177,-92945,-89010,82662,-11670,86069,53417,42702,97082,3695,-14530,-46334,17910,77999,28009,-12374,15498,-46941,97088,-35030,95040,92095,-59469,-24761,46491,67357,-66658,37446,-65130,-50416,99197,30925,27308,54122,-44719,12582,-99525,-38446,-69050,-22352,94757,-56062,33684,-40199,-46399,96842,-50881,-22380,-65021,40582,53623,-76034,77018,-97074,-84838,-22953,-74205,79715,-33920,-35794,-91369,73421,-82492,63680,-14915,-33295,37145,76852,-69442,60125,-74166,74308,-1900,-30195,-16267,-60781,-27760,5852,38917,25742,-3765,49097,-63541,98612,-92865,-30248,9612,-8798,53262,95781,-42278,-36529,7252,-27394,-5021,59178,80934,-48480,-75131,-54439,-19145,-48140,98457,-6601,-51616,-89730,78028,32083,-48904,16822,-81153,-8832,48720,-80728,-45133,-86647,-4259,-40453,2590,28613,50523,-4105,-27790,-74579,-17223,63721,33489,-47921,97628,-97691,-14782,-65644,18008,-93651,-71266,80990,-76732,-47104,35368,28632,59818,-86269,-89753,34557,-92230,-5933,-3487,-73557,-13174,-43981,-43630,-55171,30254,-83710,-99583,-13500,71787,5017,-25117,-78586,86941,-3251,-23867,-36315,75973,86272,-45575,77462,-98836,-10859,70168,-32971,-38739,-12761,93410,14014,-30706,-77356,-85965,-62316,63918,-59914,-64088,1591,-10957,38004,15129,-83602,-51791,34381,-89382,-26056,8942,5465,71458,-73805,-87445,-19921,-80784,69150,-34168,28301,-68955,18041,6059,82342,9947,39795,44047,-57313,48569,81936,-2863,-80932,32976,-86454,-84207,33033,32867,9104,-16580,-25727,80157,-70169,53741,86522,84651,68480,84018,61932,7332,-61322,-69663,76370,41206,12326,-34689,17016,82975,-23386,39417,72793,44774,-96259,3213,79952,29265,-61492,-49337,14162,65886,3342,-41622,-62659,-90402,-24751,88511,54739,-21383,-40161,-96610,-24944,-602,-76842,-21856,69964,43994,-15121,-85530,12718,13170,-13547,69222,62417,-75305,-81446,-38786,-52075,-23110,97681,-82800,-53178,11474,35857,94197,-58148,-23689,32506,92154,-64536,-73930,-77138,97446,-83459,70963,22452,68472,-3728,-25059,-49405,95129,-6167,12808,99918,30113,-12641,-26665,86362,-33505,50661,26714,33701,89012,-91540,40517,-12716,-57185,-87230,29914,-59560,13200,-72723,58272,23913,-45586,-96593,-26265,-2141,31087,81399,92511,-34049,20577,2803,26003,8940,42117,40887,-82715,38269,40969,-50022,72088,21291,-67280,-16523,90535,18669,94342,-39568,-88080,-99486,-20716,23108,-28037,63342,36863,-29420,-44016,75135,73415,16059,-4899,86893,43136,-7041,33483,-67612,25327,40830,6184,61805,4247,81119,-22854,-26104,-63466,63093,-63685,60369,51023,51644,-16350,74438,-83514,99083,10079,-58451,-79621,48471,67131,-86940,99093,11855,-22272,-67683,-44371,9541,18123,37766,-70922,80385,-57513,-76021,-47890,36154,72935,84387,-92681,-88303,-7810,59902,-90,-64704,-28396,-66403,8860,13343,33882,85680,7228,28160,-14003,54369,-58893,92606,-63492,-10101,64714,58486,29948,-44679,-22763,10151,-56695,4031,-18242,-36232,86168,-14263,9883,47124,47271,92761,-24958,-73263,-79661,-69147,-18874,29546,-92588,-85771,26451,-86650,-43306,-59094,-47492,-34821,-91763,-47670,33537,22843,67417,-759,92159,63075,94065,-26988,55276,65903,30414,-67129,-99508,-83092,-91493,-50426,14349,-83216,-76090,32742,-5306,-93310,-60750,-60620,-45484,-21108,-58341,-28048,-52803,69735,78906,81649,32565,-86804,-83202,-65688,-1760,89707,93322,-72750,84134,71900,-37720,19450,-78018,22001,-23604,26276,-21498,65892,-72117,-89834,-23867,55817,-77963,42518,93123,-83916,63260,-2243,-97108,85442,-36775,17984,-58810,99664,-19082,93075,-69329,87061,79713,16296,70996,13483,-74582,49900,-27669,-40562,1209,-20572,34660,83193,75579,7344,64925,88361,60969,3114,44611,-27445,53049,-16085,-92851,-53306,13859,-33532,86622,-75666,-18159,-98256,51875,-42251,-27977,-18080,23772,38160,41779,9147,94175,99905,-85755,62535,-88412,-52038,-68171,93255,-44684,-11242,-104,31796,62346,-54931,-55790,-70032,46221,56541,-91947,90592,93503,4071,20646,4856,-63598,15396,-50708,32138,-85164,38528,-89959,53852,57915,-42421,-88916,-75072,67030,-29066,49542,-71591,61708,-53985,-43051,28483,46991,-83216,80991,-46254,-48716,39356,-8270,-47763,-34410,874,-1186,-7049,28846,11276,21960,-13304,-11433,-4913,55754,79616,70423,-27523,64803,49277,14906,-97401,-92390,91075,70736,21971,-3303,55333,-93996,76538,54603,-75899,98801,46887,35041,48302,-52318,55439,24574,14079,-24889,83440,14961,34312,-89260,-22293,-81271,-2586,-71059,-10640,-93095,-5453,-70041,66543,74012,-11662,-52477,-37597,-70919,92971,-17452,-67306,-80418,7225,-89296,24296,86547,37154,-10696,74436,-63959,58860,33590,-88925,-97814,-83664,85484,-8385,-50879,57729,-74728,-87852,-15524,-91120,22062,28134,80917,32026,49707,-54252,-44319,-35139,13777,44660,85274,25043,58781,-89035,-76274,6364,-63625,72855,43242,-35033,12820,-27460,77372,-47578,-61162,-70758,-1343,-4159,64935,56024,-2151,43770,19758,-30186,-86040,24666,-62332,-67542,73180,-25821,-27826,-45504,-36858,-12041,20017,-24066,-56625,-52097,-47239,-90694,8959,7712,-14258,-5860,55349,61808,-4423,-93703,64681,-98641,-25222,46999,-83831,-54714,19997,-68477,66073,51801,-66491,52061,-52866,79907,-39736,-68331,68937,91464,98892,910,93501,31295,-85873,27036,-57340,50412,21,-2445,29471,71317,82093,-94823,-54458,-97410,39560,-7628,66452,39701,54029,37906,46773,58296,60370,-61090,85501,-86874,71443,-72702,-72047,14848,34102,77975,-66294,-36576,31349,52493,-70833,-80287,94435,39745,-98291,84524,-18942,10236,93448,50846,94023,-6939,47999,14740,30165,81048,84935,-19177,-13594,32289,62628,-90612,-542,-66627,64255,71199,-83841,-82943,-73885,8623,-67214,-9474,-35249,62254,-14087,-90969,21515,-83303,94377,-91619,19956,-98810,96727,-91939,29119,-85473,-82153,-69008,44850,74299,-76459,-86464,8315,-49912,-28665,59052,-69708,76024,-92738,50098,18683,-91438,18096,-19335,35659,91826,15779,-73070,67873,-12458,-71440,-46721,54856,97212,-81875,35805,36952,68498,81627,-34231,81712,27100,-9741,-82612,18766,-36392,2759,41728,69743,26825,48355,-17790,17165,56558,3295,-24375,55669,-16109,24079,73414,48990,-11931,-78214,90745,19878,35673,-15317,-89086,94675,-92513,88410,-93248,-19475,-74041,-19165,32329,-26266,-46828,-18747,45328,8990,-78219,-25874,-74801,-44956,-54577,-29756,-99822,-35731,-18348,-68915,-83518,-53451,95471,-2954,-13706,-8763,-21642,-37210,16814,-60070,-42743,27697,-36333,-42362,11576,85742,-82536,68767,-56103,-63012,71396,-78464,-68101,-15917,-11113,-3596,77626,-60191,-30585,-73584,6214,-84303,18403,23618,-15619,-89755,-59515,-59103,-74308,-63725,-29364,-52376,-96130,70894,-12609,50845,-2314,42264,-70825,64481,55752,4460,-68603,-88701,4713,-50441,-51333,-77907,97412,-66616,-49430,60489,-85262,-97621,-18980,44727,-69321,-57730,66287,-92566,-64427,-14270,11515,-92612,-87645,61557,24197,-81923,-39831,-10301,-23640,-76219,-68025,92761,-76493,68554,-77734,-95620,-11753,-51700,98234,-68544,-61838,29467,46603,-18221,-35441,74537,40327,-58293,75755,-57301,-7532,-94163,18179,-14388,-22258,-46417,-48285,18242,-77551,82620,250,-20060,-79568,-77259,82052,-98897,-75464,48773,-79040,-11293,45941,-67876,-69204,-46477,-46107,792,60546,-34573,-12879,-94562,20356,-48004,-62429,96242,40594,2099,99494,25724,-39394,-2388,-18563,-56510,-83570,-29214,3015,74454,74197,76678,-46597,60630,-76093,37578,-82045,-24077,62082,-87787,-74936,58687,12200,-98952,70155,-77370,21710,-84625,-60556,-84128,925,65474,-15741,-94619,88377,89334,44749,22002,-45750,-93081,-14600,-83447,46691,85040,-66447,-80085,56308,44310,24979,-29694,57991,4675,-71273,-44508,13615,-54710,23552,-78253,-34637,50497,68706,81543,-88408,-21405,6001,-33834,-21570,-46692,-25344,20310,71258,-97680,11721,59977,59247,-48949,98955,-50276,-80844,-27935,-76102,55858,-33492,40680,66691,-33188,8284,64893,-7528,6019,-85523,8434,-64366,-56663,26862,30008,-7611,-12179,-70076,21426,-11261,-36864,-61937,-59677,929,-21052,3848,-20888,-16065,98995,-32293,-86121,-54564,77831,68602,74977,31658,40699,29755,98424,80358,-69337,26339,13213,-46016,-18331,64713,-46883,-58451,-70024,-92393,-4088,70628,-51185,71164,-75791,-1636,-29102,-16929,-87650,-84589,-24229,-42137,-15653,94825,13042,88499,-47100,-90358,-7180,29754,-65727,-42659,-85560,-9037,-52459,20997,-47425,17318,21122,20472,-23037,65216,-63625,-7877,-91907,24100,-72516,22903,-85247,-8938,73878,54953,87480,-31466,-99524,35369,-78376,89984,-15982,94045,-7269,23319,-80456,-37653,-76756,2909,81936,54958,-12393,60560,-84664,-82413,66941,-26573,-97532,64460,18593,-85789,-38820,-92575,-43663,-89435,83272,-50585,13616,-71541,-53156,727,-27644,16538,34049,57745,34348,35009,16634,-18791,23271,-63844,95817,21781,16590,59669,15966,-6864,48050,-36143,97427,-59390,96931,78939,-1958,50777,43338,-51149,39235,-27054,-43492,67457,-83616,37179,10390,85818,2391,73635,87579,-49127,-81264,-79023,-81590,53554,-74972,-83940,-13726,-39095,29174,78072,76104,47778,25797,-29515,-6493,-92793,22481,-36197,-65560,42342,15750,97556,99634,-56048,-35688,13501,63969,-74291,50911,39225,93702,-3490,-59461,-30105,-46761,-80113,92906,-68487,50742,36152,-90240,-83631,24597,-50566,-15477,18470,77038,40223,-80364,-98676,70957,-63647,99537,13041,31679,86631,37633,-16866,13686,-71565,21652,-46053,-80578,-61382,68487,-6417,4656,20811,67013,-30868,-11219,46,74944,14627,56965,42275,-52480,52162,-84883,-52579,-90331,92792,42184,-73422,-58440,65308,-25069,5475,-57996,59557,-17561,2826,-56939,14996,-94855,-53707,99159,43645,-67719,-1331,21412,41704,31612,32622,1919,-69333,-69828,22422,-78842,57896,-17363,27979,-76897,35008,46482,-75289,65799,20057,7170,41326,-76069,90840,-81253,-50749,3649,-42315,45238,-33924,62101,96906,58884,-7617,-28689,-66578,62458,50876,-57553,6739,41014,-64040,-34916,37940,13048,-97478,-11318,-89440,-31933,-40357,-59737,-76718,-14104,-31774,28001,4103,41702,-25120,-31654,63085,-3642,84870,-83896,-76422,-61520,12900,88678,85547,33132,-88627,52820,63915,-27472,78867,-51439,33005,-23447,-3271,-39308,39726,-74260,-31874,-36893,93656,910,-98362,60450,-88048,99308,13947,83996,-90415,-35117,70858,-55332,-31721,97528,82982,-86218,6822,25227,36946,97077,-4257,-41526,56795,89870,75860,-70802,21779,14184,-16511,-89156,-31422,71470,69600,-78498,74079,-19410,40311,28501,26397,-67574,-32518,68510,38615,19355,-6088,-97159,-29255,-92523,3023,-42536,-88681,64255,41206,44119,52208,39522,-52108,91276,-70514,83436,63289,-79741,9623,99559,12642,85950,83735,-21156,-67208,98088,-7341,-27763,-30048,-44099,-14866,-45504,-91704,19369,13700,10481,-49344,-85686,33994,19672,36028,60842,66564,-24919,33950,-93616,-47430,-35391,-28279,56806,74690,39284,-96683,-7642,-75232,37657,-14531,-86870,-9274,-26173,98640,88652,64257,46457,37814,-19370,9337,-22556,-41525,39105,-28719,51611,-93252,98044,-90996,21710,-47605,-64259,-32727,53611,-31918,-3555,33316,-66472,21274,-37731,-2919,15016,48779,-88868,1897,41728,46344,-89667,37848,68092,-44011,85354,-43776,38739,-31423,-66330,65167,-22016,59405,34328,-60042,87660,-67698,-59174,-1408,-46809,-43485,-88807,-60489,13974,22319,55836,-62995,-37375,-4185,32687,-36551,-75237,58280,26942,-73756,71756,78775,-40573,14367,-71622,-77338,24112,23414,-7679,-51721,87492,85066,-21612,57045,10673,-96836,52461,-62218,-9310,65862,-22748,89906,-96987,-98698,26956,-43428,46141,47456,28095,55952,67323,-36455,-60202,-43302,-82932,42020,77036,10142,60406,70331,63836,58850,-66752,52109,21395,-10238,-98647,-41962,27778,69060,98535,-28680,-52263,-56679,66103,-42426,27203,80021,10153,58678,36398,63112,34911,20515,62082,-15659,-40785,27054,43767,-20289,65838,-6954,-60228,-72226,52236,-35464,25209,-15462,-79617,-41668,-84083,62404,-69062,18913,46545,20757,13805,24717,-18461,-47009,-25779,68834,64824,34473,39576,31570,14861,-15114,-41233,95509,68232,67846,84902,-83060,17642,-18422,73688,77671,-26930,64484,-99637,73875,6428,21034,-73471,19664,-68031,15922,-27028,48137,54955,-82793,-41144,-10218,-24921,-28299,-2288,68518,-54452,15686,-41814,66165,-72207,-61986,80020,50544,-99500,16244,78998,40989,14525,-56061,-24692,-94790,21111,37296,-90794,72100,70550,-31757,17708,-74290,61910,78039,-78629,-25033,73172,-91953,10052,64502,99585,-1741,90324,-73723,68942,28149,30218,24422,16659,10710,-62594,94249,96588,46192,34251,73500,-65995,-81168,41412,-98724,-63710,-54696,-52407,19746,45869,27821,-94866,-76705,-13417,-61995,-71560,43450,67384,-8838,-80293,-28937,23330,-89694,-40586,46918,80429,-5475,78013,25309,-34162,37236,-77577,86744,26281,-29033,-91813,35347,13033,-13631,-24459,3325,-71078,-75359,81311,19700,47678,-74680,-84113,45192,35502,37675,19553,76522,-51098,-18211,89717,4508,-82946,27749,85995,89912,-53678,-64727,-14778,32075,-63412,-40524,86440,-2707,-36821,63850,-30883,67294,-99468,-23708,34932,34386,98899,29239,-23385,5897,54882,98660,49098,70275,17718,88533,52161,63340,50061,-89457,19491,-99156,24873,-17008,64610,-55543,50495,17056,-10400,-56678,-29073,-42960,-76418,98562,-88104,-96255,10159,-90724,54011,12052,45871,-90933,-69420,67039,37202,78051,-52197,-40278,-58425,65414,-23394,-1415,6912,-53447,7352,17307,-78147,63727,98905,55412,-57658,-32884,-44878,22755,39730,3638,35111,39777,74193,38736,-11829,-61188,-92757,55946,-71232,-63032,-83947,39147,-96684,-99233,25131,-32197,24406,-55428,-61941,25874,-69453,64483,-19644,-68441,12783,87338,-48676,66451,-447,-61590,50932,-11270,29035,65698,-63544,10029,80499,-9461,86368,91365,-81810,-71914,-52056,-13782,44240,-30093,-2437,24007,67581,-17365,-69164,-8420,-69289,-29370,48010,90439,13141,69243,50668,39328,61731,78266,-81313,17921,-38196,55261,9948,-24970,75712,-72106,28696,7461,31621,61047,51476,56512,11839,-96916,-82739,28924,-99927,58449,37280,69357,11219,-32119,-62050,-48745,-83486,-52376,42668,82659,68882,38773,46269,-96005,97630,25009,-2951,-67811,99801,81587,-79793,-18547,-83086,69512,33127,-92145,-88497,47703,59527,1909,88785,-88882,69188,-46131,-5589,-15086,36255,-53238,-33009,82664,53901,35939,-42946,-25571,33298,69291,53199,74746,-40127,-39050,91033,51717,-98048,87240,36172,65453,-94425,-63694,-30027,59004,88660,3649,-20267,-52565,-67321,34037,4320,91515,-56753,60115,27134,68617,-61395,-26503,-98929,-8849,-63318,10709,-16151,61905,-95785,5262,23670,-25277,90206,-19391,45735,37208,-31992,-92450,18516,-90452,-58870,-58602,93383,14333,17994,82411,-54126,-32576,35440,-60526,-78764,-25069,-9022,-394,92186,-38057,55328,-61569,67780,77169,19546,-92664,-94948,44484,-13439,83529,27518,-48333,72998,38342,-90553,-98578,-76906,81515,-16464,78439,92529,35225,-39968,-10130,-7845,-32245,-74955,-74996,67731,-13897,-82493,33407,93619,59560,-24404,-57553,19486,-45341,34098,-24978,-33612,79058,71847,76713,-95422,6421,-96075,-59130,-28976,-16922,-62203,69970,68331,21874,40551,89650,51908,58181,66480,-68177,34323,-3046,-49656,-59758,43564,-10960,-30796,15473,-20216,46085,-85355,41515,-30669,-87498,57711,56067,63199,-83805,62042,91213,-14606,4394,-562,74913,10406,96810,-61595,32564,31640,-9732,42058,98052,-7908,-72330,1558,-80301,34878,32900,3939,-8824,88316,20937,21566,-3218,-66080,-31620,86859,54289,90476,-42889,-15016,-18838,75456,30159,-67101,42328,-92703,85850,-5475,23470,-80806,68206,17764,88235,46421,-41578,74005,-81142,80545,20868,-1560,64017,83784,68863,-97516,-13016,-72223,79630,-55692,82255,88467,28007,-34686,-69049,-41677,88535,-8217,68060,-51280,28971,49088,49235,26905,-81117,-44888,40623,74337,-24662,97476,79542,-72082,-35093,98175,-61761,-68169,59697,-62542,-72965,59883,-64026,-37656,-92392,-12113,-73495,98258,68379,-21545,64607,-70957,-92254,-97460,-63436,-8853,-19357,-51965,-76582,12687,-49712,45413,-60043,33496,31539,-57347,41837,67280,-68813,52088,-13155,-86430,-15239,-45030,96041,18749,-23992,46048,35243,-79450,85425,-58524,88781,-39454,53073,-48864,-82289,39086,82540,-11555,25014,-5431,-39585,-89526,2705,31953,-81611,36985,-56022,68684,-27101,11422,64655,-26965,-63081,-13840,-91003,-78147,-8966,41488,1988,99021,-61575,-47060,65260,-23844,-21781,-91865,-19607,44808,2890,63692,-88663,-58272,15970,-65195,-45416,-48444,-78226,-65332,-24568,42833,-1806,-71595,80002,-52250,30952,48452,-90106,31015,-22073,62339,63318,78391,28699,77900,-4026,-76870,-45943,33665,9174,-84360,-22684,-16832,-67949,-38077,-38987,-32847,51443,-53580,-13505,9344,-92337,26585,70458,-52764,-67471,-68411,-1119,-2072,-93476,67981,40887,-89304,-12235,41488,1454,5355,-34855,-72080,24514,-58305,3340,34331,8731,77451,-64983,-57876,82874,62481,-32754,-39902,22451,-79095,-23904,78409,-7418,77916]`),
		// 	Expected: JsonToSliceSlice[int](`[]`), // 答案不正確，用於測試超時
		// },
		// {
		// 	Arg1:     JsonToSlice[int](`[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]`),
		// 	Expected: JsonToSliceSlice[int](`[0,0,0]`), // 用於測試記憶體
		// },
	}

	Decorate(func() {
		for _, d := range data {
			for i := range d.Expected {
				sort.Ints(d.Expected[i])
			}
			actual := Q153Sum(d.Arg1)
			for i := range actual {
				sort.Ints(actual[i])
			}
			assert.Equal(t, d.Expected, actual, d)
		}
	}, 1, 0)
}

// TestQ1214TwoSumBSTs Medium
func TestQ1214TwoSumBSTs(t *testing.T) {
	type param struct {
		Arg1     *TreeNode
		Arg2     *TreeNode
		Arg3     int
		Expected bool
	}
	data := []param{
		{
			Arg1:     JSONArrayToTreeNode(`[2,1,4]`),
			Arg2:     JSONArrayToTreeNode(`[1,0,3]`),
			Arg3:     5,
			Expected: true,
		},
		{
			Arg1:     JSONArrayToTreeNode(`[0,-10,10]`),
			Arg2:     JSONArrayToTreeNode(`[5,1,7,0,2]`),
			Arg3:     18,
			Expected: false,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q1214TwoSumBSTs(d.Arg1, d.Arg2, d.Arg3), d)
		}
	}, 1, 0)

}

// TestQ653TwoSumIV-InputIsABST Easy
func TestQ653TwoSumIV(t *testing.T) {
	type param struct {
		Arg1     *TreeNode
		Arg2     int
		Expected bool
	}
	data := []param{
		{
			Arg1:     JSONArrayToTreeNode(`[5,3,6,2,4,null,7]`),
			Arg2:     9,
			Expected: true,
		},
		{
			Arg1:     JSONArrayToTreeNode(`[5,3,6,2,4,null,7]`),
			Arg2:     28,
			Expected: false,
		},
		{
			Arg1:     JSONArrayToTreeNode(`[2,1,3]`),
			Arg2:     4,
			Expected: true,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q653TwoSumIV(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)

}

// TestQ1099TwoSumLessThanK Easy
func TestQ1099TwoSumLessThanK(t *testing.T) {
	type param struct {
		Arg1     []interface{}
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonStringToSliceAny(`[34,23,1,24,75,33,54,8]`),
			Arg2:     60,
			Expected: 58,
		},
		{
			Arg1:     JsonStringToSliceAny(`[10,20,30]`),
			Arg2:     15,
			Expected: -1,
		},
		{
			Arg1:     JsonStringToSliceAny(`[254,914,110,900,147,441,209,122,571,942,136,350,160,127,178,839,201,386,462,45,735,467,153,415,875,282,204,534,639,994,284,320,865,468,1,838,275,370,295,574,309,268,415,385,786,62,359,78,854,944]`),
			Arg2:     200,
			Expected: 198,
		},
	}

	Decorate(func() {
		for _, d := range data {
			var param1 []int
			Arg1 := ConvertSlice[float64](d.Arg1)
			param1 = make([]int, 0, len(Arg1))
			for _, v := range Arg1 {
				param1 = append(param1, int(v))
			}

			assert.Equal(t, d.Expected, Q1099TwoSumLessThanK(param1, d.Arg2), d)
		}
	}, 1, 0)

}

// TestQ184Sum Medium
func TestQ184Sum(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Expected [][]int
	}
	data := []param{
		{
			Arg1:     JsonToSlice[int](`[1,0,-1,0,-2,2]`),
			Arg2:     0,
			Expected: JsonToSliceSlice[int](`[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]`),
		},
		{
			Arg1:     JsonToSlice[int](`[2,2,2,2,2]`),
			Arg2:     8,
			Expected: JsonToSliceSlice[int](`[[2,2,2,2]]`),
		},
		{
			Arg1:     JsonToSlice[int](`[0,0,0,0]`),
			Arg2:     0,
			Expected: JsonToSliceSlice[int](`[[0,0,0,0]]`),
		},
		{
			Arg1:     JsonToSlice[int](`[-1,0,-5,-2,-2,-4,0,1,-2]`),
			Arg2:     -9,
			Expected: JsonToSliceSlice[int](`[[-5,-4,-1,1],[-5,-4,0,0],[-5,-2,-2,0],[-4,-2,-2,-1]]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			for i := range d.Expected {
				sort.Ints(d.Expected[i])
			}
			actual := Q184Sum(d.Arg1, d.Arg2)
			for i := range actual {
				sort.Ints(actual[i])
			}
			assert.Equal(t, d.Expected, actual, d)
		}
	}, 1, 0)

}
func TestQ170TwoSumIIIDataStructureDesign(t *testing.T) {
	type param struct {
		Arg1     []interface{}
		Arg2     []interface{}
		Expected []interface{}
	}
	data := []param{
		{
			Arg1:     JsonStringToSliceAny(`["TwoSum", "add", "add", "add", "find", "find"]`),
			Arg2:     JsonStringToSliceAny(`[[], [1], [3], [5], [4], [7]]`),
			Expected: JsonStringToSliceAny(`[null, null, null, null, true, false]`),
		},
	}

	Decorate(func() {
		for _, d := range data {
			var obj q170twosumiiidatastructuredesign.TwoSum
			var actual interface{}
			for idx, action := range d.Arg1 {
				switch action {
				case "TwoSum":
					obj = q170twosumiiidatastructuredesign.Constructor()
					actual = nil
				case "add":
					arg, _ := d.Arg2[idx].([]interface{})
					Arg2 := ConvertSlice[float64](arg)
					obj.Add(int(Arg2[0]))
					actual = nil
				case "find":
					arg, _ := d.Arg2[idx].([]interface{})
					Arg2 := ConvertSlice[float64](arg)
					actual = obj.Find(int(Arg2[0]))
				}
				assert.Equal(t, d.Expected[idx], actual, "d")
			}
		}

	}, 1, 0)

}

// TestQ7ReverseInteger Medium
func TestQ7ReverseInteger(t *testing.T) {
	type param struct {
		Arg1     int
		Expected int
	}
	data := []param{
		{
			Arg1:     123,
			Expected: 321,
		},
		{
			Arg1:     -123,
			Expected: -321,
		},
		{
			Arg1:     120,
			Expected: 21,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q7ReverseInteger(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ1531StringCompressionII Hard *fail*
func TestQ1531StringCompressionII(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     "aaaabbab",
			Arg2:     3,
			Expected: 2,
		},
		// {
		// 	Arg1:     "aaabcccd",
		// 	Arg2:     2,
		// 	Expected: 4,
		// },
		// {
		// 	Arg1:     "aabbaa",
		// 	Arg2:     2,
		// 	Expected: 2,
		// },
		// {
		// 	Arg1:     "aaaaaaaaaaa",
		// 	Arg2:     0,
		// 	Expected: 3,
		// },
		// {
		// 	Arg1:     "abc",
		// 	Arg2:     0,
		// 	Expected: 3,
		// },
		// {
		// 	Arg1:     "abbbbbbbbbba",
		// 	Arg2:     2,
		// 	Expected: 3,
		// },
		// {
		// 	Arg1:     "llllllllllttttttttt",
		// 	Arg2:     1,
		// 	Expected: 4,
		// },
		// {
		// 	Arg1:     "bababbaba",
		// 	Arg2:     1,
		// 	Expected: 7,
		// },
		// {
		// 	Arg1:     "bbabbbabbbbcbb",
		// 	Arg2:     4,
		// 	Expected: 3,
		// },
		// {
		// 	Arg1:     "baaaaa",
		// 	Arg2:     2,
		// 	Expected: 2,
		// },
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q1531StringCompressionII(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)

}

// TestQ331VerifyPreorderSerializationOfABinaryTree Medium
func TestQ331VerifyPreorderSerializationOfABinaryTree(t *testing.T) {
	type param struct {
		Arg1     string
		Expected bool
	}
	data := []param{
		{
			Arg1:     "9,3,4,#,#,1,#,#,2,#,6,#,#",
			Expected: true,
		},
		{
			Arg1:     "1,#",
			Expected: false,
		},
		{
			Arg1:     "9,#,#,1",
			Expected: false,
		},
		{
			Arg1:     "9,#,92,#,#",
			Expected: true,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q331VerifyPreorderSerializationOfABinaryTree(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ27RemoveElement Easy
func TestQ27RemoveElement(t *testing.T) {
	type param struct {
		Arg1     []interface{}
		Arg2     int
		Expected int
	}
	data := []param{
		{
			Arg1:     JsonStringToSliceAny(`[3,2,2,3]`),
			Arg2:     3,
			Expected: 2,
		},
		{
			Arg1:     JsonStringToSliceAny(`[0,1,2,2,3,0,4,2]`),
			Arg2:     2,
			Expected: 5,
		},
	}

	Decorate(func() {
		for _, d := range data {
			var param1 []int
			Arg1 := ConvertSlice[float64](d.Arg1)
			param1 = make([]int, 0, len(Arg1))
			for _, v := range Arg1 {
				param1 = append(param1, int(v))
			}

			assert.Equal(t, d.Expected, Q27RemoveElement(param1, d.Arg2), d)
		}
	}, 1, 0)

}

// TestQ2231LargestNumberAfterDigitSwapsByParity Easy
func TestQ2231LargestNumberAfterDigitSwapsByParity(t *testing.T) {
	type param struct {
		Arg1     int
		Expected int
	}
	data := []param{
		{1234, 3412},
		{65875, 87655},
		{247, 427},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q2231LargestNumberAfterDigitSwapsByParity(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ4MedianOfTwoSortedArrays Hard
func TestQ4MedianOfTwoSortedArrays(t *testing.T) {

	type param struct {
		Arg1     []interface{}
		Arg2     []interface{}
		Expected float64
	}
	data := []param{
		{
			Arg1:     JsonStringToSliceAny(`[1,3]`),
			Arg2:     JsonStringToSliceAny(`[2]`),
			Expected: 2.00000,
		},
		{
			Arg1:     JsonStringToSliceAny(`[1,2]`),
			Arg2:     JsonStringToSliceAny(`[3,4]`),
			Expected: 2.50000,
		},
	}

	Decorate(func() {
		for _, d := range data {
			var param1 []int
			Arg1 := ConvertSlice[float64](d.Arg1)
			param1 = make([]int, 0, len(Arg1))
			for _, v := range Arg1 {
				param1 = append(param1, int(v))
			}

			var param2 []int
			Arg2 := ConvertSlice[float64](d.Arg2)
			param2 = make([]int, 0, len(Arg2))
			for _, v := range Arg2 {
				param2 = append(param2, int(v))
			}
			assert.Equal(t, d.Expected, Q4MedianOfTwoSortedArrays(param1, param2), d)
		}
	}, 1, 0)

}

// TestQ8StringToInteger Medium
func TestQ8StringToInteger(t *testing.T) {

	type param struct {
		Arg1     string
		Expected int
	}
	data := []param{
		{"", 0},
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"3.14159", 3},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q8StringToInteger(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ14LongestCommonPrefix Easy
func TestQ14LongestCommonPrefix(t *testing.T) {

	type param struct {
		Arg1     []any
		Expected string
	}
	data := []param{
		{
			Arg1:     JsonStringToSliceAny(`["flower","flow","flight"]`),
			Expected: "fl",
		},
		{
			Arg1:     JsonStringToSliceAny(`["dog","racecar","car"]`),
			Expected: "",
		},
		{
			Arg1:     JsonStringToSliceAny(`["a","aaa","aa"]`),
			Expected: "a",
		},
		{
			Arg1:     JsonStringToSliceAny(`["ab","a"]`),
			Expected: "a",
		},
		{
			Arg1:     JsonStringToSliceAny(`["baab","bacb","b","cbc"]`),
			Expected: "",
		},
		{
			Arg1:     JsonStringToSliceAny(`["abca","aba","aaab"]`),
			Expected: "a",
		},
	}

	Decorate(func() {
		for _, d := range data {
			param1 := ConvertSlice[string](d.Arg1)
			assert.Equal(t, d.Expected, Q14LongestCommonPrefix(param1), d)
		}
	}, 1, 0)

}

// TestQ13RomanToInteger Easy
func TestQ13RomanToInteger(t *testing.T) {

	type param struct {
		Arg1     string
		Expected int
	}
	data := []param{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q13RomanToInteger(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ564FindtheClosestPalindrome Hard
func TestQ564FindTheClosestPalindrome(t *testing.T) {

	type param struct {
		Arg1     string
		Expected string
	}
	data := []param{
		{"1234", "1221"},
		{"123", "121"},
		{"1", "0"},
		{"100", "99"},
		{"1000", "999"},
		{"11", "9"},
		{"99", "101"},
		{"9", "8"},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q564FindTheClosestPalindrome(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ2217FindPalindromeWithFixedLength Medium
func TestQ2217FindPalindromeWithFixedLength(t *testing.T) {
	type param struct {
		Arg1     []interface{}
		Arg2     int
		Expected []interface{}
	}
	data := []param{
		// {
		// 	Arg1:     JsonStringToSliceAny(`[1,2,3,4,5,90]`),
		// 	Arg2:     3,
		// 	Expected: JsonStringToSliceAny(`[101,111,121,131,141,999]`),
		// },
		// {
		// 	Arg1:     JsonStringToSliceAny(`[2,4,6]`),
		// 	Arg2:     4,
		// 	Expected: JsonStringToSliceAny(`[1111,1331,1551]`),
		// },
		// {
		// 	Arg1:     JsonStringToSliceAny(`[2,201429812,8,520498110,492711727,339882032,462074369,9,7,6]`),
		// 	Arg2:     1,
		// 	Expected: JsonStringToSliceAny(`[2,-1,8,-1,-1,-1,-1,9,7,6]`),
		// },
		// {
		// 	Arg1:     JsonStringToSliceAny(`[928053739,72,86059860,90,622074924,831263840,8,10,43,13,54,870906009,64]`),
		// 	Arg2:     3,
		// 	Expected: JsonStringToSliceAny(`[-1,818,-1,999,-1,-1,171,191,525,222,636,-1,737]`),
		// },
	}

	Decorate(func() {
		for _, d := range data {

			c := ConvertSlice[float64](d.Expected) // 還原unmarshal
			var expected []int64
			expected = make([]int64, 0, len(c))
			for _, v := range c {
				expected = append(expected, int64(v))
			}

			Arg1 := ConvertSlice[float64](d.Arg1) // 還原unmarshal
			var param1 []int
			param1 = make([]int, 0, len(Arg1))
			for _, v := range Arg1 {
				param1 = append(param1, int(v))
			}

			assert.Equal(t, expected, Q2217FindPalindromeWithFixedLength(param1, d.Arg2), d)
		}
	}, 1, 0)
}
func TestQ9PalindromeNumber(t *testing.T) {

	type param struct {
		Arg1     int
		Expected bool
	}
	data := []param{
		{121, true},
		{-121, false},
		{10, false},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q9PalindromeNumber(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ2353DesignAFoodRatingSystem Medium #heap #maxHeap
func TestQ2353DesignAFoodRatingSystem(t *testing.T) {
	type param struct {
		Arg1     []interface{}
		Arg2     []interface{}
		Expected []interface{}
	}
	// var slice []interface{}
	// fmt.Println("ori:", []byte("[null, 1, 2, null, 2, 3, 4, 5, null]"))
	// json.Unmarshal([]byte("[null, 1, 2, null, 2, 3, 4, 5, null]"), &slice)
	// fmt.Println("rst:", slice[0] == nil)
	data := []param{
		// {
		// 	Arg1:     JsonStringToSliceAny(`["FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"]`),
		// 	Arg2:     JsonStringToSliceAny(`[[["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]], ["korean"], ["japanese"], ["sushi", 16], ["japanese"], ["ramen", 16], ["japanese"]]`),
		// 	Expected: JsonStringToSliceAny(`[null, "kimchi", "ramen", null, "sushi", null, "ramen"]`),
		// },
		{
			Arg1:     JsonStringToSliceAny(`["FoodRatings","highestRated","changeRating","changeRating","changeRating","highestRated","highestRated","highestRated","highestRated","changeRating","changeRating","changeRating","changeRating"]`),
			Arg2:     JsonStringToSliceAny(`[[["ixoldpvcl","bmdzu","zmazdit","wdz","yxsoc","jyxxdmeqpy","hxvyjar","jktdotax","kgdct","kxuhddwif"],["uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc"],[5,9,4,6,8,6,17,9,11,4]],["uudduznsjc"],["jyxxdmeqpy",3],["hxvyjar",19],["bmdzu",12],["uudduznsjc"],["uudduznsjc"],["uudduznsjc"],["uudduznsjc"],["hxvyjar",10],["yxsoc",6],["hxvyjar",14],["yxsoc",2]]`),
			Expected: JsonStringToSliceAny(`[null,"hxvyjar",null,null,null,"hxvyjar","hxvyjar","hxvyjar","hxvyjar",null,null,null,null]`),
		},
	}
	Decorate(func() {
		for _, d := range data {
			var obj q2353designafoodratingsystem.FoodRatings
			var actual interface{}
			for idx, action := range d.Arg1 {
				switch action {
				case "FoodRatings":
					arg, _ := d.Arg2[idx].([]interface{})
					// fmt.Println()
					c := ConvertSlice[float64](arg[2].([]interface{})) // 還原unmarshal
					var param3 []int
					param3 = make([]int, 0, len(c))
					for _, v := range c {
						// fmt.Println(v)
						param3 = append(param3, int(v))
					}
					obj = q2353designafoodratingsystem.Constructor(ConvertSlice[string](arg[0].([]interface{})), ConvertSlice[string](arg[1].([]interface{})), param3)
					actual = nil
				case "highestRated":
					arg, _ := d.Arg2[idx].([]interface{})
					// fmt.Println(arg[0].(string))
					actual = obj.HighestRated(arg[0].(string))
				case "changeRating":
					arg, _ := d.Arg2[idx].([]interface{})
					// fmt.Println(arg[1])
					obj.ChangeRating(arg[0].(string), int(arg[1].(float64)))
					actual = nil
				}
				assert.Equal(t, d.Expected[idx], actual, "tttttttttttt")
			}
		}
	}, 1, 0)
}
func TestQ1636SortArrayByIncreasingFrequency(t *testing.T) {
	type param struct {
		Arg1     []int
		Expected []int
	}
	data := []param{
		{[]int{1, 1, 2, 2, 2, 3}, []int{3, 1, 1, 2, 2, 2}},
		{[]int{2, 3, 1, 3, 2}, []int{1, 3, 3, 2, 2}},
		{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Expected, Q1636SortArrayByIncreasingFrequency(d.Arg1), d)
		}
	}, 1, 0)
}
func TestQ595BigCountries(t *testing.T) {
	// skip
}
func TestQ237DeleteNodeInALinkedList(t *testing.T) {
	// skip
}

func TestQ2349DesignANumberContainerSystem(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     []int
		Expected interface{}
	}
	data := []param{
		// {"NumberContainers", []int{}, nil},
		// 1
		// {"find", []int{10}, -1},
		// {"change", []int{2, 10}, nil},
		// {"change", []int{1, 10}, nil},
		// {"change", []int{3, 10}, nil},
		// {"change", []int{5, 10}, nil},
		// {"find", []int{10}, 1},
		// {"change", []int{1, 20}, nil},
		// {"find", []int{10}, 2},

		// 2
		{"change", []int{1, 10}, nil},
		{"find", []int{10}, 1},
		{"change", []int{1, 20}, nil},
		{"find", []int{10}, -1},
		{"find", []int{20}, 1},
		{"find", []int{30}, -1},
	}

	Decorate(func() {
		obj := q2349designanumbercontainersystem.Constructor()
		for _, d := range data {
			// fmt.Println(d, obj)
			switch d.Arg1 {
			case "find":
				assert.Equal(t, d.Expected, obj.Find(d.Arg2[0]))
			case "change":
				obj.Change(d.Arg2[0], d.Arg2[1])
				assert.Equal(t, d.Expected, nil)
			}
		}
	}, 1, 0)
}

// TestQ167TwoSumIIInputArrayIsSorted Easy
func TestQ167TwoSumIIInputArrayIsSorted(t *testing.T) {
	type param struct {
		Arg1     []interface{}
		Arg2     int
		Expected []int
	}
	data := []param{
		{
			Arg1:     JsonStringToSliceAny(`[2,7,11,15]`),
			Arg2:     9,
			Expected: []int{1, 2},
		},
		{
			Arg1:     JsonStringToSliceAny(`[2,3,4]`),
			Arg2:     6,
			Expected: []int{1, 3},
		},
		{
			Arg1:     JsonStringToSliceAny(`[-1,0]`),
			Arg2:     -1,
			Expected: []int{1, 2},
		},
	}

	Decorate(func() {
		for _, d := range data {
			var param1 []int
			Arg1 := ConvertSlice[float64](d.Arg1)
			param1 = make([]int, 0, len(Arg1))
			for _, v := range Arg1 {
				param1 = append(param1, int(v))
			}

			assert.Equal(t, d.Expected, Q167TwoSumIIInputArrayIsSorted(param1, d.Arg2), d)
		}
	}, 1, 0)

}

// TestQ1TwoSum Easy
func TestQ1TwoSum(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Expected []int
	}
	data := []param{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	Decorate(func() {
		for _, d := range data {
			nums := d.Arg1
			target := d.Arg2
			actual := Q1TwoSum(nums, target)
			excepted := d.Expected
			assert.Equal(t, excepted, actual, nums)
		}
	}, 3000, 10000)
}

func TestQ2180CountIntegersWithEvenDigitSum(t *testing.T) {
	type param struct {
		Arg1     int
		Expected int
	}
	data := []param{
		{4, 2},
		{30, 14},
	}

	Decorate(func() {
		for _, d := range data {
			actual := Q2180CountIntegersWithEvenDigitSum(d.Arg1)
			excepted := d.Expected
			assert.Equal(t, excepted, actual, d)
		}
	}, 1, 0)
}
