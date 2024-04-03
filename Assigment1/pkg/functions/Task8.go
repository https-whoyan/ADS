package functions

import "fmt"

func Task8() error {
	// O(1)
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		return err
	}
	isOk := true
	for _, myRune := range s {
		if !(myRune >= '0' && myRune <= '9') {
			isOk = false
			break
		}
	}
	if isOk {
		fmt.Printf("%v contains only digits", s)
	} else {
		fmt.Printf("%v not contains only digits", s)
	}
	return nil
}
