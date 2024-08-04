package binary_tree

type TreeNode[T comparable] struct {
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type BinaryTree[T comparable] struct {
	Root *TreeNode[T]
}

func (t *BinaryTree[T]) Insert(a T) {

}
