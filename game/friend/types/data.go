package types

//表白记录数据
type MarryDevelopLogData struct {
	SendId     int64
	RecvId     int64
	SendName   string
	RecvName   string
	ItemId     int32
	ItemNum    int32
	CharmNum   int32
	DevelopExp int32
	ContextStr string
}

func NewAllMarryDevelopData(sendId int64, recvId int64, sendName string, recvName string, itemId int32, itemNum int32, charmNum int32, developExp int32, contextStr string) *MarryDevelopLogData {
	d := &MarryDevelopLogData{
		SendId:     sendId,
		RecvId:     recvId,
		SendName:   sendName,
		RecvName:   recvName,
		ItemId:     itemId,
		ItemNum:    itemNum,
		CharmNum:   charmNum,
		DevelopExp: developExp,
		ContextStr: contextStr,
	}
	return d
}
