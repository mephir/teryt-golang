package binarytree

import "fmt"

type BinaryTree[T any] struct {
	root *TreeNode[T]
}

func (t *BinaryTree[T]) GetRoot() *TreeNode[T] {
	return t.root
}
func (t *BinaryTree[T]) Insert(value T) error {
	t.root = insertNode(t.root, value, nil)

	return nil
}

func (t *BinaryTree[T]) Print(noParents bool) {
	if t.root != nil {
		printNode(t.root, "", false, noParents)
	}
}

func insertNode[T any](node *TreeNode[T], value T, parent *TreeNode[T]) *TreeNode[T] {
	if node == nil {
		return &TreeNode[T]{
			Value:  value,
			Parent: parent,
			Height: 1,
		}
	}

	if fmt.Sprintf("%v", value) < fmt.Sprintf("%v", node.Value) {
		node.Left = insertNode(node.Left, value, node)
	} else if fmt.Sprintf("%v", value) > fmt.Sprintf("%v", node.Value) {
		node.Right = insertNode(node.Right, value, node)
	} else {
		panic(fmt.Errorf("duplicate values are not allowed"))
	}

	node.Height = max(getHeight(node.Left), getHeight(node.Right)) + 1

	return balance(node)
}

func leftRotate[T any](x *TreeNode[T]) *TreeNode[T] {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Left = x
	y.Parent = x.Parent
	x.Parent = y
	x.Height = max(getHeight(x.Left), getHeight(x.Right)) + 1
	y.Height = max(getHeight(y.Left), getHeight(y.Right)) + 1
	return y
}

func rightRotate[T any](y *TreeNode[T]) *TreeNode[T] {
	x := y.Left
	y.Left = x.Right
	if x.Right != nil {
		x.Right.Parent = y
	}
	x.Right = y
	x.Parent = y.Parent
	y.Parent = x
	y.Height = max(getHeight(y.Left), getHeight(y.Right)) + 1
	x.Height = max(getHeight(x.Left), getHeight(x.Right)) + 1
	return x
}

func balance[T any](node *TreeNode[T]) *TreeNode[T] {
	balanceFactor := getBalanceFactor(node)

	if balanceFactor > 1 {
		if getBalanceFactor(node.Left) < 0 {
			node.Left = leftRotate(node.Left) // Left-Right case
		}
		return rightRotate(node) // Left-Left case
	}

	if balanceFactor < -1 {
		if getBalanceFactor(node.Right) > 0 {
			node.Right = rightRotate(node.Right) // Right-Left case
		}
		return leftRotate(node) // Right-Right case
	}

	return node
}

func getHeight[T any](node *TreeNode[T]) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func getBalanceFactor[T any](node *TreeNode[T]) int {
	if node == nil {
		return 0
	}
	return getHeight(node.Left) - getHeight(node.Right)
}

func printNode[T any](node *TreeNode[T], prefix string, isLeft bool, noParents bool) {
	if node != nil {
		// Print the current node value
		fmt.Printf("%s", prefix)
		if isLeft {
			fmt.Printf("├── ")
			prefix += "│   "
		} else {
			fmt.Printf("└── ")
			prefix += "    "
		}

		if node.Parent == nil || noParents {
			fmt.Println(node.Value)
		} else {
			fmt.Printf("%v(%v)\n", node.Value, node.Parent.Value)
		}
		// Print the parent of the current node
		// if node.Parent != nil {
		// 	fmt.Printf("%sParent: %v\n", prefix, node.Parent.Value)
		// }

		printNode(node.Left, prefix, true, noParents)
		printNode(node.Right, prefix, false, noParents)
	}
}
