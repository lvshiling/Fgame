package types

type GuaJiTypeUnrealBossOptionType int32

const (
	GuaJiTypeUnrealBossOptionTypeAutoBuy GuaJiTypeUnrealBossOptionType = iota
)

func (t GuaJiTypeUnrealBossOptionType) GetType() int32 {
	return int32(t)
}

func (t GuaJiTypeUnrealBossOptionType) Valid() bool {
	switch t {
	case GuaJiTypeUnrealBossOptionTypeAutoBuy:
		return true
	}
	return false
}

var (
	guaJiTypeUnrealBossOptionTypeMap = map[GuaJiTypeUnrealBossOptionType]string{
		GuaJiTypeUnrealBossOptionTypeAutoBuy: "自动购买疲劳值",
	}
)

func (t GuaJiTypeUnrealBossOptionType) String() string {
	return guaJiTypeUnrealBossOptionTypeMap[t]
}

func CreateGuaJiTypeUnrealBossOptionType(typ int32) GuaJiOptionType {
	return GuaJiTypeUnrealBossOptionType(typ)
}

func init() {
	RegisterGuaJiOptionTypeFactory(GuaJiTypeUnrealBoss, GuaJiOptionTypeFactoryFunc(CreateGuaJiTypeUnrealBossOptionType))
}
