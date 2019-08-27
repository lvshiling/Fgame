package types

import (
	godsiegetypes "fgame/fgame/game/godsiege/types"
)

type GodSiegeEventType string

const (
	//Boss状态刷新
	EventTypeGodSiegeBossStatusRefresh GodSiegeEventType = "GodSiegeBossStatusRefresh"
	//玩家进入神兽攻城场景
	EventTypeGodSiegePlayerEnter GodSiegeEventType = "GodSiegePlayerEnter"
	//神兽攻城场景完成
	EventTypeGodSiegeSceneFinish GodSiegeEventType = "GodSiegeSceneFinish"
	//玩家取消排队
	EventTypeGodSiegeCancleLineUp GodSiegeEventType = "GodSiegeCancleLineUp"
	//玩家退出神兽攻城场景
	EventTypeGodSiegePlayerExit GodSiegeEventType = "GodSiegePlayerExit"
	//玩家排队完成
	EventTypeGodSiegePlayerLineUpFinish GodSiegeEventType = "GodSiegePlayerLineUpFinish"
	//采集物变化
	EventTypeGodSiegeCollectNpcChanged GodSiegeEventType = "GodSiegeCollectNpcChanged"
)

type DenseWatEventType string

const (
	//金银密窟物品变化
	EventTypeDenseWatItemChanged DenseWatEventType = "DenseWatItemChanged"
)

type GodSiegeCancleLineUpEventData struct {
	pos     int32
	godType godsiegetypes.GodSiegeType
}

func (d *GodSiegeCancleLineUpEventData) GetPos() int32 {
	return d.pos
}

func (d *GodSiegeCancleLineUpEventData) GetGodType() godsiegetypes.GodSiegeType {
	return d.godType
}

func CreateGodSiegeCancleLineUpEventData(pos int32, godType godsiegetypes.GodSiegeType) *GodSiegeCancleLineUpEventData {
	d := &GodSiegeCancleLineUpEventData{
		pos:     pos,
		godType: godType,
	}
	return d
}

type GodSiegeLineUpFinishEventData struct {
	playerId int64
	godType  godsiegetypes.GodSiegeType
}

func (d *GodSiegeLineUpFinishEventData) GetPlayerId() int64 {
	return d.playerId
}

func (d *GodSiegeLineUpFinishEventData) GetGodType() godsiegetypes.GodSiegeType {
	return d.godType
}

func CreateGodSiegeFinishLineUpEventData(playerId int64, godType godsiegetypes.GodSiegeType) *GodSiegeLineUpFinishEventData {
	d := &GodSiegeLineUpFinishEventData{
		playerId: playerId,
		godType:  godType,
	}
	return d
}
