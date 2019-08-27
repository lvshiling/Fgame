package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancescene "fgame/fgame/game/alliance/scene"
	alliancetemplate "fgame/fgame/game/alliance/template"
	"fgame/fgame/game/center/center"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gamecommon "fgame/fgame/game/common/common"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

const (
	batchLimit = 100
)

//检查是否可以加入仙盟
func CheckPlayerIfCanBatchJoinAlliance(pl player.Player) (flag bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeAlliance) {
		return false
	}
	if pl.GetAllianceId() != 0 {
		return false
	}
	//CD
	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	if allianceManager.IsOnBatchJoinCD() {
		return false
	}

	alList := alliance.GetAllianceService().GetAllianceList()
	if len(alList) <= 0 {
		return false
	}
	return true
}

//批量加入仙盟
func HandleAllianceJoinBatch(pl player.Player) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeAlliance) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理批量加入仙盟,仙盟功能未开放")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	//CD
	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	if allianceManager.IsOnBatchJoinCD() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理批量加入仙盟,一键加入CD中")
		playerlogic.SendSystemMessage(pl, lang.AllianceBatchJoinCD)
		return
	}
	allianceManager.UpdateLastBatchJoinTime()
	alList := alliance.GetAllianceService().GetAllianceList()
	limit := batchLimit
	if len(alList) < limit {
		limit = len(alList)
	}
	batchList := alList[:limit]
	for _, al := range batchList {
		isJoinIn, err := HandleApplyJoinAlliance(pl, al.GetAllianceId())
		if err != nil {
			continue
		}
		if isJoinIn {
			break
		}
	}

	scMsg := pbutil.BuildSCAllianceJoinApplyBatch()
	pl.SendMsg(scMsg)
	return
}

func HandleApplyJoinAlliance(pl player.Player, allianceId int64) (isJoinIn bool, err error) {
	//申请加入仙盟
	alreadyApplay := false
	name := pl.GetName()
	role := pl.GetRole()
	sex := pl.GetSex()
	level := pl.GetLevel()
	force := pl.GetForce()
	// al, applyObj, err := alliance.GetAllianceService().ApplyJoinAlliance(allianceId, pl.GetId(), pl.GetCamp(), name, role, sex, level, force)
	al, applyObj, err := alliance.GetAllianceService().ApplyJoinAlliance(allianceId, pl.GetId(), name, role, sex, level, force)
	if err != nil {
		terr, ok := err.(gamecommon.Error)
		if !ok {
			return
		}
		if terr.Code() == lang.AllianceAlreadyApply {
			alreadyApplay = true
		} else {
			return
		}
	}

	//第一次申请广播管理层
	if !alreadyApplay && applyObj != nil {
		scAllianceJoinApplyBroadcast := pbutil.BuildSCAllianceJoinApplyBroadcast(applyObj)
		for _, manager := range al.GetAllManagers() {
			managerPlayer := player.GetOnlinePlayerManager().GetPlayerById(manager.GetMemberId())
			if managerPlayer == nil {
				continue
			}

			managerPlayer.SendMsg(scAllianceJoinApplyBroadcast)
		}
	}

	// 免审核，直接加入仙盟
	if applyObj == nil && err == nil {
		isJoinIn = true
		//通知用户
		joinId := pl.GetId()
		joinPlayer := player.GetOnlinePlayerManager().GetPlayerById(joinId)
		if joinPlayer != nil {
			//同步用户数据
			joinName := joinPlayer.GetName()
			joinSex := joinPlayer.GetSex()
			joinLevel := joinPlayer.GetLevel()
			joinForce := joinPlayer.GetForce()
			joinZhuanSheng := joinPlayer.GetZhuanSheng()
			//TODO xzk:优化不要用GetLingyuInfo
			joinLingyu := joinPlayer.GetLingyuInfo().AdvanceId
			joinVip := joinPlayer.GetVip()
			alliance.GetAllianceService().SyncMemberInfo(joinId, joinName, joinSex, joinLevel, joinForce, joinZhuanSheng, joinLingyu, joinVip)

			mem := alliance.GetAllianceService().GetAllianceMember(joinId)
			if mem == nil {
				panic("alliance agree join: 成员应该存在")
			}
			scAllianceInfo := pbutil.BuildSCAllianceInfo(al, mem)
			joinPlayer.SendMsg(scAllianceInfo)

			allianceId := al.GetAllianceId()
			allianceName := al.GetAllianceObject().GetName()
			scMsg := pbutil.BuildSCAllianceAgreeJoinApplyToApply(allianceId, allianceName, true)
			joinPlayer.SendMsg(scMsg)
		}
	}
	return
}

