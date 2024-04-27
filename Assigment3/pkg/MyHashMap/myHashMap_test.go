package MyHashMap

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

const (
	testingCap         = 200
	testingAddElsCount = 1123
)

// Random Formula, IDK :)
func randomFormula(x int) int {
	sqrtX := int(math.Sqrt(float64(x)))
	a := x * (x + 3)
	b := sqrtX - 5*x
	myAbs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}
	c := myAbs(a-b) - b/2
	d := sqrtX * (sqrtX - 6)
	e := c*(sqrtX-6) - d
	randomSqrt := int(math.Sqrt(float64(sqrtX + 8)))
	f := e - x*b - c*randomSqrt
	return f
}

func TestMyHashMap_Put(t *testing.T) {
	t.Run("Put, not error?", func(t *testing.T) {
		var err error
		defer func() {
			if r := recover(); r != nil {
				err = errors.New("put func getError")
			}
		}()
		mp := NewMyHashMap[int, int](testingCap)
		// Add Some values
		for i := 1; i <= testingAddElsCount; i++ {
			mp.Put(i, randomFormula(i))
		}
		assert.Equal(t, nil, err)
	})
	t.Run("Checking size", func(t *testing.T) {
		mp := NewMyHashMap[int, int](testingCap)
		exceptedSize := 0
		// Add Some values
		for i := 1; i <= testingAddElsCount; i++ {
			haveACollision, _, _ := mp.Put(i, randomFormula(i))
			if !haveACollision {
				exceptedSize++
			}
		}

		currSize := mp.Size()
		assert.Equal(t, exceptedSize, currSize)
	})
}

func TestMyHashMap_Get(t *testing.T) {
	// Add values
	mp := NewMyHashMap[int, int](testingCap)
	collisionsMp := make(map[int]bool)
	for i := 1; i <= testingAddElsCount; i++ {
		// if a have a collision, I'll stand that 'ok, oldKey have an incorrect val, then
		// I'll skip this key for checking correctly check.
		haveACollision, oldKey, _ := mp.Put(i, randomFormula(i))
		if haveACollision {
			collisionsMp[oldKey] = true
		}
	}
	t.Run("Test all not collision keys", func(t *testing.T) {
		exceptedMp := make(map[int]int)
		for i := 1; i <= testingAddElsCount; i++ {
			if _, contains := collisionsMp[i]; !contains {
				exceptedMp[i] = randomFormula(i)
			}
		}
		actualMp := make(map[int]int)
		for i := 1; i <= testingAddElsCount; i++ {
			if _, contains := collisionsMp[i]; !contains {
				val, isContains := mp.Get(i)
				if !isContains {
					assert.Error(t, errors.New("have a unexpected not key contains err"))
				}
				actualMp[i] = val
			}
		}
		assert.Equal(t, exceptedMp, actualMp)
	})
	t.Run("Get value of not contains key", func(t *testing.T) {
		mp.Clear()
		mp.Put(5, 6)
		_, isContains := mp.Get(19)
		assert.Equal(t, false, isContains)
	})
}

func TestMyHashMap_Contains(t *testing.T) {
	// Add Values
	mp := NewMyHashMap[int, int](testingCap)
	collisionMp := make(map[int]bool)
	for i := 1; i <= testingAddElsCount; i++ {
		isCollision, oldKey, _ := mp.Put(i, randomFormula(i))
		if isCollision {
			collisionMp[oldKey] = true
		}
	}

	FGetExceptedContains := func(testingKey int, isCollision bool) bool {
		if isCollision {
			return false
		}
		return (testingKey >= 1) && (testingKey <= testingAddElsCount)
	}

	universalTest := func(mp *MyHashMap[int, int], testingKey int) (excepted, actual bool) {
		isContains := mp.Contains(testingKey)
		_, containsInCollisionMp := collisionMp[testingKey]
		exceptedContains := FGetExceptedContains(testingKey, containsInCollisionMp)
		return exceptedContains, isContains
	}

	t.Run("Contains key, test1", func(t *testing.T) {
		testingKey := 1
		ext, act := universalTest(mp, testingKey)
		assert.Equal(t, ext, act)
	})
	t.Run("Contains key, test2", func(t *testing.T) {
		testingKey := 124
		ext, act := universalTest(mp, testingKey)
		assert.Equal(t, ext, act)
	})
	t.Run("Contains key, test3", func(t *testing.T) {
		testingKey := 43221
		ext, act := universalTest(mp, testingKey)
		assert.Equal(t, ext, act)
	})
	t.Run("Contains key, test4", func(t *testing.T) {
		testingKey := -19284
		ext, act := universalTest(mp, testingKey)
		assert.Equal(t, ext, act)
	})
}

