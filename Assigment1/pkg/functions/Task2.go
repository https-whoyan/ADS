package functions

import (
	"errors"
	"fmt"
)

func Task2() error {
	// O(1)
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("can't find avg value of zero items")
	}
	// O(n)
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		_, err = fmt.Scan(&arr[i])
		if err != nil {
			return err
		}
	}
	// O(n)
	sumOfNums := 0
	for _, num := range arr {
		sumOfNums += num
	}
	avgValue := float64(sumOfNums) / float64(n)
	fmt.Printf("A avg value of %v arr: %v", arr, avgValue)
	return nil
}
