package types

const (
	RECORD_MAXLEN = 50
)

type OneArenaLevelType int32

const (
	//黄级灵池
	OneArenaLevelTypeHuangJi OneArenaLevelType = 1 + iota
	//玄级灵池
	OneArenaLevelTypeXuanJi
	//地级灵池
	OneArenaLevelTypeDiJi
	//天级灵池
	OneArenaLevelTypeTianJi
	//太清灵池
	OneArenaLevelTypeTaiQing
)

const (
	OneArenaLevelMin = OneArenaLevelTypeHuangJi
	OneArenaLevelMax = OneArenaLevelTypeTaiQing
)

func (o OneArenaLevelType) Vaild() bool {
	switch o {
	case OneArenaLevelTypeHuangJi,
		OneArenaLevelTypeXuanJi,
		OneArenaLevelTypeDiJi,
		OneArenaLevelTypeTianJi,
		OneArenaLevelTypeTaiQing:
		return true
	}
	return false
}

//灵池被抢状态
type OneArenaRobbedStatusType int32

const (
	//被抢走了
	OneArenaRobbedStatusTypeSucess OneArenaRobbedStatusType = 1 + iota
	//没被抢走
	OneArenaRobbedStatusTypeFail
)
