package types

//vip信息
type VipInfo struct {
	VipLevel int32 `json:"viplevel"`
	VipStar  int32 `json:"vipStar"`
}

// 升阶系统升阶系数
type AdvancedRateData struct {
	maxRate int32 //最大次数系数
	minRate int32 //最小次数系数
}

func CreateAdvancedRateData(min, max int32) *AdvancedRateData {
	d := &AdvancedRateData{
		maxRate: max,
		minRate: min,
	}

	return d
}

func (d *AdvancedRateData) GetMaxRate() int32 {
	return d.maxRate
}

func (d *AdvancedRateData) GetMinRate() int32 {
	return d.minRate
}
