package types

type GuaJiTypeExpFuBenOptionType int32

const (
	GuaJiTypeExpFuBenOptionTypeAutoBuy GuaJiTypeExpFuBenOptionType = iota
)

func (t GuaJiTypeExpFuBenOptionType) GetType() int32 {
	return int32(t)
}

func (t GuaJiTypeExpFuBenOptionType) Valid() bool {
	switch t {
	case GuaJiTypeExpFuBenOptionTypeAutoBuy:
		return true
	}
	return false
}

var (
	guaJiTypeExpFuBenOptionTypeMap = map[GuaJiTypeExpFuBenOptionType]string{
		GuaJiTypeExpFuBenOptionTypeAutoBuy: "自动购买材料",
	}
)

func (t GuaJiTypeExpFuBenOptionType) String() string {
	return guaJiTypeExpFuBenOptionTypeMap[t]
}

func CreateGuaJiTypeExpFuBenOptionType(typ int32) GuaJiOptionType {
	return GuaJiTypeExpFuBenOptionType(typ)
}

func init() {
	RegisterGuaJiOptionTypeFactory(GuaJiTypeXianFuExp, GuaJiOptionTypeFactoryFunc(CreateGuaJiTypeExpFuBenOptionType))
}
