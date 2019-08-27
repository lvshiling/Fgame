package listener

import (
	"fgame/fgame/core/event"
	teamcopyeventtypes "fgame/fgame/cross/teamcopy/event/types"
	teamcopylogic "fgame/fgame/cross/teamcopy/logic"
	"fgame/fgame/cross/teamcopy/pbutil"
	"fgame/fgame/cross/teamcopy/scene"
	"fgame/fgame/cross/teamcopy/teamcopy"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/robot/robot"
)

//组队副本场景结束
func teamCopySceneEnd(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(scene.TeamCopySceneData)
	sucess := data.(bool)
	teamObj := sd.GetTeamObj()
	purpose := int32(teamObj.GetTeamPurpose())

	isTeamCopyBattleResult := pbutil.BuildISTeamCopyBattleResult(purpose, sucess)
	teamcopylogic.BroadcastTeamCopy(sd, isTeamCopyBattleResult)
	teamcopy.GetTeamCopyService().TeamCopyFinish(sd)
	for _, mem := range sd.GetTeamObj().GetMemberList() {
		robotPl := robot.GetRobotService().GetRobot(mem.GetPlayerId())
		if robotPl == nil {
			continue
		}
		teamcopylogic.RobotExit(robotPl)
	}
	return
}

func init() {
	gameevent.AddEventListener(teamcopyeventtypes.EventTypeTeamCopySceneFinish, event.EventListenerFunc(teamCopySceneEnd))
}
