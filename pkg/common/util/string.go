package util

import (
	"strconv"
)

func StringToUint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}
