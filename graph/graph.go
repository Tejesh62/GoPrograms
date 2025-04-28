package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// Map to track original node -> cloned node
	visited := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if v, ok := visited[n]; ok {
			return v
		}

		clone := &Node{Val: n.Val} // <-- this must be := instead of just clone
		visited[n] = clone

		for _, neighbor := range n.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
		}

		return clone
	}

	return dfs(node)
}


// func main() {
	
// 	node1 := &graph.Node{Val: 1}
// 	node2 := &graph.Node{Val: 2}
// 	node1.Neighbors = []*graph.Node{node2}
// 	node2.Neighbors = []*graph.Node{node1}

// 	cloned := graph.CloneGraph(node1)

// 	fmt.Println("Original Node:", node1.Val)
// 	fmt.Println("Cloned Node:", cloned.Val)

	
// 	fmt.Println("Original Neighbor:", node1.Neighbors[0].Val)
// 	fmt.Println("Cloned Neighbor:", cloned.Neighbors[0].Val)
// }
