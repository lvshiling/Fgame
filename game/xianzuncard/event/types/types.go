package types

import (
	xianzuncardtypes "fgame/fgame/game/xianzuncard/types"
)

type EventTypeXianZunCard string

const (
	EventTypeXianZunCardCrossDay   EventTypeXianZunCard = "玩家天跨天未领取奖励"
	EventTypeXianZunCardExpire                          = "仙尊会员过期"
	EventTypeXianZunCardDataChange                      = "玩家仙尊会员信息变化"
)

type PlayerXianZunCardCrossDayEventData struct {
	typ     xianzuncardtypes.XianZunCardType
	diffDay int32
}

func CreatePlayerXianZunCardCrossDayEventData(typ xianzuncardtypes.XianZunCardType, diffDay int32) *PlayerXianZunCardCrossDayEventData {
	data := &PlayerXianZunCardCrossDayEventData{
		typ:     typ,
		diffDay: diffDay,
	}
	return data
}

func (d *PlayerXianZunCardCrossDayEventData) GetType() xianzuncardtypes.XianZunCardType {
	return d.typ
}

func (d *PlayerXianZunCardCrossDayEventData) GetDiffDay() int32 {
	return d.diffDay
}
