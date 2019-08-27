package types

type TianShuType int32

const (
	TianShuTypeSilver    TianShuType = iota + 1 //财富天书
	TianShuTypeExp                              //经验天书
	TianShuTypeDrop                             //打宝天书
	TianShuTypeAdvanced                         //祝福天书
	TianShuTypeBoss                             //BOSS天书
	TianShuTypeBindGold                         //绑元天书
	TianShuTypeGold                             //元宝天书
	TianShuTypeChuangShi                        //创世天书
)

func (t TianShuType) Valid() bool {
	switch t {
	case TianShuTypeSilver,
		TianShuTypeExp,
		TianShuTypeDrop,
		TianShuTypeAdvanced,
		TianShuTypeBoss,
		TianShuTypeBindGold,
		TianShuTypeGold,
		TianShuTypeChuangShi:
		return true
	}

	return false
}

var (
	tianshuMap = map[TianShuType]string{
		TianShuTypeSilver:    "财富天书",
		TianShuTypeExp:       "经验天书",
		TianShuTypeDrop:      "打宝天书",
		TianShuTypeAdvanced:  "祝福天书",
		TianShuTypeBoss:      "BOSS天书",
		TianShuTypeBindGold:  "绑元天书",
		TianShuTypeGold:      "元宝天书",
		TianShuTypeChuangShi: "创世天书",
	}
)

func (t TianShuType) String() string {
	return tianshuMap[t]
}
