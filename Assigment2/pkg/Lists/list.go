package Lists

import (
	"errors"
	"reflect"

	arrList "github.com/https-whoyan/ADS/tree/main/Assigment2/pkg/Lists/MyArrayList"
	linkList "github.com/https-whoyan/ADS/tree/main/Assigment2/pkg/Lists/MyLinkedList"
)

var incorrectIndexErr = errors.New("incorrect index")
var incorrectElementType = errors.New("incorrect added element type")
var dataIsEmptyErr = errors.New("data is empty, cannot pick/delete any element")
var incorrectLessFunction = errors.New("incorrect less function")

type List interface {
	Add(newEl any, index int) error
	AddFirst(newEl any) error
	AddLast(newEl any) error
	Set(setEl any, index int) error
	Get(index int) (any, error)
	GetFirst() (any, error)
	GetLast() (any, error)
	Remove(index int) error
	RemoveFirst() error
	RemoveLast() error
	Sort(cmp func(i, j any) bool) error
	IndexOf(searchedEl any) (int, error)
	LastIndexOf(searchedEl any) (int, error)
	Exists(searchedEl any) bool
	ToArray() []any
	Clear()
	GetType() reflect.Type
	Size() int
}

func NewList(listType string) List {
	if listType == "ArrayList" {
		return arrList.NewArrayList()
	} else if listType == "LinkedList" {
		return linkList.NewLinkedList()
	}
	return nil
}
