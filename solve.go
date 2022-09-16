package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	q237deletenodeinalinkedlist "github.com/limiu82214/GoLeetcode/exam/q237deletenodeinalinkedlist"
)

func main() {
	fmt.Println("hello world")
}

func Q564FindTheClosestPalindrome(n string) string {
	Abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	ReverseString := func(s string) string {
		ans := ""
		for _, v := range s {
			ans = string(v) + ans
		}
		return ans
	}
	halfIdx := len(n)/2 + len(n)%2
	isOdd := len(n)%2 != 0
	halfs := n[0:halfIdx]
	halfn, _ := strconv.Atoi(halfs)
	halfnAdd := halfn + 1
	halfnMinus := halfn - 1
	halfsLength := len(halfs)
	ca := strconv.Itoa(halfnAdd)
	cn := halfs
	cm := strconv.Itoa(halfnMinus)
	if isOdd {
		if len(ca) > halfsLength && halfsLength > 1 {
			ca += strings.Repeat("0", halfsLength-2) + "1"
		} else {
			ca += ReverseString(ca[:halfsLength-1])
		}
		cn += ReverseString(cn[:halfsLength-1])
		if len(cm) >= halfsLength {
			cm += ReverseString(cm[:halfsLength-1])
		} else {
			cm += strings.Repeat("9", halfsLength-1)
		}
	} else {
		if len(ca) > halfsLength {
			ca += strings.Repeat("0", halfsLength-1) + "1"
		} else {
			ca += ReverseString(ca)
		}
		cn += ReverseString(cn)
		if len(cm) >= halfsLength {
			if cm == "0" {
				cm += "9"
			} else {
				cm += ReverseString(cm[:halfsLength])
			}
		} else {
			cm += strings.Repeat("9", halfsLength)
		}
	}

	candidateList := []int{}
	var t int
	t, _ = strconv.Atoi(ca)
	candidateList = append(candidateList, t)
	t, _ = strconv.Atoi(cm)
	candidateList = append(candidateList, t)
	if cn != n {
		t, _ = strconv.Atoi(cn)
		candidateList = append(candidateList, t)
	}

	nn, _ := strconv.Atoi(n)
	closetLength := -1
	var ans int
	for _, v := range candidateList {
		if closetLength == -1 || Abs(nn-v) <= closetLength {
			if Abs(nn-v) == closetLength {
				if ans > v {
					ans = v
				}
			} else {
				closetLength = Abs(nn - v)
				ans = v
			}
		}
	}
	// fmt.Println(ca, cn, cm, closetLength, ans)

	return strconv.Itoa(ans)
}

func Q2217FindPalindromeWithFixedLength(queries []int, intLength int) []int64 {
	ans := make([]int64, 0, len(queries))
	solveOne := func(n int, intLength int) int64 {
		switch intLength {
		case 1:
			if n > 9 {
				return -1
			}
			return int64(n)
		case 2:
			if n > 9 {
				return -1
			}
			return int64(n*10 + n)
		}
		n1 := n - 1
		b := make([]byte, intLength)
		midIdx := intLength / 2
		var l, r int
		if intLength%2 == 0 {
			l, r = midIdx-1, midIdx

		} else {
			l, r = midIdx, midIdx
		}
		// n1--, 最中間 and intLength >2  直接拿
		// 最外面 +1
		for l >= 0 {
			if l == 0 {
				b[l] = byte(n1%10+1) + 48
				b[r] = byte(n1%10+1) + 48
				break
			}
			b[l] = byte(n1%10) + 48
			b[r] = byte(n1%10) + 48
			l--
			r++
			n1 /= 10
		}
		if n1 > 8 {
			return -1
		}

		v, _ := strconv.ParseInt(string(b), 10, 64)
		return v
	}
	for _, v := range queries {
		ans = append(ans, solveOne(v, intLength))
	}

	return ans
}

func Q9PalindromeNumber(x int) bool {
	s := strconv.Itoa(x)
	b := []byte(s)
	l := 0
	r := len(b) - 1
	for l < r {
		if b[l] == b[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func Q1636SortArrayByIncreasingFrequency(nums []int) []int {
	type idxCountMap struct {
		numsCount map[int]int
		uniqNum   []int
	}
	var icm = idxCountMap{
		make(map[int]int),
		[]int{},
	}
	for _, v := range nums {
		_, ok := icm.numsCount[v]
		if !ok {
			icm.numsCount[v] = 0
			icm.uniqNum = append(icm.uniqNum, v)
		}
		icm.numsCount[v]++

		sort.Slice(icm.uniqNum, func(i, j int) bool {
			f := icm.numsCount[icm.uniqNum[i]]
			e := icm.numsCount[icm.uniqNum[j]]
			if f < e {
				return true
			} else if f == e {
				if icm.uniqNum[i] > icm.uniqNum[j] {
					return true
				} else {
					return false
				}
			} else {
				return false
			}

		})
	}

	ans := []int{}
	for _, num := range icm.uniqNum {
		for i := 0; i < icm.numsCount[num]; i++ {
			ans = append(ans, num)
		}
	}
	return ans
}

func Q595BigCountries() {
	// # Write your MySQL query statement below
	// SELECT name, population, area FROM World WHERE ( World.area > 2999999 OR World.population > 24999999)
}

func Q237DeleteNodeInALinkedList(node *q237deletenodeinalinkedlist.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
func Q1TwoSum(nums []int, target int) []int {
	// fmt.Println(nums)
	alreadyTryInt := map[int]bool{}
	for i := 0; i < len(nums); i++ {

		if alreadyTryInt[i] == true {
			continue
		}
		for j := (i + 1); j < len(nums); j++ {
			// fmt.Println(nums[i]+nums[j], i, j)
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
		alreadyTryInt[i] = true
	}
	return nil
}

func Q2180CountIntegersWithEvenDigitSum(num int) int {
	isFit := func(num int) bool {
		s := strconv.Itoa(num)
		b := strings.Split(s, "")
		sum := 0
		for _, v := range b {
			sum = sum + int([]byte(v)[0]-48)
		}
		return sum%2 == 0
	}

	sum := 0
	for i := 1; i <= num; i++ {
		if isFit(i) == true {
			sum++
		}
	}
	return sum
}
