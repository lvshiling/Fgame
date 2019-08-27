package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	propertytypes "fgame/fgame/game/property/types"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
)

//玩家最大血量变化
func playerMaxHpChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	teamId := pl.GetTeamId()
	if teamId == 0 {
		return
	}

	playerId := pl.GetId()
	maxHp := pl.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	teamData := team.GetTeamService().GetTeam(teamId)
	if teamData == nil {
		return
	}
	scTeamBroadcast := pbutil.BuildSCTeamMaxHpChange(playerId, maxHp)
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerMaxHPChanged, event.EventListenerFunc(playerMaxHpChanged))
}
