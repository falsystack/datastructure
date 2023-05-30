package tree

type TreeNode[T any] struct {
	Value    T
	Children []*TreeNode[T]
}

func (t *TreeNode[T]) Add(val T) *TreeNode[T] {
	n := &TreeNode[T]{
		Value: val,
	}

	t.Children = append(t.Children, n)
	return n
}

func (t *TreeNode[T]) DFS(fn func(val T)) {
	stack := make([]*TreeNode[T], 0)
	stack = append(stack, t)

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fn(last.Value)

		stack = append(stack, last.Children...)
	}
}

func (t *TreeNode[T]) Preorder(fn func(val T)) {
	if t == nil {
		return
	}
	fn(t.Value)

	for _, child := range t.Children {
		child.Preorder(fn)
	}
}

func (t *TreeNode[T]) Postorder(fn func(val T)) {
	if t == nil {
		return
	}

	for _, child := range t.Children {
		child.Postorder(fn)
	}

	fn(t.Value)
}

func (t *TreeNode[T]) BFS(fn func(val T)) {
	queue := make([]*TreeNode[T], 0)
	queue = append(queue, t)

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		fn(front.Value)

		for _, child := range front.Children {
			queue = append(queue, child)
		}
	}
}
