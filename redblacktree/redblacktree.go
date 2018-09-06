package redblacktree

import (
	"strconv"
)

// http://wwwa.pikara.ne.jp/okojisan/rb-tree/index.html

const (
	red = iota
	black
)

type RedBlackTree struct {
	Parent   *RedBlackTree
	Right    *RedBlackTree
	Left     *RedBlackTree
	Value    *int64
	redblack int
}

func (tree *RedBlackTree) isRoot() bool {
	return tree.Parent == nil
}

func (tree *RedBlackTree) root() *RedBlackTree {
	if tree.isRoot() {
		return tree
	}
	return tree.Parent.root()
}

// Color 色を返す
func (tree *RedBlackTree) Color() int {
	return tree.redblack
}

// Push 挿入処理
func (tree *RedBlackTree) Push(i int64) *RedBlackTree {
	if tree.Value == nil {
		tree.Value = &i
		tree.redblack = red

		tree.rotate()
	} else {
		if *tree.Value < i {
			if tree.Right == nil {
				tree.Right = new(RedBlackTree)
				tree.Right.Parent = tree
			}
			tree.Right.Push(i)
		} else {
			if tree.Left == nil {
				tree.Left = new(RedBlackTree)
				tree.Left.Parent = tree
			}
			tree.Left.Push(i)
		}
	}
	return tree.root()
}

func (tree *RedBlackTree) rotate() {
	// 自分が根なら、黒に変えて終わる
	if tree.isRoot() {
		tree.redblack = black
		return
	}

	// 親が黒なら無条件で終わる
	if tree.Parent.Color() == black {
		return
	}

	if tree.Parent.Parent.Left == tree.Parent {
		// Lパターン
		if tree.Parent.Left == tree {
			// LLパターン
			tree.Parent.rotateLL()
		} else {
			// LRパターン
			tree.rotateLR()
		}
	} else {
		// Rパターン
		if tree.Parent.Left == tree {
			// RLパターン
			tree.rotateRL()
		} else {
			// RRパターン
			tree.Parent.rotateRR()
		}
	}
}

func (tree *RedBlackTree) rotateLL() {
	if tree.Parent.Parent != nil {
		if tree.Parent.Parent.Left == tree.Parent {
			tree.Parent.Parent.Left,
				tree.Right, tree.Parent.Left, tree.Parent.Parent, tree.Parent =
				tree,
				tree.Parent, tree.Right, tree, tree.Parent.Parent
		} else {
			tree.Parent.Parent.Right,
				tree.Right, tree.Parent.Left, tree.Parent.Parent, tree.Parent =
				tree,
				tree.Parent, tree.Right, tree, tree.Parent.Parent
		}
	} else {
		tree.Right, tree.Parent.Left, tree.Parent.Parent, tree.Parent =
			tree.Parent, tree.Right, tree, tree.Parent.Parent
	}

	tree.redblack = red
	tree.Left.redblack = black
	tree.Right.redblack = black

	tree.rotate()
}

func (tree *RedBlackTree) rotateLR() {

	if tree.Parent.Parent.Parent != nil {
		if tree.Parent.Parent.Parent.Left == tree.Parent.Parent {
			tree.Parent.Parent.Parent.Left,
				tree.Left, tree.Right, tree.Parent,
				tree.Parent.Right, tree.Parent.Parent.Left,
				tree.Parent.Parent, tree.Parent.Parent.Parent =
				tree,
				tree.Parent, tree.Parent.Parent, tree.Parent.Parent.Parent,
				tree.Left, tree.Right,
				tree, tree
		} else {
			tree.Parent.Parent.Parent.Right,
				tree.Left, tree.Right, tree.Parent,
				tree.Parent.Right, tree.Parent.Parent.Left,
				tree.Parent.Parent, tree.Parent.Parent.Parent =
				tree,
				tree.Parent, tree.Parent.Parent, tree.Parent.Parent.Parent,
				tree.Left, tree.Right,
				tree, tree
		}
	} else {
		tree.Left, tree.Right, tree.Parent,
			tree.Parent.Right, tree.Parent.Parent.Left,
			tree.Parent.Parent, tree.Parent.Parent.Parent =
			tree.Parent, tree.Parent.Parent, tree.Parent.Parent.Parent,
			tree.Left, tree.Right,
			tree, tree
	}

	tree.redblack = red
	tree.Left.redblack = black
	tree.Right.redblack = black

	tree.rotate()
}

func (tree *RedBlackTree) rotateRL() {
	if tree.Parent.Parent.Parent != nil {
		if tree.Parent.Parent.Parent.Left == tree.Parent.Parent {
			tree.Parent.Parent.Parent.Left,
				tree.Left, tree.Right, tree.Parent,
				tree.Parent.Left, tree.Parent.Parent.Right,
				tree.Parent.Parent, tree.Parent.Parent.Parent =
				tree,
				tree.Parent.Parent, tree.Parent, tree.Parent.Parent.Parent,
				tree.Right, tree.Left,
				tree, tree
		} else {
			tree.Parent.Parent.Parent.Right,
				tree.Left, tree.Right, tree.Parent,
				tree.Parent.Left, tree.Parent.Parent.Right,
				tree.Parent.Parent, tree.Parent.Parent.Parent =
				tree,
				tree.Parent.Parent, tree.Parent, tree.Parent.Parent.Parent,
				tree.Right, tree.Left,
				tree, tree
		}
	} else {
		tree.Left, tree.Right, tree.Parent,
			tree.Parent.Left, tree.Parent.Parent.Right,
			tree.Parent.Parent, tree.Parent.Parent.Parent =
			tree.Parent.Parent, tree.Parent, tree.Parent.Parent.Parent,
			tree.Right, tree.Left,
			tree, tree
	}

	tree.redblack = red
	tree.Left.redblack = black
	tree.Right.redblack = black

	tree.rotate()
}

func (tree *RedBlackTree) rotateRR() {
	if tree.Parent.Parent != nil {
		if tree.Parent.Parent.Left == tree.Parent {
			tree.Parent.Parent.Left,
				tree.Left, tree.Parent, tree.Parent.Right, tree.Parent.Parent =
				tree,
				tree.Parent, tree.Parent.Parent, tree.Left, tree
		} else {
			tree.Parent.Parent.Right,
				tree.Left, tree.Parent, tree.Parent.Right, tree.Parent.Parent =
				tree,
				tree.Parent, tree.Parent.Parent, tree.Left, tree
		}
	} else {
		tree.Left, tree.Parent, tree.Parent.Right, tree.Parent.Parent =
			tree.Parent, tree.Parent.Parent, tree.Left, tree
	}

	tree.redblack = red
	tree.Left.redblack = black
	tree.Right.redblack = black

	tree.rotate()
}

func (tree *RedBlackTree) toString() (ret string) {
	if tree == nil {
		return ""
	}
	ret = "("
	ret += tree.Left.toString()
	ret += strconv.FormatInt(*tree.Value, 10)
	ret += tree.Right.toString()

	ret += ")"
	return
}
