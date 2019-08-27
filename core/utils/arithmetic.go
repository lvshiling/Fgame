package utils

// x的n次方
func Pow(x int32, n uint32) int32 {
	ret := int32(1) // 结果初始为0次方的值，整数0次方为1。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
