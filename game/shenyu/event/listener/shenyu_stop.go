package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	shenyueventtypes "fgame/fgame/game/shenyu/event/types"
	shenyuscene "fgame/fgame/game/shenyu/scene"
	shenyutypes "fgame/fgame/game/shenyu/types"

	log "github.com/Sirupsen/logrus"
)

//神域之战关闭（一轮）
func shenYuStop(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(shenyuscene.ShenYuSceneData)
	if !ok {
		return
	}

	s := sd.GetScene()
	shenYuTemp := sd.GetShenYuTemplate()

	// 进入下一轮
	if shenYuTemp.GetNextTemp() == nil {
		return
	}

	rankList := s.GetAllRankList(shenyutypes.ShenYuSceneRankTypeKey)
	winIdList := make([]int64, 0, shenYuTemp.WinRank)
	for _, rankInfo := range rankList {
		winIdList = append(winIdList, rankInfo.GetPlayerId())
	}
	if len(winIdList) > int(shenYuTemp.WinRank) {
		winIdList = winIdList[:shenYuTemp.WinRank]
	}

	nextShenYuTemp := shenYuTemp.GetNextTemp()
	sceneEndTime := s.GetEndTime() + int64(shenYuTemp.RoundTime)
	nextSd := shenyuscene.CreateShenYuSceneData(nextShenYuTemp, sd.GetActivityEndTime())
	nextS := scene.CreateActivityScene(nextShenYuTemp.MapId, sceneEndTime, nextSd)
	if nextS == nil {
		log.WithFields(
			log.Fields{
				"Mapid":        nextShenYuTemp.MapId,
				"sceneEndTime": sceneEndTime,
			}).Warn("shenyu:下一轮神域之战场景不存在")
		return
	}

	bornPos := s.MapTemplate().GetBornPos()
	for _, plId := range winIdList {
		pl := s.GetPlayer(plId)
		if pl == nil {
			continue
		}
		if !scenelogic.PlayerEnterScene(pl, nextS, bornPos) {
			return
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(shenyueventtypes.EventTypeShenYuStop, event.EventListenerFunc(shenYuStop))
}
