package types

type GuaJiTypeMaterialFuBenOptionType int32

const (
	GuaJiTypeMaterialFuBenOptionTypeAutoBuy GuaJiTypeMaterialFuBenOptionType = iota
)

func (t GuaJiTypeMaterialFuBenOptionType) GetType() int32 {
	return int32(t)
}

func (t GuaJiTypeMaterialFuBenOptionType) Valid() bool {
	switch t {
	case GuaJiTypeMaterialFuBenOptionTypeAutoBuy:
		return true
	}
	return false
}

var (
	guaJiTypeMaterialFuBenOptionTypeMap = map[GuaJiTypeMaterialFuBenOptionType]string{
		GuaJiTypeMaterialFuBenOptionTypeAutoBuy: "自动购买材料",
	}
)

func (t GuaJiTypeMaterialFuBenOptionType) String() string {
	return guaJiTypeMaterialFuBenOptionTypeMap[t]
}

func CreateGuaJiTypeMaterialFuBenOptionType(typ int32) GuaJiOptionType {
	return GuaJiTypeMaterialFuBenOptionType(typ)
}

func init() {
	RegisterGuaJiOptionTypeFactory(GuaJiTypeMaterial, GuaJiOptionTypeFactoryFunc(CreateGuaJiTypeMaterialFuBenOptionType))
}
