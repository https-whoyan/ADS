package Lists

import (
	rnd "github.com/https-whoyan/ADS/tree/main/Assigment2/pkg/Random"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestNewMyArrayList(t *testing.T) {
	newArrList := NewMyArrayList()
	t.Run("NewArrayList, is Empty Arr", func(t *testing.T) {
		arr := newArrList.ToArray()
		assert.Equal(t, []any{}, arr)
	})
	t.Run("NewArrayList, size", func(t *testing.T) {
		arrSize := newArrList.Size()
		assert.Equal(t, 0, arrSize)
	})
}

func TestAl_AddFirst(t *testing.T) {
	newArrList := NewMyArrayList()
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

func TestAl_AddLast(t *testing.T) {
	newArrList := NewMyArrayList()
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

// Add. With int's
func TestAl_Add_Test1(t *testing.T) {
	newArrList := NewMyArrayList()
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

// Add. With struct's
func TestAl_Add_Test2(t *testing.T) {
	newArrList := NewMyArrayList()
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
		err := newArrList.AddLast(rnd.GetExampleA())
		if err != nil {
			assert.Error(t, err)
		}
	}
	// Add Str
	t.Run("Addition some incorrect type val. Test 1", func(t *testing.T) {
		err := newArrList.Add("5", 4)
		assert.Equal(t, incorrectElementType, err)
	})
	// Add Not RandStructA type struct
	t.Run("Addition some incorrect type val. Test 2", func(t *testing.T) {
		err := newArrList.Add(rnd.GetExampleC(), 4)
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

func TestAl_GetFirst(t *testing.T) {
	newArrList := NewMyArrayList()
	t.Run("Get first of empty err", func(t *testing.T) {
		val, err := newArrList.GetFirst()
		assert.Equal(t, []any{nil, dataIsEmptyErr}, []any{val, err})
	})
	// Add One Val
	randNumber := rnd.GetRandomNumber()
	newArrList.AddFirst(randNumber)
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

func TestAl_GetLast(t *testing.T) {
	newArrList := NewMyArrayList()
	t.Run("Get last of empty err", func(t *testing.T) {
		val, err := newArrList.GetLast()
		assert.Equal(t, []any{nil, dataIsEmptyErr}, []any{val, err})
	})
	// Add One Val
	randNumber := rnd.GetRandomNumber()
	newArrList.AddLast(randNumber)
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

func TestAl_Get(t *testing.T) {
	newArrList := NewMyArrayList()
	t.Run("Get of empty err", func(t *testing.T) {
		val, err := newArrList.Get(0)
		assert.Equal(t, []any{nil, dataIsEmptyErr}, []any{val, err})
	})
	// Add One Val
	randNumber := rnd.GetRandomNumber()
	newArrList.AddLast(randNumber)
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
	t.Run("Another val's check", func(t *testing.T) {
		for i := 0; i <= 200; i++ {
			err := newArrList.AddLast(rnd.GetRandomNumber())
			if err != nil {
				assert.Error(t, err)
			}
		}

		arr := newArrList.ToArray()
		// Get, Test 1
		getVal1, getErr1 := newArrList.Get(6)
		if getErr1 != nil {
			assert.Error(t, getErr1)
		}
		arrVal1 := arr[6]
		assert.Equal(t, arrVal1, getVal1)
		// Get, Test 2
		getVal2, getErr2 := newArrList.Get(81)
		if getErr2 != nil {
			assert.Error(t, getErr2)
		}
		arrVal2 := arr[81]
		assert.Equal(t, arrVal2, getVal2)
		// Get, Test 3
		getVal3, getErr3 := newArrList.Get(162)
		if getErr3 != nil {
			assert.Error(t, getErr3)
		}
		arrVal3 := arr[162]
		assert.Equal(t, arrVal3, getVal3)
	})
}

func TestAl_Set(t *testing.T) {
	newArrList := NewMyArrayList()
	t.Run("Set of empty err", func(t *testing.T) {
		err := newArrList.Set(5, 0)
		assert.Equal(t, dataIsEmptyErr, err)
	})
	newArrList.AddLast(rnd.GetRandomNumber())

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
	t.Run("Another set val's check", func(t *testing.T) {
		for i := 0; i <= 200; i++ {
			err := newArrList.AddLast(rnd.GetRandomNumber())
			if err != nil {
				assert.Error(t, err)
			}
		}

		// Set, Test 1
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
		// Set, Test 2
		randomNum = rnd.GetRandomNumber()
		err = newArrList.Set(randomNum, 113)
		if err != nil {
			assert.Error(t, err)
		}
		arrVal = newArrList.ToArray()[113]
		newVal, err = newArrList.Get(113)
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, arrVal, newVal)
		// Set, Test 3, wrong type of set el
		err = newArrList.Set("5", 113)
		assert.Equal(t, incorrectElementType, err)
	})
}

// Sort. Wits ints
func TestAl_Sort(t *testing.T) {
	newArrList := NewMyArrayList()
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
