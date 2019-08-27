package common

type PlayerXueChiObject struct {
	blood     int64
	bloodLine int32
}

func (o *PlayerXueChiObject) GetBlood() int64 {
	return o.blood
}

func (o *PlayerXueChiObject) GetBloodLine() int32 {
	return o.bloodLine
}

func CreatePlayerXueChiObject(blood int64, bloodLine int32) *PlayerXueChiObject {
	o := &PlayerXueChiObject{}
	o.blood = blood
	o.bloodLine = bloodLine
	return o
}
