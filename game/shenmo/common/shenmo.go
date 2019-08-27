package common

type PlayerShenMoObject struct {
	gongXunNum int32
	killNum    int32
	endTime    int64
}

func (o *PlayerShenMoObject) GetGongXunNum() int32 {
	return o.gongXunNum
}

func (o *PlayerShenMoObject) GetKillNum() int32 {
	return o.killNum
}

func (o *PlayerShenMoObject) GetEndTime() int64 {
	return o.endTime
}

func CreatePlayerShenMoObject(gongXunNum int32, killNum int32, endTime int64) *PlayerShenMoObject {
	o := &PlayerShenMoObject{}
	o.gongXunNum = gongXunNum
	o.killNum = killNum
	o.endTime = endTime
	return o
}
