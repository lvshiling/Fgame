package types

type GuaJiTypeDailyQuestOptionType int32

const (
	GuaJiTypeDailyQuestOptionDouble GuaJiTypeDailyQuestOptionType = iota
)

func (t GuaJiTypeDailyQuestOptionType) GetType() int32 {
	return int32(t)
}

func (t GuaJiTypeDailyQuestOptionType) Valid() bool {
	switch t {
	case GuaJiTypeDailyQuestOptionDouble:
		return true
	}
	return false
}

var (
	guaJiTypeDailyQuestOptionTypeMap = map[GuaJiTypeDailyQuestOptionType]string{
		GuaJiTypeDailyQuestOptionDouble: "双倍领取",
	}
)

func (t GuaJiTypeDailyQuestOptionType) String() string {
	return guaJiTypeDailyQuestOptionTypeMap[t]
}

func CreateGuaJiTypeDailyQuestOptionType(typ int32) GuaJiOptionType {
	return GuaJiTypeDailyQuestOptionType(typ)
}

func init() {
	RegisterGuaJiOptionTypeFactory(GuaJiTypeDailyQuest, GuaJiOptionTypeFactoryFunc(CreateGuaJiTypeDailyQuestOptionType))
}
