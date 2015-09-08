package gllrb

// Node is a leaf of our LLRB
type Node struct {
	// Red = True and Black = False
	color  bool
	Key    Comparer
	Value  interface{}
	Number uint64
	left   *Node
	right  *Node
}

func isRed(n *Node) bool {
	if n == nil {
		return false // BLACK
	}

	return n.color == true // RED is true
}

func (n *Node) flipColors() {
	n.color = !n.color             // RED
	n.left.color = !n.left.color   // BLACK
	n.right.color = !n.right.color // BLACK
}

func rotateLeft(n *Node) *Node {
	x := n.right
	n.right = x.left
	x.left = n
	x.color = n.color
	n.color = true // RED
	x.Number = n.Number
	n.Number = calculateNumber(n) + 1
	return x
}

func rotateRight(n *Node) *Node {
	x := n.left
	n.left = x.right
	x.right = n
	x.color = n.color
	n.color = true // RED
	x.Number = n.Number
	n.Number = calculateNumber(n) + 1
	return x
}

func calculateNumber(n *Node) uint64 {
	if n.right != nil && n.left != nil {
		return n.right.Number + n.left.Number
	}

	rightNumber := uint64(0)
	leftNumber := uint64(0)

	if n.right != nil && n.left == nil {
		rightNumber = n.right.Number
	}

	if n.left != nil && n.right == nil {
		leftNumber = n.left.Number
	}

	return rightNumber + leftNumber

}

func moveRedLeft(n *Node) *Node {
	n.flipColors()
	if isRed(n.right.left) {
		n.right = rotateRight(n.right)
		n = rotateLeft(n)
		n.flipColors()
	}
	return n
}

func moveRedRight(n *Node) *Node {
	n.flipColors()
	if isRed(n.left.left) {
		n = rotateRight(n)
		n.flipColors()
	}
	return n
}

func min(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func max(node *Node) *Node {
	for node.right != nil {
		node = node.right
	}
	return node
}

func deleteMin(n *Node) *Node {
	if n.left == nil {
		return nil
	}

	if !isRed(n.left) && !isRed(n.left.left) {
		n = moveRedLeft(n)
	}

	n.left = deleteMin(n.left)
	return fixUp(n)
}

// fixUp ensures that our red black tree's RED nodes lean left
func fixUp(n *Node) *Node {
	// Do we have a right leaning RED node?
	if isRed(n.right) && !isRed(n.left) {
		n = rotateLeft(n)
	}

	// Do we have 2 consecutive left leaning red nodes?
	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}

	// Are we perfectly balanced but need to flip colors?
	if isRed(n.right) && isRed(n.left) {
		n.flipColors()
	}

	return n
}
