package utils

import (
	"strconv"
)

func ConverStringToInt64(p_string string) int64 {
	rst, _ := strconv.ParseInt(p_string, 10, 64)
	return rst
}

func ConverStringToInt64Error(p_string string) (int64, error) {
	rst, err := strconv.ParseInt(p_string, 10, 64)
	return rst, err
}

func ConverInt64ToString(p_id int64) string {
	return strconv.FormatInt(p_id, 10)
}

func ConverStringToInt32Error(p_string string) (int32, error) {
	rst, err := strconv.ParseInt(p_string, 10, 64)
	return int32(rst), err
}
