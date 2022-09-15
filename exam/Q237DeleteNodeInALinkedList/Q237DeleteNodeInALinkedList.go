package q237deletenodeinalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
