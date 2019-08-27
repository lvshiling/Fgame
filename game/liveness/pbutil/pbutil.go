package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerliveness "fgame/fgame/game/liveness/player"
)

func BuildSCLivenessGet(livenessObj *playerliveness.PlayerLivenessObject, livenessMap map[int32]*playerliveness.PlayerLivenessQuestObject) *uipb.SCLivenessGet {
	livenessGet := &uipb.SCLivenessGet{}
	liveness := livenessObj.GetLiveness()
	livenessGet.Liveness = &liveness
	for _, boxId := range livenessObj.GetOpenBoxs() {
		livenessGet.OpenBoxList = append(livenessGet.OpenBoxList, boxId)
	}

	for _, livenessQuest := range livenessMap {
		livenessGet.LivenessList = append(livenessGet.LivenessList, buildLivenessQuest(livenessQuest))
	}
	return livenessGet
}

func buildLivenessQuest(livenessQuest *playerliveness.PlayerLivenessQuestObject) *uipb.LivenessInfo {
	livenessInfo := &uipb.LivenessInfo{}
	questId := livenessQuest.GetQuestId()
	num := livenessQuest.GetNum()
	livenessInfo.QuestId = &questId
	livenessInfo.Num = &num
	return livenessInfo
}

func BuildSCLivenessNumChanged(livenessQuest *playerliveness.PlayerLivenessQuestObject, liveness int64) *uipb.SCLivenessNumChanged {
	livenessNumChanged := &uipb.SCLivenessNumChanged{}
	livenessNumChanged.Liveness = &liveness
	livenessNumChanged.LivenessInfo = buildLivenessQuest(livenessQuest)
	return livenessNumChanged
}

func BuildSCLivenessOpen(livenessObj *playerliveness.PlayerLivenessObject, openBoxId int32) *uipb.SCLivenessOpen {
	livenessOpen := &uipb.SCLivenessOpen{}
	liveness := livenessObj.GetLiveness()
	livenessOpen.Liveness = &liveness
	livenessOpen.BoxId = &openBoxId
	for _, boxId := range livenessObj.GetOpenBoxs() {
		livenessOpen.OpenBoxList = append(livenessOpen.OpenBoxList, boxId)
	}
	return livenessOpen
}
