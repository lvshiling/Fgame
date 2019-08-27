package quest_guaji

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	_ "fgame/fgame/game/quest/quest_guaji/system"
	xianfulogic "fgame/fgame/game/xianfu/logic"
	xianfuplayer "fgame/fgame/game/xianfu/player"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeUpgradeSpecialXianFu, guaji.QuestGuaJiFunc(upgradeSpecialXianFu))
}

//升级仙府
func upgradeSpecialXianFu(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {
	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}
	//获取当前任务数据
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))

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
		flag := doUpgradeSpecialXianFu(p, xianfuType)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId":   p.GetId(),
					"xianfuType": xianfuType,
				}).Warn("quest_guaji:升级仙府失败")
			return false
		}
	}
	return true
}

func doUpgradeSpecialXianFu(pl player.Player, xianfuType xianfutypes.XianfuType) bool {

	xianfuManager := pl.GetPlayerDataManager(playertypes.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	xianfuId := xianfuManager.GetXianfuId(xianfuType)
	xfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xianfuId, xianfuType)
	if xfTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("quest_guaji:秘境仙府升级，模板不存在")

		return false
	}

	if xianfuManager.IsUpgrading(xianfuType) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("quest_guaji:秘境仙府升级请求，仙府升级中")
		return false
	}

	//是否达到等级上限
	if xfTemplate.GetNextId() == 0 {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("quest_guaji:秘境仙府升级请求，当前建筑已达最高级，无法升级")
		return false
	}

	nextXfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xfTemplate.GetNextId(), xianfuType)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	upgradeNeedGold := nextXfTemplate.GetUpgradeGold()
	upgradeNeedBindGold := nextXfTemplate.GetUpgradeBindGold()
	totalUpgradeNeedGold := upgradeNeedGold + upgradeNeedBindGold
	upgradeNeedSilver := nextXfTemplate.GetUpgradeYinliang()
	upgradeNeedItemId := nextXfTemplate.GetUpgradeItemId()
	upgradeNeedItemNum := nextXfTemplate.GetUpgradeItemNum()

	//元宝是否足够
	if upgradeNeedGold > 0 {
		if !propertyManager.HasEnoughGold(int64(upgradeNeedGold), false) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("quest_guaji:秘境仙府升级请求，当前元宝不足，无法升级")

			return false
		}
	}

	if totalUpgradeNeedGold > 0 {
		if !propertyManager.HasEnoughGold(int64(totalUpgradeNeedGold), true) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("quest_guaji:秘境仙府升级请求，当前绑定元宝不足，无法升级")

			return false
		}
	}

	//银两是否足够
	if upgradeNeedSilver > 0 {
		if !propertyManager.HasEnoughSilver(upgradeNeedSilver) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
					"needSilver": upgradeNeedSilver,
				}).Warn("quest_guaji:秘境仙府升级请求，银两不足，无法升级")

			return false
		}
	}

	//物品是否足够
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if upgradeNeedItemId > 0 && upgradeNeedItemNum > 0 {
		if !inventoryManager.HasEnoughItem(upgradeNeedItemId, upgradeNeedItemNum) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("quest_guaji:秘境仙府升级请求，物品不足，无法升级")

			return false
		}
	}

	xianfulogic.HandleXianfuUpgrade(pl, xianfuType)
	return true

}
