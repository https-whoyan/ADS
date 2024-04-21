package MyHeap

import (
	"errors"
	"fmt"
	"github.com/https-whoyan/ADS/tree/main/Assigment2/pkg/Lists/MyArrayList"
	"reflect"
)

type MyHeap struct {
	data     *MyArrayList.MyArrayList
	size     int
	heapType reflect.Type
	cmp      func(x, y any) bool
}

var incorrectElementType = errors.New("incorrect added element type")
var dataIsEmptyErr = errors.New("data is empty, cannot pick any element")
var incorrectLessFunction = errors.New("incorrect less function. recreate your minHeap")

func NewMyMinHeap(cmp func(x, y any) bool) *MyHeap {
	return &MyHeap{
		data: MyArrayList.NewArrayList(),
		size: 0,
		cmp:  cmp,
	}
}

// minHeap array represent of binary tree functions
func parent(i int) int {
	return i / 2
}

func leftChild(i int) int {
	return 2 * i
}

func rightChild(i int) int {
	return 2*i + 1
}

func (h *MyHeap) validateType(el any) bool {
	if h.heapType == nil {
		return true
	}
	return !(reflect.ValueOf(el).Type() != h.heapType)
}

func (h *MyHeap) standListType(refEl any) {
	h.heapType = reflect.ValueOf(refEl).Type()
}

func (h *MyHeap) needStandListType() bool {
	return h.size == 0 && h.heapType == nil
}

// cmp check
func isCorrectLessFunction(x any, cmp func(i, j any) bool) bool {
	ok := true
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	cmp(x, x)
	return ok
}

func (h *MyHeap) initRoot(root any) error {
	// Fill the 0-index element in arrayList
	// (Heap work from 1-th index)
	err := h.data.AddLast(root)
	if err != nil {
		return err
	}
	err = h.data.AddLast(root)
	h.size++
	return err
}

func (h *MyHeap) heapify(i int) {
	if i == 1 {
		return
	}

	parentIndex := parent(i)
	parentNode, _ := h.data.Get(parentIndex)
	currNode, _ := h.data.Get(i)
	if h.cmp(currNode.(any), parentNode.(any)) {
		h.data.Set(currNode, parentIndex)
		h.data.Set(parentNode, i)
		h.heapify(parentIndex)
	}
}

func (h *MyHeap) Add(newEl any) error {
	if h.needStandListType() {
		h.standListType(newEl)
	}
	if !isCorrectLessFunction(newEl, h.cmp) {
		return incorrectLessFunction
	}
	if ok := h.validateType(newEl); !ok {
		return incorrectElementType
	}

	if h.data.Size() == 0 {
		return h.initRoot(newEl)
	}
	h.data.AddLast(newEl)
	h.size++
	h.heapify(h.size - 1)
	return nil
}

func (h *MyHeap) GetType() reflect.Type {
	return h.heapType
}

func (h *MyHeap) Size() int {
	return h.size
}

func (h *MyHeap) GetRoot() (any, error) {
	if h.size == 0 {
		return nil, dataIsEmptyErr
	}
	root, _ := h.data.Get(1)
	if !isCorrectLessFunction(root, h.cmp) {
		return nil, incorrectLessFunction
	}
	return root, nil
}

func (h *MyHeap) down(i int) {
	minChild, _ := h.data.Get(0)
	minChildIndex := -1
	leftChildI := leftChild(i)
	rightChildI := rightChild(i)
	if leftChildI <= h.size {
		minChild, _ = h.data.Get(leftChildI)
		minChildIndex = leftChildI
	}
	if rightChildI <= h.size {
		if minChildIndex == -1 {
			minChild, _ = h.data.Get(rightChildI)
			minChildIndex = rightChildI
		} else {
			lVal, _ := h.data.Get(leftChildI)
			rVal, _ := h.data.Get(rightChildI)
			if h.cmp(rVal.(any), lVal.(any)) {
				minChild = rVal
				minChildIndex = rightChildI
			}
		}
	}

	currVal, _ := h.data.Get(i)
	if minChildIndex == -1 {
		return
	}
	if !h.cmp(currVal.(any), minChild.(any)) {
		h.data.Set(currVal, minChildIndex)
		h.data.Set(minChild, i)
		h.down(minChildIndex)
	}
}

func (h *MyHeap) ExtractRoot() (any, error) {
	root, err := h.GetRoot()
	if err != nil {
		return nil, err
	}
	lastEl, _ := h.data.GetLast()
	h.data.RemoveLast()
	h.size--
	if h.size == 0 {
		return root, nil
	}
	h.data.Set(lastEl, 1)
	h.down(1)

	return root, nil
}

func (h *MyHeap) printData() {
	fmt.Println(h.data.ToArray())
}
