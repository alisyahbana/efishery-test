package lock

import (
	"fmt"
	"time"
)

//output args1 bool => want to retry?
//output args2 error => error message
type TransactionFunc func() (bool, error)

func Lock(numRetry int, f TransactionFunc) error {
	i := 1
	for {
		retry, err := f()

		if retry {
			if i >= numRetry {
				return fmt.Errorf("Exceeded %v retries with error %s", retry, err.Error())
			}
		} else {
			return err
		}

		i++
		time.Sleep(10 * time.Millisecond)
	}
}
