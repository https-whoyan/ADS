package functions

import (
	"fmt"
)

func comb(s string) []string {
	// O(1)
	n := len(s)
	if n == 1 {
		return []string{s}
	}
	nextCombs := comb(s[1:])
	ans := []string{}
	// O(len(nextCombs)) = O((len(s) - 1)!)
	for _, nextComb := range nextCombs {
		// O(n)
		for indexToAdd := 0; indexToAdd <= n-1; indexToAdd++ {
			strToAdd := nextComb[:indexToAdd] + string(s[0]) + nextComb[indexToAdd:]
			ans = append(ans, strToAdd)
		}
	}
	return ans
}

func Task7() error {
	// O(1)
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		return err
	}
	// O(N!)
	allPermutations := comb(s)
	fmt.Printf("All permutations of %v:\n", s)
	for _, myPermutation := range allPermutations {
		fmt.Println(myPermutation)
	}
	return nil
}
