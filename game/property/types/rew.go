package types

type RewData struct {
	RewExp      int32
	RewExpPoint int32
	RewSilver   int32
	RewGold     int32
	RewBindGold int32
}

func CreateRewData(rewExp int32, rewExpPoint int32, rewSilver int32, rewGold int32, rewBindGold int32) *RewData {
	rd := &RewData{
		RewExp:      rewExp,
		RewExpPoint: rewExpPoint,
		RewSilver:   rewSilver,
		RewGold:     rewGold,
		RewBindGold: rewBindGold,
	}
	return rd
}

func (rd *RewData) GetRewExp() int32 {
	return rd.RewExp
}

func (rd *RewData) GetRewExpPoint() int32 {
	return rd.RewExpPoint
}

func (rd *RewData) GetRewSilver() int32 {
	return rd.RewSilver
}

func (rd *RewData) GetRewGold() int32 {
	return rd.RewGold
}

func (rd *RewData) GetRewBindGold() int32 {
	return rd.RewBindGold
}

func (rd *RewData) AddRewData(d *RewData) {
	if d == nil {
		return
	}
	rd.RewExp += d.RewExp
	rd.RewBindGold += d.RewBindGold
	rd.RewGold += d.RewGold
	rd.RewSilver += d.RewSilver
	rd.RewExpPoint += d.RewExpPoint
}

func (rd *RewData) MultTimes(needTimes int32) {
	if needTimes < 0 {
		return
	}
	rd.RewExp *= needTimes
	rd.RewBindGold *= needTimes
	rd.RewGold *= needTimes
	rd.RewSilver *= needTimes
	rd.RewExpPoint *= needTimes
}

func (rd *RewData) Valid() bool {
	if rd.RewExp < 0 {
		return false
	}
	if rd.RewExpPoint < 0 {
		return false
	}
	if rd.RewSilver < 0 {
		return false
	}
	if rd.RewBindGold < 0 {
		return false
	}
	if rd.RewGold < 0 {
		return false
	}
	return true
}
