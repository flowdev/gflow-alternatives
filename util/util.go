package util

import "fmt"

var failNumber, failCount int

// FailUtil returns an error every third time it is called with the same i.
func FailUtil(i int) error {
	if failNumber != i {
		failNumber = i
		failCount = 0
	}
	failCount++

	if failCount == (i % 3) {
		return fmt.Errorf("Its time for an error: %d", i)
	}
	return nil
}
