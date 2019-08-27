package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arena/arena"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	"fgame/fgame/cross/arena/pbutil"
	arenascene "fgame/fgame/cross/arena/scene"
	"fgame/fgame/cross/player/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/robot/robot"
	teamtypes "fgame/fgame/game/team/types"
)

//竞技场队伍创建
func arenaTeamCreate(target event.EventTarget, data event.EventData) (err error) {
	t := target.(*arenascene.ArenaTeam)
	eventData := data.(*arena.TeamCreateEventData)

	//更新机器人队伍
	for _, mem := range t.GetTeam().GetMemberList() {
		robotPl := robot.GetRobotService().GetRobot(mem.GetPlayerId())
		if robotPl == nil {
			continue
		}
		robotPl.SetArenaTeam(t.GetTeam().GetTeamId(), t.GetTeam().GetTeamName(), teamtypes.TeamPurposeTypeNormal)
	}

	//推送匹配到了
	isArenaMatchResult := pbutil.BuildISArenaMatchResult(true)
	for _, memList := range eventData.GetTeamList() {
		captainId := memList[0].GetPlayerId()
		captainPlayer := player.GetOnlinePlayerManager().GetPlayerById(captainId)
		if captainPlayer == nil {
			continue
		}
		captainPlayer.SendMsg(isArenaMatchResult)
	}
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaTeamCreate, event.EventListenerFunc(arenaTeamCreate))
}
