package main

import (
	"fmt"
	"testing"

	q2349designanumbercontainersystem "github.com/limiu82214/GoLeetcode/exam/Q2349DesignANumberContainerSystem"
	q2353designafoodratingsystem "github.com/limiu82214/GoLeetcode/exam/Q2353DesignAFoodRatingSystem"
	"github.com/stretchr/testify/assert"
)

// TestForCopy Easy ============= 複製一個copy並開始你的表演
func TestForCopy(t *testing.T) {
	type param struct {
		Arg1     []interface{}
		Arg2     []interface{}
		Excepted float64
	}
	data := []param{
		{
			Arg1:     JsonStringToSlice(`[1,3]`),
			Arg2:     JsonStringToSlice(`[2]`),
			Excepted: 2.00000,
		},
		{
			Arg1:     JsonStringToSlice(`[1,2]`),
			Arg2:     JsonStringToSlice(`[3,4]`),
			Excepted: 2.50000,
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
			assert.Equal(t, d.Excepted, Q4MedianOfTwoSortedArrays(param1, param2), d)
		}
	}, 1, 0)

}

// TestQ184Sum Medium
// func TestQ184Sum(t *testing.T) {
// 	type param struct {
// 		Arg1     []interface{}
// 		Arg2     int
// 		Excepted [][]int
// 	}
// 	data := []param{
// 		{
// 			Arg1: JsonStringToSlice(`[1,0,-1,0,-2,2]`),
// 			Arg2: 0,
// 			Excepted: [][]int{
// 				[]int{-2, -1, 1, 2},
// 				[]int{-2, 0, 0, 2},
// 				[]int{-1, 0, 0, 1},
// 			},
// 		},
// 		{
// 			Arg1: JsonStringToSlice(`[2,2,2,2,2]`),
// 			Arg2: 8,
// 			Excepted: [][]int{
// 				[]int{2, 2, 2, 2},
// 			},
// 		},
// 	}

// 	Decorate(func() {
// 		for _, d := range data {
// 			var param1 []int
// 			Arg1 := ConvertSlice[float64](d.Arg1)
// 			param1 = make([]int, 0, len(Arg1))
// 			for _, v := range Arg1 {
// 				param1 = append(param1, int(v))
// 			}

// 			assert.Equal(t, d.Excepted, Q184Sum(param1, d.Arg2), d)
// 		}
// 	}, 1, 0)

// }

