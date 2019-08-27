package common

import (
	"fmt"
	"strconv"
)

func ConverStringToInt64(p_string string) int64 {
	rst, _ := strconv.ParseInt(p_string, 10, 64)
	return rst
}

func CombinInt64Array(p_array []int64) string {
	if len(p_array) == 0 {
		return ""
	}
	result := ""
	for index, value := range p_array {
		if index != 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", value)
	}
	return result
}

func CombinIntArray(p_array []int) string {
	if len(p_array) == 0 {
		return ""
	}
	result := ""
	for index, value := range p_array {
		if index != 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", value)
	}
	return result
}
