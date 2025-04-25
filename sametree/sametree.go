package sametree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func SameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {

		return true
	}

	if p == nil || q == nil ||p.Val != q.Val {

		return false
	}

	return SameTree(p.Left, q.Left) && SameTree(p.Right, q.Right)
}

// func main() {
// 	p := &TreeNode{Val: 1}
// 	p.Left = &TreeNode{Val: 2}
// 	p.Right = &TreeNode{Val: 3}

// 	q := &TreeNode{Val: 1}
// 	q.Left = &TreeNode{Val: 2}
// 	q.Right = &TreeNode{Val: 3}

// 	fmt.Println(SameTree(p, q)) 

// 	r := &TreeNode{Val: 1}
// 	r.Right = &TreeNode{Val: 2}

// 	fmt.Println(SameTree(p, r)) 
// }