// TestQ7ReverseInteger Medium
func TestQ7ReverseInteger(t *testing.T) {
	type param struct {
		Arg1     int
		Excepted int
	}
	data := []param{
		{
			Arg1:     123,
			Excepted: 321,
		},
		{
			Arg1:     -123,
			Excepted: -321,
		},
		{
			Arg1:     120,
			Excepted: 21,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Excepted, Q7ReverseInteger(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ1531StringCompressionII Hard
func TestQ1531StringCompressionII(t *testing.T) {
	type param struct {
		Arg1     string
		Arg2     int
		Excepted int
	}
	data := []param{
		{
			Arg1:     "aaaabbab",
			Arg2:     3,
			Excepted: 2,
		},
		// {
		// 	Arg1:     "aaabcccd",
		// 	Arg2:     2,
		// 	Excepted: 4,
		// },
		// {
		// 	Arg1:     "aabbaa",
		// 	Arg2:     2,
		// 	Excepted: 2,
		// },
		// {
		// 	Arg1:     "aaaaaaaaaaa",
		// 	Arg2:     0,
		// 	Excepted: 3,
		// },
		// {
		// 	Arg1:     "abc",
		// 	Arg2:     0,
		// 	Excepted: 3,
		// },
		// {
		// 	Arg1:     "abbbbbbbbbba",
		// 	Arg2:     2,
		// 	Excepted: 3,
		// },
		// {
		// 	Arg1:     "llllllllllttttttttt",
		// 	Arg2:     1,
		// 	Excepted: 4,
		// },
		// {
		// 	Arg1:     "bababbaba",
		// 	Arg2:     1,
		// 	Excepted: 7,
		// },
		// {
		// 	Arg1:     "bbabbbabbbbcbb",
		// 	Arg2:     4,
		// 	Excepted: 3,
		// },
		// {
		// 	Arg1:     "baaaaa",
		// 	Arg2:     2,
		// 	Excepted: 2,
		// },
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Excepted, Q1531StringCompressionII(d.Arg1, d.Arg2), d)
		}
	}, 1, 0)

}

// TestQ331VerifyPreorderSerializationOfABinaryTree Medium
func TestQ331VerifyPreorderSerializationOfABinaryTree(t *testing.T) {
	type param struct {
		Arg1     string
		Excepted bool
	}
	data := []param{
		{
			Arg1:     "9,3,4,#,#,1,#,#,2,#,6,#,#",
			Excepted: true,
		},
		{
			Arg1:     "1,#",
			Excepted: false,
		},
		{
			Arg1:     "9,#,#,1",
			Excepted: false,
		},
		{
			Arg1:     "9,#,92,#,#",
			Excepted: true,
		},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Excepted, Q331VerifyPreorderSerializationOfABinaryTree(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ27RemoveElement Easy
func TestQ27RemoveElement(t *testing.T) {
	type param struct {
		Arg1     []interface{}
		Arg2     int
		Excepted int
	}
	data := []param{
		{
			Arg1:     JsonStringToSlice(`[3,2,2,3]`),
			Arg2:     3,
			Excepted: 2,
		},
		{
			Arg1:     JsonStringToSlice(`[0,1,2,2,3,0,4,2]`),
			Arg2:     2,
			Excepted: 5,
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

			assert.Equal(t, d.Excepted, Q27RemoveElement(param1, d.Arg2), d)
		}
	}, 1, 0)

}

// TestQ2231LargestNumberAfterDigitSwapsByParity Easy
func TestQ2231LargestNumberAfterDigitSwapsByParity(t *testing.T) {
	type param struct {
		Arg1     int
		Excepted int
	}
	data := []param{
		{1234, 3412},
		{65875, 87655},
		{247, 427},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Excepted, Q2231LargestNumberAfterDigitSwapsByParity(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ4MedianOfTwoSortedArrays Hard
func TestQ4MedianOfTwoSortedArrays(t *testing.T) {

	type param struct {
		Arg1     []interface{}
		Arg2     []interface{}
		Excepted float64
	}
	data := []param{
		{
			Arg1:     JsonStringToSlice(`[1,3]`),
			Arg2:     JsonStringToSlice(`[2]`),
			Excepted: 2.00000,
		},
		{
			Arg1:     JsonStringToSlice(`[1,2]`),
			Arg2:     JsonStringToSlice(`[3,4]`),
			Excepted: 2.50000,
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
			assert.Equal(t, d.Excepted, Q4MedianOfTwoSortedArrays(param1, param2), d)
		}
	}, 1, 0)

}

// TestQ8StringToInteger Medium
func TestQ8StringToInteger(t *testing.T) {

	type param struct {
		Arg1     string
		Excepted int
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
			assert.Equal(t, d.Excepted, Q8StringToInteger(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ14LongestCommonPrefix Easy
func TestQ14LongestCommonPrefix(t *testing.T) {

	type param struct {
		Arg1     []any
		Excepted string
	}
	data := []param{
		{
			Arg1:     JsonStringToSlice(`["flower","flow","flight"]`),
			Excepted: "fl",
		},
		{
			Arg1:     JsonStringToSlice(`["dog","racecar","car"]`),
			Excepted: "",
		},
		{
			Arg1:     JsonStringToSlice(`["a","aaa","aa"]`),
			Excepted: "a",
		},
		{
			Arg1:     JsonStringToSlice(`["ab","a"]`),
			Excepted: "a",
		},
		{
			Arg1:     JsonStringToSlice(`["baab","bacb","b","cbc"]`),
			Excepted: "",
		},
		{
			Arg1:     JsonStringToSlice(`["abca","aba","aaab"]`),
			Excepted: "a",
		},
	}

	Decorate(func() {
		for _, d := range data {
			param1 := ConvertSlice[string](d.Arg1)
			assert.Equal(t, d.Excepted, Q14LongestCommonPrefix(param1), d)
		}
	}, 1, 0)

}

// TestQ13RomanToInteger Easy
func TestQ13RomanToInteger(t *testing.T) {

	type param struct {
		Arg1     string
		Excepted int
	}
	data := []param{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Excepted, Q13RomanToInteger(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ564FindtheClosestPalindrome Hard
func TestQ564FindTheClosestPalindrome(t *testing.T) {

	type param struct {
		Arg1     string
		Excepted string
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
			assert.Equal(t, d.Excepted, Q564FindTheClosestPalindrome(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ2217FindPalindromeWithFixedLength Medium
func TestQ2217FindPalindromeWithFixedLength(t *testing.T) {
	type param struct {
		Arg1     []interface{}
		Arg2     int
		Excepted []interface{}
	}
	data := []param{
		// {
		// 	Arg1:     JsonStringToSlice(`[1,2,3,4,5,90]`),
		// 	Arg2:     3,
		// 	Excepted: JsonStringToSlice(`[101,111,121,131,141,999]`),
		// },
		// {
		// 	Arg1:     JsonStringToSlice(`[2,4,6]`),
		// 	Arg2:     4,
		// 	Excepted: JsonStringToSlice(`[1111,1331,1551]`),
		// },
		// {
		// 	Arg1:     JsonStringToSlice(`[2,201429812,8,520498110,492711727,339882032,462074369,9,7,6]`),
		// 	Arg2:     1,
		// 	Excepted: JsonStringToSlice(`[2,-1,8,-1,-1,-1,-1,9,7,6]`),
		// },
		// {
		// 	Arg1:     JsonStringToSlice(`[928053739,72,86059860,90,622074924,831263840,8,10,43,13,54,870906009,64]`),
		// 	Arg2:     3,
		// 	Excepted: JsonStringToSlice(`[-1,818,-1,999,-1,-1,171,191,525,222,636,-1,737]`),
		// },
	}

	Decorate(func() {
		for _, d := range data {

			c := ConvertSlice[float64](d.Excepted) // 還原unmarshal
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
		Excepted bool
	}
	data := []param{
		{121, true},
		{-121, false},
		{10, false},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Excepted, Q9PalindromeNumber(d.Arg1), d)
		}
	}, 1, 0)

}

// TestQ2353DesignAFoodRatingSystem Medium #heap #maxHeap
func TestQ2353DesignAFoodRatingSystem(t *testing.T) {
	type param struct {
		Arg1     []interface{}
		Arg2     []interface{}
		Excepted []interface{}
	}
	// var slice []interface{}
	// fmt.Println("ori:", []byte("[null, 1, 2, null, 2, 3, 4, 5, null]"))
	// json.Unmarshal([]byte("[null, 1, 2, null, 2, 3, 4, 5, null]"), &slice)
	// fmt.Println("rst:", slice[0] == nil)
	data := []param{
		// {
		// 	Arg1:     JsonStringToSlice(`["FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"]`),
		// 	Arg2:     JsonStringToSlice(`[[["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]], ["korean"], ["japanese"], ["sushi", 16], ["japanese"], ["ramen", 16], ["japanese"]]`),
		// 	Excepted: JsonStringToSlice(`[null, "kimchi", "ramen", null, "sushi", null, "ramen"]`),
		// },
		{
			Arg1:     JsonStringToSlice(`["FoodRatings","highestRated","changeRating","changeRating","changeRating","highestRated","highestRated","highestRated","highestRated","changeRating","changeRating","changeRating","changeRating"]`),
			Arg2:     JsonStringToSlice(`[[["ixoldpvcl","bmdzu","zmazdit","wdz","yxsoc","jyxxdmeqpy","hxvyjar","jktdotax","kgdct","kxuhddwif"],["uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc","uudduznsjc"],[5,9,4,6,8,6,17,9,11,4]],["uudduznsjc"],["jyxxdmeqpy",3],["hxvyjar",19],["bmdzu",12],["uudduznsjc"],["uudduznsjc"],["uudduznsjc"],["uudduznsjc"],["hxvyjar",10],["yxsoc",6],["hxvyjar",14],["yxsoc",2]]`),
			Excepted: JsonStringToSlice(`[null,"hxvyjar",null,null,null,"hxvyjar","hxvyjar","hxvyjar","hxvyjar",null,null,null,null]`),
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
				assert.Equal(t, d.Excepted[idx], actual, "tttttttttttt")
			}
		}
	}, 1, 0)
}
func TestQ1636SortArrayByIncreasingFrequency(t *testing.T) {
	type param struct {
		Arg1     []int
		Excepted []int
	}
	data := []param{
		{[]int{1, 1, 2, 2, 2, 3}, []int{3, 1, 1, 2, 2, 2}},
		{[]int{2, 3, 1, 3, 2}, []int{1, 3, 3, 2, 2}},
		{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
	}

	Decorate(func() {
		for _, d := range data {
			assert.Equal(t, d.Excepted, Q1636SortArrayByIncreasingFrequency(d.Arg1), d)
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
		Excepted interface{}
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
			fmt.Println(d, obj)
			switch d.Arg1 {
			case "find":
				assert.Equal(t, d.Excepted, obj.Find(d.Arg2[0]))
			case "change":
				obj.Change(d.Arg2[0], d.Arg2[1])
				assert.Equal(t, d.Excepted, nil)
			}
		}
	}, 1, 0)
}

func TestQ1TwoSum(t *testing.T) {
	type param struct {
		Arg1     []int
		Arg2     int
		Excepted []int
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
			excepted := d.Excepted
			assert.Equal(t, excepted, actual, nums)
		}
	}, 3000, 10000)
}

func TestQ2180CountIntegersWithEvenDigitSum(t *testing.T) {
	type param struct {
		Arg1     int
		Excepted int
	}
	data := []param{
		{4, 2},
		{30, 14},
	}

	Decorate(func() {
		for _, d := range data {
			actual := Q2180CountIntegersWithEvenDigitSum(d.Arg1)
			excepted := d.Excepted
			assert.Equal(t, excepted, actual, d)
		}
	}, 1, 0)
}
