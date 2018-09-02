package tree

// Tree Library
type Tree struct {
	parent *Tree
	left   *Tree
	right  *Tree
	value  interface{}
}
