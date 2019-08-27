package common

type PlayerDenseWatObject struct {
	num     int32
	endTime int64
}

func (o *PlayerDenseWatObject) GetNum() int32 {
	return o.num
}

func (o *PlayerDenseWatObject) GetEndTime() int64 {
	return o.endTime
}

func CreatePlayerDenseWatObject(num int32, endTime int64) *PlayerDenseWatObject {
	o := &PlayerDenseWatObject{}
	o.num = num
	o.endTime = endTime
	return o
}
