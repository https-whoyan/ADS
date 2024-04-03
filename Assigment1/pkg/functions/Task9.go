package functions

import (
	"errors"
	"fmt"
)

func mathBinomialCoefficient(k, n int) int {
	if k == 0 || k == n {
		return 1
	}
	fPart := mathBinomialCoefficient(k-1, n-1)
	sPart := mathBinomialCoefficient(k, n-1)
	return fPart + sPart
}

func Task9() error {
	// O(1)
	var (
		k, n int
	)
	_, err := fmt.Scan(&n)
	if err != nil {
		return err
	}
	_, err = fmt.Scan(&k)
	if err != nil {
		return err
	}
	if n < 0 || k < 0 {
		return errors.New(fmt.Sprintf("both of n and k must be more, than 0. Got n = %v, k = %v", n, k))
	}
	// O(n + k)
	fmt.Printf("C (k = %v, n = %v) = %v", k, n, mathBinomialCoefficient(k, n))
	return nil
}
