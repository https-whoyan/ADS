package functions

import (
	"errors"
	"fmt"
)

func binPow(p int, st int) int {
	if st == 0 {
		return 1
	}
	if st%2 == 0 {
		sqrtAns := binPow(p, st/2)
		return sqrtAns * sqrtAns
	}
	return binPow(p, st-1) * p
}

func Task6() error {
	// O(1)
	var (
		a, n int
	)
	_, err := fmt.Scan(&a)
	if err != nil {
		return err
	}
	_, err = fmt.Scan(&n)
	if err != nil {
		return err
	}
	if n < 0 {
		return errors.New(fmt.Sprintf("power of %v must be >= than zero! Got %v", a, n))
	}
	// Binary Pow - Runtime: O(log(N))
	// Binary Pow - Memory: O(log(N))
	fmt.Printf("%v^%v: %v", a, n, binPow(a, n))
	return nil
}
