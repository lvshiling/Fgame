package types

type LingTongFashionType int32

const (
	//普通类型
	LingTongFashionTypeNormal LingTongFashionType = 1 + iota
	//时效性时装
	LingTongFashionTypeEffective
)

func (ft LingTongFashionType) Valid() bool {
	switch ft {
	case LingTongFashionTypeNormal,
		LingTongFashionTypeEffective:
		break
	default:
		return false
	}
	return true
}

// 试用卡过期类型
type LingTongFashionTrialOverdueType int32

const (
	LingTongFashionTrialOverdueTypeActivate LingTongFashionTrialOverdueType = iota + 1 //激活过期
	LingTongFashionTrialOverdueTypeExpire                                              //时效过期
)
