package signal

import (
	"strconv"
	"strings"
)

func Solution(inputString string) bool {
	array := strings.Split(inputString, ".")

	if len(array) != 4 {
		return false
	}

	for i := 0; i < len(array); i++ {
		n, err := strconv.Atoi(array[i])

		if err != nil {
			return false
		}

		if n < 0 || n > 255 {
			return false
		}

		if n < 10 && len(array[i]) > 1 {
			return false
		}

		if n < 100 && len(array[i]) > 2 {
			return false
		}

		if n > 100 && len(array[i]) != 3 {
			return false
		}

	}

	return true

}
