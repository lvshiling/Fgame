package utils

//乘系数
func MultMap(data map[int32]int32, ratio int32) map[int32]int32 {
	newMap := make(map[int32]int32)
	for key, val := range data {
		newMap[key] = val * ratio
	}

	return newMap
}

// 合并
func MergeMap(parent, child map[int32]int32) map[int32]int32 {
	for key, val := range child {
		_, ok := parent[key]
		if !ok {
			parent[key] = val
		} else {
			parent[key] += val
		}
	}

	return parent
}
