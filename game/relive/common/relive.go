package common

type PlayerReliveObject struct {
	culTime        int32
	lastReliveTime int64
}

func (o *PlayerReliveObject) GetCulTime() int32 {
	return o.culTime
}

func (o *PlayerReliveObject) GetLastReliveTime() int64 {
	return o.lastReliveTime
}

func CreatePlayerReliveObject(culTime int32, lastReliveTime int64) *PlayerReliveObject {
	o := &PlayerReliveObject{
		culTime:        culTime,
		lastReliveTime: lastReliveTime,
	}
	return o
}
