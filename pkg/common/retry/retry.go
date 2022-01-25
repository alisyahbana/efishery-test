package retry

import (
	"fmt"
	"time"
)

//output args1 bool => want to retry?
//output args2 error => error message
type RetryFunc func() (bool, error)

func Retry(numRetry int, delay time.Duration, f RetryFunc) error {
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
		time.Sleep(delay)
	}
}
