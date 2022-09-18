package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	q237deletenodeinalinkedlist "github.com/limiu82214/GoLeetcode/exam/q237deletenodeinalinkedlist"
)

func main() {
	fmt.Println("hello world")
}

func Q1531StringCompressionII(s string, k int) int {
	// 目前的方法是先找出刪除最優解，再依序刪除
	// 遇到的問題: aaaabbab, 3
	// 問題描述: 依照目前邏輯順序
	// 	aaaabbab => a4b2ab, 3
	// 	del a => a4b3, 2
	// 	k 剩2不足以讓長度變短，於是跳出
	// 	// ---------------
	// 	但其實可以刪除b2，使結果變成a5b，再刪除b變成a2
	b := []byte(s)

	type BESTKILL int
	const (
		NONE BESTKILL = iota
		ONEKILL
		FULLKILL
	)
	type node struct {
		letter       byte
		n            int
		length       int
		bestKillWay  BESTKILL // fullKill or oneKill
		bestKillCost int
	}
	CalLength := func(n int) int {
		if n == 1 {
			return 1
		}
		l := 1
		for n > 0 {
			n /= 10
			l++
		}
		return l
	}

	nlist := []*node{}

	// forr b put to node
	var tmpNode *node
	for i, v := range b {
		if i == 0 || (i-1 >= 0 && b[i] != b[i-1]) {
			tmpNode = &node{
				letter:       v,
				n:            1,
				length:       0,
				bestKillWay:  NONE,
				bestKillCost: -1,
			}
			tmpNode.length = CalLength(tmpNode.n)
			nlist = append(nlist, tmpNode)
		} else {
			tmpNode.n++
			tmpNode.length = CalLength(tmpNode.n)
		}
	}

	// do it if k > 0
	for k > 0 {
		// // calKillCost keep best one
		biggestCP := -1
		biggestCPIdx := -1
		for i, n := range nlist {
			nlist[i].bestKillWay = NONE
			nlist[i].bestKillCost = -1
			oneKillCP := 0
			if n.n >= 10 { // oneKill
				oneLengthKillCost := n.n%10 + 1
				// cp = 1/oneLengthKillCost // 1 because cost 1 length
				oneKillCP = 1 / oneLengthKillCost
				if k >= oneLengthKillCost {
					nlist[i].bestKillWay = ONEKILL
					nlist[i].bestKillCost = oneLengthKillCost
				}
			}

			bestCP := oneKillCP
			fullKillCP := 0
			{ // fullKill
				fullKillCost := n.n

				// ori left + n + right 's length
				leftLen := 0
				if i-1 < 0 {
					leftLen = 0
				} else {
					leftLen = nlist[i-1].length
				}
				rightLen := 0
				if i >= (len(nlist) - 1) {
					rightLen = 0
				} else {
					rightLen = nlist[i+1].length
				}
				originLen := leftLen + n.length + rightLen
				// after left + right 's Length (notice if letter same must be join)
				afterLen := -1
				if leftLen != 0 && rightLen != 0 { // both side has num
					if nlist[i-1].letter != nlist[i+1].letter {
						afterLen = nlist[i-1].length + nlist[i+1].length
					} else {
						afterLen = CalLength(nlist[i-1].length + nlist[i+1].length)
					}
				} else if leftLen == 0 && rightLen != 0 {
					afterLen = nlist[i+1].length
				} else if leftLen != 0 && rightLen == 0 {
					afterLen = nlist[i-1].length
				}

				// cp = after - ori / fullKillCost
				if afterLen < 0 { // no afterLen exist
					fullKillCP = originLen / fullKillCost
				} else {
					fullKillCP = (originLen - afterLen) / fullKillCost
				}

				if k >= fullKillCost {
					if fullKillCP >= oneKillCP {
						bestCP = fullKillCP
						nlist[i].bestKillWay = FULLKILL
						nlist[i].bestKillCost = fullKillCost
					}
				}
			}

			// ----------

			// save biggest cp
			if k >= nlist[i].bestKillCost && nlist[i].bestKillWay != NONE {
				if biggestCP == -1 {
					biggestCP = bestCP
					biggestCPIdx = i
				} else if bestCP >= biggestCP && nlist[i].bestKillWay >= nlist[biggestCPIdx].bestKillWay {
					biggestCP = bestCP
					biggestCPIdx = i
				}
			}
		}

		//kill best cp, and cost k
		tmpNode = nil
		if biggestCPIdx != -1 {
			k -= nlist[biggestCPIdx].bestKillCost
			switch nlist[biggestCPIdx].bestKillWay {
			case FULLKILL:
				if biggestCPIdx-1 >= 0 && biggestCPIdx+1 < len(nlist) { // both
					if nlist[biggestCPIdx-1].letter == nlist[biggestCPIdx+1].letter {
						nlist[biggestCPIdx-1].n = nlist[biggestCPIdx-1].n + nlist[biggestCPIdx+1].n
						nlist[biggestCPIdx-1].length = CalLength(nlist[biggestCPIdx-1].n)
						nlist = append(nlist[:biggestCPIdx], nlist[biggestCPIdx+2:]...)
					} else {
						nlist = append(nlist[:biggestCPIdx], nlist[biggestCPIdx+1:]...)
					}
				} else if biggestCPIdx-1 >= 0 && biggestCPIdx+1 >= len(nlist) { // only left
					nlist = append(nlist[:biggestCPIdx], nlist[biggestCPIdx+1:]...)
				} else if biggestCPIdx-1 < 0 && biggestCPIdx+1 < len(nlist) { // only right
					nlist = append(nlist[:biggestCPIdx], nlist[biggestCPIdx+1:]...)
				} else { //kill will be empty
					return 0
				}

			case ONEKILL:
				nlist[biggestCPIdx].n -= nlist[biggestCPIdx].bestKillCost
				nlist[biggestCPIdx].length = CalLength(nlist[biggestCPIdx].n)
			}
		} else {
			break // cannot find
		}

	}
	ans := 0
	for _, n := range nlist {
		ans += n.length
	}
	return ans
}
func Q331VerifyPreorderSerializationOfABinaryTree(preorder string) bool {
	/* readable but slow
	heap := 1
	q := strings.Split(preorder, ",")
	for _, v := range q {
		heap--
		if heap < 0 {
			return false
		}
		if v != "#" {
			heap += 2
		}
	}
	return heap == 0
	*/
	heap := 1
	q := []byte(preorder)
	l := len(q)
	i := 0
	for i < l {
		heap--
		if heap < 0 {
			return false
		}
		if q[i] != '#' {
			for i < l && q[i] != ',' {
				i++
			}
			heap += 2
		} else {
			i++
		}
		i++
	}
	return heap == 0
}

