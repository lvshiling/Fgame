package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shenmo/pbutil"
	playershenmo "fgame/fgame/game/shenmo/player"
	"fgame/fgame/game/shenmo/shenmo"
	"fgame/fgame/game/shenmo/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	manager := p.GetPlayerDataManager(playertypes.PlayerShenMoWarDataManagerType).(*playershenmo.PlayerShenMoDataManager)
	shenMoObj := manager.GetShenMoInfo()
	rankTime := shenMoObj.GetRewTime()
	scShenMoGetReward := pbutil.BuildSCShenMoGetReward(rankTime)
	p.SendMsg(scShenMoGetReward)

	gongXunNum := shenMoObj.GetGongXunNum()
	scPlayerGongXunChanged := pbutil.BuildSCPlayerGongXunChanged(gongXunNum)
	p.SendMsg(scPlayerGongXunChanged)

	//我的上周排名
	rankType := types.RankTimeTypeLast
	lastRankTime := shenmo.GetShenMoService().GetRankTime(rankType)
	allianceId := p.GetAllianceId()
	if allianceId == 0 {
		scShenMoMyRank := pbutil.BuildSCShenMoMyRank(false, 0, lastRankTime)
		p.SendMsg(scShenMoMyRank)
		return
	}
	serverId := global.GetGame().GetServerIndex()
	pos, _ := shenmo.GetShenMoService().GetMyRank(rankType, serverId, allianceId)
	scShenMoMyRank := pbutil.BuildSCShenMoMyRank(false, pos, lastRankTime)
	p.SendMsg(scShenMoMyRank)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
