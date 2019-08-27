package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	playerarena "fgame/fgame/game/arena/player"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	commontypes "fgame/fgame/game/common/types"
	pbutildianxing "fgame/fgame/game/dianxing/pbutil"
	playerdianxing "fgame/fgame/game/dianxing/player"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	"fgame/fgame/game/fourgod/fourgod"
	pbuitl "fgame/fgame/game/fourgod/pbutil"
	playerfourgod "fgame/fgame/game/fourgod/player"
	fourgodscene "fgame/fgame/game/fourgod/scene"
	"fgame/fgame/game/gem/pbutil"
	playergem "fgame/fgame/game/gem/player"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	playerjieyi "fgame/fgame/game/jieyi/player"
	pbutilmassacre "fgame/fgame/game/massacre/pbutil"
	playermassacre "fgame/fgame/game/massacre/player"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	playerqixue "fgame/fgame/game/qixue/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	playershenqi "fgame/fgame/game/shenqi/player"
	playershenyu "fgame/fgame/game/shenyu/player"
	gametemplate "fgame/fgame/game/template"
	playerxiantao "fgame/fgame/game/xiantao/player"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func PlayerEnterFourGodScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) (flag bool, err error) {
	return PlayerEnterFourGodSceneArgs(pl, activityTemplate, "")
}

func PlayerEnterFourGodSceneArgs(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	s := pl.GetScene()
	mapType := s.MapTemplate().GetMapType()
	if mapType == scenetypes.SceneTypeFourGodWar {
		return
	}

	sd := fourgod.GetFourGodService().GetFourGodSceneData()
	if sd == nil {
		now := global.GetGame().GetTimeService().Now()
		openTime := global.GetGame().GetServerTime()
		mergeTime := merge.GetMergeService().GetMergeTime()
		act := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeFourGod)
		timeTemplate, err := act.GetActivityTimeTemplate(now, openTime, mergeTime)
		if err != nil {
			return false, err
		}
		if timeTemplate == nil {
			return false, err
		}
		endTime, err := timeTemplate.GetEndTime(now)
		if err != nil {
			return false, err
		}
		sd = fourgod.GetFourGodService().CreateFourGodSceneData(act.Mapid, endTime)
		if sd == nil {
			return false, err
		}

		//记录活动结束时间
		mananger := pl.GetPlayerDataManager(types.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
		mananger.EndTime(endTime)
	}

	warScene := sd.GetScene()
	pos := warScene.MapTemplate().GetBornPos()
	if !scenelogic.PlayerEnterScene(pl, warScene, pos) {
		return
	}

	flag = true
	return
}

//广播消息
func BroadcastMsgInScene(allPlayer map[int64]scene.Player, msg proto.Message) {
	for _, spl := range allPlayer {
		pl := spl.(player.Player)
		if pl.GetScene().MapTemplate().GetMapType() != scenetypes.SceneTypeFourGodWar {
			continue
		}
		pl.SendMsg(msg)
	}
	return
}

