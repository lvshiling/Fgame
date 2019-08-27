package types

type TransportationType int32

const (
	TransportationTypeSilver TransportationType = iota + 1
	TransportationTypeGold
	TransportationTypeAlliance
)

var (
	transportationTypeMap = map[TransportationType]string{
		TransportationTypeSilver:   "银两",
		TransportationTypeGold:     "绑元",
		TransportationTypeAlliance: "仙盟",
	}
)

func (t TransportationType) Valid() bool {
	switch t {
	case TransportationTypeSilver,
		TransportationTypeGold,
		TransportationTypeAlliance:
		return true
	default:
		return false
	}
}

func (t TransportationType) String() string {
	return transportationTypeMap[t]
}

//镖车状态
type TransportStateType int32

const (
	TransportStateTypeRuning TransportStateType = iota //运输中
	TransportStateTypeFail                             //失败
	TransportStateTypeFinish                           //完成
)
