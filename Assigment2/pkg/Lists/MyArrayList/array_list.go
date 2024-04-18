package MyArrayList

import (
	"errors"
	"reflect"
	"sort"
)

var incorrectIndexErr = errors.New("incorrect index")
var incorrectElementType = errors.New("incorrect added element type")
var dataIsEmptyErr = errors.New("data is empty, cannot pick/delete any element")
var incorrectLessFunction = errors.New("incorrect less function")

const defaultCap = 10

type myArrayList struct {
	data     []any
	listType reflect.Type
	size     int
}

type al myArrayList

func NewArrayList() *al {
	return &al{
		data:     make([]any, defaultCap, defaultCap),
		listType: nil,
		size:     0,
	}
}

func (l *al) increaseCap() {
	newCap := l.size * 2
	newData := make([]any, newCap, newCap)

	for i, el := range l.data {
		newData[i] = el
	}

	l.data = newData
}

// true, if l.listType == el.type, or l.listType == nil
func (l *al) validateType(el any) bool {
	if l.listType == nil {
		return true
	}
	return !(reflect.ValueOf(el).Type() != l.listType)
}

// Standing and checking standing list type
func (l *al) standListType(refEl any) {
	l.listType = reflect.ValueOf(refEl).Type()
}

func (l *al) needStandListType() bool {
	return l.size == 0 && l.listType == nil
}

// Add functions
func (l *al) Add(newEl any, indexToAdd int) error {
	if indexToAdd > l.size || indexToAdd < 0 {
		return incorrectIndexErr
	}
	if !(l.validateType(newEl)) {
		return incorrectElementType
	}
	if l.needStandListType() {
		l.standListType(newEl)
	}
	if l.size == cap(l.data)-1 {
		l.increaseCap()
	}
	dataArr := l.ToArray()
	currDataLen := len(dataArr)
	for index := 0; index <= currDataLen; index++ {
		if index < indexToAdd {
			l.data[index] = dataArr[index]
		} else if index == indexToAdd {
			l.data[index] = newEl
		} else {
			l.data[index] = dataArr[index-1]
		}
	}
	l.size++
	return nil
}

func (l *al) AddLast(newEl any) error {
	if ok := l.validateType(newEl); !ok {
		return incorrectElementType
	}
	if l.needStandListType() {
		l.standListType(newEl)
	}
	if l.size == cap(l.data)-1 {
		l.increaseCap()
	}
	l.data[l.size] = newEl

	l.size++
	return nil
}

func (l *al) AddFirst(newEl any) error {
	if ok := l.validateType(newEl); !ok {
		return incorrectElementType
	}
	if l.needStandListType() {
		l.standListType(newEl)
	}
	if l.size == cap(l.data)-1 {
		l.increaseCap()
	}
	newData := append([]any{newEl}, l.data...)
	l.data = newData
	l.size++
	return nil
}

// Get functions
func (l *al) Get(index int) (any, error) {
	if l.size == 0 {
		return nil, dataIsEmptyErr
	}
	if index >= l.size || index < 0 {
		return nil, incorrectIndexErr
	}
	return l.data[index], nil
}

func (l *al) GetFirst() (any, error) {
	if l.size == 0 {
		return nil, dataIsEmptyErr
	}
	return l.data[0], nil
}

func (l *al) GetLast() (any, error) {
	if l.size == 0 {
		return nil, dataIsEmptyErr
	}
	return l.data[l.size-1], nil
}

// Remove functions
func (l *al) Remove(index int) error {
	if l.size == 0 {
		return dataIsEmptyErr
	}
	if index >= l.size || index < 0 {
		return incorrectIndexErr
	}
	for i := index + 1; i <= l.size-1; i++ {
		l.data[i-1] = l.data[i]
	}
	// Clear ell
	l.data[l.size-1] = nil
	l.size--
	return nil
}

func (l *al) RemoveFirst() error {
	err := l.Remove(0)
	return err
}

func (l *al) RemoveLast() error {
	removedIndex := l.size - 1
	l.data[removedIndex] = nil
	l.size--
	return nil
}

// ToArray func
func (l *al) ToArray() []any {
	if l.size == 0 {
		return []any{}
	}
	arrSize := l.size
	arr := make([]any, arrSize, arrSize)
	for i := 0; i <= arrSize-1; i++ {
		arr[i] = l.data[i]
	}
	return arr
}

// Sort func and utils
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

func (l *al) Sort(cmp func(i, j any) bool) error {
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
	for i := 0; i <= l.size-1; i++ {
		l.data[i] = elsMsv[i]
	}
	// Total Runtime - O(N * logN)
	return nil
}

// Set function
func (l *al) Set(val any, index int) error {
	if l.size == 0 {
		return dataIsEmptyErr
	}
	if ok := l.validateType(val); !ok {
		return incorrectElementType
	}
	if index >= l.size || index < 0 {
		return incorrectIndexErr
	}
	l.data[index] = val
	return nil
}

// Finding function
func (l *al) IndexOf(searchedEl any) (int, error) {
	if l.size == 0 {
		return -1, dataIsEmptyErr
	}
	if !l.validateType(searchedEl) {
		return -1, incorrectElementType
	}

	for index, el := range l.data {
		if reflect.DeepEqual(el, searchedEl) {
			return index, nil
		}
	}

	return -1, nil
}

func (l *al) LastIndexOf(searchedEl any) (int, error) {
	if l.size == 0 {
		return -1, dataIsEmptyErr
	}
	if !l.validateType(searchedEl) {
		return -1, incorrectElementType
	}

	returnedIndex := -1
	for index, el := range l.data {
		if reflect.DeepEqual(el, searchedEl) {
			returnedIndex = index
		}
	}

	return returnedIndex, nil
}

func (l *al) Exists(searchedEl any) bool {
	if l.size == 0 {
		return false
	}
	if !l.validateType(searchedEl) {
		return false
	}

	for i, el := range l.data {
		if i == l.size {
			break
		}
		if reflect.DeepEqual(el, searchedEl) {
			return true
		}
	}

	return false
}

// Other functions
func (l *al) Clear() {
	*l = al{
		data: make([]any, defaultCap, defaultCap),
		size: 0,
	}
}

func (l *al) GetType() reflect.Type {
	return l.listType
}

func (l *al) Size() int {
	return l.size
}
