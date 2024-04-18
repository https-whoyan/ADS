package Lists

import (
	"errors"
	"fmt"
	rnd "github.com/https-whoyan/ADS/tree/main/Assigment2/pkg/Random"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
)

// Err, which tells you that the test was not expecting
// this error at all, for its cause is not the current function under test
var unexceptionalErr = errors.New("unexpected error not related to case testing")

// View below
var internalTestErr = errors.New(" internal error in getTestingIntArr function")

// / !!!!
// / !!!!
var testingType = "ArrayList" // Each one of LinkedList or ArrayList. It types of list will be testing
// I don't know why, but Goland frowns on the function
// GetExampleA function, IDE does not see it, but everything goes smoothly when running tests.

func TestNewList(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("NewList, is Empty Arr", func(t *testing.T) {
		arr := newArrList.ToArray()
		assert.Equal(t, []any{}, arr)
	})
	t.Run("NewList, size", func(t *testing.T) {
		arrSize := newArrList.Size()
		assert.Equal(t, 0, arrSize)
	})
}

func Test_AddFirst(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("First Add First, it is ok? ", func(t *testing.T) {
		err := newArrList.AddFirst(rnd.GetRandomNumber())
		if err != nil {
			assert.Error(t, err)
		}
		assert.True(t, true)
	})
	t.Run("IncorrectType", func(t *testing.T) {
		err := newArrList.AddFirst("56")
		assert.Equal(t, incorrectElementType, err)
	})
	t.Run("Expedition capacity", func(t *testing.T) {
		for i := 0; i <= 200; i++ {
			err := newArrList.AddFirst(rnd.GetRandomNumber())
			if err != nil {
				assert.Error(t, err)
			}
		}
		assert.True(t, true)
	})
}

func Test_AddLast(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("First Add Last, it is ok? ", func(t *testing.T) {
		err := newArrList.AddFirst(rnd.GetRandomNumber())
		if err != nil {
			assert.Error(t, err)
		}
		assert.True(t, true)
	})
	t.Run("IncorrectType", func(t *testing.T) {
		err := newArrList.AddFirst("56")
		assert.Equal(t, incorrectElementType, err)
	})
	t.Run("Expedition capacity", func(t *testing.T) {
		for i := 0; i <= 200; i++ {
			err := newArrList.AddLast(rnd.GetRandomNumber())
			if err != nil {
				assert.Error(t, err)
			}
		}
		assert.True(t, true)
	})
}

// Add. With int type els
func Test_Add_Test1(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("Add to minus index", func(t *testing.T) {
		err := newArrList.Add(rnd.GetRandomNumber(), -1)
		assert.Equal(t, incorrectIndexErr, err)
	})
	t.Run("Add to incorrect index", func(t *testing.T) {
		err := newArrList.Add(rnd.GetRandomNumber(), 5)
		assert.Equal(t, incorrectIndexErr, err)
	})
	//Addition some values
	for i := 0; i <= 8; i++ {
		err := newArrList.AddLast(rnd.GetRandomNumber())
		if err != nil {
			assert.Error(t, err)
		}
	}
	t.Run("Addition some incorrect type val", func(t *testing.T) {
		err := newArrList.Add("5", 3)
		assert.Equal(t, incorrectElementType, err)
	})
	t.Run("Expedition capacity", func(t *testing.T) {
		for i := 0; i <= 200; i++ {
			currLength := newArrList.Size()
			err := newArrList.Add(rnd.GetRandomNumber(), currLength)
			if err != nil {
				assert.Error(t, err)
			}
		}
		assert.True(t, true)
	})
}