func TestMyHashMap_Remove(t *testing.T) {
	// Add Values
	mp := NewMyHashMap[int, int](testingCap)
	for i := 1; i <= testingAddElsCount; i++ {
		mp.Put(i, randomFormula(i))
	}

	// I'll be checking before containing because I don't sure
	// That mp don't have a collisions.
	universalTest := func(hashMap *MyHashMap[int, int], removedKey int) (notContains bool) {
		mp.Remove(removedKey)
		_, afterIsContaining := mp.Get(removedKey)
		notContains = afterIsContaining == false
		return
	}
	// universalTest was written for less size code.
	t.Run("Remove, test1", func(t *testing.T) {
		act := universalTest(mp, 1)
		assert.Equal(t, true, act)
	})
	t.Run("Remove, test2", func(t *testing.T) {
		act := universalTest(mp, 2)
		assert.Equal(t, true, act)
	})
	t.Run("Remove, test3", func(t *testing.T) {
		act := universalTest(mp, -23498)
		assert.Equal(t, true, act)
	})
	t.Run("Remove, test4", func(t *testing.T) {
		act := universalTest(mp, 2)
		assert.Equal(t, true, act)
	})
	t.Run("Remove, test5", func(t *testing.T) {
		act := universalTest(mp, 239298)
		assert.Equal(t, true, act)
	})
}

func TestWithStructs(t *testing.T) {
	type testStruct struct {
		name      string
		age       int
		isStudent bool
	}
	getStruct := func(i int) *testStruct {
		return &testStruct{
			name:      string(rune(i)),
			age:       randomFormula(i),
			isStudent: (i+randomFormula(i))%2 == 0,
		}
	}

	getMp := func() *MyHashMap[int, *testStruct] {
		return NewMyHashMap[int, *testStruct](testingCap)
	}

	t.Run("Put struct elements, isOk?", func(t *testing.T) {
		mp := getMp()
		recoveredPanicErr := false
		defer func() {
			if r := recover(); r != nil {
				recoveredPanicErr = true
			}
		}()
		for i := 1; i <= testingAddElsCount; i++ {

			mp.Put(i, getStruct(i))
		}

		assert.Equal(t, false, recoveredPanicErr)
	})
	t.Run("Get struct elements", func(t *testing.T) {
		mp := getMp()
		collisionsMp := make(map[int]bool)
		for i := 1; i <= testingAddElsCount; i++ {
			haveACollision, oldKey, _ := mp.Put(i, getStruct(i))
			if haveACollision {
				collisionsMp[oldKey] = true
			}
		}
		exceptedMp := make(map[int]*testStruct)
		for i := 1; i <= testingAddElsCount; i++ {
			if _, contains := collisionsMp[i]; !contains {
				exceptedMp[i] = getStruct(i)
			}
		}
		actualMp := make(map[int]*testStruct)
		for i := 1; i <= testingAddElsCount; i++ {
			if _, contains := collisionsMp[i]; !contains {
				val, isContains := mp.Get(i)
				if !isContains {
					assert.Error(t, errors.New("have a unexpected not key contains err"))
				}
				actualMp[i] = val
			}
		}
		assert.Equal(t, exceptedMp, actualMp)
	})
	t.Run("Contains with struct", func(t *testing.T) {
		mp := getMp()
		collisionsMp := make(map[int]bool)
		for i := 1; i <= testingAddElsCount; i++ {
			haveACollision, oldKey, _ := mp.Put(i, getStruct(i))
			if haveACollision {
				collisionsMp[oldKey] = true
			}
		}

		FGetExceptedContains := func(testingKey int, isCollision bool) bool {
			if isCollision {
				return false
			}
			return (testingKey >= 1) && (testingKey <= testingAddElsCount)
		}

		universalTest := func(mp *MyHashMap[int, *testStruct], testingKey int) (excepted, actual bool) {
			isContains := mp.Contains(testingKey)
			_, containsInCollisionMp := collisionsMp[testingKey]
			exceptedContains := FGetExceptedContains(testingKey, containsInCollisionMp)
			return exceptedContains, isContains
		}

		t.Run("Contains key, test1", func(t *testing.T) {
			testingKey := 1
			ext, act := universalTest(mp, testingKey)
			assert.Equal(t, ext, act)
		})
		t.Run("Contains key, test2", func(t *testing.T) {
			testingKey := 124
			ext, act := universalTest(mp, testingKey)
			assert.Equal(t, ext, act)
		})
		t.Run("Contains key, test3", func(t *testing.T) {
			testingKey := 43221
			ext, act := universalTest(mp, testingKey)
			assert.Equal(t, ext, act)
		})
		t.Run("Contains key, test4", func(t *testing.T) {
			testingKey := -19284
			ext, act := universalTest(mp, testingKey)
			assert.Equal(t, ext, act)
		})
	})
}