//仙盟仓库存入
func CheckPlayerIfCanSaveAllianceDepot(pl player.Player, index, num int32) (flag bool) {
	if !center.GetCenterService().IsAllianceOpen() {
		return false
	}
	if pl.GetAllianceId() == 0 {
		return false
	}
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	it := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
	if it == nil || it.IsEmpty() {
		return
	}

	itemId := it.ItemId
	curNum := it.Num
	level := it.Level
	bind := it.BindType
	propertyData := it.PropertyData

	if curNum < num {
		return
	}

	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {

		return
	}

	//是否能存入
	if !itemTemp.IsCanSaveInAllianceDepot() {
		return
	}

	// 存入
	itemData := droptemplate.CreateItemData(itemId, num, level, bind)
	return alliance.GetAllianceService().HasEnoughDepotSlot(pl.GetAllianceId(), itemData, propertyData)
}

//仙盟仓库存入
func HandleSaveAllianceDepot(pl player.Player, index, num int32) (err error) {
	if !center.GetCenterService().IsAllianceOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:仙盟仓库关闭中")
		playerlogic.SendSystemMessage(pl, lang.AllianceDepotClose)
		return
	}
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	it := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("alliance:处理仙盟仓库存入,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	itemId := it.ItemId
	curNum := it.Num
	level := it.Level
	bind := it.BindType
	propertyData := it.PropertyData

	if curNum < num {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"curNum":   curNum,
				"num":      num,
			}).Warn("alliance:处理仙盟仓库存入,超过当前位置最大数量")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("alliance:处理仙盟仓库存入,物品数据模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//是否能存入
	if !itemTemp.IsCanSaveInAllianceDepot() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("alliance:处理仙盟仓库存入,该物品不能放入仙盟仓库")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// 存入
	itemData := droptemplate.CreateItemData(itemId, num, level, bind)
	al, err := alliance.GetAllianceService().SaveInDepot(pl.GetAllianceId(), itemData, propertyData)
	if err != nil {
		return
	}

	// 背包移除
	reason := commonlog.InventoryLogReasonSaveInAllianceDepot
	flag, err := inventoryManager.RemoveIndex(inventorytypes.BagTypePrim, index, num, reason, reason.String())
	if err != nil {
		return
	}
	if !flag {
		panic(fmt.Errorf("alliance:通过索引移除物品应该成功,index:%d,num:%d", index, num))
	}

	// 获取积分
	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	beforePoint := allianceManager.GetDepotPoint()
	point := itemTemp.UnionGet * num
	allianceManager.AddDepotPoint(point)
	inventorylogic.SnapInventoryChanged(pl)

	// 仓库放入日志
	saveItemReason := commonlog.AllianceLogReasonDepotSaveItem
	saveItemReasonText := fmt.Sprintf(saveItemReason.String(), pl.GetId())
	depotLogEventData := allianceeventtypes.CreateAllianceDepotItemChangedLogEventData(itemId, num, saveItemReason, saveItemReasonText)
	gameevent.Emit(allianceeventtypes.EventTypeAllianceDepotItemChangedLog, al, depotLogEventData)

	//玩家积分变化日志
	pointReason := commonlog.AllianceLogReasonPlayerDepotPointChanged
	pointReasonText := fmt.Sprintf(pointReason.String(), itemId)
	pointLogEventData := allianceeventtypes.CreatePlayerAllianceDepotPointLogEventData(beforePoint, point, pointReason, pointReasonText)
	gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceDepotPointChangedLog, pl, pointLogEventData)

	//广播帮派
	if itemTemp.GetQualityType() >= itemtypes.ItemQualityTypeOrange {
		format := lang.GetLangService().ReadLang(lang.AllianceDepotSaveNotice)
		memName := coreutils.FormatColor(chattypes.ColorTypePlayerName, pl.GetName())
		itemName := coreutils.FormatColor(itemTemp.GetQualityType().GetColor(), coreutils.FormatNoticeStrUnderline(itemTemp.FormateItemNameOfNum(num)))
		linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
		itemNameLink := coreutils.FormatLink(itemName, linkArgs)

		args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(funcopentypes.FuncOpenTypeAlliance), al.GetAllianceId()}
		link := coreutils.FormatLink(chattypes.ButtonTypeToLook, args)
		content := fmt.Sprintf(format, memName, itemNameLink, point, link)
		chatlogic.BroadcastAllianceSystem(al.GetAllianceId(), pl.GetId(), pl.GetName(), chattypes.MsgTypeText, []byte(content), "")
	}

	curPoint := allianceManager.GetDepotPoint()
	scMsg := pbutil.BuildSCSaveInAllianceDepot(curPoint)
	pl.SendMsg(scMsg)
	return
}

