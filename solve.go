package main

import (
	"bytes"
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

func Q2074ReverseNodesInEvenLengthGroups(head *ListNode) *ListNode {
	// 注意 最後一個Group的節點可能是偶數，也要反轉Q2074ReverseNodesInEvenLengthGroups
	reverseListNode := func(preNode *ListNode, postNode *ListNode) {
		head := preNode.Next
		last := postNode
		for head != postNode {
			t := head.Next
			head.Next = last
			last = head
			head = t
		}
		preNode.Next = last
	}
	nowGroupLen := 1
	calLen := 1
	now := head
	var pre, preG *ListNode
	for now != nil { // 還沒處理最後一組
		if calLen > nowGroupLen { // 計算現在在哪個Group
			nowGroupLen++
			calLen = 1
			preG = pre
		}

		pre = now
		now = now.Next
		if calLen == nowGroupLen { // 走到一個Group的末端拉
			if nowGroupLen%2 == 0 {
				pre = preG.Next
				reverseListNode(preG, now)

			}
		} else if now == nil && calLen%2 == 0 { // 末端的狀況
			pre = preG.Next
			reverseListNode(preG, now)
		}
		calLen++
	}
	return head
}

func Q680ValidPalindromeII(s string) bool {
	var recursive func(s string, chance int) bool
	recursive = func(s string, chance int) bool {
		for i, j := 0, len(s)-1; i < j; {
			if s[i] != s[j] {
				if chance > 0 {
					chance--
					return recursive(s[i+1:j+1], chance) || recursive(s[i:j], chance)
				} else {
					return false
				}
			}
			i++
			j--
		}
		return true
	}
	return recursive(s, 1)
}

func Q125ValidPalindrome(s string) bool {
	solve1 := func(s string) bool {
		str := strings.ToLower(s)
		str = regexp.MustCompile(`[^a-z0-9]`).ReplaceAllString(str, "")
		for i, j := 0, len(str)-1; i < j; {
			if str[i] != str[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	solve2 := func(s string) bool {
		str := []byte(s)
		fitQuiz := func(b byte) byte {
			if b >= 'A' && b <= 'Z' {
				return byte(int(b) + 32)
			}
			if b >= 'a' && b <= 'z' ||
				b >= '0' && b <= '9' {
				return b
			}
			return ' '
		}
		for i, j := 0, len(str)-1; i < j; {
			a := fitQuiz(str[i])
			b := fitQuiz(str[j])
			if a == ' ' {
				i++
				continue
			}
			if b == ' ' {
				j--
				continue
			}
			if a != b {
				return false
			} else {
				i++
				j--
			}
		}
		return true
	}
	_ = solve1(s)
	return solve2(s)
}

func Q234PalindromeLinkedList(head *ListNode) bool {
	// solve1 := func(head *ListNode) bool {
	// 	arrange := []byte{}
	// 	var last *ListNode
	// 	for head != nil {
	// 		t := head.Next
	// 		head.Next = last
	// 		last = head
	// 		head = t
	// 		arrange = append(arrange, byte(last.Val))
	// 	}

	// 	for _, v := range arrange {
	// 		if v != byte(last.Val) {
	// 			return false
	// 		}
	// 		last = last.Next
	// 	}
	// 	return true
	// }
	solve2 := func(head *ListNode) bool {
		mid, tail := head, head
		// get mid and tail
		for tail != nil && tail.Next != nil {
			mid = mid.Next
			tail = tail.Next.Next
		}
		// reverse mid to tail
		now := mid.Next
		mid.Next = nil
		last := mid
		for now != nil {
			t := now.Next
			now.Next = last
			last = now
			now = t
		}
		// compare head to mid and head to tail
		for head != nil && last != nil {
			if head.Val != last.Val {
				return false
			}
			head = head.Next
			last = last.Next
		}
		return true
	}
	// _ = solve1(head)
	return solve2(head)
}

func Q156BinaryTreeUpsideDown(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil {
		return root
	}
	var newRoot *TreeNode

	var re func(left *TreeNode, father *TreeNode, fatherRight *TreeNode)
	re = func(left *TreeNode, father *TreeNode, fatherRight *TreeNode) {
		if left.Left != nil {
			re(left.Left, left, left.Right)
		}
		if newRoot == nil {
			newRoot = left
		}
		if left.Right != nil {
			re(left.Right, left, nil)
		}
		if father != nil && left != father.Right {
			left.Left = father.Right
			left.Right = father
		}

	}
	re(root, nil, nil)
	root.Left = nil
	root.Right = nil
	return newRoot
}

func Q92ReverseLinkedListII(head *ListNode, left int, right int) *ListNode {
	var preLp, preRp, tp *ListNode
	tp = head
	i := 0
	for tp != nil {
		i++
		if i+1 == left {
			preLp = tp
		}
		if i == right {
			preRp = tp.Next
			break
		}
		tp = tp.Next
	}
	// reverse left to right
	var h *ListNode
	if preLp != nil {
		h = preLp.Next
	} else {
		h = head
	}
	last := preRp
	for h != preRp {
		t := h.Next
		h.Next = last
		last = h
		h = t
	}
	if preLp != nil {
		preLp.Next = last
	} else {
		head = last
	}
	return head
}

func Q206ReverseLinkedList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}

func Q445AddTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	// reverse
	var pre, ll1, ll2 *ListNode
	pre = nil
	for l1 != nil {
		ll1 = &ListNode{
			Val:  l1.Val,
			Next: nil,
		}
		if pre != nil {
			ll1.Next = pre
		}
		pre = ll1
		l1 = l1.Next
	}
	pre = nil
	for l2 != nil {
		ll2 = &ListNode{
			Val:  l2.Val,
			Next: nil,
		}
		if pre != nil {
			ll2.Next = pre
		}
		pre = ll2
		l2 = l2.Next
	}

	// add
	var ans *ListNode
	carry := 0
	pre = nil
	for ll1 != nil || ll2 != nil {
		t := 0
		if ll1 != nil {
			t += ll1.Val
			ll1 = ll1.Next
		}
		if ll2 != nil {
			t += ll2.Val
			ll2 = ll2.Next
		}
		t += carry
		carry = t / 10
		ans = &ListNode{
			Val:  t % 10,
			Next: nil,
		}
		if pre != nil {
			ans.Next = pre
		}
		pre = ans
	}
	if carry != 0 {
		ans = &ListNode{
			Val:  carry % 10,
			Next: nil,
		}
		if pre != nil {
			ans.Next = pre
		}
		pre = ans
	}

	for ans != nil && ans.Next != nil && ans.Val == 0 {
		ans = ans.Next
	}
	return ans
}
func Q371SumOfTwoIntegers(a int, b int) int {
	// without using the operators + and -.
	b1 := a
	b2 := b

	sum := b1 ^ b2
	carry := b1 & b2
	for carry != 0 {
		b1, b2 = sum, carry<<1
		sum = b1 ^ b2
		carry = b1 & b2
	}
	return sum
}
func Q2AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var ans, preN, nowN *ListNode
	for l1 != nil || l2 != nil {
		t := 0
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		t += carry
		carry = t / 10

		nowN = &ListNode{
			Val:  t % 10,
			Next: nil,
		}
		if preN == nil {
			preN, ans = nowN, nowN

		} else {
			preN.Next = nowN
			preN = nowN
		}
	}
	if carry != 0 {
		nowN = &ListNode{
			Val:  carry % 10,
			Next: nil,
		}
		preN.Next = nowN
	}
	return ans
}

func Q415AddStrings(num1 string, num2 string) string {
	// You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.
	reverseAns := []byte{}
	carry := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; {
		t := 0
		if i >= 0 {
			t += int(num1[i]) - 48
			i--
		}
		if j >= 0 {
			t += int(num2[j]) - 48
			j--
		}
		t += carry
		carry = t / 10
		reverseAns = append(reverseAns, byte(t%10+48))
	}
	if carry != 0 {
		reverseAns = append(reverseAns, byte(carry+48))
	}

	l := len(reverseAns)
	for i := 0; i < l-i-1; i++ {
		reverseAns[i], reverseAns[l-i-1] = reverseAns[l-i-1], reverseAns[i]
	}

	return string(reverseAns)
}
func Q989AddToArray_FormOfInteger(num []int, k int) []int {
	carry := 0
	for i := len(num) - 1; i >= 0 && (k > 0 || carry > 0); i-- {
		t := num[i] + carry + k%10
		num[i], carry, k = t%10, t/10, k/10
	}
	carry += k
	for carry != 0 {
		num = append([]int{carry % 10}, num...)
		carry /= 10
	}
	return num
}
func Q369PlusOneLinkedList(head *ListNode) *ListNode {
	query := []*ListNode{}

	for now := head; now != nil; now = now.Next {
		query = append(query, now)
	}

	carry := 1
	for i := len(query) - 1; i >= 0; i-- {
		t := query[i].Val + carry
		query[i].Val, carry = t%10, t/10
	}
	if carry > 0 {
		head = &ListNode{
			Val:  carry,
			Next: head,
		}
	}
	return head
}

func Q43MultiplyStrings(num1 string, num2 string) string {
	// Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
	Itb := func(i int) byte {
		return byte(i) + 48
	}
	Bti := func(b byte) int {
		return int(b) - 48
	}
	ans := []byte{}
	aa := []byte(num1)
	bb := []byte(num2)

	sumList := [][]byte{}
	for i := len(aa) - 1; i >= 0; i-- {
		carry := 0
		row := []byte{}
		for j := len(bb) - 1; j >= 0; j-- {
			t := Bti(aa[i])*Bti(bb[j]) + carry
			carry = t / 10
			row = append(row, Itb(t%10))
		}
		if carry != 0 {
			row = append(row, Itb(carry))
		}
		sumList = append(sumList, row)
	}

	popIdx := 0
	carry := 0
	for hasVal := true; hasVal; {
		hasVal = false
		t := 0
		for i := range sumList {
			if popIdx-i < 0 {
				continue
			}
			if popIdx-i > len(sumList[i])-1 {
				continue
			}
			t += Bti(sumList[i][popIdx-i])
			hasVal = true
		}
		if hasVal {
			t += carry
			carry = t / 10
			ans = append(ans, Itb(t%10))
			popIdx++
		}
	}
	if carry != 0 {
		ans = append(ans, Itb(carry))
	}

	l := len(ans) - 1
	for l >= 0 && ans[l] == '0' {
		l--
	}
	if l < 0 {
		return "0"
	}
	for i := 0; i < l-i; i++ {
		ans[i], ans[l-i] = ans[l-i], ans[i]
	}

	return string(ans)
}

func Q66PlusOne(digits []int) []int {

	carry := 1 // basic 1
	ans := []int{}
	for i := len(digits) - 1; i >= 0; i-- {
		t := digits[i] + carry
		ans = append(ans, t%10)
		carry = t / 10
	}
	if carry == 1 {
		ans = append(ans, 1)
	}

	// leetcode use will not change
	// sort.Slice(ans, func(i, j int) bool {
	// 	return i > j
	// })

	l := len(ans)
	for i := 0; i < l-i-1; i++ {
		ans[i], ans[l-i-1] = ans[l-i-1], ans[i]
	}

	return ans
}

func Q67AddBinary(a string, b string) string {
	aa := []byte(a)
	bb := []byte(b)
	i := len(aa) - 1
	j := len(bb) - 1
	carry := '0'
	ans := []byte{}
	for i >= 0 || j >= 0 {
		cnt := 0
		if i >= 0 {
			if aa[i] == '1' {
				cnt++
			}
			i--
		}
		if j >= 0 {
			if bb[j] == '1' {
				cnt++
			}
			j--
		}
		if carry == '1' {
			cnt++
		}
		if cnt/2 > 0 {
			carry = '1'
		} else {
			carry = '0'
		}
		if cnt%2 == 0 {
			ans = append(ans, '0')
		} else {
			ans = append(ans, '1')
		}
	}
	if carry == '1' {
		ans = append(ans, '1')
	}

	// leetcode use will not change
	// sort.Slice(ans, func(i, j int) bool {
	// 	return i > j
	// })

	l := len(ans)
	for i := 0; i < l-i-1; i++ {
		ans[i], ans[l-i-1] = ans[l-i-1], ans[i]
	}

	return string(ans)
}

func Q718MaximumLengthOfRepeatedSubarray(nums1 []int, nums2 []int) int {
	// time limit
	solve1 := func(nums1 []int, nums2 []int) int {
		if len(nums1) > len(nums2) {
			nums1, nums2 = nums2, nums1
		}
		l1 := len(nums1)
		l2 := len(nums2)
		maxLength := 0
		for baseIdx := range nums1 {
			if l1-(baseIdx+1) < maxLength {
				break
			}
			for i := baseIdx; i < l1; i++ {
				if l1-(i+1) < maxLength {
					break
				}
				for j := 0; j < l2; j++ {
					if l2-(j+1) < maxLength {
						break
					}
					nowLength := 0
					for k, l := i, j; k < l1 && l < l2 && nums1[k] == nums2[l]; {
						k++
						l++
						nowLength++

					}
					if nowLength > maxLength {
						maxLength = nowLength
					}
				}
			}
		}
		return maxLength
	}

	// solve2 := func(nums1 []int, nums2 []int) int {
	// 	dp := make(map[int]int)
	// 	list := []int{}
	// 	for

	// 	for i := range nums1 {
	// 		for j := range nums1[:i+1] {
	// 			nums1[j:] // all nums1 sub
	// 		}
	// 	}
	// 	for i := range nums2 {
	// 		for j := range nums2[:i+1] {
	// 			nums2[j:] // all nums2 sub
	// 		}
	// 	}
	// }
	// _ = solve2(nums1, nums2)
	return solve1(nums1, nums2)
	// note---
}

func Q1995CountSpecialQuadruplets(nums []int) int {
	// solve1 := func(nums []int) int {
	// 	tmp := []int{}
	// 	for i := len(nums) - 1; i >= 0; i-- {
	// 		tmp = append(tmp, nums[i])
	// 	}
	// 	nums = tmp
	// 	// fmt.Println(nums)
	// 	n := 4
	// 	var find func(nums []int, n int, target int) int
	// 	find = func(nums []int, n int, target int) int {
	// 		l := len(nums)
	// 		if n == 1 {
	// 			// fmt.Println("n==1")
	// 			ans := 0
	// 			for _, v := range nums {
	// 				if v == target {
	// 					ans++
	// 				}
	// 			}
	// 			// fmt.Printf("find: %v times: %v in nums:%v \n", target, ans, nums)
	// 			return ans
	// 		}
	// 		ans := 0
	// 		for i, v := range nums {
	// 			// fmt.Printf("nums:%v, ", nums)
	// 			if i > l-n {
	// 				// fmt.Printf("overflow nums, break i:%v, l:%v, n:%v\n", i, l, n)
	// 				break

	// 			}
	// 			goal := target - v
	// 			if goal < 0 { // every num should be positive
	// 				// fmt.Printf("choose %v ,every num should be positive, continue i:%v, l:%v, n:%v\n", v, i, l, n)
	// 				continue
	// 			}
	// 			// fmt.Printf("base:%v, goal %v, find %v num in %v, -- i:%v, l:%v, n:%v\n", v, goal, n-1, nums[i+1:], i, l, n)
	// 			ans += find(nums[i+1:], n-1, goal)
	// 		}
	// 		return ans
	// 	}
	// 	ans := 0
	// 	l := len(nums)
	// 	for i, v := range nums {
	// 		if i > l-n {
	// 			break
	// 		}
	// 		goal := v
	// 		// fmt.Printf("base:%v, find %v num in %v\n", v, n-1, nums[i+1:])
	// 		ans += find(nums[i+1:], n-1, goal)
	// 	}
	// 	return ans
	// }
	// _ = solve1(nums)

	solve2 := func(nums []int) int {
		// a + b + c = d  // value
		// a < b < c < d  // idx
		// ---
		// a + b + c - d = 0
		len := len(nums)
		ans := 0
		for i := 0; i < len-1; i++ {
			for j := i + 1; j < len; j++ {
				for k := j + 1; k < len; k++ {
					for l := k + 1; l < len; l++ {
						tSum := nums[i] + nums[j] + nums[k] - nums[l]
						if tSum == 0 {
							ans++
						}
					}
				}
			}
		}

		return ans
	}
	return solve2(nums)
}

func Q4544SumII(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// solve1 超時
	solve1 := func(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
		return 0
		/*
			twoSum := func(nums1 []int, nums2 []int, target int) int {
				rMap := make(map[int]int, 200)
				ans := 0
				for _, v := range nums2 {
					rMap[v]++
				}
				for i := range nums1 {
					goal := target - nums1[i]
					ans += rMap[goal]
				}
				return ans
			}
			rNMap := make(map[int]map[int]int, 5) // number mt target mt combine
			rNMap[2] = make(map[int]int, 4000)
			rNMap[3] = make(map[int]int, 4000)
			rNMap[4] = make(map[int]int, 200)
			var nSum func(n int, nums [][]int, target int) int
			nSum = func(n int, nums [][]int, target int) int {
				ans := 0
				if n > 2 {
					for _, v := range nums[0] {
						goal := target - v
						if _, ok := rNMap[n][goal]; ok {
							ans += rNMap[n][goal]
						} else {
							rNMap[n] = make(map[int]int)
							rNMap[n][goal] = nSum(n-1, nums[1:], goal)
							ans += rNMap[n][goal]
						}
					}
				} else if n == 2 {
					if _, ok := rNMap[2][target]; ok {
						ans += rNMap[2][target]
					} else {
						rNMap[2] = make(map[int]int)
						rNMap[2][target] = twoSum(nums[0], nums[1], target)
						ans += rNMap[2][target]
					}
				}
				return ans
			}
			sort.Ints(nums1)
			sort.Ints(nums2)
			sort.Ints(nums3)
			sort.Ints(nums4)
			nsums := append([][]int{}, nums1, nums2, nums3, nums4)
			return nSum(4, nsums, 0) // because 4sum
		*/
	}

	// 原本考慮 前面的組合會影響到後面的，但後來思考發現兩組兩組並不會影響，因為已經列出
	// 兩組所有的可能了
	solve2 := func(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
		target := 0
		sort.Ints(nums1)
		sort.Ints(nums2)
		sort.Ints(nums3)
		sort.Ints(nums4)
		nSum := func(nums [][]int) (goalCntMap map[int]int) {
			goalCntMap = make(map[int]int)
			for _, v := range nums[0] {
				for _, v2 := range nums[1] {
					goalCntMap[v+v2]++
				}
			}
			return goalCntMap
		}
		A := nSum([][]int{nums1, nums2}) //A是nums1和nums2組合次數表
		ans := 0
		// for _, v := range nums3 {
		// 	for _, v2 := range nums4 {
		// 		ans += A[target-(v+v2)]
		// 	}
		// }
		B := nSum([][]int{nums3, nums4})
		for aGoal, aCnt := range A {
			ans += aCnt * B[target-aGoal]
		}
		return ans
	}
	_ = solve1(nums1, nums2, nums3, nums4)
	return solve2(nums1, nums2, nums3, nums4)
}

func Q153Sum(nums []int) [][]int {
	solve1 := func(nums []int) [][]int {
		if len(nums) < 1 {
			return [][]int{}
		}
		sort.Ints(nums)
		ansMap := make(map[string][]int)
		findTwoSum := func(basicV int, numsSorted []int, k int) {
			gMap := make(map[int]struct{})
			for _, v := range numsSorted {
				goal := k - v
				if _, ok := gMap[goal]; ok {
					tmp := []int{basicV, v, goal}
					sort.Ints(tmp)
					key := strconv.Itoa(tmp[0]) + "," + strconv.Itoa(tmp[1]) + "," + strconv.Itoa(tmp[2])
					ansMap[key] = tmp

				} else {
					gMap[v] = struct{}{}
				}
			}
		}
		for i, v := range nums {
			if i+1 > len(nums)-1 {
				break
			}
			goal := 0 - v // 0 is quiz target
			findTwoSum(v, nums[i+1:], goal)
		}

		ans := [][]int{}
		for _, v := range ansMap {
			ans = append(ans, v)
		}
		return ans
	}

	// solve2 有一個進步是使用ij逼近的時候，若要把剩下的找完i和j可以繼續下去
	// * 有可能相同的數組，但是index不一樣，但由於題目要的答案不要求索引，只要
	//      求值並且不重複，所以有排序的數組可以很好的避免重複查找的問題
	// 		由於是部分index重複不會涵蓋到所有，所以要把重複的可能避免掉
	solve2 := func(nums []int) [][]int {
		// [-2,0,0,2,2]
		// [[-2,0,2]]
		ans := [][]int{}
		l := len(nums)
		if l < 3 {
			return [][]int{}
		}
		sort.Ints(nums)

		for i, j, k := 0, 0, 0; i < l-2; i++ {
			j = i + 1
			k = l - 1
			if i > 0 && nums[i] == nums[i-1] {
				continue // 跳過重複的數字
			}
			for j < k {
				if k+1 < l && nums[k] == nums[k+1] {
					k--
					continue // 跳過重複的數字
				}
				if j-1 > i && nums[j] == nums[j-1] {
					j++
					continue // 跳過重複的數字
				}
				goal := 0 - nums[i] // 0 is quiz target
				tSum := nums[j] + nums[k]
				if tSum == goal {
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
					// 繼續循找
					k--
					j++
				} else if tSum > goal {
					k--
				} else {
					j++
				}
			}
		}
		return ans
	}

	_ = solve1(nums)
	return solve2(nums)
}

func Q1214TwoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	r1Map := make(map[int]struct{})
	if root1 == nil || root2 == nil {
		return false
	}

	query1 := make([]*TreeNode, 0)
	query1 = append(query1, root1)
	var tNode *TreeNode
	for len(query1) != 0 { // go over root1 and map to r1Map O(n)
		tNode, query1 = query1[0], query1[1:]
		if tNode.Left != nil {
			query1 = append(query1, tNode.Left)
		}
		if tNode.Right != nil {
			query1 = append(query1, tNode.Right)
		}
		r1Map[tNode.Val] = struct{}{}
	}

	query2 := make([]*TreeNode, 0)
	query2 = append(query2, root2)
	for len(query2) != 0 { // go over root2 and O(n)
		tNode, query2 = query2[0], query2[1:]
		goal := target - tNode.Val
		if _, ok := r1Map[goal]; ok {
			return true
		}
		if tNode.Left != nil {
			query2 = append(query2, tNode.Left)
		}
		if tNode.Right != nil {
			query2 = append(query2, tNode.Right)
		}
	}
	return false
}

