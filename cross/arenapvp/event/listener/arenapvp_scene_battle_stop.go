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
func arenapvpBattleSceneStop(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(arenapvpscene.ArenapvpBattleSceneData)
	if !ok {
		return
	}
	s := sd.GetScene()
	activityEndTime := sd.GetActivityEndTime()
	nextPvpTemp := sd.GetPvpTemp().GetNextTemp()
	if nextPvpTemp == nil {
		return
	}
	winnerId := sd.GetWinnerId()
	if winnerId == 0 {
		return
	}
	//进入的位置
	nextS := arenapvp.GetArenapvpService().CreateArenapvpSceneBattle(nextPvpTemp, winnerId, activityEndTime)
	if nextS == nil {
		return
	}
	p := s.GetPlayer(winnerId)
	if p != nil {
		//重置复活次数/血量
		arenapvplogic.ResetBattleInfo(p)

		sd := nextS.SceneDelegate().(arenapvpscene.ArenapvpBattleSceneData)
		bornPos := sd.GetEnterPos(winnerId)
		if !scenelogic.PlayerEnterScene(p, nextS, bornPos) {
			log.WithFields(
				log.Fields{
					"playerId": winnerId,
				}).Warn("arenapvp:处理进入下一场对战失败，进入场景失败")
			return
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpBattleSceneStop, event.EventListenerFunc(arenapvpBattleSceneStop))
}
