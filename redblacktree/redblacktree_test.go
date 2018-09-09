package redblacktree

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
	if r.Replace(tree.String()) != `0123456789` {
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
	if r.Replace(tree.String()) != `0123456789` {
		t.Error("error anything")
	}
}

func Test3(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(0)
	tree = tree.Push(2)
	tree = tree.Push(1)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.String()) != `012` {
		t.Error("error anything")
	}
}

func Test4(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(2)
	tree = tree.Push(0)
	tree = tree.Push(1)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.String()) != `012` {
		t.Error("error anything")
	}
}

func Test5(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(2)
	tree = tree.Push(1)
	tree = tree.Push(0)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.String()) != `012` {
		t.Error("error anything")
	}
}

func Test6(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(1)
	tree = tree.Push(2)
	tree = tree.Push(0)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.String()) != `012` {
		t.Error("error anything")
	}
}

func Test7(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(1)
	tree = tree.Push(0)
	tree = tree.Push(2)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.String()) != `012` {
		t.Error("error anything")
	}
}

func TestRandomData(t *testing.T) {
	r := strings.NewReplacer("(", "", ")", "")
	testCaseNum := 10
	for i := 0; i < 10; i++ {
		tree := new(RedBlackTree)
		testData := make([]int64, 0, testCaseNum)
		for j := 0; j < testCaseNum; j++ {
			n := rand.Int63()
			testData = append(testData, n)
			tree = tree.Push(n)
		}
		sort.Slice(testData,
			func(i, j int) bool {
				return testData[i] < testData[j]
			},
		)
		testDataString := make([]string, testCaseNum)
		for j := 0; j < testCaseNum; j++ {
			testDataString[j] = strconv.FormatInt(testData[j], 10)
		}
		if r.Replace(tree.String()) != strings.Join(testDataString, "") {
			t.Error("error anything")
		}
	}
}
