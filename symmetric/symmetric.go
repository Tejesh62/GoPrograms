package symmetric

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode

}

func Symmetric(root *TreeNode) bool {
	if root == nil {
			return true
		}

		return Mirror(root.Left, root.Right)
	}

func Mirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true

	}
	if t1 == nil || t2 == nil {
		return false
	}

	return (t1.Val == t2.Val) &&
		Mirror(t1.Left, t2.Right) &&
		Mirror(t1.Right, t2.Left)

}

// func main() {

// 	root := &TreeNode {
// 		Val: 1,
// 		Left: &TreeNode {
// 			Val: 2,
// 			Left: &TreeNode{Val: 3},
// 			Right: &TreeNode{Val: 4},
// 		},

// 		Right: &TreeNode {
// 			Val: 2,
// 			Left: &TreeNode{Val: 4},
// 			Right: &TreeNode{Val: 3},

// 		},
// 	}

// 	fmt.Println("is tree Symmetric?", Symmetric(root))

// }
