package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/template"
	coreutils "fgame/fgame/core/utils"
	playerarena "fgame/fgame/game/arena/player"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	collectnpc "fgame/fgame/game/collect/npc"
	pbuitl "fgame/fgame/game/collect/pbutil"
	collecttypes "fgame/fgame/game/collect/types"
	commontypes "fgame/fgame/game/common/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	pbutildianxing "fgame/fgame/game/dianxing/pbutil"
	playerdianxing "fgame/fgame/game/dianxing/player"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	"fgame/fgame/game/gem/pbutil"
	playergem "fgame/fgame/game/gem/player"
	"fgame/fgame/game/global"
	godsiegetemplate "fgame/fgame/game/godsiege/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	playerjieyi "fgame/fgame/game/jieyi/player"
	pbutilmassacre "fgame/fgame/game/massacre/pbutil"
	playermassacre "fgame/fgame/game/massacre/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	playerqixue "fgame/fgame/game/qixue/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	playershenqi "fgame/fgame/game/shenqi/player"
	playershenyu "fgame/fgame/game/shenyu/player"
	gametemplate "fgame/fgame/game/template"
	playerxiantao "fgame/fgame/game/xiantao/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//处理采集信息逻辑
func HandleCollectNpc(pl scene.Player, collectType collecttypes.CollectType, npcId int64) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,场景为空")
		scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
		pl.SendMsg(scSceneCollect)
		return
	}

	mapType := s.MapTemplate().GetMapType()
	if mapType == scenetypes.SceneTypeCrossDenseWat {
		godSiegeTemplate := godsiegetemplate.GetGodSiegeTemplateService().GetConstantTemplate()
		num := pl.GetDenseWatNum()
		if num >= godSiegeTemplate.CaiJiCountLimit {
			scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
			pl.SendMsg(scSceneCollect)
			return
		}
	}
	var so scene.SceneObject
	switch collectType {
	case collecttypes.CollectTypeClient:
		so = s.GetNPCByIdx(npcId)
		if so == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"npcId":    npcId,
				}).Warn("collect:处理采集消息,npc不存在")
			scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
			pl.SendMsg(scSceneCollect)
			return
		}
		if so.GetScene() == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"npcId":    npcId,
				}).Warn("collect:处理采集消息,npc场景不存在")
			scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
			pl.SendMsg(scSceneCollect)
			return
		}
		if so.GetScene() != s {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"npcId":    npcId,
				}).Warn("collect:处理采集消息,人和npc场景不一样")
			scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
			pl.SendMsg(scSceneCollect)
			return
		}
		break
	case collecttypes.CollectTypeServer:
		so = s.GetSceneObject(npcId)
		if so == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"npcId":    npcId,
				}).Warn("collect:处理采集消息,生物不存在")
			scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
			pl.SendMsg(scSceneCollect)
			return
		}
		break
	}

	n, ok := so.(scene.CollectNPC)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,不是采集物")
		scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
		pl.SendMsg(scSceneCollect)
		return
	}

	distance := coreutils.Distance(n.GetPosition(), pl.GetPos())
	collectDistance := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCollectDistance)) / float64(1000)
	if distance > float64(collectDistance) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Warn("collect:不在采集范围内")
		playerlogic.SendSystemMessage(pl, lang.CommonCollectNoDistance)
		scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
		pl.SendMsg(scSceneCollect)
		return
	}

	if n.IsDead() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Warn("collect:该采集物已被采集过,请等待重生")
		playerlogic.SendSystemMessage(pl, lang.CommonCollectIsDead)
		scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
		pl.SendMsg(scSceneCollect)
		return
	}

	conditionOk := collectnpc.IsSceneCollectCondition(pl, n, mapType)
	if !conditionOk {
		return
	}

	_, isCollecting := pl.HasCollect()
	if isCollecting {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("collect:处理采集消息,已经在采集中")
		playerlogic.SendSystemMessage(pl, lang.CollectNowCollecting)
		scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
		pl.SendMsg(scSceneCollect)
		return
	}

	flag, isMax := n.IfCanCollect(pl.GetId())
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"ncpId":    npcId,
		}).Warn("collect:其它玩家正在采集, 采集点次数不足, 已经在采集了")
		if isMax {
			playerlogic.SendSystemMessage(pl, lang.CollectPointNoNum)
		} else {
			playerlogic.SendSystemMessage(pl, lang.CommonCollectOtherExist)
		}
		scSceneCollect := pbuitl.BuildSCSceneCollect(false, npcId, 0)
		pl.SendMsg(scSceneCollect)
		return
	}

	startTime, flag := n.StartCollect(pl)
	if !flag {
		panic(fmt.Errorf("collect: 采集应该是ok的"))
	}

	scSceneCollect := pbuitl.BuildSCSceneCollect(true, npcId, startTime)
	pl.SendMsg(scSceneCollect)
	return
}

