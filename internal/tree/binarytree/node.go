package binarytree

type TreeNode[T any] struct {
	Value  T
	Left   *TreeNode[T]
	Right  *TreeNode[T]
	Parent *TreeNode[T]
	Height int
}
