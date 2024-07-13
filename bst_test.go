package main

import (
	"testing"
)

func TestInsertOnEmpty(t *testing.T) {
	var tree Node[int] = &Tree[int]{5, nil, nil}

	if tree.Get() != 5 {
		t.Errorf("Didn't find expected value")
	}
}

func TestInsertOnNonEmpty(t *testing.T) {
	var tree Node[int] = &Tree[int]{5, nil, nil}

	tree.Insert(7)
	if tree.Right().Get() != 7 {
		t.Errorf("Didn't find expected value")
	}

	tree.Insert(1)
	if tree.Left().Get() != 1 {
		t.Errorf("Didn't find expected value")
	}

	tree.Insert(5)
	if tree.Left().Right().Get() != 5 {
		t.Errorf("Didn't find expected value")
	}
}
