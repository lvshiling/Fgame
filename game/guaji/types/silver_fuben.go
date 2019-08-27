package types

type GuaJiTypeSilverFuBenOptionType int32

const (
	GuaJiTypeSilverFuBenOptionTypeAutoBuy GuaJiTypeSilverFuBenOptionType = iota
)

func (t GuaJiTypeSilverFuBenOptionType) GetType() int32 {
	return int32(t)
}

func (t GuaJiTypeSilverFuBenOptionType) Valid() bool {
	switch t {
	case GuaJiTypeSilverFuBenOptionTypeAutoBuy:
		return true
	}
	return false
}

var (
	guaJiTypeSilverFuBenOptionTypeMap = map[GuaJiTypeSilverFuBenOptionType]string{
		GuaJiTypeSilverFuBenOptionTypeAutoBuy: "自动购买材料",
	}
)

func (t GuaJiTypeSilverFuBenOptionType) String() string {
	return guaJiTypeSilverFuBenOptionTypeMap[t]
}

func CreateGuaJiTypeSilverFuBenOptionType(typ int32) GuaJiOptionType {
	return GuaJiTypeSilverFuBenOptionType(typ)
}

func init() {
	RegisterGuaJiOptionTypeFactory(GuaJiTypeXianFuSilver, GuaJiOptionTypeFactoryFunc(CreateGuaJiTypeSilverFuBenOptionType))
}
