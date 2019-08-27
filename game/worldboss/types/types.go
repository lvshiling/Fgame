package types

import (
	crosstypes "fgame/fgame/game/cross/types"
)

//世界boss类型
type WorldBossType int32

const (
	WorldBossTypeLocal WorldBossType = 1 //本服世界boss
	WorldBossTypeCross               = 2 //跨服世界boss
)

func (t WorldBossType) Valid() bool {
	switch t {
	case WorldBossTypeCross,
		WorldBossTypeLocal:
		return true
	default:
		return false
	}
}

//Boss类型
type BossType int32

const (
	BossTypeWorldBoss   BossType = iota //0世界boss
	BossTypeShareBoss                   //1跨服boss
	BossTypeUnrealBoss                  //2幻境boss
	BossTypeOutlandBoss                 //3外域boss
	BossTypeCangJingGe                  //4藏经阁boss
	BossTypeZhenXi                      //5珍惜
	BossTypeDingShi                     //6定时boss
	BossTypeArena                       //7圣兽boss
)

const (
	MinBossType = BossTypeWorldBoss
	MaxBossType = BossTypeArena
)

func (t BossType) Valid() bool {
	switch t {
	case BossTypeWorldBoss,
		BossTypeShareBoss,
		BossTypeUnrealBoss,
		BossTypeOutlandBoss,
		BossTypeCangJingGe,
		BossTypeZhenXi,
		BossTypeDingShi,
		BossTypeArena:
		return true
	default:
		return false
	}
}

var (
	bossMap = map[BossType]string{
		BossTypeWorldBoss:   "世界boss",
		BossTypeShareBoss:   "跨服世界boss",
		BossTypeUnrealBoss:  "幻境boss",
		BossTypeOutlandBoss: "外域boss",
		BossTypeCangJingGe:  "藏经阁boss",
		BossTypeZhenXi:      "珍稀boss",
		BossTypeDingShi:     "定时boss",
		BossTypeArena:       "圣兽boss",
	}
)

func (t BossType) String() string {
	return bossMap[t]
}

var (
	bossCrossMap = map[BossType]crosstypes.CrossType{
		BossTypeWorldBoss:   crosstypes.CrossTypeNone,
		BossTypeShareBoss:   crosstypes.CrossTypeWorldboss,
		BossTypeUnrealBoss:  crosstypes.CrossTypeNone,
		BossTypeOutlandBoss: crosstypes.CrossTypeNone,
		BossTypeCangJingGe:  crosstypes.CrossTypeNone,
		BossTypeZhenXi:      crosstypes.CrossTypeZhenXi,
		BossTypeDingShi:     crosstypes.CrossTypeNone,
		BossTypeArena:       crosstypes.CrossTypeShengShou,
	}
)

func (t BossType) CrossType() crosstypes.CrossType {
	return bossCrossMap[t]
}
