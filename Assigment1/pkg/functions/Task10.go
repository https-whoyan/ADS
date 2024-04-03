package functions

import (
	"errors"
	"fmt"
)

func getGCD(a, b int) int {
	if a*b == 0 {
		return a + b
	}
	maxNum, minNum := max(a, b), min(a, b)
	return getGCD(maxNum%minNum, minNum)
}

func Task10() error {
	// O(1)
	var (
		a, b int
	)
	_, err := fmt.Scan(&a)
	if err != nil {
		return err
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		return err
	}
	if a < 0 || b < 0 {
		return errors.New(fmt.Sprintf("both of n and k must be more, than 0. Got a = %v, b = %v", a, b))
	}
	// O(log(min(a, b))
	fmt.Printf("GCD of %v and %v: %v", a, b, getGCD(a, b))
	return nil
}
