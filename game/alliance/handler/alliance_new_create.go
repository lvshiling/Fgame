package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coredirty "fgame/fgame/core/dirty"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/center/center"
	chargetemplate "fgame/fgame/game/charge/template"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	"fgame/fgame/game/inventory/inventory"
	inventorylogic "fgame/fgame/game/inventory/logic"
	inventoryplayer "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_CREATE_TYPE), dispatch.HandlerFunc(handleAllianceNewCreate))
}

func handleAllianceNewCreate(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance: 处理仙盟创建消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAllianceCreate := msg.(*uipb.CSAllianceCreate)
	name := csAllianceCreate.GetName()
	versionType := alliancetypes.AllianceVersionType(csAllianceCreate.GetVersionType())
	if !versionType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"name":        name,
				"versionType": versionType,
				"error":       err,
			}).Warn("alliance: 处理仙盟创建,仙盟版本类型不合法")
		return
	}
	allianceType := alliancetypes.AllianceNewType(0)
	if versionType == alliancetypes.AllianceVersionTypeNew {
		allianceType = alliancetypes.AllianceNewType(csAllianceCreate.GetAllianceType())
		if !allianceType.Valid() {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"name":         name,
					"allianceType": allianceType,
					"error":        err,
				}).Warn("alliance: 处理仙盟创建,新版本仙盟类型不合法")
			return
		}
	}

	err = allianceNewCreate(tpl, name, versionType, allianceType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"name":         name,
				"versionType":  versionType,
				"allianceType": allianceType,
				"error":        err,
			}).Error("alliance: 处理仙盟创建,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"name":         name,
			"versionType":  versionType,
			"allianceType": allianceType,
		}).Debug("alliance: 处理仙盟创建,完成")
	return nil
}

const (
	minAllianceNameLen = 1
	maxAllianceNameLen = 6
)