const (
	numOfDoor = 3
)

func ChenZhanCheckMove(p scene.Player, pos coretypes.Position) (flag bool, fixPos coretypes.Position) {
	s := p.GetScene()
	if s == nil {
		flag = true
		return
	}
	sd, ok := s.SceneDelegate().(alliancescene.AllianceSceneData)
	if !ok {
		flag = true
		return
	}

	destArea := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().GetArea(pos)
	//TODO 优化3个门
	//城门全破了
	isDefend := sd.GetCurrentDefendAllianceId() == p.GetAllianceId()
	currentDoor := sd.GetCurrentDoor()
	if currentDoor >= numOfDoor {
		flag = true
		return
	}
	switch currentDoor {
	case 0:
		if isDefend && (destArea <= 1) {
			flag = true
			return
		}
		if destArea <= 0 {
			flag = true
			return
		}
		flag = false
		sourceArea := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().GetArea(p.GetPosition())
		if sourceArea != 0 {
			fixPos = s.MapTemplate().GetBornPos()
		} else {
			tempFixPos, tflag := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().GetFixPos(currentDoor)
			if tflag {
				if s.MapTemplate().GetMap().IsMask(tempFixPos.X, tempFixPos.Z) {
					fixPos = tempFixPos
					return
				}
			}
			fixPos = tempFixPos
			fixPos.Y = s.MapTemplate().GetMap().GetHeight(fixPos.X, fixPos.Z)
			return
		}
		return
	default:
		if destArea <= currentDoor {
			flag = true
			return
		}
		sourceArea := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().GetArea(p.GetPosition())
		if sourceArea >= destArea {
			fixPos = s.MapTemplate().GetBornPos()
		} else {
			tempFixPos, tflag := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().GetFixPos(currentDoor)
			if tflag {
				if s.MapTemplate().GetMap().IsMask(tempFixPos.X, tempFixPos.Z) {
					fixPos = tempFixPos
					return
				}
			}
			fixPos = tempFixPos
			fixPos.Y = s.MapTemplate().GetMap().GetHeight(fixPos.X, fixPos.Z)
			return
		}
		return
	}
	flag = true
	return
}

func IfHuGongClose(s scene.Scene) (flag bool) {
	//城战特殊处理
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeHuangGong {
		now := global.GetGame().GetTimeService().Now()
		if (s.GetEndTime() - now) <= int64(alliancetemplate.GetAllianceTemplateService().GetWarTemplate().HuanggongTime) {
			return true
		}
	}
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeChengZhan {
		now := global.GetGame().GetTimeService().Now()
		if (s.GetEndTime() - now) <= int64(alliancetemplate.GetAllianceTemplateService().GetWarTemplate().HuanggongTime) {
			return true
		}
	}
	return false
}
