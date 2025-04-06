package tree

type Tree[N any] interface {
	GetRoot() *TreeNode[N]
	Insert(parent *TreeNode[N], value any) error
}

type TreeNode[M any] interface {
	GetValue() any
}
