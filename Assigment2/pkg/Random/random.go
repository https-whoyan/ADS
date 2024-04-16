package Random

import (
	"math/rand"
	"strconv"
)

// Random number

const maxRandN = 200
const minRandN = -200

func GetRandomNumber() int {
	return rand.Intn(maxRandN-minRandN) + minRandN
}

// A Some random structs

type RandStructA struct {
	RandFieldA_str         string
	RandFieldB_float64     float64
	RandFieldC_ptr_structB *RandStructB
}

type RandStructB struct {
	RandFieldA_int         int
	RandFieldB_slice_ptr_B []*RandStructB
	RandFieldC_struct_C    RandStructC
}

type RandStructC struct {
	RandFieldA_str string
}

// GetExample Func's

func GetExampleA() RandStructA {
	strRandomNum := strconv.Itoa(GetRandomNumber())
	exampleStructB := GetExampleB()
	return RandStructA{
		strRandomNum,
		float64(GetRandomNumber()),
		&exampleStructB,
	}
}

func GetExampleB() RandStructB {
	exampleStructC := GetExampleC()
	return RandStructB{
		GetRandomNumber(),
		[]*RandStructB{},
		exampleStructC,
	}
}

func GetExampleC() RandStructC {
	strRandomNum := strconv.Itoa(GetRandomNumber())
	return RandStructC{
		strRandomNum,
	}
}
