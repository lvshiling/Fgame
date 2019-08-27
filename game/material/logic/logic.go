package logic

import (
	"fgame/fgame/common/lang"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"

	// commonlog "fgame/fgame/common/log"
	// emaillogic "fgame/fgame/game/email/logic"
	// inventorylogic "fgame/fgame/game/inventory/logic"
	commonlog "fgame/fgame/common/log"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	materialeventtypes "fgame/fgame/game/material/event/types"
	"fgame/fgame/game/material/pbutil"
	playermaterial "fgame/fgame/game/material/player"
	materialtemplate "fgame/fgame/game/material/template"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertytypes "fgame/fgame/game/property/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func CheckIfCanEnterMaterial(pl player.Player, materialType materialtypes.MaterialType) (flag bool) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return false
	}

	//TODO 添加功能开启判断
	if !pl.IsFuncOpen(materialType.GetFuncOpenType()) {
		return
	}

	materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(materialType)
	if materialTemplate == nil {
		return false
	}

	//刷新数据
	materialManager := pl.GetPlayerDataManager(playertypes.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	materialManager.RefreshData()

	//是否免费次数
	if !materialManager.IsFreeTimes(materialType) {
		//挑战次数是否足够
		if !materialManager.IsEnoughAttendTimes(materialType, 1) {
			return false
		}

		//挑战所需物品是否足够
		needItemId := materialTemplate.NeedItemId
		needItemNum := materialTemplate.NeedItemCount
		if !inventoryManager.HasEnoughItem(needItemId, needItemNum) {
			return false
		}

	}

	return true
}

//材料副本挑战逻辑
func HandlePlayerMaterialChallenge(pl player.Player, materialType materialtypes.MaterialType) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	if !pl.IsFuncOpen(materialType.GetFuncOpenType()) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"materialType": materialType,
			}).Warn("material:材料副本挑战请求,功能未开始")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(materialType)
	if materialTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"materialType": materialType,
			}).Warn("material:材料副本挑战请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//刷新数据
	materialManager := pl.GetPlayerDataManager(playertypes.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	materialManager.RefreshData()

	//是否免费次数
	if !materialManager.IsFreeTimes(materialType) {
		//挑战次数是否足够
		if !materialManager.IsEnoughAttendTimes(materialType, 1) {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"materialType": materialType,
				}).Warn("material:材料副本挑战请求，副本次数不足，无法挑战")
			playerlogic.SendSystemMessage(pl, lang.MaterialNotEnoughChallengeTimes)
			return
		}

		//挑战所需物品是否足够
		needItemId := materialTemplate.NeedItemId
		needItemNum := materialTemplate.NeedItemCount
		if !inventoryManager.HasEnoughItem(needItemId, needItemNum) {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"materialType": materialType,
					"needItemId":   needItemId,
					"needItemNum":  needItemNum,
				}).Warn("material:材料副本挑战请求，道具不足，无法挑战")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		//消耗挑战所需物品
		itemReason := commonlog.InventoryLogReasonMaterialChallenge
		itemReasonText := fmt.Sprintf(itemReason.String(), materialType.String())
		if flag := inventoryManager.UseItem(needItemId, needItemNum, itemReason, itemReasonText); !flag {
			panic(fmt.Errorf("material: materialChallenge use item should be ok"))
		}
	}

	//进入场景
	flag := PlayerEnterMaterial(pl, materialType)
	if !flag {
		panic("enter material scene should be ok!")
	}

	//更新玩家挑战记录
	materialManager.UseTimes(materialType, 1)

	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCMaterialChallenge(materialType, flag)
	pl.SendMsg(scMsg)

	return
}

