package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marryscene "fgame/fgame/game/marry/scene"
	marrytemplate "fgame/fgame/game/marry/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

//婚礼场景结束
func marrySceneEnd(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(marryscene.MarrySceneData)
	if !ok {
		return
	}

	marryData := sd.GetMarryData()
	playerId := marryData.PlayerId
	spouseId := marryData.SpouseId
	//场景内广播婚礼结束
	scMarryWedEnd := pbuitl.BuildSCMarryWedEnd(marryData)
	sd.GetScene().BroadcastMsg(scMarryWedEnd)

	sceneAllPlayer := sd.GetScene().GetAllPlayers()
	//场景内清空豪气值
	for _, spl := range sceneAllPlayer {
		pl := spl.(player.Player)
		if pl == nil {
			continue
		}
		manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
		manager.RefreshHeriosm()
	}

	itemMap := marrytemplate.GetMarryTemplateService().GetMarryEndRewMap(marryData.Grade, marryData.HunCheGrade, marryData.SugarGrade)
	sd.FinishClearData()
	pairIdList := make([]int64, 0, 2)
	pairIdList = append(pairIdList, playerId, spouseId)
	distributionGift(pairIdList, itemMap)
	return
}

func distributionGift(pairIdList []int64, itemMap map[int32]int32) {
	for _, playerId := range pairIdList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl != nil {
			ctx := scene.WithPlayer(context.Background(), pl)
			playerReceiveGiftMsg := message.NewScheduleMessage(onWedGiftEmail, ctx, itemMap, nil)
			pl.Post(playerReceiveGiftMsg)
			continue
		}
		if len(itemMap) == 0 {
			continue
		}
		//离线邮件
		emailTitle := lang.GetLangService().ReadLang(lang.MarryWedGiftTitle)
		emailContent := lang.GetLangService().ReadLang(lang.MarryWedGiftContent)
		emaillogic.AddOfflineEmail(playerId, emailTitle, emailContent, itemMap)
	}
	return
}

func onWedGiftEmail(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	itemMap := result.(map[int32]int32)

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.EndWedding()

	if len(itemMap) == 0 {
		return nil
	}
	//写邮件
	emailTitle := lang.GetLangService().ReadLang(lang.MarryWedGiftTitle)
	emailContent := lang.GetLangService().ReadLang(lang.MarryWedGiftContent)
	emaillogic.AddEmail(pl, emailTitle, emailContent, itemMap)
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarrySceneWedEnd, event.EventListenerFunc(marrySceneEnd))
}
