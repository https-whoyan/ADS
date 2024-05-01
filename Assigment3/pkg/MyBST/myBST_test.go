package MyBST

import (
	"errors"
	"fmt"
	rnd "github.com/https-whoyan/ADS/Assigment3/pkg/Random"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

const testingElementsCount = 1000

var notExceptedErr = errors.New("not excepted key contains err, err")
var panicCatch = errors.New("catch panic")

func TestMyBST_Put(t *testing.T) {
	t.Run("Put, with int key and struct key", func(t *testing.T) {
		bst := NewMyBST[int, rnd.TestStruct]()
		defer func() {
			if r := recover(); r != nil {
				assert.Error(t, panicCatch)
			}
		}()
		currSlice := make([]TraversalNode[int, rnd.TestStruct], 0, testingElementsCount)
		for i := 1; i <= testingElementsCount; i++ {
			key := i
			value := rnd.GetStruct(i)
			currSlice = append(currSlice, TraversalNode[int, rnd.TestStruct]{
				key:   key,
				value: *value,
			})
			err := bst.Put(key, *value)
			if err != nil {
				assertionErr := errors.Join(notExceptedErr, err)
				assert.Error(t, assertionErr)
			}
		}
		assert.True(t, true)
	})
	t.Run("Put, exist key (excepted error)", func(t *testing.T) {
		bst := NewMyBST[int, int]()
		err := bst.Put(1, 1)
		if err != nil {
			assertionErr := errors.Join(notExceptedErr, err)
			assert.Error(t, assertionErr)
		}
		err = bst.Put(1, 8)
		exceptedErr := errors.New(fmt.Sprintf("Key %v is already exist!", 1))
		assert.Equal(t, exceptedErr, err)
	})
}

func TestMyBST_Get(t *testing.T) {
	t.Run("Get, with int key and values", func(t *testing.T) {
		bst := NewMyBST[int, int]()
		defer func() {
			if r := recover(); r != nil {
				assert.Error(t, panicCatch)
			}
		}()
		mp := make(map[int]int)
		for i := 1; i <= testingElementsCount; i++ {
			key := i
			value := rnd.RandomFormula(i)
			mp[key] = value
			err := bst.Put(key, value)
			if err != nil {
				assertionErr := errors.Join(notExceptedErr, err)
				assert.Error(t, assertionErr)
			}
		}
		actualMp := make(map[int]int)
		for i := 1; i <= testingElementsCount; i++ {
			bstVal, err := bst.Get(i)
			if err != nil {
				assertionErr := errors.Join(notExceptedErr, err)
				assert.Error(t, assertionErr)
			}
			actualMp[i] = bstVal
		}
		assert.Equal(t, mp, actualMp)
	})
}

func TestMyBST_Delete(t *testing.T) {
	t.Run("Delete, with int key and values", func(t *testing.T) {
		bst := NewMyBST[int, int]()

		// this function declarated to construct norlmal leveled BST,
		// (I'm so lazy to write balancing tree
		var getNormalOrderToAdd func(l, r int) []int
		getNormalOrderToAdd = func(l, r int) []int {
			if l > r {
				return []int{}
			}
			if l == r {
				return []int{l}
			}
			mid := (l + r) / 2
			currAns := []int{mid}
			leftAns := getNormalOrderToAdd(l, mid-1)
			rightAns := getNormalOrderToAdd(mid+1, r)
			return append(append(leftAns, currAns...), rightAns...)
		}

		normalAddOrder := getNormalOrderToAdd(1, testingElementsCount)
		for _, i := range normalAddOrder {
			value := rnd.RandomFormula(i)
			err := bst.Put(i, value)
			if err != nil {
				assertionErr := errors.Join(notExceptedErr, err)
				assert.Error(t, assertionErr)
			}
		}
		currArr := make(map[int]int)
		var exceptedInOrderTraversals [][]TraversalNode[int, int]
		var actualInOrderTraversals [][]TraversalNode[int, int]
		for i := 1; i <= testingElementsCount; i++ {
			currArr[i] = rnd.RandomFormula(i)
		}

		getCurrentTraversal := func(mp map[int]int) []TraversalNode[int, int] {
			arr := make([]TraversalNode[int, int], 0, len(mp))
			for key, value := range mp {
				arr = append(arr, TraversalNode[int, int]{
					key:   key,
					value: value,
				})
			}
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].key < arr[j].key
			})
			return arr
		}
		// Ok, created.
		// Let's delete keys from 1 to testingElementsCount
		for i := 1; i <= testingElementsCount; i++ {
			bst.Delete(i)
			delete(currArr, i)
			currentTraversal := bst.InOrderTraversal()
			actualInOrderTraversals = append(actualInOrderTraversals, currentTraversal)
			exceptedTraversal := getCurrentTraversal(currArr)
			exceptedInOrderTraversals = append(exceptedInOrderTraversals, exceptedTraversal)
		}
		assert.Equal(t, exceptedInOrderTraversals, exceptedInOrderTraversals)
	})
}