func Q653TwoSumIV(root *TreeNode, k int) bool {
	var recursive func(now *TreeNode, k int, root *TreeNode) bool
	recursive = func(now *TreeNode, k int, root *TreeNode) bool {
		if now == nil {
			return false
		}
		goal := k - now.Val
		tNode := root
		for tNode != nil {
			if goal == tNode.Val && tNode != now {
				return true
			}
			if goal > tNode.Val {
				tNode = tNode.Right
				continue
			}
			if goal <= tNode.Val {
				tNode = tNode.Left
				continue
			}
		}
		return recursive(now.Left, k, root) || recursive(now.Right, k, root)
	}
	return recursive(root, k, root)
}

func Q1099TwoSumLessThanK(nums []int, k int) int {
	sort.Ints(nums)
	if len(nums) < 1 {
		return -1
	}
	low := 0
	high := len(nums) - 1
	closestK := -1
	for low != high {
		tSum := nums[low] + nums[high]
		if tSum == k {
			high--
		} else if tSum > k {
			high--
		} else if tSum < k {
			if tSum > closestK {
				closestK = tSum
			}
			low++
		}
	}
	return closestK
}

func Q184Sum(nums []int, target int) [][]int {
	solve1 := func(nums []int, target int) [][]int {
		sort.Ints(nums)
		twoSum := func(nums []int, target int) [][]int {
			ans := [][]int{}
			l := len(nums)
			for i, j := 0, l-1; i < j; {
				if j+1 < l && nums[j] == nums[j+1] {
					j--
					continue
				}
				if i-1 > 0 && nums[i] == nums[i-1] {
					i++
					continue
				}
				// goal := target - nums[i]
				tSum := nums[i] + nums[j]
				if tSum == target {
					ans = append(ans, []int{nums[i], nums[j]})
					i++
					j--
				} else if tSum > target {
					j--
				} else {
					i++
				}
			}
			return ans
		}
		threeSum := func(nums []int, target int) [][]int {
			ans := [][]int{}
			l := len(nums)
			for i := 0; i < l-2; i++ {
				if i > 0 && nums[i] == nums[i-1] {
					continue
				}
				goal := target - nums[i]
				sumSlice := twoSum(nums[i+1:], goal)
				for _, intSlice := range sumSlice {
					intSlice = append(intSlice, nums[i])
					ans = append(ans, intSlice)
				}
			}
			return ans
		}
		fourSum := func(nums []int, target int) [][]int {
			ans := [][]int{}
			l := len(nums)
			for i := 0; i < l-3; i++ {
				if i > 0 && nums[i] == nums[i-1] {
					continue
				}
				goal := target - nums[i]
				sumSlice := threeSum(nums[i+1:], goal)
				for _, intSlice := range sumSlice {
					intSlice = append(intSlice, nums[i])
					ans = append(ans, intSlice)
				}
			}
			return ans
		}
		return fourSum(nums, target)
	}
	_ = solve1(nums, target)
	// solve2 是 solve1的簡化重複版本，可以處理nSum
	solve2 := func(nums []int, target int) [][]int {
		sort.Ints(nums)
		twoSum := func(nums []int, target int) [][]int {
			ans := [][]int{}
			l := len(nums)
			for i, j := 0, l-1; i < j; {
				if j+1 < l && nums[j] == nums[j+1] {
					j--
					continue
				}
				if i-1 > 0 && nums[i] == nums[i-1] {
					i++
					continue
				}
				// goal := target - nums[i]
				tSum := nums[i] + nums[j]
				if tSum == target {
					ans = append(ans, []int{nums[i], nums[j]})
					i++
					j--
				} else if tSum > target {
					j--
				} else {
					i++
				}
			}
			return ans
		}
		var nSum func(nsum int, nums []int, target int) [][]int
		nSum = func(nsum int, nums []int, target int) [][]int {
			ans := [][]int{}
			l := len(nums)
			for i := 0; i < l-(nsum-1); i++ {
				if i > 0 && nums[i] == nums[i-1] {
					continue
				}
				goal := target - nums[i]
				var sumSlice [][]int
				if nsum-1 == 2 {
					sumSlice = twoSum(nums[i+1:], goal)
				} else {
					sumSlice = nSum(nsum-1, nums[i+1:], goal)
				}
				for _, intSlice := range sumSlice {
					intSlice = append(intSlice, nums[i])
					ans = append(ans, intSlice)
				}
			}
			return ans
		}
		return nSum(4, nums, target) // 4 for quiz
	}
	return solve2(nums, target)
}

