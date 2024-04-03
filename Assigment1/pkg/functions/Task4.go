package functions

import (
	"errors"
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return factorial(n-1) * n
}

func Task4() error {
	// O(1)
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return err
	}
	if n < 0 {
		return errors.New(fmt.Sprintf("%v is less then zero!", n))
	}
	// Factorial - O(N) (Runtime)
	// Factorial - O(N) (Memory)
	fmt.Printf("Factorial of %v, %v", n, factorial(n))
	return nil
}
