package buff

type BuffExpEventData struct {
	buffId   int32
	exp      int32
	expPoint int32
}

func (d *BuffExpEventData) GetBuffId() int32 {
	return d.buffId
}

func (d *BuffExpEventData) GetExp() int32 {
	return d.exp
}

func (d *BuffExpEventData) GetExpPoint() int32 {
	return d.expPoint
}

func CreateBuffExpEventData(buffId, exp, expPoint int32) *BuffExpEventData {
	d := &BuffExpEventData{
		buffId:   buffId,
		exp:      exp,
		expPoint: expPoint,
	}
	return d
}
