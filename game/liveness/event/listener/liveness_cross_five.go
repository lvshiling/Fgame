package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	"fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	livenesseventtypes "fgame/fgame/game/liveness/event/types"
	playerliveness "fgame/fgame/game/liveness/player"
	livenesstemplate "fgame/fgame/game/liveness/template"
	"fgame/fgame/game/player"
)

//活跃值改变
func playerLivenessCrossFive(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	livenessObj := data.(*playerliveness.PlayerLivenessObject)
	liveness := livenessObj.GetLiveness()
	if liveness <= 0 {
		return
	}
	openBoxs := livenessObj.GetOpenBoxs()
	rewardBoxs := make([]int32, 0, 3)
	huoYueBoxList := livenesstemplate.GetHuoYueTempalteService().GetHuoYueBoxList()
	for _, huoYueBox := range huoYueBoxList {
		if liveness <= 0 {
			break
		}
		flag := utils.ContainInt32(openBoxs, int32(huoYueBox.Id))
		if flag {
			continue
		}

		if liveness >= int64(huoYueBox.NeedStar) {
			rewardBoxs = append(rewardBoxs, int32(huoYueBox.Id))
		}
	}

	if len(rewardBoxs) == 0 {
		return
	}
	rewardItemMap := make(map[int32]int32)
	for _, boxId := range rewardBoxs {
		huoYueBoxTemplate := livenesstemplate.GetHuoYueTempalteService().GetHuoYueBoxTemplate(int32(boxId))
		if huoYueBoxTemplate == nil {
			continue
		}
		rewData := huoYueBoxTemplate.GetRewData()
		if rewData != nil {
			if rewData.RewSilver != 0 {
				rewardItemMap[constanttypes.SilverItem] += rewData.RewSilver
			}
			if rewData.RewGold != 0 {
				rewardItemMap[constanttypes.GoldItem] += rewData.RewGold
			}
			if rewData.RewBindGold != 0 {
				rewardItemMap[constanttypes.BindGoldItem] += rewData.RewBindGold
			}

		}
		rewItemMap := huoYueBoxTemplate.GetRewItemMap()
		for itemId, num := range rewItemMap {
			rewardItemMap[itemId] += num
		}
	}

	if len(rewardItemMap) == 0 {
		return
	}

	emailTitle := lang.GetLangService().ReadLang(lang.LivenessRewardTitle)
	emailContent := lang.GetLangService().ReadLang(lang.LivenessRewardContent)
	emaillogic.AddEmail(pl, emailTitle, emailContent, rewardItemMap)
	return
}

func init() {
	gameevent.AddEventListener(livenesseventtypes.EventTypeLivenessCrossFive, event.EventListenerFunc(playerLivenessCrossFive))
}
