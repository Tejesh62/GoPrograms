package traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func OrderTraversal(root *TreeNode) []int {
	var  result []int
	Order(root, &result)
	return result
}

func Order(node *TreeNode, result *[] int) {
	if node == nil {
		return
	}

	Order(node.Left, result)
	*result = append(*result, node.Val)
	Order(node.Right, result)
}

// func main() {
// 	root := &TreeNode{Val: 1}
// 	root.Right = &TreeNode{Val: 2}
// 	root.Right.Left = &TreeNode{Val: 3}

// 	fmt.Println(inorderTraversal(root)) 

// }
