package tree

import (
	"encoding/json"
	"testing"
)

func TestStructCheck1(t *testing.T) {
	tree := &Tree{}
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
	tree := &Tree{value: 1}
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
	tree := &Tree{value: int64(1)}
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

func TestMarshalJSONEmpty(t *testing.T) {
	tree := &Tree{}
	res, err := json.Marshal(tree)
	if err != nil {
		t.Error(err)
	}
	if string(res) != "{}" {
		t.Error("Error empty Tree to JSON")
	}
}

func TestMarshalJSONValueIsInt(t *testing.T) {
	tree := &Tree{value: 1}
	res, err := json.Marshal(tree)
	if err != nil {
		t.Error(err)
	}
	if string(res) != "{\"value\":1}" {
		t.Errorf("Error JSON %v", string(res))
	}
}

func TestMarshalJSONExistLeft(t *testing.T) {
	tree := &Tree{left: &Tree{}}
	res, err := json.Marshal(tree)
	if err != nil {
		t.Error(err)
	}
	if string(res) != "{\"left\":{}}" {
		t.Errorf("Error JSON %v", string(res))
	}
}
