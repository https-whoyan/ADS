package functions

import (
	"errors"
	"fmt"
)

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func Task5() error {
	// O(1)
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return err
	}
	if n < 0 {
		return errors.New(fmt.Sprintf("%v is less then zero!", n))
	}
	// Fibonacci - Runtime: O(1,6 ^ (N))
	// Fibonacci - Memory: O(1,6 ^ (N))
	fmt.Printf("Fibonacci's %v number: %v", n, fibonacci(n))
	return nil
}