func PlayerEnterMaterial(pl player.Player, typ materialtypes.MaterialType) (flag bool) {
	materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(typ)
	if materialTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"MaterialType": typ.String(),
			}).Warn("material:处理跳转材料副本,材料副本不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	sh := createMaterialSceneData(pl.GetId(), materialTemplate)
	s := scene.CreateFuBenScene(materialTemplate.MapId, sh)
	if s == nil {
		panic(fmt.Errorf("material:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)

	flag = true
	return
}

//挑战完成奖励
func onMaterialFinish(pl player.Player, curTemp *gametemplate.MaterialTemplate, itemMap map[int32]int32, success bool, useTime int64, group int32) {
	materialType := curTemp.GetMaterialType()
	//记录第几波
	s := pl.GetScene()
	if s != nil {
		data := materialeventtypes.CreateRefreshGroupEventData(group, materialType)
		gameevent.Emit(materialeventtypes.EventTypeMaterialRefreshGroup, pl, data)
	}

	//发送提示
	scMsg := pbutil.BuildSCMaterialChallengeResult(int32(materialType), success, itemMap, group)
	pl.SendMsg(scMsg)
	return
}

//下发场景信息
func onPushSceneInfo(playerId int64, curTemp *gametemplate.MaterialTemplate, s scene.Scene) {
	if s == nil {
		return
	}
	pl := s.GetPlayer(playerId)
	if pl == nil {
		return
	}
	materialType := int32(curTemp.GetMaterialType())
	currentGroup := s.GetCurrentGroup()
	totalGroup := curTemp.GetMapTemplate().GetNumGroup()
	startTime := s.GetStartTime()

	scMsg := pbutil.BuildSCMaterialSceneInfo(materialType, currentGroup, totalGroup, startTime)
	pl.SendMsg(scMsg)
	return
}

//下发怪物刷新波数
func onPushBiologyGroupInfo(pl player.Player, groupIndex int32, curTemp *gametemplate.MaterialTemplate) {

	showGroup := groupIndex + 1
	scMsg := pbutil.BuildSCMaterialRefreshBiology(showGroup)
	pl.SendMsg(scMsg)

	data := materialeventtypes.CreateRefreshGroupEventData(groupIndex, curTemp.GetMaterialType())
	gameevent.Emit(materialeventtypes.EventTypeMaterialRefreshGroup, pl, data)
}

//材料副本扫荡
func GetSaoDangDrop(pl player.Player, saoDangNum int32, materialType materialtypes.MaterialType, group int32) (showItemList [][]*droptemplate.DropItemData, rewardsItemList []*droptemplate.DropItemData, rewardsResMap map[itemtypes.ItemAutoUseResSubType]int32, totalRewData *propertytypes.RewData, flag bool) {
	if group <= 0 {
		panic("material:扫荡波数不能为空")
	}

	materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(materialType)

	var totalItemList []*droptemplate.DropItemData
	dropIdList := materialTemplate.GetSaodangRewardDropArr()
	for index := int32(0); index < saoDangNum; index++ {
		var itemList []*droptemplate.DropItemData
		// 固定奖励
		for itemId, num := range materialTemplate.GetSaodangRewardItemMap(1) {
			num *= group
			newData := droptemplate.CreateItemData(itemId, num, 0, itemtypes.ItemBindTypeUnBind)
			itemList = append(itemList, newData)
		}

		// 掉落包
		dropItemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)
		itemList = append(itemList, dropItemList...)

		showItemList = append(showItemList, itemList)
		totalItemList = append(totalItemList, itemList...)
	}

	if len(totalItemList) != 0 {
		rewardsItemList, rewardsResMap = droplogic.SeperateItemDatas(totalItemList)
	}

	rewSilver := int32(materialTemplate.RawSilver) * saoDangNum * group
	rewBindGold := materialTemplate.RawBindGold * saoDangNum * group
	rewGold := materialTemplate.RawGold * saoDangNum * group
	rewExp := int32(materialTemplate.RawExp) * saoDangNum * group
	rewExpPoint := int32(materialTemplate.RawExpPoint) * saoDangNum * group
	totalRewData = propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)
	flag = true
	return
}