//采集被打断
func CollectInterrupt(pl scene.Player, n scene.CollectNPC) {
	n.CollectInterrupt(pl)
	// npc := n.GetCollectNPC()
	pl.ClearCollect()

	scSceneCollectStop := pbuitl.BuildSCSceneCollectStop(n.GetId())
	pl.SendMsg(scSceneCollectStop)
}

//采集物掉落直接入背包
func CollectDropToInventory(pl player.Player, biologyId int32) (itemDataList []*droptemplate.DropItemData) {
	to := template.GetTemplateService().Get(int(biologyId), (*gametemplate.BiologyTemplate)(nil))
	if to == nil {
		return
	}
	biologyTempalte := to.(*gametemplate.BiologyTemplate)
	dropIdList := biologyTempalte.GetDropIdList()
	dropItemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)

	var rewItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(dropItemList) > 0 {
		rewItemList, resMap = droplogic.SeperateItemDatas(dropItemList)
	}
	GiveCollectDropReward(pl, rewItemList, resMap)
	return dropItemList
}

//采集物选择掉落直接入背包
func CollectChooseDropToInventory(pl player.Player, biologyId int32, typ collecttypes.CollectChooseFinishType) (itemDataList []*droptemplate.DropItemData) {
	to := template.GetTemplateService().Get(int(biologyId), (*gametemplate.BiologyTemplate)(nil))
	if to == nil {
		return
	}
	biologyTempalte := to.(*gametemplate.BiologyTemplate)
	dropIdList := biologyTempalte.GetDropIdList()
	chooseDropId, ok := biologyTempalte.GetCaiJiChooseDropId(typ)
	if !ok {
		panic("collect choose drop should be ok")
	}
	dropIdList = append(dropIdList, chooseDropId)
	dropItemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)

	var rewItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(dropItemList) > 0 {
		rewItemList, resMap = droplogic.SeperateItemDatas(dropItemList)
	}
	GiveCollectDropReward(pl, rewItemList, resMap)
	return dropItemList
}

//采集掉落奖励
func GiveCollectDropReward(pl player.Player, rewItemList []*droptemplate.DropItemData, resMap map[itemtypes.ItemAutoUseResSubType]int32) {
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
		curSilver, curGold, curBindGold, curExp, curStone, curShaQi, curXingChen, curGongde, curEquipBaoKuPoint, curLingQiNum, curShenYuKey, curQXianTao, curBXianTao, curShaLuXin, curArenaPoint, curShengWei, curArenapvpJiFen, curMaterialJiFen := droplogic.GetAddAutoRes(subType, num)
		silver += curSilver
		gold += curGold
		bindGold += curBindGold
		exp += curExp
		stone += curStone
		shaqiNum += curShaQi
		xingchenNum += curXingChen
		gongdeNum += curGongde
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
			//添加邮件
			collectTitle := lang.GetLangService().ReadLang(lang.CollectTitle)
			collectContent := lang.GetLangService().ReadLang(lang.CollectContent)
			now := global.GetGame().GetTimeService().Now()
			emaillogic.AddEmailItemLevel(pl, collectTitle, collectContent, now, rewItemList)
		} else {
			//物品奖励
			reasonText := fmt.Sprintf(commonlog.InventoryLogReasonCollect.String())
			flag := inventoryManager.BatchAddOfItemLevel(rewItemList, commonlog.InventoryLogReasonCollect, reasonText)
			if !flag {
				panic(fmt.Errorf("collect: CollectDropToInventory BatchAdd  should be ok"))
			}
			inventorylogic.SnapInventoryChanged(pl)
		}

	}

	var totalRewData *propertytypes.RewData
	if silver != 0 || gold != 0 || bindGold != 0 || exp != 0 {
		totalRewData = propertytypes.CreateRewData(exp, 0, silver, gold, bindGold)
		reasonGold := commonlog.GoldLogReasonCollectRew
		reasonSilver := commonlog.SilverLogReasonCollectRew
		reasonLevel := commonlog.LevelLogReasonCollectRew
		flag := propertyManager.AddRewData(totalRewData, reasonGold, reasonGold.String(), reasonSilver, reasonSilver.String(), reasonLevel, reasonLevel.String())
		if !flag {
			panic(fmt.Errorf("collect: CollectDropToInventory AddRewData  should be ok"))
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
