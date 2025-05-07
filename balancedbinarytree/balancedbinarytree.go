package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHeight, leftBalanced := checkBalance(root.Left)
	rightHeight, rightBalanced := checkBalance(root.Right)

	
	balanced := leftBalanced && rightBalanced && abs(leftHeight-rightHeight) <= 1
	height := 1 + max(leftHeight, rightHeight)

	return height, balanced
}

func IsBalanced(root *TreeNode) bool {
	_, balanced := checkBalance(root)
	return balanced
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// func main() {
// 	root1 := &balancedbinarytree.TreeNode{
// 		Val: 3,
// 		Left: &balancedbinarytree.TreeNode{Val: 9},
// 		Right: &balancedbinarytree.TreeNode{
// 			Val: 20,
// 			Left:  &balancedbinarytree.TreeNode{Val: 15},
// 			Right: &balancedbinarytree.TreeNode{Val: 7},
// 		},
// 	}
// 	fmt.Println("Is tree 1 balanced? Expected: true ->", balancedbinarytree.IsBalanced(root1))

// 	root2 := &balancedbinarytree.TreeNode{
// 		Val: 1,
// 		Left: &balancedbinarytree.TreeNode{
// 			Val: 2,
// 			Left: &balancedbinarytree.TreeNode{
// 				Val: 3,
// 				Left: &balancedbinarytree.TreeNode{Val: 4},
// 			},
// 		},
// 		Right: &balancedbinarytree.TreeNode{Val: 2},
// 	}
// 	fmt.Println("Is tree 2 balanced? Expected: false ->", balancedbinarytree.IsBalanced(root2))
// }
