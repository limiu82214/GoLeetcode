package Q21MergeTwoSortedListsPkg

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solve(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy

	for list1 != nil || list2 != nil {
		switch {
		case list1 == nil:
			pre.Next = list2
			list2 = list2.Next
		case list2 == nil:
			pre.Next = list1
			list1 = list1.Next
		case list1.Val < list2.Val:
			pre.Next = list1
			list1 = list1.Next
		default:
			pre.Next = list2
			list2 = list2.Next
		}

		pre = pre.Next
	}

	return dummy.Next
}
