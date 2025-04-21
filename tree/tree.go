package tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Tree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		size := len(queue)
		prev := 0

		if level%2 == 0 {
			prev = -1 
		} else {
			prev = 1 << 31 - 1
		}

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				if node.Val%2 == 0 || node.Val <= prev {
					return false
				}
			} else {
				if node.Val%2 == 1 || node.Val >= prev {
					return false
				}
			}

			prev = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}

	return true
}

// func main() {
// 	root := &tree.TreeNode{
	// 	Val: 1,
	// 	Left: &tree.TreeNode{
	// 		Val: 10,
	// 		Left: &tree.TreeNode{
	// 			Val: 3,
	// 			Left:  &tree.TreeNode{Val: 12},
	// 			Right: &tree.TreeNode{Val: 8},
	// 		},
	// 	},
	// 	Right: &tree.TreeNode{
	// 		Val: 4,
	// 		Left: &tree.TreeNode{
	// 			Val: 7,
	// 			Left: &tree.TreeNode{Val: 6},
	// 		},
	// 		Right: &tree.TreeNode{
	// 			Val: 9,
	// 			Right: &tree.TreeNode{Val: 2},
	// 		},
	// 	},
	// }

	// fmt.Println(tree.Tree(root))
// }