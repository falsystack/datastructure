package tree

import (
	"fmt"
	"strings"
	"testing"
)

func TestTreeNode_Add(t *testing.T) {
	root := &TreeNode[string]{
		Value: "A",
	}

	b := root.Add("B")
	root.Add("C")
	d := root.Add("D")

	b.Add("E")
	b.Add("F")

	d.Add("G")

	var sb strings.Builder
	root.Preorder(func(val string) {
		sb.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(sb.String())

	sb.Reset()
	root.Postorder(func(val string) {
		sb.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(sb.String())
	/*
	   Preorder -> tree_test.go:27: A - B - E - F - C - D - G
	   Postorder -> tree_test.go:33: E - F - B - C - G - D - A
	*/

	sb.Reset()
	root.BFS(func(val string) {
		sb.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(sb.String())
	// BFS -> tree_test.go:43: A - B - C - D - E - F - G

	sb.Reset()
	root.DFS(func(val string) {
		sb.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(sb.String())
}
