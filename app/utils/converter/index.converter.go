package converter

import (
	"strconv"
)

func StringToInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}