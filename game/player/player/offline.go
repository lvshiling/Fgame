package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player/dao"
	playerentity "fgame/fgame/game/player/entity"
	"fgame/fgame/game/player/types"
)

type OfflinePlayer struct {
	pe *playerentity.PlayerEntity
}

func (p *OfflinePlayer) Forbid(forbidReason string, forbidName string, forbidTime int64) {
	now := global.GetGame().GetTimeService().Now()
	p.pe.Forbid = 1
	p.pe.ForbidText = forbidReason
	p.pe.ForbidName = forbidName
	if forbidTime == 0 {
		p.pe.ForbidEndTime = 0
	} else {
		p.pe.ForbidEndTime = forbidTime + now
	}
	p.pe.ForbidTime = now
	p.pe.UpdateTime = now
	global.GetGame().GetGlobalUpdater().AddChangedObject(p.pe)
}

func (p *OfflinePlayer) Unforbid() {
	now := global.GetGame().GetTimeService().Now()
	p.pe.Forbid = 0
	p.pe.UpdateTime = now
	global.GetGame().GetGlobalUpdater().AddChangedObject(p.pe)
}

func (p *OfflinePlayer) ForbidChat(forbidChatReason string, forbidChatName string, forbidTime int64) {
	now := global.GetGame().GetTimeService().Now()
	p.pe.ForbidChat = 1
	p.pe.ForbidChatText = forbidChatReason
	p.pe.ForbidChatName = forbidChatName
	p.pe.ForbidChatTime = now
	if forbidTime == 0 {
		p.pe.ForbidChatEndTime = 0
	} else {
		p.pe.ForbidChatEndTime = forbidTime + now
	}
	p.pe.UpdateTime = now
	global.GetGame().GetGlobalUpdater().AddChangedObject(p.pe)
}

func (p *OfflinePlayer) UnforbidChat() {
	now := global.GetGame().GetTimeService().Now()
	p.pe.ForbidChat = 0
	p.pe.UpdateTime = now
	global.GetGame().GetGlobalUpdater().AddChangedObject(p.pe)
}

func (p *OfflinePlayer) IgnoreChat(ignoreChatReason string, ignoreChatName string, ignoreTime int64) {
	now := global.GetGame().GetTimeService().Now()
	p.pe.IgnoreChat = 1
	p.pe.IgnoreChatText = ignoreChatReason
	p.pe.IgnoreChatName = ignoreChatName
	p.pe.IgnoreChatTime = now
	if ignoreTime == 0 {
		p.pe.IgnoreChatEndTime = 0
	} else {
		p.pe.IgnoreChatEndTime = ignoreTime + now
	}

	p.pe.UpdateTime = now
	global.GetGame().GetGlobalUpdater().AddChangedObject(p.pe)
}

func (p *OfflinePlayer) UnignoreChat() {
	now := global.GetGame().GetTimeService().Now()
	p.pe.IgnoreChat = 0
	p.pe.UpdateTime = now
	global.GetGame().GetGlobalUpdater().AddChangedObject(p.pe)
}

func (p *OfflinePlayer) SetPrivilege(privilegeType types.PrivilegeType) {
	now := global.GetGame().GetTimeService().Now()
	p.pe.PrivilegeType = int32(privilegeType)
	p.pe.UpdateTime = now
	global.GetGame().GetGlobalUpdater().AddChangedObject(p.pe)
}

func CreateOfflinePlayer(playerId int64) (pl *OfflinePlayer, err error) {
	pl = &OfflinePlayer{}
	pe, err := dao.GetPlayerDao().QueryById(playerId)
	if err != nil {
		return
	}
	if pe == nil {
		return nil, nil
	}
	pl.pe = pe
	return pl, nil
}
