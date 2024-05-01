package Random

import (
	"math"
	"math/rand"
)

// Random Formula, IDK :)
func RandomFormula(x int) int {
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

type TestStruct struct {
	Name      string
	Age       int
	IsStudent bool
}

func GetStruct(i int) *TestStruct {
	return &TestStruct{
		Name:      string(rune(i)),
		Age:       RandomFormula(i),
		IsStudent: (i+RandomFormula(i))%2 == 0,
	}
}

func ShuffleMp(mp map[int]int) []int {
	n := len(mp)
	arr := make([]int, 0, n)
	for key, _ := range mp {
		arr = append(arr, key)
	}
	rand.Shuffle(n, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}
