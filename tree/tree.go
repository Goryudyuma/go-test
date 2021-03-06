package tree

import (
	"encoding/json"
)

// Tree Library
type Tree struct {
	parent *Tree
	left   *Tree
	right  *Tree
	value  interface{}
}

// UnmarshalJSON
func (tree *Tree) UnmarshalJSON(data []byte) error {
	var v map[string]json.RawMessage
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v["right"] != nil {
		if err := json.Unmarshal(v["right"], &tree.right); err != nil {
			return err
		}
	}
	if v["left"] != nil {
		if err := json.Unmarshal(v["left"], &tree.left); err != nil {
			return err
		}
	}
	if v["value"] != nil {
		if err := json.Unmarshal(v["value"], &tree.value); err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON
func (tree *Tree) MarshalJSON() ([]byte, error) {
	ret := []byte("{")
	if tree.left != nil {
		leftJSON, err := json.Marshal(tree.left)
		if err != nil {
			return ret, err
		}
		if len(ret) != 1 {
			ret = append(ret, ',')
		}
		ret = append(ret, []byte(`"left":`)...)
		ret = append(ret, leftJSON...)
	}
	if tree.right != nil {
		rightJSON, err := json.Marshal(tree.right)
		if err != nil {
			return ret, err
		}
		if len(ret) != 1 {
			ret = append(ret, ',')
		}
		ret = append(ret, []byte(`"right":`)...)
		ret = append(ret, rightJSON...)
	}
	if tree.value != nil {
		valueJSON, err := json.Marshal(tree.value)
		if err != nil {
			return ret, err
		}
		if len(ret) != 1 {
			ret = append(ret, ',')
		}
		ret = append(ret, []byte(`"value":`)...)
		ret = append(ret, valueJSON...)
	}

	return append(ret, '}'), nil
}
