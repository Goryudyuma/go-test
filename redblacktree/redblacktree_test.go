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

func Test8(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(0)
	tree = tree.Push(1)
	tree = tree.Push(2)
	tree = tree.Push(3)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.String()) != `0123` {
		t.Error("error anything")
	}
}

func Test9(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(0)
	tree = tree.Push(3)
	tree = tree.Push(1)
	tree = tree.Push(2)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.String()) != `0123` {
		t.Error("error anything")
	}
}

func Test10(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(3)
	tree = tree.Push(0)
	tree = tree.Push(1)
	tree = tree.Push(2)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.String()) != `0123` {
		t.Error("error anything")
	}
}

func Test11(t *testing.T) {
	tree := new(RedBlackTree)
	tree = tree.Push(3)
	tree = tree.Push(2)
	tree = tree.Push(0)
	tree = tree.Push(1)
	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.String()) != `0123` {
		t.Error("error anything")
	}
}

func Test13(t *testing.T) {
	tree := new(RedBlackTree)
	testData := []int64{
		1324408774600613285,
		1508371146504897700,
		4605201821391833555,
		4169652878687413376,
		781282587162764562,
		4553344659976848263,
		4555980653633010339,
		6067279116940164439,
		7020951801917454432,
	}
	for _, v := range testData {
		tree = tree.Push(v)
	}

	testDataString := make([]string, 9)
	for j := 0; j < 9; j++ {
		testDataString[j] = strconv.FormatInt(testData[j], 10)
	}

	r := strings.NewReplacer("(", "", ")", "")
	if r.Replace(tree.String()) != `6052018213918335557812825871627645621324408774600613285150837114650489770041696528786874133764553344659976848263455598065363301033960672791169401644397020951801917454432` {
		t.Errorf("%s %s\n", r.Replace(tree.String()), strings.Join(testDataString, ""))
		t.Error("error anything")
	}
}

func TestRandomData(t *testing.T) {
	r := strings.NewReplacer("(", "", ")", "")
	testCaseNum := 9
	for i := 0; i < 10000; i++ {
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
			spew.Dump(testData)
			fmt.Println(strings.Join(testDataString, ""))
		}
	}
}