// Add. With struct type els
func Test_Add_Test2(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("Add to minus index", func(t *testing.T) {
		err := newArrList.Add(rnd.GetExampleA(), -1)
		assert.Equal(t, incorrectIndexErr, err)
	})
	t.Run("Add to incorrect index", func(t *testing.T) {
		err := newArrList.Add(rnd.GetExampleA(), 3)
		assert.Equal(t, incorrectIndexErr, err)
	})
	//Addition some values
	for i := 0; i <= 8; i++ {
		currLength := newArrList.Size()
		err := newArrList.Add(rnd.GetExampleA(), currLength)
		if err != nil {
			assert.Error(t, err)
		}
	}
	// Add Str
	t.Run("Addition some incorrect type val. Test 1", func(t *testing.T) {
		err := newArrList.Add("5", 0)
		assert.Equal(t, incorrectElementType, err)
	})
	// Add Not RandStructA type struct
	t.Run("Addition some incorrect type val. Test 2", func(t *testing.T) {
		err := newArrList.Add(rnd.GetExampleC(), 0)
		assert.Equal(t, incorrectElementType, err)
	})
	t.Run("Expedition capacity", func(t *testing.T) {
		for i := 0; i <= 200; i++ {
			currLength := newArrList.Size()
			err := newArrList.Add(rnd.GetExampleA(), currLength)
			if err != nil {
				assert.Error(t, err)
			}
		}
		assert.True(t, true)
	})
}

func Test_GetFirst(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("Get first of empty err", func(t *testing.T) {
		val, err := newArrList.GetFirst()
		assert.Equal(t, []any{nil, dataIsEmptyErr}, []any{val, err})
	})
	// Add One Val
	randNumber := rnd.GetRandomNumber()
	maybeErr := newArrList.AddFirst(randNumber)
	if maybeErr != nil {
		assert.Error(t, errors.Join(unexceptionalErr, maybeErr))
	}
	t.Run("Get first ell of one size arr", func(t *testing.T) {
		val, err := newArrList.GetFirst()
		assert.Equal(t, []any{randNumber, nil}, []any{val, err})
	})
	// Add another val's
	t.Run("Another val's check", func(t *testing.T) {
		for i := 0; i <= 200; i++ {
			err := newArrList.AddLast(rnd.GetRandomNumber())
			if err != nil {
				assert.Error(t, err)
			}
		}
		val, err := newArrList.GetFirst()
		assert.Equal(t, []any{randNumber, nil}, []any{val, err})
	})
}

func Test_GetLast(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("Get last of empty err", func(t *testing.T) {
		val, err := newArrList.GetLast()
		assert.Equal(t, []any{nil, dataIsEmptyErr}, []any{val, err})
	})
	// Add One Val
	randNumber := rnd.GetRandomNumber()
	maybeErr := newArrList.AddLast(randNumber)
	if maybeErr != nil {
		assert.Error(t, errors.Join(unexceptionalErr, maybeErr))
	}
	t.Run("Get first ell of one size arr", func(t *testing.T) {
		val, err := newArrList.GetLast()
		assert.Equal(t, []any{randNumber, nil}, []any{val, err})
	})
	// Add another val's
	t.Run("Another val's check", func(t *testing.T) {
		for i := 0; i <= 200; i++ {
			randNumber = rnd.GetRandomNumber()
			err := newArrList.AddLast(randNumber)
			if err != nil {
				assert.Error(t, err)
			}
		}
		val, err := newArrList.GetLast()
		assert.Equal(t, []any{randNumber, nil}, []any{val, err})
	})
}

func Test_Get(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("Get of empty err", func(t *testing.T) {
		val, err := newArrList.Get(0)
		assert.Equal(t, []any{nil, dataIsEmptyErr}, []any{val, err})
	})
	// Add One Val
	randNumber := rnd.GetRandomNumber()
	maybeErr := newArrList.AddLast(randNumber)
	if maybeErr != nil {
		assert.Error(t, errors.Join(unexceptionalErr, maybeErr))
	}
	t.Run("Get of incorrect index < 0", func(t *testing.T) {
		val, err := newArrList.Get(-3)
		assert.Equal(t, []any{nil, incorrectIndexErr}, []any{val, err})
	})
	t.Run("Get of incorrect index > size", func(t *testing.T) {
		val, err := newArrList.Get(8)
		assert.Equal(t, []any{nil, incorrectIndexErr}, []any{val, err})
	})
	t.Run("Get first ell of one size arr", func(t *testing.T) {
		val, err := newArrList.Get(0)
		assert.Equal(t, []any{randNumber, nil}, []any{val, err})
	})
	// Add another val's
	for i := 0; i <= 200; i++ {
		err := newArrList.AddLast(rnd.GetRandomNumber())
		if err != nil {
			assert.Error(t, err)
		}
	}
	arr := newArrList.ToArray()
	t.Run("Another val's check, Test 1", func(t *testing.T) {
		getVal1, getErr1 := newArrList.Get(6)
		if getErr1 != nil {
			assert.Error(t, getErr1)
		}
		arrVal1 := arr[6]
		assert.Equal(t, arrVal1, getVal1)
	})
	t.Run("Another val's check, Test 1", func(t *testing.T) {
		getVal2, getErr2 := newArrList.Get(81)
		if getErr2 != nil {
			assert.Error(t, getErr2)
		}
		arrVal2 := arr[81]
		assert.Equal(t, arrVal2, getVal2)
	})
	t.Run("Another val's check, Test 1", func(t *testing.T) {
		getVal3, getErr3 := newArrList.Get(162)
		if getErr3 != nil {
			assert.Error(t, getErr3)
		}
		arrVal3 := arr[162]
		assert.Equal(t, arrVal3, getVal3)
	})
}

