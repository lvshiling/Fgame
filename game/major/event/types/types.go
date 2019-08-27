package types

import (
	majortemplate "fgame/fgame/game/major/template"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
)

type MajorEventType string

const (

	//玩家双修决策
	EventTypeMajorInviteDeal MajorEventType = "MajorInviteDeal"
	//玩家取消双修邀请
	EventTypeMajorInviteCancle MajorEventType = "MajorInviteCancle"
	//邀请过期(邀请发出 配偶下线未回复)
	EventTypeMajorInviteNoAnswer MajorEventType = "MajorInviteNoAnswer"
	//玩家进入双修场景
	EventTypePlayerEnterMajorScene MajorEventType = "PlayerEnterMajorScene"
	//玩家双修副本成功
	EventTypePlayerMajorSuccess MajorEventType = "PlayerMajorSuccess"
	//扫荡
	EventTypePlayerMajorSweep MajorEventType = "PlayerMajorSweep"
)

type MajorInviteDealEventData struct {
	playerId  int64
	spl       player.Player
	agree     bool
	majorType majortypes.MajorType
	fubenId   int32
}

func (m *MajorInviteDealEventData) GetPlayerId() int64 {
	return m.playerId
}

func (m *MajorInviteDealEventData) GetSpousePlayer() player.Player {
	return m.spl
}

func (r *MajorInviteDealEventData) GetAgree() bool {
	return r.agree
}

func (r *MajorInviteDealEventData) GetMajorType() majortypes.MajorType {
	return r.majorType
}

func (r *MajorInviteDealEventData) GetFubenId() int32 {
	return r.fubenId
}

func CreateMajorInviteDealEventData(playerId int64, spl player.Player, agree bool, majorType majortypes.MajorType, fubenId int32) *MajorInviteDealEventData {
	d := &MajorInviteDealEventData{
		playerId:  playerId,
		spl:       spl,
		agree:     agree,
		majorType: majorType,
		fubenId:   fubenId,
	}
	return d
}

type MajorSuccessEventData struct {
	temp majortemplate.MajorTemplate
	num  int32
}

func (m *MajorSuccessEventData) GetTemp() majortemplate.MajorTemplate {
	return m.temp
}

func (m *MajorSuccessEventData) GetNum() int32 {
	return m.num
}

func CreateMajorSuccessEventData(temp majortemplate.MajorTemplate, num int32) *MajorSuccessEventData {
	d := &MajorSuccessEventData{
		temp: temp,
		num:  num,
	}
	return d
}

type MajorSweepEventData struct {
	temp majortemplate.MajorTemplate
	num  int32
}

func (m *MajorSweepEventData) GetTemp() majortemplate.MajorTemplate {
	return m.temp
}

func (m *MajorSweepEventData) GetNum() int32 {
	return m.num
}

func CreateMajorSweepEventData(temp majortemplate.MajorTemplate, num int32) *MajorSweepEventData {
	d := &MajorSweepEventData{
		temp: temp,
		num:  num,
	}
	return d
}
