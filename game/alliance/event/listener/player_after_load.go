package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	alliancelogic "fgame/fgame/game/alliance/logic"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAllianceAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)

	//加载玩家仙盟信息
	err = loadPlayerAllianceInfo(p)
	if err != nil {
		return
	}

	//发送霸主信息
	alliancelogic.SendAllianceHegemonInfo(p)

	return
}

func loadPlayerAllianceInfo(p player.Player) (err error) {
	mem := alliance.GetAllianceService().GetAllianceMember(p.GetId())
	allianceManager := p.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	if mem == nil {
		//退出仙盟
		allianceManager.SyncAlliance(0, 0, "", 0, 0)
		return
	}

	//同步成员信息
	alliance.GetAllianceService().SyncMemberInfo(p.GetId(), p.GetName(), p.GetSex(), p.GetLevel(), p.GetForce(), p.GetZhuanSheng(), p.GetLingyuInfo().AdvanceId, p.GetVip())
	//加入仙盟
	allianceManager.SyncAlliance(mem.GetAllianceId(), mem.GetAlliance().GetAllianceMengZhuId(), mem.GetAlliance().GetAllianceName(), mem.GetAlliance().GetAllianceLevel(), mem.GetPosition())
	scAllianceInfo := pbutil.BuildSCAllianceInfo(mem.GetAlliance(), mem)
	p.SendMsg(scAllianceInfo)

	//发送仙盟个人信息
	scAlliancePlayerInfo := pbutil.BuildSCAlliancePlayerInfo(allianceManager.GetPlayerAllianceObject(), allianceManager.GetPlayerAllianceSkillMap())
	p.SendMsg(scAlliancePlayerInfo)

	//加载仙术
	err = alliancelogic.ReloadAllianceSkill(p)

	//加载成员上线时间
	alliance.GetAllianceService().OnlineMember(p.GetId())
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAllianceAfterLoad))
}
