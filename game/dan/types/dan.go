package types

type AlchemyState int32

const (
	//领取丹药
	AlchemyStateReceive AlchemyState = iota
	//炼丹进行中
	AlchemyStateStart
	//炼丹完成
	AlchemyStateEnd
)
