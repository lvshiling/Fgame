package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arenapvp/arenapvp"
	arenapvpeventtypes "fgame/fgame/cross/arenapvp/event/types"
	arenapvplogic "fgame/fgame/cross/arenapvp/logic"
	arenapvpscene "fgame/fgame/cross/arenapvp/scene"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"

	log "github.com/Sirupsen/logrus"
)

//竞技场场景结束
func arenapvpElectionSceneStop(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(arenapvpscene.ArenapvpSceneData)
	if !ok {
		return
	}

	s := sd.GetScene()
	activityEndTime := sd.GetActivityEndTime()
	nextPvpTemp := sd.GetPvpTemp().GetNextTemp()
	if nextPvpTemp == nil {
		return
	}

	for _, spl := range s.GetAllPlayers() {
		//重置复活次数/血量
		arenapvplogic.ResetBattleInfo(spl)

		//进入的位置
		nextS := arenapvp.GetArenapvpService().CreateArenapvpSceneBattle(nextPvpTemp, spl.GetId(), activityEndTime)
		if nextS == nil {
			continue
		}
		sd := nextS.SceneDelegate().(arenapvpscene.ArenapvpBattleSceneData)
		bornPos := sd.GetEnterPos(spl.GetId())
		if !scenelogic.PlayerEnterScene(spl, nextS, bornPos) {
			log.WithFields(
				log.Fields{
					"playerId": spl.GetId(),
				}).Warn("arenapvp:处理进入下一场对战失败，进入场景失败")
			return
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpElectionSceneStop, event.EventListenerFunc(arenapvpElectionSceneStop))
}
