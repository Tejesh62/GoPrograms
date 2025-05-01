package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
}

// func main() {
// 	root := &TreeNode{Val: 5,
// 		Left: &TreeNode{Val: 4,
// 			Left: &TreeNode{Val: 11,
// 				Left:  &TreeNode{Val: 7},
// 				Right: &TreeNode{Val: 2},
// 			},
// 		},
// 		Right: &TreeNode{Val: 8,
// 			Left:  &TreeNode{Val: 13},
// 			Right: &TreeNode{Val: 4,
// 				Right: &TreeNode{Val: 1},
// 			},
// 		},
// 	}

// 	target := 22
// 	fmt.Println("Has path sum:", hasPathSum(root, target)) // Output: true
// }
