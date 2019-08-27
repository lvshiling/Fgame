package types

type FourGodSpecialPosType int32

const (
	//特殊怪默认出生点
	FourGodSpecialPosTypeBorn FourGodSpecialPosType = 1 + iota
	//特殊怪终点
	FourGodSpecialPosTypeEnd
	//Boss出生点
	FourGodBossBorn
)

func (f FourGodSpecialPosType) Valid() bool {
	switch f {
	case FourGodSpecialPosTypeBorn,
		FourGodSpecialPosTypeEnd,
		FourGodBossBorn:
		return true
	}
	return false
}

type FourGodBoxType int32

const (
	//小宝箱
	FourGodBoxTypeSmall FourGodBoxType = iota
	//中宝箱
	FourGodBoxTypeMedium
	//大宝箱
	FourGodBoxTypeBig
)

func (f FourGodBoxType) Valid() bool {
	switch f {
	case FourGodBoxTypeSmall,
		FourGodBoxTypeMedium,
		FourGodBoxTypeBig:
		return true
	}
	return false
}
