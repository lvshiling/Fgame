package listener

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	shenyueventtypes "fgame/fgame/game/shenyu/event/types"
	"fgame/fgame/game/shenyu/pbutil"
	shenyuscene "fgame/fgame/game/shenyu/scene"
	"fgame/fgame/game/shenyu/shenyu"
	shenyutemplate "fgame/fgame/game/shenyu/template"
	shenyutypes "fgame/fgame/game/shenyu/types"
	"fmt"
)

//神域之战结束（一轮）
func shenYuFinish(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(shenyuscene.ShenYuSceneData)
	if !ok {
		return
	}
	//场景结束
	defer shenyu.GetShenYuService().ShenYuSceneFinish()
	s := sd.GetScene()
	shenYuTemp := sd.GetShenYuTemplate()
	s.Sort()
	rankList := s.GetAllRankList(shenyutypes.ShenYuSceneRankTypeKey)
	for ranking, rankInfo := range rankList {
		spl := s.GetPlayer(rankInfo.GetPlayerId())
		pl, ok := spl.(player.Player)
		if !ok {
			continue
		}
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		rankTemp := shenyutemplate.GetShenYuTemplateService().GetShenYuRankTemplate(shenYuTemp.RoundType, int32(ranking+1))
		if rankTemp == nil {
			continue
		}
		addItemMap := rankTemp.GetRewItemMap()
		totalItemMap := make(map[int32]int32)
		rewEmailItemMap := rankTemp.GetRewEmailItemMap()

		totalItemMap = coreutils.MergeMap(totalItemMap, rewEmailItemMap)
		pointExp := propertyManager.GetExpPoint(int64(rankTemp.RewExpPoint))
		if pointExp > 0 {
			totalItemMap[constanttypes.ExpItem] += int32(pointExp)
		}

		exp := int64(rankTemp.RewExp) + pointExp
		if exp > 0 {
			levelReason := commonlog.LevelLogReasonShenYuRoundRew
			levelReasonText := fmt.Sprintf(levelReason.String(), shenYuTemp.RoundType)
			propertyManager.AddExp(exp, levelReason, levelReasonText)
		}

		goldReason := commonlog.GoldLogReasonShenYuRoundRew
		goldReasonText := fmt.Sprintf(goldReason.String(), shenYuTemp.RoundType)
		silverReason := commonlog.SilverLogReasonShenYuRoundRew
		silverReasonText := fmt.Sprintf(silverReason.String(), shenYuTemp.RoundType)
		propertyManager.AddMoney(int64(rankTemp.RewBindGold), int64(rankTemp.RewGold), goldReason, goldReasonText, int64(rankTemp.RewSilver), silverReason, silverReasonText)
		//if !flag {
		//	panic(fmt.Errorf("shenyu:添加银两元宝应该成功"))
		//}
		if inventoryManager.HasEnoughSlots(addItemMap) {
			itemGetReason := commonlog.InventoryLogReasonShenYuRoundRew
			itemGetReasonText := fmt.Sprintf(itemGetReason.String(), shenYuTemp.RoundType)
			flag := inventoryManager.BatchAdd(addItemMap, itemGetReason, itemGetReasonText)
			if !flag {
				panic(fmt.Errorf("shenyu:添加物品应该成功"))
			}
		} else {
			mailName := lang.GetLangService().ReadLang(lang.ShenYuRankMailTitle)
			mailContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ShenYuRankMailContent), ranking+1)
			emaillogic.AddEmail(pl, mailName, mailContent, addItemMap)
		}

		propertylogic.SnapChangedProperty(pl)
		inventorylogic.SnapInventoryChanged(pl)

		scMsg := pbutil.BuildSCShenYuFinishRew(shenYuTemp.RoundType, int32(ranking), totalItemMap)
		spl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(shenyueventtypes.EventTypeShenYuFinish, event.EventListenerFunc(shenYuFinish))
}
