package gllrb

// LLRB is our left leaning red black tree
type LLRB struct {
	root *Node
}

// NewLLRB Creates a new Left Leaning red black tree
func NewLLRB() *LLRB {
	return &LLRB{root: nil}
}

// Put places a value into our LLRB
func (l *LLRB) Put(Key Comparer) {
	l.root = put(l.root, Key)
	l.root.color = false // BLACK
}

// Delete removes a value from our LLRB
func (l *LLRB) Delete(Key Comparer) {
	l.root = deletellrb(l.root, Key)

	if l.root != nil {
		l.root.color = false // BLACK
	}

}

// Get will return a value from our key LLRB
func (l *LLRB) Get(key Comparer) interface{} {

	if l.root == nil {
		return nil
	}

	val := l.root
	for val != nil {
		cmp := val.Key.Compare(key)
		switch {
		case cmp < 0:
			val = val.left
		case cmp > 0:
			val = val.right
		case cmp == 0:
			return key.Value()
		}
	}
	return nil
}

// Min returns the "lowest" item on the tree
func (l *LLRB) Min() interface{} {
	if l.root == nil {
		return nil
	}

	return min(l.root).Value
}

// Max returns the "highest" item on the tree
func (l *LLRB) Max() interface{} {
	if l.root == nil {
		return nil
	}

	return max(l.root).Value
}

// Size returns the length of the LLRB
func (l *LLRB) Size() uint64 {
	return l.root.Number
}

func deletellrb(n *Node, key Comparer) *Node {
	if n == nil {
		return nil
	}

	if cmp := n.Key.Compare(key); cmp < 0 {
		if !isRed(n.left) && !isRed(n.left.left) {
			n = moveRedLeft(n)
		}
		n.left = deletellrb(n.left, key)
		return fixUp(n)
	}

	if isRed(n.left) {
		n = rotateRight(n)
	}

	if n.Key.Compare(key) == 0 && n.right == nil {
		return nil
	}

	if !isRed(n.right) && !isRed(n.right.left) {
		n = moveRedRight(n)
	}

	if n.Key.Compare(key) == 0 {
		node := min(n.right)
		n.Value = node.Value
		n.Key = node.Key
		n.right = deleteMin(n.right)
	} else {
		n.right = deletellrb(n.right, key)
	}

	return fixUp(n)
}

func put(n *Node, key Comparer) *Node {
	if n == nil {
		return &Node{
			Key:    key,
			Value:  key.Value(),
			Number: uint64(1),
			color:  true, // RED
		}
	}

	cmp := n.Key.Compare(key)
	switch {
	case cmp < 0:
		n.left = put(n.left, key)
	case cmp > 0:
		n.right = put(n.right, key)
	case cmp == 0:
		n.Value = key.Value()
	}

	n = fixUp(n)

	n.Number = calculateNumber(n) + 1

	return n
}
