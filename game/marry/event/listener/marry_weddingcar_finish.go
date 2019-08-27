package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
	marrynpc "fgame/fgame/game/marry/npc/hunche"
	pbuitl "fgame/fgame/game/marry/pbutil"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//婚车巡游结束
func marryWeddingCarEnd(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := target.(*marrynpc.HunCheNPC)
	if !ok {
		return
	}
	marryScene := marry.GetMarryService().GetScene()

	marry.GetMarryService().WeddingCarEnd()

	playerId := eventData.GetHunCheObject().GetPlayerId()
	spouseId := eventData.GetHunCheObject().GetSpouseId()
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)

	playerAllianceId := int64(0)
	playerInfo, _ := player.GetPlayerService().GetPlayerInfo(playerId)
	if playerInfo != nil {
		playerAllianceId = playerInfo.AllianceId
	}
	spouseAllianceId := int64(0)
	spousePlayerInfo, _ := player.GetPlayerService().GetPlayerInfo(spouseId)
	if spousePlayerInfo != nil {
		spouseAllianceId = spousePlayerInfo.AllianceId
	}
	//婚礼酒席开始
	sd := marry.GetMarryService().GetMarrySceneData()
	scMarryWedPushStatus := pbuitl.BuildSCMarryWedPushStatus(sd, true)
	marryPairMsgRelated := marrytypes.CreateMarryPushWedRelated(sd.Id, sd.PlayerId, sd.SpouseId, scMarryWedPushStatus, playerAllianceId, spouseAllianceId)
	player.GetOnlinePlayerManager().BroadcastMsgRelated(marrylogic.OnPlayerMarryPushWedBanquet, marryPairMsgRelated)

	if pl == nil && spl == nil {
		return
	}
	//进入结婚场景
	if pl != nil {
		PlayerEnterMarryScene(pl, marryScene)
	}
	if spl != nil {
		PlayerEnterMarryScene(spl, marryScene)
	}
	return
}

//进入结婚场景
func PlayerEnterMarryScene(pl player.Player, marryScene scene.Scene) {
	scMarryWorish := pbuitl.BuildSCMarryWorish()
	pl.SendMsg(scMarryWorish)
	if pl.GetScene() != marryScene {
		plCtx := scene.WithPlayer(context.Background(), pl)
		playerEnterSceneMsg := message.NewScheduleMessage(onPlayerEnterMarryScene, plCtx, marryScene, nil)
		pl.Post(playerEnterSceneMsg)
	}
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryHunCheEnd, event.EventListenerFunc(marryWeddingCarEnd))
}
