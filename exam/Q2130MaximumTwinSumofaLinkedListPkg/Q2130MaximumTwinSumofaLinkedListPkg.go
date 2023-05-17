package Q2130MaximumTwinSumofaLinkedListPkg

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solve(head *ListNode) int {
	s := []int{}
	helf := head
	for head != nil {
		s = append(s, helf.Val)
		helf = helf.Next
		head = head.Next.Next
	}

	// fmt.Println(s)
	// fmt.Println(helf)
	l := len(s) - 1
	max := 0
	var t int
	for helf != nil {
		t = s[l] + helf.Val
		if t > max {
			max = t
		}
		l--
		helf = helf.Next
	}

	return max
}