func Test_Set(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("Set of empty err", func(t *testing.T) {
		err := newArrList.Set(5, 0)
		assert.Equal(t, dataIsEmptyErr, err)
	})
	maybeErr := newArrList.AddLast(rnd.GetRandomNumber())
	if maybeErr != nil {
		assert.Error(t, errors.Join(unexceptionalErr, maybeErr))
	}

	t.Run("Set the singleton el in arr", func(t *testing.T) {
		err := newArrList.Set(rnd.GetRandomNumber(), 0)
		if err != nil {
			assert.Error(t, err)
		}
	})
	t.Run("Set to incorrect index", func(t *testing.T) {
		err := newArrList.Set(rnd.GetRandomNumber(), 5)
		assert.Equal(t, incorrectIndexErr, err)
	})
	// Add some values
	for i := 0; i <= 200; i++ {
		err := newArrList.AddLast(rnd.GetRandomNumber())
		if err != nil {
			assert.Error(t, err)
		}
	}
	t.Run("Another set val's check", func(t *testing.T) {
		randomNum := rnd.GetRandomNumber()
		err := newArrList.Set(randomNum, 6)
		if err != nil {
			assert.Error(t, err)
		}
		arrVal := newArrList.ToArray()[6]
		newVal, err := newArrList.Get(6)
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, arrVal, newVal)
	})
	t.Run("Another set val's check", func(t *testing.T) {
		randomNum := rnd.GetRandomNumber()
		err := newArrList.Set(randomNum, 113)
		if err != nil {
			assert.Error(t, err)
		}
		arrVal := newArrList.ToArray()[113]
		newVal, err := newArrList.Get(113)
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, arrVal, newVal)
	})
	t.Run("Incorrect type val set", func(t *testing.T) {
		err := newArrList.Set("5", 113)
		assert.Equal(t, incorrectElementType, err)
	})
}

// Sort. With int type elements
func Test_Sort1(t *testing.T) {
	newArrList := NewList(testingType)
	// Add some values
	for i := 0; i <= 200; i++ {
		err := newArrList.AddLast(rnd.GetRandomNumber())
		if err != nil {
			assert.Error(t, err)
		}
	}

	t.Run("Incorrect sort function", func(t *testing.T) {
		err := newArrList.Sort(func(x, y any) bool {
			return x.(string) < y.(string)
		})
		assert.Equal(t, incorrectLessFunction, err)
	})
	t.Run("Normal sort", func(t *testing.T) {
		err := newArrList.Sort(func(x, y any) bool {
			return x.(int) < y.(int)
		})
		if err != nil {
			assert.Error(t, err)
		}
	})

	sortedArr := newArrList.ToArray()
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i].(int) < sortedArr[j].(int)
	})
	currArr := newArrList.ToArray()
	assert.Equal(t, sortedArr, currArr)
}

