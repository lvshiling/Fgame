package types

type EmperorEventType string

const (
	//帝王被抢
	EmperorEventTypeRobed EmperorEventType = "EmperorRobed"
	//抢龙椅合服
	EmperorMergeServer EmperorEventType = "EmperorMergeServer"
)
