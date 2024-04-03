package functions

import (
	"fmt"
	"math"
)

func Task1() error {
	// O(1)
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return err
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
	minNum := math.MaxInt
	for _, num := range arr {
		minNum = min(minNum, num)
	}
	fmt.Printf("Minumum number of %v arr: %v", arr, minNum)
	return nil
}
