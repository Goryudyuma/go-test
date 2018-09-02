package tree

import "testing"

func TestStructCheck(t *testing.T) {
	tree := Tree{}
	if tree.parent != nil {
		t.Error("error tree parent")
	}
	if tree.right != nil {
		t.Error("error tree right")
	}
	if tree.left != nil {
		t.Error("error tree left")
	}
	if tree.value != nil {
		t.Error("error tree left")
	}
}
