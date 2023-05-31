# Tree
- NodeとEdgeで構成されているデータ構造
- 子ノードがないノードをleafと呼ぶ
- 同じdepthのノードをsiblingsだと呼ぶ
- File Systemでよく使われる
```go
type TreeNode[T any] struct {
	Value T
	Children []*TreeNode[T]
}
```

## Traversal(巡回)
### Depth first traversal, Depth first search(深さ優先探索, DFS)
- Inorder(Left -> Right)
- Preorder(現 -> 子)
- Postorder(子 -> 現)
### Breadth first traversal, Breadth first search(幅優先探索, BFS)

## Binary Tree
- 子ノードを二つのみ持つTree
- **Binary Search Tree**, <- **一番重要**
- Heap Tree

```go
type TreeNode[T any] struct {
	Value T
	Left *TreeNode[T]
    Right *TreeNode[T]
}
```
