package listener

import (
	"fgame/fgame/common/lang"
	commonlang "fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/event"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/outlandboss/outlandboss"
	"fgame/fgame/game/outlandboss/pbutil"
	playeroutlandboss "fgame/fgame/game/outlandboss/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//外域boss死亡
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	attackId, ok := data.(int64)
	if !ok {
		return
	}

	//校验怪物类型
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeOutlandBoss {
		return
	}

	//boss死亡（boss信息场景推送）
	scMsg := pbutil.BuildSCOutlandBossInfoBroadcast(n)
	for _, pl := range n.GetScene().GetAllPlayers() {
		pl.SendMsg(scMsg)
	}

	s := n.GetScene()
	if s == nil {
		return
	}

	attackPl := s.GetPlayer(attackId)
	if attackPl == nil {
		return
	}

	pl := attackPl.(player.Player)
	manager := pl.GetPlayerDataManager(playertypes.PlayerOutlandBossDataManagerType).(*playeroutlandboss.PlayerOutlandBossDataManager)
	manager.RefreshZhuoQi()

	// 浊气值上限后继续击杀boss将无法 获得收益
	//needZhuoQi := n.GetBiologyTemplate().ZhuoQi
	if pl.IsZhuoQiLimit() {
		playerlogic.SendSystemMessage(pl, lang.OutlandBossZhuoQiNumNotEnough)
		return
	}

	biologyId := int32(n.GetBiologyTemplate().Id)
	// 浊气值增加
	manager.KillZhuoQiBoss(biologyId)
	// 掉落
	dropIdList := n.GetBiologyTemplate().GetDropIdList()
	itemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)
	showItemDataList := droplogic.MergeItemLevel(itemList)

	//掉落记录
	now := global.GetGame().GetTimeService().Now()
	outlandboss.GetOutlandBossService().AddDropRecords(attackPl.GetName(), biologyId, attackPl.GetMapId(), now, itemList)

	var itemDataList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(itemList) > 0 {
		itemDataList, resMap = droplogic.SeperateItemDatas(itemList)
	}

	// boss死亡（所有攻击者推送击杀奖励信息）
	scEnemiesNoticeMsg := pbutil.BuildSCOutlandBossEnemiesNotice(attackId, attackPl.GetName(), n.GetBiologyTemplate().Name, showItemDataList, biologyId)
	for _, so := range n.GetEnemies() {
		if so.GetHate() <= 1 {
			continue
		}
		switch hurtObj := so.BattleObject.(type) {
		case player.Player:
			{
				hurtObj.SendMsg(scEnemiesNoticeMsg)
			}
			break
		}
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlotsOfItemLevel(itemDataList) {
		title := commonlang.GetLangService().ReadLang(commonlang.EmailInventorySlotNoEnough)
		content := commonlang.GetLangService().ReadLang(commonlang.EmailOutlandBossDropItems)
		emaillogic.AddEmailItemLevel(pl, title, content, now, showItemDataList)
	} else {
		if len(itemDataList) > 0 {
			getReason := commonlog.InventoryLogReasonOutlandBossGet
			flag := inventoryManager.BatchAddOfItemLevel(itemDataList, getReason, getReason.String())
			if !flag {
				panic(fmt.Errorf("outlandboss:击杀外域boss奖励应该添加成功"))
			}
		}

		if len(resMap) > 0 {
			reasonGold := commonlog.GoldLogReasonOutlandBossDrop
			reasonSliver := commonlog.SilverLogReasonOutlandBossDrop
			reasonLevel := commonlog.LevelLogReasonOutlandBossDrop
			err = droplogic.AddRes(pl, resMap, reasonGold, reasonGold.String(), reasonSliver, reasonSliver.String(), reasonLevel, reasonLevel.String())
			if err != nil {
				return
			}
		}

		inventorylogic.SnapInventoryChanged(pl)
		propertylogic.SnapChangedProperty(pl)
	}

	eventData := sceneeventtypes.CreateBattleObjectDropIntoInventoryData(pl.GetId(), showItemDataList)
	gameevent.Emit(sceneeventtypes.EventTypeBattleObjectDropIntoInventory, n, eventData)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
