package types

const (
	//首冲激活矿工
	RechargeActive = int32(2)
)

type GemGambleType int32

const (
	//初级赌石
	GemGambleTypePrimary GemGambleType = 1 + iota
	//高级赌石
	GemGambleTypeSenior
)

func (gt GemGambleType) Valid() bool {
	switch gt {
	case GemGambleTypePrimary,
		GemGambleTypeSenior:
		break
	default:
		return false
	}
	return true
}
