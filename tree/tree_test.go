package tree

import "testing"

func TestStructCheck1(t *testing.T) {
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
		t.Error("error tree value")
	}
}

func TestStructCheckValueIsInt(t *testing.T) {
	tree := Tree{value: 1}
	if tree.parent != nil {
		t.Error("error tree parent")
	}
	if tree.right != nil {
		t.Error("error tree right")
	}
	if tree.left != nil {
		t.Error("error tree left")
	}
	if tree.value.(int) != 1 {
		t.Error("error tree value")
	}
}

func TestStructCheckValueIsInt64(t *testing.T) {
	tree := Tree{value: int64(1)}
	if tree.parent != nil {
		t.Error("error tree parent")
	}
	if tree.right != nil {
		t.Error("error tree right")
	}
	if tree.left != nil {
		t.Error("error tree left")
	}
	if tree.value.(int64) != 1 {
		t.Error("error tree value")
	}
}
