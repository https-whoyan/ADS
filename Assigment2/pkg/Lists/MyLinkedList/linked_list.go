package MyLinkedList

import (
	"errors"
	"reflect"
	"sort"
)

type MyLinkedList struct {
	head     *node
	tail     *node
	listType reflect.Type
	size     int
}

type node struct {
	val  any
	next *node
	prev *node
}

var incorrectIndexErr = errors.New("incorrect index")
var incorrectElementType = errors.New("incorrect added element type")
var dataIsEmptyErr = errors.New("data is empty, cannot pick/delete any element")
var incorrectLessFunction = errors.New("incorrect less function")

func NewLinkedList() *MyLinkedList {
	return &MyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// true, if l.listType == el.type, or l.listType == nil
func (l *MyLinkedList) validateType(el any) bool {
	if l.listType == nil {
		return true
	}
	return !(reflect.ValueOf(el).Type() != l.listType)
}

func (l *MyLinkedList) standListType(refEl any) {
	l.listType = reflect.ValueOf(refEl).Type()
}

func (l *MyLinkedList) needStandListType() bool {
	return l.size == 0 && l.listType == nil
}

func (l *MyLinkedList) initNewLinkedListWithOneNode(newEl any) error {
	if !(l.validateType(newEl)) {
		return incorrectElementType
	}
	if l.needStandListType() {
		l.standListType(newEl)
	}

	newNode := &node{
		val: newEl,
	}
	l.head = newNode
	l.tail = newNode
	l.size = 1

	return nil
}

func (l *MyLinkedList) Add(newEl any, indexToAdd int) error {
	if indexToAdd > l.size || indexToAdd < 0 {
		return incorrectIndexErr
	}
	if !(l.validateType(newEl)) {
		return incorrectElementType
	}
	if l.needStandListType() {
		l.standListType(newEl)
	}

	if l.size == 0 {
		return l.initNewLinkedListWithOneNode(newEl)
	}
	currDataLen := l.size
	if indexToAdd == currDataLen {
		return l.AddLast(newEl)
	}
	if indexToAdd == 0 {
		return l.AddFirst(newEl)
	}

	copyHead := l.head
	for index := 0; index <= indexToAdd-1; index++ {
		copyHead = copyHead.next
	}

	newNode := &node{
		val:  newEl,
		next: copyHead.next,
		prev: copyHead,
	}
	copyHead.next.prev = newNode
	copyHead.next = copyHead
	l.size++

	return nil
}

func (l *MyLinkedList) AddLast(newEl any) error {
	if ok := l.validateType(newEl); !ok {
		return incorrectElementType
	}
	if l.needStandListType() {
		l.standListType(newEl)
	}

	if l.size == 0 {
		return l.initNewLinkedListWithOneNode(newEl)
	}

	newNode := &node{
		val:  newEl,
		prev: l.tail,
	}
	l.tail.next = newNode
	l.tail = newNode
	l.size++

	return nil
}

func (l *MyLinkedList) AddFirst(newEl any) error {
	if ok := l.validateType(newEl); !ok {
		return incorrectElementType
	}
	if l.needStandListType() {
		l.standListType(newEl)
	}

	if l.size == 0 {
		return l.initNewLinkedListWithOneNode(newEl)
	}

	newNode := &node{
		val:  newEl,
		next: l.head,
	}
	l.head.prev = newNode
	l.head = newNode
	l.size++

	return nil
}

// Get functions
func (l *MyLinkedList) Get(index int) (any, error) {
	if l.size == 0 {
		return nil, dataIsEmptyErr
	}
	if index < 0 || index >= l.size {
		return nil, incorrectIndexErr
	}
	copyHead := l.head
	for i := 0; i <= index-1; i++ {
		copyHead = copyHead.next
	}
	return copyHead.val, nil
}

func (l *MyLinkedList) GetFirst() (any, error) {
	if l.size == 0 {
		return nil, dataIsEmptyErr
	}
	return l.head.val, nil
}

func (l *MyLinkedList) GetLast() (any, error) {
	if l.size == 0 {
		return nil, dataIsEmptyErr
	}
	return l.tail.val, nil
}

// Remove functions

func (l *MyLinkedList) removeOneEl() error {
	if l.size != 1 {
		return errors.New("unexpected function call that deletes a single value")
	}
	l.Clear()
	return nil
}

func (l *MyLinkedList) Remove(index int) error {
	if l.size == 0 {
		return dataIsEmptyErr
	}
	if index < 0 || index >= l.size {
		return incorrectIndexErr
	}
	if index == 0 {
		return l.RemoveFirst()
	} else if l.size == index {
		return l.RemoveLast()
	}

	currNode := l.head
	for i := 0; i <= index-2; i++ {
		currNode = currNode.next
	}
	nextCurrNodeNextNode := currNode.next.next
	nextCurrNodeNextNode.prev = currNode
	currNode.next = nextCurrNodeNextNode
	l.size--

	return nil
}

func (l *MyLinkedList) RemoveFirst() error {
	if l.size == 0 {
		return dataIsEmptyErr
	}
	if l.size == 1 {
		return l.removeOneEl()
	}

	newHead := l.head.next
	newHead.prev = nil
	l.head = newHead
	l.size--
	return nil
}

func (l *MyLinkedList) RemoveLast() error {
	if l.size == 0 {
		return dataIsEmptyErr
	}
	if l.size == 1 {
		return l.removeOneEl()
	}

	// update old tail prev node.next
	l.tail.prev.next = nil
	// replace tail
	l.tail = l.tail.prev
	l.size--
	return nil
}

// ToArray Func
func (l *MyLinkedList) ToArray() []any {
	if l.size == 0 {
		return []any{}
	}
	msv := []any{}
	copyHead := l.head
	for copyHead != nil {
		msv = append(msv, copyHead.val)
		copyHead = copyHead.next
	}
	return msv
}

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

func (l *MyLinkedList) Sort(cmp func(i, j any) bool) error {
	// Copying array to not sorting interface{} (empty) els
	elsMsv := l.ToArray()
	// If length is 0, return
	if len(elsMsv) == 0 {
		return nil
	}
	// Checking it is a correct cmp function
	if !isCorrectLessFunction(elsMsv[0], cmp) {
		return incorrectLessFunction
	}

	// Create a sort.Slice Less Function
	lessSliceFunc := func(i, j int) bool {
		return cmp(elsMsv[i], elsMsv[j])
	}
	// Sort it
	sort.Slice(elsMsv, lessSliceFunc)
	// Replace All Elements, O(N)
	l.Clear()
	for _, el := range elsMsv {
		err := l.AddLast(el)
		if err != nil {
			return err
		}
	}
	// Total Runtime - O(N * logN)
	return nil
}

// Set func
func (l *MyLinkedList) Set(setEl any, index int) error {
	if l.size == 0 {
		return dataIsEmptyErr
	}
	if index < 0 || index >= l.size {
		return incorrectIndexErr
	}
	if ok := l.validateType(setEl); !ok {
		return incorrectElementType
	}

	copyHead := l.head
	for i := 0; i <= index-1; i++ {
		copyHead = copyHead.next
	}
	copyHead.val = setEl
	return nil
}

// Searching all functions
func (l *MyLinkedList) IndexOf(searchedEl any) (int, error) {
	if l.size == 0 {
		return -1, dataIsEmptyErr
	}
	if !l.validateType(searchedEl) {
		return -1, incorrectElementType
	}
	copyHead := l.head
	counter := 0
	for copyHead != nil {
		if reflect.DeepEqual(copyHead.val, searchedEl) {
			return counter, nil
		}
		copyHead = copyHead.next
		counter++
	}

	return -1, nil
}

func (l *MyLinkedList) LastIndexOf(searchedEl any) (int, error) {
	if l.size == 0 {
		return -1, dataIsEmptyErr
	}
	if !l.validateType(searchedEl) {
		return -1, incorrectElementType
	}
	copyHead := l.head
	counter, mxCounter := 0, -1
	for copyHead != nil {
		if reflect.DeepEqual(copyHead.val, searchedEl) {
			mxCounter = counter
		}
		copyHead = copyHead.next
		counter++
	}

	return mxCounter, nil
}

func (l *MyLinkedList) Exists(searchedEl any) bool {
	if l.size == 0 {
		return false
	}
	if !l.validateType(searchedEl) {
		return false
	}
	copyHead := l.head
	for copyHead != nil {
		if reflect.DeepEqual(copyHead.val, searchedEl) {
			return true
		}
		copyHead = copyHead.next
	}

	return false
}

// Other functions
func (l *MyLinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *MyLinkedList) GetType() reflect.Type {
	return l.listType
}

func (l *MyLinkedList) Size() int {
	return l.size
}
