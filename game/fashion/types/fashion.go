package types

type FashionType int32

const (
	//普通类型
	FashionTypeNormal FashionType = 1 + iota
	//时效性时装
	FashionTypeEffective
	//阵营时装
	FashionTypeCamp
)

func (ft FashionType) Valid() bool {
	switch ft {
	case FashionTypeNormal,
		FashionTypeEffective,
		FashionTypeCamp:
		break
	default:
		return false
	}
	return true
}

// 试用卡过期类型
type FashionTrialOverdueType int32

const (
	FashionTrialOverdueTypeActivate FashionTrialOverdueType = iota + 1 //激活过期
	FashionTrialOverdueTypeExpire                                      //时效过期
)
