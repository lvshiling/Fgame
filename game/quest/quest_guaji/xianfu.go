package quest_guaji

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	xianfulogic "fgame/fgame/game/xianfu/logic"
	xianfuplayer "fgame/fgame/game/xianfu/player"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"

	log "github.com/Sirupsen/logrus"
)

//进入指定X次秘境仙府
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeEnterSpecialXianFu, guaji.QuestGuaJiFunc(specifiedXianFu))
}

//进入指定X次秘境仙府
func specifiedXianFu(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}
	//获取当前任务数据
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))
	if q == nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("quest_guaji:任务不存在")
		return false
	}
	if q.QuestState != questtypes.QuestStateAccept {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("quest_guaji:任务没在进行中")
		return false
	}
	for k, v := range demandMap {
		xianfuType := xianfutypes.XianfuType(k)
		if !xianfuType.Valid() {
			log.WithFields(
				log.Fields{
					"playerId":   p.GetId(),
					"xianfuType": xianfuType,
				}).Warn("quest_guaji:仙府类型无效")
			return false
		}
		num := q.QuestDataMap[k]
		if num >= v {
			continue
		}
		flag := doXianFu(p, xianfuType)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId":   p.GetId(),
					"xianfuType": xianfuType,
				}).Warn("quest_guaji:进入仙府失败")
			return false
		}
		break
	}
	return true
}

func doXianFu(pl player.Player, xianfuType xianfutypes.XianfuType) bool {

	xianfuManager := pl.GetPlayerDataManager(playertypes.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	xianfuId := xianfuManager.GetXianfuId(xianfuType)
	xfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xianfuId, xianfuType)
	if xfTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("quest_guaji:秘境仙府挑战请求，模板不存在")
		return false
	}

	//是否免费次数
	freeTimes := xianfulogic.FreeTimesCount(pl, xianfuType)
	if freeTimes < 1 {
		//挑战次数是否足够
		leftTimes := xianfuManager.GetChallengeTimes(xianfuType)
		if leftTimes < 1 {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("quest_guaji:秘境仙府挑战请求，副本次数不足，无法挑战")
			return false
		}

		attendNeedItemId := xfTemplate.GetNeedItemId()
		attendNeedItemNum := xfTemplate.GetNeedItemCount()

		//挑战所需物品是否足够
		if !inventoryManager.HasEnoughItem(attendNeedItemId, attendNeedItemNum) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("quest_guaji:秘境仙府挑战请求，道具不足，无法挑战")
			return false
		}
	}

	xianfulogic.HandleXianfuChallenge(pl, xianfuType)

	return true
}