// Sort. With struct type elements
func Test_Sort2(t *testing.T) {
	newArrList := NewList(testingType)
	// Add some values
	for i := 0; i <= 200; i++ {
		err := newArrList.AddLast(rnd.GetExampleA())
		if err != nil {
			assert.Error(t, err)
		}
	}

	t.Run("Incorrect sort function", func(t *testing.T) {
		err := newArrList.Sort(func(x, y any) bool {
			return x.(rnd.RandStructB).RandFieldA_int < y.(rnd.RandStructB).RandFieldA_int
		})
		assert.Equal(t, incorrectLessFunction, err)
	})
	t.Run("Normal sort", func(t *testing.T) {
		err := newArrList.Sort(func(x, y any) bool {
			return x.(rnd.RandStructA).RandFieldB_float64 < y.(rnd.RandStructA).RandFieldB_float64
		})
		if err != nil {
			assert.Error(t, err)
		}
	})

	sortedArr := newArrList.ToArray()
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i].(rnd.RandStructA).RandFieldB_float64 < sortedArr[j].(rnd.RandStructA).RandFieldB_float64
	})
	currArr := newArrList.ToArray()
	assert.Equal(t, sortedArr, currArr)
}

// Remove function's
func Test_RemoveFirst(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("Remove of empty err", func(t *testing.T) {
		err := newArrList.RemoveFirst()
		assert.Equal(t, dataIsEmptyErr, err)
	})
	maybeErr := newArrList.AddLast(5)
	if maybeErr != nil {
		assert.Error(t, errors.Join(unexceptionalErr, maybeErr))
	}
	t.Run("Remove is one unique el", func(t *testing.T) {
		err := newArrList.RemoveFirst()
		if err != nil {
			assert.Error(t, err)
		}
		currMsv := newArrList.ToArray()
		currLen := newArrList.Size()
		assert.Equal(t, [][]any{[]any{}, {0}}, [][]any{currMsv, {currLen}})
	})
	t.Run("Remove, normal test", func(t *testing.T) {
		// Add some values
		for i := 0; i <= 8; i++ {
			maybeErr = newArrList.AddLast(rnd.GetRandomNumber())
			if maybeErr != nil {
				assert.Error(t, errors.Join(unexceptionalErr, maybeErr))
			}
		}
		exceptedLen := newArrList.Size() - 1
		exceptedNextFirstEl := newArrList.ToArray()[1]
		err := newArrList.RemoveFirst()
		if err != nil {
			assert.Error(t, err)
		}
		currFirstElAfterDelete := newArrList.ToArray()[0]
		currLen := newArrList.Size()

		// Checking
		exceptedArr := []any{exceptedNextFirstEl, exceptedLen}
		currArr := []any{currFirstElAfterDelete, currLen}
		assert.Equal(t, exceptedArr, currArr)
	})
}

func Test_RemoveLast(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("Remove of empty err", func(t *testing.T) {
		err := newArrList.RemoveFirst()
		assert.Equal(t, dataIsEmptyErr, err)
	})
	maybeErr := newArrList.AddLast(5)
	if maybeErr != nil {
		assert.Error(t, errors.Join(unexceptionalErr, maybeErr))
	}
	t.Run("Remove is one unique el", func(t *testing.T) {
		err := newArrList.RemoveLast()
		if err != nil {
			assert.Error(t, err)
		}
		currMsv := newArrList.ToArray()
		currLen := newArrList.Size()
		assert.Equal(t, [][]any{[]any{}, {0}}, [][]any{currMsv, {currLen}})
	})
	t.Run("Remove, normal test", func(t *testing.T) {
		// Add some values
		for i := 0; i <= 8; i++ {
			maybeErr = newArrList.AddLast(rnd.GetRandomNumber())
			if maybeErr != nil {
				assert.Error(t, errors.Join(unexceptionalErr, maybeErr))
			}
		}
		//fmt.Println(newArrList.ToArray())
		firstLen, exceptedLen := newArrList.Size(), newArrList.Size()-1
		exceptedNextLastEl := newArrList.ToArray()[firstLen-2]
		err := newArrList.RemoveLast()
		if err != nil {
			assert.Error(t, err)
		}
		currLen := newArrList.Size()
		currLastElAfterDelete := newArrList.ToArray()[currLen-1]

		// Checking
		exceptedArr := []any{exceptedNextLastEl, exceptedLen}
		currArr := []any{currLastElAfterDelete, currLen}
		assert.Equal(t, exceptedArr, currArr)
	})
}

