package types

type MarryWedCodeType int32

const (
	//成功
	MarryCodeTypeWeddingSucess MarryWedCodeType = 1 + iota
	//该婚期场次已被抢先一步预定
	MarryCodeTypeWeddingExist
	//预定的举办场次距离开始时间小于1分钟
	MarryCodeTypeWeddingTimeLimit
)