func Q27RemoveElement(nums []int, val int) int {
	k := 0
	i := 0
	l := len(nums)
	for i < l {
		if nums[i] == val {
			nums[i] = -1
		} else {
			nums[k] = nums[i]
			k++
		}
		i++
	}
	return k
}

func Q4MedianOfTwoSortedArrays(nums1 []int, nums2 []int) float64 {
	q := []int{}
	n1Idx := 0
	n1Len := len(nums1)
	n2Idx := 0
	n2Len := len(nums2)

	for n1Idx < n1Len || n2Idx < n2Len { // 若還沒取完
		if n1Idx == n1Len {
			q = append(q, nums2[n2Idx])
			n2Idx++
			continue
		}
		if n2Idx == n2Len {
			q = append(q, nums1[n1Idx])
			n1Idx++
			continue
		}

		if nums1[n1Idx] <= nums2[n2Idx] {
			q = append(q, nums1[n1Idx])
			n1Idx++
		} else {
			q = append(q, nums2[n2Idx])
			n2Idx++
		}
	}

	// -----------------
	l := len(q)
	isEven := l%2 == 0
	var mid float64
	if isEven { // even get len(q)/ 2  /2-1  ) /2
		mid = (float64(q[l/2]) + float64(q[l/2-1])) / 2
	} else { // odd get len(q)/ 2
		mid = float64(q[l/2])
	}

	return mid
}

func Q2231LargestNumberAfterDigitSwapsByParity(num int) int {
	b := []byte(strconv.Itoa(num))
	odds := []byte{}
	evens := []byte{}
	sequence := []byte{} // e, o will be appended
	for _, c := range b {
		if (c-48)%2 == 0 { // even
			sequence = append(sequence, 'e')
			evens = append(evens, c)
		} else {
			sequence = append(sequence, 'o')
			odds = append(odds, c)
		}
	}

	sort.Slice(odds, func(i, j int) bool {
		return odds[i] > odds[j]
	})
	sort.Slice(evens, func(i, j int) bool {
		return evens[i] > evens[j]
	})

	ans := []byte{}
	for _, v := range sequence {
		switch v {
		case 'e':
			// pop evens
			c := evens[0:1][0]
			evens = evens[1:]
			ans = append(ans, c)
		case 'o':
			// pop odds
			c := odds[0:1][0]
			odds = odds[1:]
			ans = append(ans, c)
		}
	}

	t, _ := strconv.Atoi(string(ans))
	return t
}

func Q175CombineTwoTables() {
	// # Write your MySQL query statement below
	// SELECT P.firstName, P.lastName, A.city, A.state
	// FROM Person AS P LEFT JOIN Address AS A
	// ON P.personId = A.personId
}

func Q8StringToInteger(s string) int {
	re := regexp.MustCompile(`^[ ]*((?:[+-])?(?:(?:\d+[.]\d+)|(?:\d))*)`)
	t := re.FindStringSubmatch(s)[1]
	if t == "" {
		return 0
	}
	// ans, _ := strconv.Atoi(t)
	ans, _ := strconv.ParseFloat(t, 64)
	if ans > (2<<30 - 1) {
		return 2<<30 - 1
	}
	if ans < -2<<30 {
		return -2 << 30
	}
	return int(ans)
}

func Q14LongestCommonPrefix(strs []string) string {
	prefix := []byte(strs[0])
	// get shortest
	var prefixIdx int
	for i, s := range strs {
		if len(s) < len(prefix) {
			prefix = []byte(s)
			prefixIdx = i
		}
	}
	l := len(prefix)

	// nil
	if string(prefix) == "" {
		return ""
	}

	for i, s := range strs {
		if i == prefixIdx {
			continue
		}
		for j, c := range s {
			if j >= l {
				break
			}
			if prefix[j] != byte(c) {
				prefix = []byte(string(prefix)[:j])
				l = len(prefix)
				break
			}
		}
	}
	return string(prefix)
}

func Q13RomanToInteger(s string) int {
	l := len(s)
	sum := 0
	for i, v := range s {
		switch v {
		case 'I':
			if i+1 < l && (s[i+1] == 'V' || s[i+1] == 'X') {
				sum -= 1
			} else {
				sum += 1
			}
		case 'V':
			sum += 5
		case 'X':
			if i+1 < l && (s[i+1] == 'L' || s[i+1] == 'C') {
				sum -= 10
			} else {
				sum += 10
			}
		case 'L':
			sum += 50
		case 'C':
			// if next is D or M Minus 100
			if i+1 < l && (s[i+1] == 'D' || s[i+1] == 'M') {
				sum -= 100
			} else {
				sum += 100
			}
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		}
	}
	return sum
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
