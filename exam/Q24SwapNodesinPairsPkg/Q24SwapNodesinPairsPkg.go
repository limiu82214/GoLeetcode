package Q24SwapNodesinPairsPkg

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solve(head *ListNode) *ListNode {
	if head == nil {
		return head
	} else if head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	pre := dummy
	for pre != nil && pre.Next != nil && pre.Next.Next != nil {
		// fmt.Println(pre, pre.Next, pre.Next.Next)
		// lop(dummy)
		swapNext2(pre)
		pre = pre.Next
		if pre.Next != nil {
			pre = pre.Next
		}
		// fmt.Println(pre)
	}
	return dummy.Next
}
func lop(h *ListNode) {
	for h.Next != nil {
		fmt.Print(h.Val)
		h = h.Next
	}
	fmt.Println(h.Val)
}

func swapNext2(n *ListNode) {
	pre := n
	mid := n.Next
	suf := n.Next.Next
	pre.Next, suf.Next, mid.Next = suf, mid, suf.Next
}
