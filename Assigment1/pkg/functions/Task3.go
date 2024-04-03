package functions

import (
	"fmt"
	"math"
)

func Task3() error {
	// O(1)
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return nil
	}
	if n <= 0 {
		fmt.Printf("%v is not natural number", n)
		return nil
	}
	if n == 1 {
		fmt.Printf("%v can't be determine, a it is prime or not", n)
		return nil
	}
	sqrtN := int(math.Sqrt(float64(n)))
	// O(Sqrt(N))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			fmt.Printf("%v is Composite", n)
			return nil
		}
	}
	fmt.Printf("%v is Prime", n)
	return nil
}
