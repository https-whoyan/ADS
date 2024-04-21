package MyHeap

import (
	rnd "github.com/https-whoyan/ADS/tree/main/Assigment2/pkg/Random"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
)

func TestNewList(t *testing.T) {
	t.Run("Incorrect cmp function + incorrect element type", func(t *testing.T) {
		newHeap := NewMyMinHeap(func(x, y any) bool {
			return x.(string) < y.(string)
		})
		err := newHeap.Add(rnd.GetExampleC())
		heapType := newHeap.heapType
		heapSize := newHeap.Size()

		exceptedArr := []any{incorrectLessFunction, reflect.ValueOf(rnd.RandStructC{}).Type(), 0}
		actualArr := []any{err, heapType, heapSize}
		assert.Equal(t, exceptedArr, actualArr)
	})
	t.Run("Correct cmp function + correct element type", func(t *testing.T) {
		newHeap := NewMyMinHeap(func(x, y any) bool {
			return x.(string) < y.(string)
		})
		err := newHeap.Add("6")
		heapType := newHeap.heapType
		heapSize := newHeap.Size()

		exceptedArr := []any{nil, reflect.ValueOf("").Type(), 1}
		actualArr := []any{err, heapType, heapSize}
		assert.Equal(t, exceptedArr, actualArr)
	})
}

func getExceptedRoot(arr []any, cmp func(x, y any) bool) any {
	lessFunc := func(i, j int) bool {
		return cmp(arr[i], arr[j])
	}
	sort.Slice(arr, lessFunc)
	return arr[0]
}

func universalTest(cmp func(x, y any) bool, newElF func() any) ([]any, []any, error) {
	heap := NewMyMinHeap(cmp)
	var testArr []any
	// add Some Values
	for i := 0; i <= 200; i++ {
		randomEl := newElF()
		testArr = append(testArr, randomEl)
		err := heap.Add(randomEl)
		if err != nil {
			return nil, nil, err
		}
	}
	sort.Slice(testArr, func(i, j int) bool {
		return cmp(testArr[i], testArr[j])
	})

	var exceptedArr []any
	var currentArr []any

	// Testing 50 values
	heap.printData()
	for i := 1; i <= 50; i++ {
		currRoot, err := heap.ExtractRoot()
		if err != nil {
			return nil, nil, err
		}
		exceptedArr = append(exceptedArr, getExceptedRoot(testArr[i-1:], cmp))
		currentArr = append(currentArr, currRoot)
	}
	// Add incorrect type val
	err := heap.Add("5")
	exceptedArr = append(exceptedArr, reflect.ValueOf(incorrectElementType).Type())
	currentArr = append(currentArr, reflect.ValueOf(err).Type())
	return exceptedArr, currentArr, nil
}

func Test_Global(t *testing.T) {
	t.Run("MinHeap, with int type values, minHeap: ", func(t *testing.T) {
		cmpTest := func(x, y any) bool {
			return x.(int) < y.(int)
		}
		newElF := func() any {
			return any(rnd.GetRandomNumber())
		}
		exceptedArr, currentArr, err := universalTest(cmpTest, newElF)
		if err != nil {
			assert.Error(t, err)
		}
		// Checking
		assert.Equal(t, exceptedArr, currentArr)
	})
	t.Run("MaxHeap, with int type values, maxHeap: ", func(t *testing.T) {
		cmpTest := func(x, y any) bool {
			return x.(int) > y.(int)
		}
		newElF := func() any {
			return any(rnd.GetRandomNumber())
		}
		exceptedArr, currentArr, err := universalTest(cmpTest, newElF)
		if err != nil {
			assert.Error(t, err)
		}
		// Checking
		assert.Equal(t, exceptedArr, currentArr)
	})
	t.Run("Min, with struct type values", func(t *testing.T) {
		cmpTest := func(x, y any) bool {
			typedX, typedY := x.(rnd.RandStructA), y.(rnd.RandStructA)
			xStr, yStr := typedX.RandFieldA_str, typedY.RandFieldA_str
			if xStr == yStr {
				return typedX.RandFieldB_float64 < typedY.RandFieldB_float64
			}
			return x.(rnd.RandStructA).RandFieldA_str < y.(rnd.RandStructA).RandFieldA_str
		}
		newElF := func() any {
			return any(rnd.GetExampleA())
		}
		exceptedArr, currentArr, err := universalTest(cmpTest, newElF)
		if err != nil {
			assert.Error(t, err)
		}
		// Checking
		assert.Equal(t, exceptedArr, currentArr)
	})
}
