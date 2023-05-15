package Q1721SwappingNodesinaLinkedListPkg

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solve(head *ListNode, k int) *ListNode {
	if head.Next == nil {
		return head
	}
	oriHead := head
	var (
		preFirst *ListNode
		first    *ListNode
		preLast  *ListNode
		last     *ListNode
		pre      *ListNode
	)

	count := 1 // 第幾個
	for head != nil {

		if count == k {
			preFirst = pre
			first = head
		}
		if count >= k {
			if last == nil { //剛進場沒有pre
				preLast = nil
				last = oriHead
			} else {
				preLast = last
				last = last.Next
			}
		}

		pre = head
		head = head.Next
		count++
	}

	// fmt.Println(preFirst)
	// fmt.Println(first)
	// fmt.Println(preLast)
	// fmt.Println(last)
	// 處理掉preFisrt是nil的情況，也就是要交換的前一個是1
	if preFirst == nil {
		oriHead = last
		if preLast == first { // 兩個的情況
			first.Next, last.Next = last.Next, first
		} else { // 三個以上的情況
			first.Next, last.Next = last.Next, first.Next
			preLast.Next = first
		}
	} else if preLast == nil {
		oriHead = first
		if preFirst == last {
			last.Next, first.Next = first.Next, last
		} else {
			last.Next, first.Next = first.Next, last.Next
			preFirst.Next = last
		}

	} else {
		swap(preFirst, first, preLast, last)
	}

	return oriHead
}

func swap(preF, f, preL, l *ListNode) {
	if preF != nil {
		preF.Next = l
		preL.Next = f
	}
	f.Next, l.Next = l.Next, f.Next
}
