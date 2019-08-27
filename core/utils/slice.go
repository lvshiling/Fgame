package utils

func ContainInt32(arr []int32, num int32) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
	}
	return false
}

func ContainInt64(arr []int64, num int64) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
	}
	return false
}

func IfRepeatElementInt32(arr []int32) bool {
	newMap := make(map[int32]int32)
	for index := 0; index < len(arr); index++ {
		val := arr[index]
		_, isExist := newMap[val]
		if isExist {
			return true
		}

		newMap[val] = val
	}

	return false
}
