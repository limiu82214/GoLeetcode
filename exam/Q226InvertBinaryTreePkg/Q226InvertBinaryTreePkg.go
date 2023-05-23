package Q226InvertBinaryTreePkg

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solve(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	deepSwap(root)

	return root
}

func deepSwap(r *TreeNode) {
	r.Left, r.Right = r.Right, r.Left
	if r.Left != nil {
		deepSwap(r.Left)
	}

	if r.Right != nil {
		deepSwap(r.Right)
	}
}

func Solve2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	for len(q) != 0 {
		pop := q[len(q)-1]
		q = q[:len(q)-1]

		if pop != nil {
			q = append(q, pop.Left, pop.Right)
			pop.Left, pop.Right = pop.Right, pop.Left
		}
	}

	return root
}
