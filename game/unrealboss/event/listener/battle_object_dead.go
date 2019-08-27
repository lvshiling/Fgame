package listener

import (
	"fgame/fgame/common/lang"
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
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/unrealboss/pbutil"
	playerunrealboss "fgame/fgame/game/unrealboss/player"
	"fmt"
)

//幻境boss死亡
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
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeUnrealBoss {
		return
	}

	//boss死亡（boss信息场景推送）
	scMsg := pbutil.BuildSCUnrealBossInfoBroadcast(n)
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
	unrealManager := pl.GetPlayerDataManager(playertypes.PlayerUnrealBossDataManagerType).(*playerunrealboss.PlayerUnrealBossDataManager)
	unrealManager.RefreshPilao()

	// 疲劳值扣光后继续击杀boss将无法 获得收益
	needPilao := n.GetBiologyTemplate().NeedPilao
	if !pl.IsEnoughPilao(needPilao) {
		playerlogic.SendSystemMessage(pl, lang.UnrealBossPilaoNumNotEnough)
		return
	}

	// 掉落
	dropIdList := n.GetBiologyTemplate().GetDropIdList()
	itemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)
	showItemDataList := droplogic.MergeItemLevel(itemList)

	var itemDataList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(itemList) > 0 {
		itemDataList, resMap = droplogic.SeperateItemDatas(itemList)
	}

	biologyId := int32(n.GetBiologyTemplate().Id)
	// boss死亡（所有攻击者推送击杀奖励信息）
	scEnemiesNoticeMsg := pbutil.BuildSCUnrealBossEnemiesNotice(attackId, attackPl.GetName(), n.GetBiologyTemplate().Name, showItemDataList, biologyId)
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

	// 疲劳值仅在获得奖励时扣除
	unrealManager.KillPiLaoBoss(biologyId)

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlotsOfItemLevel(itemDataList) {
		title := "背包空间不足"
		content := "幻境BOSS掉落"
		now := global.GetGame().GetTimeService().Now()
		emaillogic.AddEmailItemLevel(pl, title, content, now, showItemDataList)
	} else {
		if len(itemDataList) > 0 {
			getReason := commonlog.InventoryLogReasonUnrealBossGet
			flag := inventoryManager.BatchAddOfItemLevel(itemDataList, getReason, getReason.String())
			if !flag {
				panic(fmt.Errorf("unrealboss:击杀幻境boss奖励应该添加成功"))
			}
		}

		if len(resMap) > 0 {
			reasonGold := commonlog.GoldLogReasonUnrealBossDrop
			reasonSliver := commonlog.SilverLogReasonUnrealBossDrop
			reasonLevel := commonlog.LevelLogReasonUnrealBossDrop
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