func Test_Remove(t *testing.T) {
	newArrList := NewList(testingType)
	t.Run("Remove of empty err", func(t *testing.T) {
		err := newArrList.Remove(0)
		assert.Equal(t, dataIsEmptyErr, err)
	})
	maybeErr := newArrList.AddLast(5)
	if maybeErr != nil {
		assert.Error(t, errors.Join(unexceptionalErr, maybeErr))
	}
	t.Run("incorrect removed index < 0", func(t *testing.T) {
		err := newArrList.Remove(-4)
		assert.Equal(t, incorrectIndexErr, err)
	})
	t.Run("incorrect removed index == len", func(t *testing.T) {
		err := newArrList.Remove(newArrList.Size())
		assert.Equal(t, incorrectIndexErr, err)
	})
	t.Run("incorrect removed index > len", func(t *testing.T) {
		err := newArrList.Remove(173)
		assert.Equal(t, incorrectIndexErr, err)
	})
	t.Run("Remove is one unique el", func(t *testing.T) {
		err := newArrList.Remove(0)
		if err != nil {
			assert.Error(t, err)
		}
		currMsv := newArrList.ToArray()
		currLen := newArrList.Size()
		assert.Equal(t, [][]any{[]any{}, {0}}, [][]any{currMsv, {currLen}})
	})
	// Add some values
	for i := 0; i <= 8; i++ {
		maybeErr = newArrList.AddLast(rnd.GetRandomNumber())
		if maybeErr != nil {
			assert.Error(t, errors.Join(unexceptionalErr, maybeErr))
		}
	}
	t.Run("Remove, normal test 1. Removed index - 6", func(t *testing.T) {
		// I'll be deleted the 6'th index
		exceptedLen := newArrList.Size() - 1
		exceptedNext6thIndexEl := newArrList.ToArray()[7]
		err := newArrList.Remove(6)
		if err != nil {
			assert.Error(t, err)
		}
		currLen := newArrList.Size()
		curr6thIndexElAfterDelete := newArrList.ToArray()[6]

		// Checking
		exceptedArr := []any{exceptedNext6thIndexEl, exceptedLen}
		currArr := []any{curr6thIndexElAfterDelete, currLen}
		assert.Equal(t, exceptedArr, currArr)
	})
	t.Run("Remove, normal test 2. Removed index - 2", func(t *testing.T) {
		// I'll be deleted the 2'th index
		exceptedLen := newArrList.Size() - 1
		exceptedNext6thIndexEl := newArrList.ToArray()[3]
		err := newArrList.Remove(2)
		if err != nil {
			assert.Error(t, err)
		}
		currLen := newArrList.Size()
		curr6thIndexElAfterDelete := newArrList.ToArray()[2]

		// Checking
		exceptedArr := []any{exceptedNext6thIndexEl, exceptedLen}
		currArr := []any{curr6thIndexElAfterDelete, currLen}
		assert.Equal(t, exceptedArr, currArr)
	})
}

// Searching functions

// I'll be testing with 2 arrays, with
// int type of elements
// and rnd.RandStructA type of elements
func getTestingIntArr() (List, error) {
	newArrList := NewList(testingType)
	for i := 0; i <= 200; i++ {
		err := newArrList.AddLast(rnd.GetRandomNumber())
		if err != nil {
			return nil, errors.Join(unexceptionalErr, internalTestErr, err)
		}
	}
	return newArrList, nil
}

func getTestingRndStructAArr() (List, error) {
	newArrList := NewList(testingType)
	for i := 0; i <= 200; i++ {
		err := newArrList.AddLast(rnd.GetExampleA())
		if err != nil {
			return nil, errors.Join(unexceptionalErr, internalTestErr, err)
		}
	}
	return newArrList, nil
}

