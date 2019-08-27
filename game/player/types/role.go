package types

import (
	dummytypes "fgame/fgame/game/dummy/types"
	"fgame/fgame/pkg/mathutils"
)

//角色类型
type RoleType int32

const (
	//开天
	RoleTypeKaiTian RoleType = iota + 1
	//奕剑
	RoleTypeYiJian
	//破月
	RoleTypePoYue
)

func (rt RoleType) Valid() bool {
	switch rt {
	case RoleTypeKaiTian:
	case RoleTypePoYue:
	case RoleTypeYiJian:
	default:
		return false
	}
	return true
}

func RandomRole() RoleType {
	index := mathutils.RandomWeights([]int64{1, 1, 1})
	return RoleType(index + 1)
}

//性别
type SexType int32

const (
	SexTypeMan SexType = 1 + iota
	SexTypeWoman
)

func (st SexType) Valid() bool {
	switch st {
	case SexTypeMan:
	case SexTypeWoman:
	default:
		return false
	}
	return true
}

var (
	nameMap = map[SexType]dummytypes.DummyType{
		SexTypeMan:   dummytypes.DummyTypeMaleName,
		SexTypeWoman: dummytypes.DummyTypeFemaleName,
	}
)

func (st SexType) DummyType() dummytypes.DummyType {
	return nameMap[st]
}

func RandomSex() SexType {
	flag := mathutils.RandomOneHit(0.5)
	if flag {
		return SexTypeMan
	}
	return SexTypeWoman
}
