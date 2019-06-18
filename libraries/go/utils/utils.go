package utils

import (
	"strconv"
)

func Atoi(str string) int {
	id, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return id
}