func Test_IndexOf(t *testing.T) {
	newArrIntList, maybeErrInt := getTestingIntArr()
	newArrStructList, maybeErrStruct := getTestingRndStructAArr()
	if maybeErrInt != nil {
		assert.Error(t, maybeErrInt)
	}
	if maybeErrStruct != nil {
		assert.Error(t, maybeErrStruct)
	}
	t.Run("IndexOf, int arr, invalidType, excepted err", func(t *testing.T) {
		index, err := newArrIntList.IndexOf("5")
		assert.Equal(t, []any{-1, incorrectElementType}, []any{index, err})
	})
	t.Run("IndexOf, struct arr, invalidType, excepted err", func(t *testing.T) {
		index, err := newArrStructList.IndexOf("5")
		assert.Equal(t, []any{-1, incorrectElementType}, []any{index, err})
	})
	t.Run("IndexOf, int arr, not have in arr", func(t *testing.T) {
		index, err := newArrIntList.IndexOf(-1000)
		assert.Equal(t, []any{-1, nil}, []any{index, err})
	})
	t.Run("IndexOf, struct arr, not have in arr", func(t *testing.T) {
		index, err := newArrStructList.IndexOf(rnd.RandStructA{})
		assert.Equal(t, []any{-1, nil}, []any{index, err})
	})
	// Normal tests
	t.Run("IndexOf, int arr, normal test 1", func(t *testing.T) {
		indexToSearch := 19
		val, err := newArrIntList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		// Find true min Index
		for i, el := range newArrIntList.ToArray() {
			if reflect.DeepEqual(el, val) {
				indexToSearch = i
				break
			}
		}
		currIndex, err := newArrIntList.IndexOf(val)
		assert.Equal(t, []any{indexToSearch, nil}, []any{currIndex, err})
	})
	t.Run("IndexOf, struct arr, normal test 1", func(t *testing.T) {
		indexToSearch := 19
		val, err := newArrStructList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		// Find true min Index
		for i, el := range newArrStructList.ToArray() {
			if reflect.DeepEqual(el, val) {
				indexToSearch = i
				break
			}
		}
		currIndex, err := newArrStructList.IndexOf(val)
		assert.Equal(t, []any{indexToSearch, nil}, []any{currIndex, err})
	})
	t.Run("IndexOf, int arr, normal test 2", func(t *testing.T) {
		indexToSearch := 182
		val, err := newArrIntList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		// Find true min Index
		for i, el := range newArrIntList.ToArray() {
			if reflect.DeepEqual(el, val) {
				indexToSearch = i
				break
			}
		}
		currIndex, err := newArrIntList.IndexOf(val)
		assert.Equal(t, []any{indexToSearch, nil}, []any{currIndex, err})
	})
	t.Run("IndexOf, struct arr, normal test 2", func(t *testing.T) {
		indexToSearch := 182
		val, err := newArrStructList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		// Find true min Index
		for i, el := range newArrStructList.ToArray() {
			if reflect.DeepEqual(el, val) {
				indexToSearch = i
				break
			}
		}
		currIndex, err := newArrStructList.IndexOf(val)
		assert.Equal(t, []any{indexToSearch, nil}, []any{currIndex, err})
	})
}

