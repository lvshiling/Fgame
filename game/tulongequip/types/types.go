package types

type UpstarResultType int32

const (
	UpstarResultTypeSuccess UpstarResultType = iota //强化成功
	UpstarResultTypeFailed                          //强化失败
	UpstarResultTypeBack                            //强化回退
)

type TuLongSuitType int32

const (
	TuLongSuitTypeLongYa  TuLongSuitType = iota //龙牙
	TuLongSuitTypeLongLin                       //龙鳞
	TuLongSuitTypeLongXie                       //龙血
	TuLongSuitTypeLongDan                       //龙胆
	TuLongSuitTypeLongPo                        //龙魄
)

func (t TuLongSuitType) Valid() bool {
	switch t {
	case TuLongSuitTypeLongYa,
		TuLongSuitTypeLongLin,
		TuLongSuitTypeLongXie,
		TuLongSuitTypeLongDan,
		TuLongSuitTypeLongPo:
		break
	default:
		return false
	}
	return true
}

const (
	MinSuitType = TuLongSuitTypeLongYa
	MaxSuitType = TuLongSuitTypeLongPo
)
