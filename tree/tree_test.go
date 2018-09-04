package tree

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	eq, err := AreEqualJSON(string(res), `{}`)
	if err != nil {
		t.Error(err)
	}
	if !eq {
		t.Error("Error empty Tree to JSON")
	}
}

func TestMarshalJSONValueIsInt(t *testing.T) {
	tree := &Tree{value: 1}
	res, err := json.Marshal(tree)
	if err != nil {
		t.Error(err)
	}
	eq, err := AreEqualJSON(string(res), `{"value":1}`)
	if err != nil {
		t.Error(err)
	}
	if !eq {
		t.Errorf("Error JSON %v", string(res))
	}
}

func TestMarshalJSONValueIsInt64(t *testing.T) {
	tree := &Tree{value: int64(1)}
	res, err := json.Marshal(tree)
	if err != nil {
		t.Error(err)
	}
	eq, err := AreEqualJSON(string(res), `{"value":1}`)
	if err != nil {
		t.Error(err)
	}
	if !eq {
		t.Errorf("Error JSON %v", string(res))
	}
}

func TestMarshalJSONValueIsString(t *testing.T) {
	tree := &Tree{value: "abc"}
	res, err := json.Marshal(tree)
	if err != nil {
		t.Error(err)
	}
	eq, err := AreEqualJSON(string(res), `{"value":"abc"}`)
	if err != nil {
		t.Error(err)
	}
	if !eq {
		t.Errorf("Error JSON %v", string(res))
	}
}

func TestMarshalJSONExistLeft(t *testing.T) {
	tree := &Tree{left: &Tree{}}
	res, err := json.Marshal(tree)
	if err != nil {
		t.Error(err)
	}
	eq, err := AreEqualJSON(string(res), `{"left":{}}`)
	if err != nil {
		t.Error(err)
	}
	if !eq {
		t.Errorf("Error JSON %v", string(res))
	}
}

func TestUnmarshalJSONEmpty(t *testing.T) {
	var tree *Tree
	err := json.Unmarshal([]byte(`{}`), &tree)
	if err != nil {
		t.Error(err)
	}
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

func TestUnmarshalJSONValueIsInt(t *testing.T) {
	var tree *Tree
	err := json.Unmarshal([]byte(`{"value":1}`), &tree)
	if err != nil {
		t.Error(err)
	}
	if tree.parent != nil {
		t.Error("error tree parent")
	}
	if tree.right != nil {
		t.Error("error tree right")
	}
	if tree.left != nil {
		t.Error("error tree left")
	}
	if tree.value.(float64) != float64(1) {
		t.Error("error tree value")
	}
}

func TestUnmarshalJSONExistLeftValue(t *testing.T) {
	var tree *Tree
	err := json.Unmarshal([]byte(`{"left":{"value":1}}`), &tree)
	if err != nil {
		t.Error(err)
	}
	if tree.left.value != float64(1) {
		t.Error("error tree left value")
	}
}

func TestUnmarshalJSONExistLeftLeftValue(t *testing.T) {
	var tree *Tree
	err := json.Unmarshal([]byte(`{"left":{"left":{"value":1}}}`), &tree)
	if err != nil {
		t.Error(err)
	}
	if tree.left.left.value != float64(1) {
		t.Error("error tree left left value")
	}
}

func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}
