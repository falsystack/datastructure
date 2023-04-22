package main

import "fmt"

// Node Single linked list
type Node[T any] struct {
	next *Node[T]
	val  T
}

// DoubleNode Double linked list
type DoubleNode[T any] struct {
	next *DoubleNode[T]
	prev *DoubleNode[T]
	val  T
}

func Append[T int](node *Node[T], next *Node[T]) *Node[T] {
	node.next = next
	return next
}

func main() {
	root := &Node[int]{nil, 10}
	//root.next = &Node[int]{nil, 20}
	//root.next.next = &Node[int]{nil, 30}
	node := root
	for i := 0; i < 3; i++ {
		node = Append(node, &Node[int]{nil, (i + 2) * 10})
	}

	for n := root; n != nil; n = n.next {
		fmt.Printf("node val : %d\n", n.val)
	}

	fmt.Println()
	fmt.Println()

	nodeTwenteen := root.next         // 20
	nodeThirteen := nodeTwenteen.next // 30

	Append(nodeTwenteen, &Node[int]{nodeThirteen, 100})

	for n := root; n != nil; n = n.next {
		fmt.Printf("node val : %d\n", n.val)
	}
}
