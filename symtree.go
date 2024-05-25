package symtree

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

// Check if a binary tree is symmetric
func IsSymmetric(root *Node) bool {
	return isSymmetric(root, root)
}

func isSymmetric(left *Node, right *Node) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Value != right.Value {
		return false
	}

	return isSymmetric(left.Left, right.Right) && isSymmetric(left.Right, right.Left)
}
