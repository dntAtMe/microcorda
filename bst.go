package main

import (
	"cmp"
	"fmt"
)

type Node[T cmp.Ordered] interface {
	Insert(value T)
	Traverse(f func(T))
	Get() T
	Left() Node[T]
	Right() Node[T]
}

type Tree[T cmp.Ordered] struct {
	value T
	left  Node[T]
	right Node[T]
}

func (tree *Tree[T]) Get() T {
	return tree.value
}

func (tree *Tree[T]) Left() Node[T] {
	return tree.left
}

func (tree *Tree[T]) Right() Node[T] {
	return tree.right
}

func (tree *Tree[T]) Insert(value T) {
	if value <= tree.value {
		if tree.left != nil {
			tree.left.Insert(value)
		} else {
			tree.left = &Tree[T]{value, nil, nil}
		}
	} else {
		if tree.right != nil {
			tree.right.Insert(value)
		} else {
			tree.right = &Tree[T]{value, nil, nil}
		}
	}
}

func (tree *Tree[T]) Traverse(f func(T)) {
	if tree.left != nil {
		tree.left.Traverse(f)
	}
	f(tree.value)
	if tree.right != nil {
		tree.right.Traverse(f)
	}
}

func main() {
	var tree = Tree[int]{5, nil, nil}
	tree.Traverse(func(i int) { fmt.Println(i) })

	fmt.Printf("-- %d\n", tree.Get())
}
