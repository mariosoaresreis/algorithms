package TreeBreadthFirstSearch

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Traverse(root *TreeNode) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{root.Val})
	addNode(root, &result)

	return result
}

func addNode(n *TreeNode, result *[][]int) {
	newLevel := make([]int, 0)

	if n.Left != nil {
		newLevel = append(newLevel, n.Left.Val)
	}

	if n.Right != nil {
		newLevel = append(newLevel, n.Right.Val)
	}

	if len(newLevel) > 0 {
		*result = append(*result, newLevel)
	}

	if n.Left != nil {
		addNode(n.Left, result)
	}

	if n.Right != nil {
		addNode(n.Right, result)
	}

}