func Q7ReverseInteger(x int) int {
	// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
	if x == 0 {
		return 0
	}
	ans := []byte{}
	if x < 0 {
		ans = append(ans, '-')
		x = -x
	}
	isLeaderZero := true
	for ; x != 0; x /= 10 {
		mod := x%10 + 48
		if isLeaderZero && mod != 0 {
			isLeaderZero = false
		}
		if !isLeaderZero {
			ans = append(ans, byte(mod))
		}

	}
	bToInt := func(b []byte) int {
		rst, _ := strconv.Atoi(string(b))
		return rst
	}
	maxInt32b := []byte("2147483648")
	if ans[0] == '-' {
		if len(ans[1:]) > len(maxInt32b) {
			return 0
		}
		if len(ans[1:]) < len(maxInt32b) {
			return bToInt(ans)
		}
		if len(ans[1:]) == len(maxInt32b) {
			crst := bytes.Compare(ans[1:], maxInt32b)
			if crst <= 0 {
				return bToInt(ans)
			}
		}
	} else {
		if len(ans) > len(maxInt32b) {
			return 0
		}
		if len(ans) < len(maxInt32b) {
			return bToInt(ans)
		}
		if len(ans) == len(maxInt32b) {
			crst := bytes.Compare(ans, maxInt32b)
			if crst < 0 {
				return bToInt(ans)
			}
		}
	}
	return 0
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
func Q167TwoSumIIInputArrayIsSorted(numbers []int, target int) []int {
	solve1 := func(numbers []int, target int) []int {
		m := make(map[int]int)
		for i, v := range numbers {
			// target = v + x
			// x = target - v
			x := target - v
			if j, ok := m[x]; ok {
				ans := []int{i + 1, j + 1}
				sort.Ints(ans)
				return ans
			}
			m[v] = i
		}
		return nil
	}
	return solve1(numbers, target)
}
func Q1TwoSum(nums []int, target int) []int {
	solve1 := func(nums []int, target int) []int {
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

	solve2 := func(nums []int, target int) []int {
		m := make(map[int]int)
		for i, v := range nums {
			// target = v + x
			// x = target - v
			x := target - v
			if j, ok := m[x]; ok {
				ans := []int{i, j}
				sort.Ints(ans)
				return ans
			}
			m[v] = i
		}
		return nil

	}
	// solve1(nums, target)
	_ = solve1
	return solve2(nums, target)
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