//四神遗迹开箱子奖励
func GiveFourGodOpenBoxReward(pl player.Player, boxTemplate *gametemplate.FourGodBoxTemplate, rewItemList []*droptemplate.DropItemData, resMap map[itemtypes.ItemAutoUseResSubType]int32, keyNum int32) (isReturn bool) {
	//获取奖励
	rewData := boxTemplate.GetRewData()
	//奖励物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	silver := int32(0)
	gold := int32(0)
	bindGold := int32(0)
	exp := int32(0)
	stone := int32(0)
	shaqiNum := int32(0)
	xingchenNum := int32(0)
	gongdeNum := int32(0)
	equipBaoKuPoint := int32(0)
	lingqiNum := int64(0)
	shenYuKey := int32(0)
	qXianTaoNum := int32(0)
	bXianTaoNum := int32(0)
	shaLuXin := int32(0)
	arenaPoint := int32(0)
	shengWei := int32(0)
	arenapvpJiFen := int32(0)
	materialJiFen := int32(0)

	for subType, num := range resMap {
		curSilver, curGold, curBindGold, curExp, curStone, curShaQi, curXingChen, curGongDe, curEquipBaoKuPoint, curLingQiNum, curShenYuKey, curQXianTao, curBXianTao, curShaLuXin, curArenaPoint, curShengWei, curArenapvpJiFen, curMaterialJiFen := droplogic.GetAddAutoRes(subType, num)
		silver += curSilver
		gold += curGold
		bindGold += curBindGold
		exp += curExp
		stone += curStone
		shaqiNum += curShaQi
		xingchenNum += curXingChen
		gongdeNum += curGongDe
		equipBaoKuPoint += curEquipBaoKuPoint
		lingqiNum += int64(curLingQiNum)
		shenYuKey += curShenYuKey
		qXianTaoNum += curQXianTao
		bXianTaoNum += curBXianTao
		shaLuXin += curShaLuXin
		arenaPoint += curArenaPoint
		shengWei += curShengWei
		arenapvpJiFen += curArenapvpJiFen
		materialJiFen += curMaterialJiFen
	}

	if len(rewItemList) != 0 {
		//添加物品
		if !inventoryManager.HasEnoughSlotsOfItemLevel(rewItemList) {
			isReturn = true
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
		//物品奖励
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonFourGodOpenBox.String(), keyNum)
		flag := inventoryManager.BatchAddOfItemLevel(rewItemList, commonlog.InventoryLogReasonFourGodOpenBox, reasonText)
		if !flag {
			panic(fmt.Errorf("fourgod: fourGodOpenBox BatchAdd  should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	var totalRewData *propertytypes.RewData
	if rewData != nil || silver != 0 || gold != 0 || bindGold != 0 || exp != 0 {
		if rewData == nil {
			totalRewData = propertytypes.CreateRewData(exp, 0, silver, gold, bindGold)
		} else {
			totalRewData = propertytypes.CreateRewData(rewData.RewExp+exp, rewData.RewExpPoint, rewData.RewSilver+silver, rewData.RewGold+gold, rewData.RewBindGold+bindGold)
		}
		reasonGold := commonlog.GoldLogReasonFourGodOpenBox
		reasonSilver := commonlog.SilverLogResonFourGodOpenBox
		reasonLevel := commonlog.LevelLogReasonFourGodOpenBox
		reasonGoldText := fmt.Sprintf(reasonGold.String(), keyNum)
		reasonSliverText := fmt.Sprintf(reasonSilver.String(), keyNum)
		reasonlevelText := fmt.Sprintf(reasonLevel.String(), keyNum)
		flag := propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("fourgod: fourGodOpenBox AddRewData  should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	if stone != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)
		mine := manager.DropYuanShi(stone)
		scGemMineGet := pbutil.BuildSCGemMineGet(mine)
		pl.SendMsg(scGemMineGet)
	}

	if shaqiNum != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
		_ = manager.DropShaQi(shaqiNum)
		scMassacreShaQiVary := pbutilmassacre.BuildSCMassacreShaQiVary(int64(shaqiNum))
		pl.SendMsg(scMassacreShaQiVary)
	}

	if xingchenNum != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
		_ = manager.DropXingChen(xingchenNum)
		scDianxingXingchenVary := pbutildianxing.BuildSCDianxingXingchenVary(int64(xingchenNum))
		pl.SendMsg(scDianxingXingchenVary)
	}
	if gongdeNum != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
		manager.AddGongDe(int64(gongdeNum))
	}
	if equipBaoKuPoint != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
		_ = manager.AttendEquipBaoKu(0, equipBaoKuPoint, 0, commontypes.ChangeTypeItemGet, equipbaokutypes.BaoKuTypeEquip)
	}
	if materialJiFen != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
		_ = manager.AttendEquipBaoKu(0, materialJiFen, 0, commontypes.ChangeTypeItemGet, equipbaokutypes.BaoKuTypeMaterials)
	}
	if lingqiNum != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
		_ = manager.AddLingQiNum(lingqiNum)
	}
	if shenYuKey != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerShenYuDataManagerType).(*playershenyu.PlayerShenYuDataManager)
		manager.AddKeyNum(shenYuKey)
	}
	if qXianTaoNum != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
		manager.AddHighPeachCount(qXianTaoNum)
	}
	if bXianTaoNum != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
		manager.AddJuniorPeachCount(bXianTaoNum)
	}
	if shaLuXin != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerQiXueDataManagerType).(*playerqixue.PlayerQiXueDataManager)
		manager.AddShaLu(shaLuXin)
	}
	if arenaPoint != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
		manager.AddJiFen(arenaPoint)
	}
	if shengWei != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
		manager.AddShengWeiZhi(shengWei)
	}
	if arenapvpJiFen != 0 {
		manager := pl.GetPlayerDataManager(types.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
		manager.AddJiFen(arenapvpJiFen)
	}
	return
}

//四神遗迹采集宝箱被打断
func CollectBoxInterrupt(pl scene.Player) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	sd := s.SceneDelegate()
	if sd == nil {
		return
	}

	fourGodWarSceneData, ok := sd.(fourgodscene.FourGodWarSceneData)
	if !ok {
		return
	}
	npcId, hasCollect := fourGodWarSceneData.HasCollectBox(pl.GetId())
	if !hasCollect {
		return
	}
	fourGodWarSceneData.CollectBoxInterrupt(npcId)

	tpl := pl.(player.Player)
	scFourGodOpenBoxStop := pbuitl.BuildSCFourGodOpenBoxStop(npcId)
	tpl.SendMsg(scFourGodOpenBoxStop)
}
