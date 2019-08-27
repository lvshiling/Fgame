package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shenyueventtypes "fgame/fgame/game/shenyu/event/types"
	shenyuscene "fgame/fgame/game/shenyu/scene"
	"fgame/fgame/pkg/mathutils"
)

//神域之战幸运奖励
func shenYuLuckRew(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(shenyuscene.ShenYuSceneData)
	if !ok {
		return
	}

	allPlayerMap := sd.GetScene().GetAllPlayers()
	curPlNum := int32(len(allPlayerMap))
	if curPlNum == 0 {
		return
	}

	var plIdList []int64
	var weights []int64
	shenYuTemp := sd.GetShenYuTemplate()
	for plId, _ := range allPlayerMap {
		plIdList = append(plIdList, plId)
		weights = append(weights, 1)
	}

	drewNum := shenYuTemp.LuckyPalyerCount
	if curPlNum < drewNum {
		drewNum = curPlNum
	}

	randomIndexList := mathutils.RandomListFromWeights(weights, drewNum)
	for _, randomIndex := range randomIndexList {
		randomPlId := plIdList[randomIndex]
		randomPl := allPlayerMap[randomPlId]
		pl, ok := randomPl.(player.Player)
		if !ok {
			continue
		}
		title := lang.GetLangService().ReadLang(lang.ShenYuLuckyRewMailTitle)
		content := lang.GetLangService().ReadLang(lang.ShenYuLuckyRewMailContent)
		attachmentInfo := shenYuTemp.GetLuckItemMap()
		emaillogic.AddEmail(pl, title, content, attachmentInfo)
	}
	return
}

func init() {
	gameevent.AddEventListener(shenyueventtypes.EventTypeShenYuLuckRew, event.EventListenerFunc(shenYuLuckRew))
}