func allianceNewCreate(pl player.Player, name string, versionType alliancetypes.AllianceVersionType, allianceType alliancetypes.AllianceNewType) (err error) {
	// 判断创建仙盟版本
	versionType = alliancetypes.AllianceVersionTypeOld
	sdkList := center.GetCenterService().GetSdkList()
	for _, sdk := range sdkList {
		temp := chargetemplate.GetChargeTemplateService().GetQuDaoTemplateByType(sdk)
		if temp == nil {
			continue
		}
		if temp.GetAllianceVersion() == alliancetypes.AllianceVersionTypeNew {
			versionType = alliancetypes.AllianceVersionTypeNew
			break
		}
	}

	name = strings.TrimSpace(name)
	lenOfName := len([]rune(name))

	if lenOfName < minAllianceNameLen || lenOfName > maxAllianceNameLen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"name":     name,
			}).Warn("alliance:处理仙盟创建,名字不合法")
		playerlogic.SendSystemMessage(pl, lang.AllianceNameIllegal)
		return
	}

	flag := coredirty.GetDirtyService().IsLegal(name)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"name":     name,
			}).Warn("alliance:处理仙盟创建,名字含有脏字")
		playerlogic.SendSystemMessage(pl, lang.AllianceNameDirty)
		return
	}

	// 判断玩家等级
	playerLevel := pl.GetLevel()
	minLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceCreateMinLevel)
	if playerLevel < minLimit {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"allianceName": name,
			}).Warn("alliance:处理仙盟创建,玩家等级不足")

		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	// // 判断玩家阵营
	// if pl.GetCamp() == chuangshitypes.ChuangShiCampTypeNone {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":     pl.GetId(),
	// 			"allianceName": name,
	// 		}).Warn("alliance:处理仙盟创建,玩家没有阵营")
	// 	playerlogic.SendSystemMessage(pl, lang.ChuangShiNotCamp)
	// 	return
	// }

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*inventoryplayer.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	// 老版仙盟和新版低级仙盟消耗物品
	needItem := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceCreateNeedItem)
	needItemCount := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceCreateNeedItemCount)

	// 新版高级仙盟消耗元宝
	needGoldNum := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCreateNewHighAllianceUseGold))

	//判断消耗品
	if versionType == alliancetypes.AllianceVersionTypeOld ||
		(versionType == alliancetypes.AllianceVersionTypeNew && allianceType == alliancetypes.AllianceNewTypeLow) {
		// 老版本、新版本低级仙盟
		flag = inventoryManager.HasEnoughItem(needItem, needItemCount)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"allianceName": name,
				}).Warn("alliance:处理仙盟创建,仙盟令不足")

			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	} else {
		// 新版本高级仙盟
		flag = propertyManager.HasEnoughGold(needGoldNum, false)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"allianceName": name,
				}).Warn("alliance:处理仙盟创建,仙盟令不足")

			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}
	lingYuId := pl.GetLingyuInfo().AdvanceId
	//创建仙盟
	al, err := alliance.GetAllianceService().CreateAlliance(pl.GetId(), name, pl.GetRole(), pl.GetSex(), pl.GetName(), pl.GetVip(), lingYuId, pl.GetLevel(), versionType, allianceType)
	if err != nil {
		return
	}
	if al == nil {
		panic("alliance:创建仙盟应该成功")
	}

	if versionType == alliancetypes.AllianceVersionTypeOld ||
		(versionType == alliancetypes.AllianceVersionTypeNew && allianceType == alliancetypes.AllianceNewTypeLow) {
		//老版仙盟、新版低级仙盟
		itemReason := commonlog.InventoryLogReasonAllianceCreate
		flag = inventoryManager.UseItem(needItem, needItemCount, itemReason, itemReason.String())
		if !flag {
			panic(fmt.Errorf("alliance: create alliance use item should be ok"))
		}

	} else {
		// 新版高级仙盟
		goldReason := commonlog.GoldLogReasonAllianceNewCreateCost
		flag = propertyManager.CostGold(needGoldNum, false, goldReason, goldReason.String())
		if !flag {
			panic(fmt.Errorf("alliance: create alliance cost gold should be ok"))
		}

		//同步属性
		propertylogic.SnapChangedProperty(pl)
	}

	// 公告
	plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	allianceName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(name))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceCreateNotice), plName, allianceName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	// 初始化仙盟仓库物品
	itemId1 := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDepotInitItemId1)
	itemId2 := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDepotInitItemId2)
	itemId3 := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDepotInitItemId3)
	itemId4 := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDepotInitItemId4)

	itemIdList := []int32{itemId1, itemId2, itemId3, itemId4}
	for _, itemId := range itemIdList {
		itemTemp := item.GetItemService().GetItem(int(itemId))
		base := inventorytypes.CreateDefaultItemPropertyDataBase()
		propertyData := inventory.CreatePropertyDataInterface(itemTemp.GetItemType(), base)
		if itemTemp.IsGoldEquip() {
			data := propertyData.(*goldequiptypes.GoldEquipPropertyData)
			attrList := itemTemp.GetGoldEquipTemplate().RandomGoldEquipAttr()
			data.AttrList = attrList
			data.IsHadCountAttr = true
		}
		itemData := droptemplate.CreateItemData(itemId, 1, 0, itemtypes.ItemBindTypeUnBind)
		_, err := alliance.GetAllianceService().SaveInDepot(al.GetAllianceId(), itemData, propertyData)
		if err != nil {
			return err
		}
	}

	// 同步物品
	inventorylogic.SnapInventoryChanged(pl)

	//发送创建成功
	scAllianceCreate := pbutil.BuildSCAllianceCreate(al)
	pl.SendMsg(scAllianceCreate)

	// memList := al.GetDouShenList()
	// scAllianceDouShenMemberList := pbutil.BuildSCAllianceDouShenMemberList(memList)
	// pl.SendMsg(scAllianceDouShenMemberList)

	return
}
