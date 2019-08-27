package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	majoreventtypes "fgame/fgame/game/major/event/types"
	"fgame/fgame/game/major/pbutil"
	majortemplate "fgame/fgame/game/major/template"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertytypes "fgame/fgame/game/property/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func PlayerEnterMajor(pl player.Player, spouseId int64, majorType majortypes.MajorType, fubenId int32) (s scene.Scene, flag bool) {
	majorTaTemplate := majortemplate.GetMajorTemplateService().GetMajorTemplate(majorType, fubenId)
	if majorTaTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("major:处理跳转双修,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	sh := CreateMajorSceneData(pl.GetId(), spouseId, majorTaTemplate)
	s = scene.CreateFuBenScene(majorTaTemplate.GetMapId(), sh)
	if s == nil {
		panic(fmt.Errorf("major:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)
	flag = true
	return
}

//下发场景信息
func onPushSceneInfo(pl player.Player, starTime int64, ownerId int64, spouseId int64, majorType, fubenId int32) (err error) {
	//推送给前端
	scMajorScene := pbutil.BuildSCMajorScene(starTime, ownerId, spouseId, majorType, fubenId)
	pl.SendMsg(scMajorScene)
	return
}

//TODO 奖励
func onMajorFinish(sd *MajorSceneData, sucessful bool) {
	playerIdList := make([]int64, 0, 2)
	playerIdList = append(playerIdList, sd.ownerId, sd.spouseId)
	if sd.currentMajorTemplate == nil {
		return
	}
	rewItemMap := sd.currentMajorTemplate.GetRewItemMap()
	for _, playerId := range playerIdList {
		//主动退出
		if playerId == 0 {
			continue
		}
		//TODO ylz:通过场景获取
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		//玩家掉线
		if pl == nil {
			continue
		}
		s := pl.GetScene()
		if s.MapTemplate().GetMapType() != scenetypes.SceneTypeMajor &&
			s.MapTemplate().GetMapType() != scenetypes.SceneTypeFuQiFuBen {
			continue
		}
		scMajorResult := pbutil.BuildSCMajorResult(sucessful, int32(sd.currentMajorTemplate.GetMajorType()), int32(sd.currentMajorTemplate.TemplateId()))
		pl.SendMsg(scMajorResult)
		if sucessful {
			majorRewItem(pl, rewItemMap)

			//通关事件
			data := majoreventtypes.CreateMajorSuccessEventData(sd.currentMajorTemplate, 1)
			gameevent.Emit(majoreventtypes.EventTypePlayerMajorSuccess, pl, data)
		}
	}
	return
}

func majorRewItem(pl player.Player, rewItemMap map[int32]int32) {
	if len(rewItemMap) == 0 {
		return
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//奖励物品
	flag := inventoryManager.HasEnoughSlots(rewItemMap)
	if !flag {
		//写邮件
		emailTitle := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MajorRewTitle))
		emailContent := lang.GetLangService().ReadLang(lang.MajorRewContent)
		emaillogic.AddEmail(pl, emailTitle, emailContent, rewItemMap)
	} else {
		reasonInventory := commonlog.InventoryLogReasonMajorRew
		flag = inventoryManager.BatchAdd(rewItemMap, reasonInventory, reasonInventory.String())
		if !flag {
			panic(fmt.Errorf("realm: ToKillFinished BatchAdd should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}
	return

}

//副本扫荡
func GetSaoDangDrop(pl player.Player, saoDangNum int32, majorType majortypes.MajorType, fubenId int32) (showItemList [][]*droptemplate.DropItemData, rewardsItemList []*droptemplate.DropItemData, rewardsResMap map[itemtypes.ItemAutoUseResSubType]int32, totalRewData *propertytypes.RewData, flag bool) {
	majorTemp := majortemplate.GetMajorTemplateService().GetMajorTemplate(majorType, fubenId)

	var totalItemList []*droptemplate.DropItemData
	dropIdList := majorTemp.GetSaodangRewardDropArr()
	for index := int32(0); index < saoDangNum; index++ {
		var itemList []*droptemplate.DropItemData
		//TODO:cjb 优化
		// 固定奖励
		for itemId, num := range majorTemp.GetSaodangRewardItemMap(1) {
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

	rewSilver := int32(majorTemp.GetRawSilver()) * saoDangNum
	rewBindGold := majorTemp.GetRawBindGold() * saoDangNum
	rewGold := majorTemp.GetRawGold() * saoDangNum
	rewExp := int32(majorTemp.GetRawExp()) * saoDangNum
	rewExpPoint := int32(majorTemp.GetRawExpPoint()) * saoDangNum
	totalRewData = propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)
	flag = true
	return
}
