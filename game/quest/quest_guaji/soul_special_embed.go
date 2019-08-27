package quest_guaji

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	soullogic "fgame/fgame/game/soul/logic"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

//佩戴帝魂
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeSoulSpecialEmbed, guaji.QuestGuaJiFunc(soulSpecialEmbed))
}

//佩戴帝魂
func soulSpecialEmbed(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

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
				"playerId":  p.GetId(),
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务不存在")
		return false
	}
	if q.QuestState != questtypes.QuestStateAccept {
		log.WithFields(
			log.Fields{
				"playerId":  p.GetId(),
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务没在进行中")
		return false
	}

	for k, v := range demandMap {
		if v <= q.QuestDataMap[k] {
			continue
		}
		soulTag := soultypes.SoulType(k)
		if !soulTag.Valid() {
			log.WithFields(
				log.Fields{
					"playerId":  p.GetId(),
					"questType": questTemplate.GetQuestSubType().String(),
					"soulTag":   soulTag,
				}).Warn("quest_guaji:帝魂类型无效")
			return false
		}

		if !soulEmbed(p, soulTag) {
			log.WithFields(
				log.Fields{
					"playerId":  p.GetId(),
					"questType": questTemplate.GetQuestSubType().String(),
					"soulTag":   soulTag,
				}).Warn("quest_guaji:帝魂镶嵌失败")
			return false
		}
	}
	return true
}

func soulActive(p player.Player, soulTag soultypes.SoulType) bool {

	soulManager := p.GetPlayerDataManager(playertypes.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)

	flag := soulManager.IfSoulTagExist(soulTag)
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"soulTag":  soulTag,
			}).Warn("quest_guaji:帝魂,已经激活")
		return false
	}

	soulActiveTemplate := soul.GetSoulService().GetSoulActiveTemplate(soulTag)
	//激活的前置帝魂条件
	preSoulCond := soulActiveTemplate.GetPreSoulCond()
	if preSoulCond != nil {
		preSoulTag := preSoulCond.GetSoulType()
		flag := soulManager.IfPreSoul(preSoulTag, preSoulCond.Level)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": p.GetId(),
				"soulTag":  soulTag,
			}).Warn("quest_guaji:帝魂,激活的前置条件不足")
			return false
		}
	}

	//激活需要物品
	items := soulActiveTemplate.GetNeedItemMap()
	if len(items) != 0 {
		inventoryManager := p.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItems(items)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": p.GetId(),
				"soulTag":  soulTag,
			}).Warn("quest_guaji:帝魂,无法激活")
			return false
		}
	}

	soullogic.HandleSoulActive(p, soulTag)
	return true
}

func soulEmbed(p player.Player, soulTag soultypes.SoulType) bool {

	soulManager := p.GetPlayerDataManager(playertypes.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)

	flag := soulManager.IfSoulTagExist(soulTag)
	if !flag {
		//激活
		flag = soulActive(p, soulTag)
		if !flag {
			return false
		}
	}

	flag = soulManager.IfSoulTagEmemded(soulTag)
	if flag {
		log.WithFields(log.Fields{
			"playerId": p.GetId(),
			"soulTag":  soulTag,
		}).Warn("quest_guaji:帝魂,已经镶嵌")
		return false
	}

	soullogic.HandleSoulEmbed(p, soulTag)
	return true
}
