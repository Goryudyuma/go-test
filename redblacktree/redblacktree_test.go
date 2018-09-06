package redblacktree

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(9)
	tree = tree.Push(8)
	tree = tree.Push(7)
	tree = tree.Push(6)
	tree = tree.Push(5)
	tree = tree.Push(4)
	tree = tree.Push(3)
	tree = tree.Push(2)
	tree = tree.Push(1)
	tree = tree.Push(0)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.toString()) != `0123456789` {
		t.Error("error anything")
	}
}

func Test2(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(0)
	tree = tree.Push(1)
	tree = tree.Push(2)
	tree = tree.Push(3)
	tree = tree.Push(4)
	tree = tree.Push(5)
	tree = tree.Push(6)
	tree = tree.Push(7)
	tree = tree.Push(8)
	tree = tree.Push(9)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.toString()) != `0123456789` {
		t.Error("error anything")
	}
}