func Test_LastIndexOf(t *testing.T) {
	newArrIntList, maybeErrInt := getTestingIntArr()
	newArrStructList, maybeErrStruct := getTestingRndStructAArr()
	if maybeErrInt != nil {
		assert.Error(t, maybeErrInt)
	}
	if maybeErrStruct != nil {
		assert.Error(t, maybeErrStruct)
	}
	t.Run("LastIndexOf, int arr, invalidType, excepted err", func(t *testing.T) {
		index, err := newArrIntList.LastIndexOf("5")
		assert.Equal(t, []any{-1, incorrectElementType}, []any{index, err})
	})
	t.Run("LastIndexOf, struct arr, invalidType, excepted err", func(t *testing.T) {
		index, err := newArrStructList.LastIndexOf("5")
		assert.Equal(t, []any{-1, incorrectElementType}, []any{index, err})
	})
	t.Run("LastIndexOf, int arr, not have in arr", func(t *testing.T) {
		index, err := newArrIntList.LastIndexOf(-1000)
		assert.Equal(t, []any{-1, nil}, []any{index, err})
	})
	t.Run("LastIndexOf, struct arr, not have in arr", func(t *testing.T) {
		index, err := newArrStructList.LastIndexOf(rnd.RandStructA{})
		assert.Equal(t, []any{-1, nil}, []any{index, err})
	})
	// Normal tests
	t.Run("LastIndexOf, int arr, normal test 1", func(t *testing.T) {
		indexToSearch := 19
		val, err := newArrIntList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		// Find true mx Index
		for i, el := range newArrIntList.ToArray() {
			if reflect.DeepEqual(el, val) {
				indexToSearch = i
			}
		}
		currIndex, err := newArrIntList.LastIndexOf(val)
		assert.Equal(t, []any{indexToSearch, nil}, []any{currIndex, err})
	})
	t.Run("LastIndexOf, struct arr, normal test 1", func(t *testing.T) {
		indexToSearch := 19
		val, err := newArrStructList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		// Find true mx Index
		for i, el := range newArrStructList.ToArray() {
			if reflect.DeepEqual(el, val) {
				indexToSearch = i
			}
		}
		currIndex, err := newArrStructList.LastIndexOf(val)
		assert.Equal(t, []any{indexToSearch, nil}, []any{currIndex, err})
	})
	t.Run("LastIndexOf, int arr, normal test 2", func(t *testing.T) {
		indexToSearch := 182
		val, err := newArrIntList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		// Find true mx Index
		for i, el := range newArrIntList.ToArray() {
			if reflect.DeepEqual(el, val) {
				indexToSearch = i
			}
		}
		currIndex, err := newArrIntList.LastIndexOf(val)
		assert.Equal(t, []any{indexToSearch, nil}, []any{currIndex, err})
	})
	t.Run("LastIndexOf, struct arr, normal test 2", func(t *testing.T) {
		indexToSearch := 182
		val, err := newArrStructList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		// Find true mx Index
		fmt.Println(newArrStructList.ToArray())
		for i, el := range newArrStructList.ToArray() {
			if reflect.DeepEqual(el, val) {
				indexToSearch = i
			}
		}
		currIndex, err := newArrStructList.LastIndexOf(val)
		assert.Equal(t, []any{indexToSearch, nil}, []any{currIndex, err})
	})
}

func Test_Exists(t *testing.T) {
	newArrIntList, maybeErrInt := getTestingIntArr()
	newArrStructList, maybeErrStruct := getTestingRndStructAArr()
	if maybeErrInt != nil {
		assert.Error(t, maybeErrInt)
	}
	if maybeErrStruct != nil {
		assert.Error(t, maybeErrStruct)
	}
	t.Run("Exists, int arr, invalidType, excepted err", func(t *testing.T) {
		isContains := newArrIntList.Exists("5")
		assert.Equal(t, false, isContains)
	})
	t.Run("Exists, struct arr, invalidType, excepted err", func(t *testing.T) {
		isContains := newArrStructList.Exists("5")
		assert.Equal(t, false, isContains)
	})
	t.Run("Exists, int arr, not have in arr", func(t *testing.T) {
		isContains := newArrIntList.Exists(-1000)
		assert.Equal(t, false, isContains)
	})
	t.Run("Exists, struct arr, not have in arr", func(t *testing.T) {
		isContains := newArrStructList.Exists(rnd.RandStructA{})
		assert.Equal(t, false, isContains)
	})
	// Normal tests
	t.Run("Exists, int arr, normal test 1", func(t *testing.T) {
		indexToSearch := 19
		val, err := newArrIntList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		isContains := newArrIntList.Exists(val)
		assert.Equal(t, true, isContains)
	})
	t.Run("Exists, struct arr, normal test 1", func(t *testing.T) {
		indexToSearch := 19
		val, err := newArrStructList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		isContains := newArrStructList.Exists(val)
		assert.Equal(t, true, isContains)
	})
	t.Run("Exists, int arr, normal test 2", func(t *testing.T) {
		indexToSearch := 182
		val, err := newArrIntList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		isContains := newArrIntList.Exists(val)
		assert.Equal(t, true, isContains)
	})
	t.Run("Exists, struct arr, normal test 2", func(t *testing.T) {
		indexToSearch := 182
		val, err := newArrStructList.Get(indexToSearch)
		if err != nil {
			assert.Error(t, errors.Join(unexceptionalErr, err))
		}
		isContains := newArrStructList.Exists(val)
		assert.Equal(t, true, isContains)
	})
}
