package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	commonlog "fgame/fgame/common/log"
// 	coredirty "fgame/fgame/core/dirty"
// 	"fgame/fgame/core/session"
// 	coreutils "fgame/fgame/core/utils"
// 	"fgame/fgame/game/alliance/alliance"
// 	"fgame/fgame/game/alliance/pbutil"
// 	chatlogic "fgame/fgame/game/chat/logic"
// 	chattypes "fgame/fgame/game/chat/types"
// 	"fgame/fgame/game/constant/constant"
// 	constanttypes "fgame/fgame/game/constant/types"
// 	droptemplate "fgame/fgame/game/drop/template"
// 	goldequiptypes "fgame/fgame/game/goldequip/types"
// 	"fgame/fgame/game/inventory/inventory"
// 	inventorylogic "fgame/fgame/game/inventory/logic"
// 	inventoryplayer "fgame/fgame/game/inventory/player"
// 	inventorytypes "fgame/fgame/game/inventory/types"
// 	"fgame/fgame/game/item/item"
// 	itemtypes "fgame/fgame/game/item/types"
// 	noticelogic "fgame/fgame/game/notice/logic"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	"fgame/fgame/game/player/types"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"
// 	"fmt"
// 	"strings"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_CREATE_TYPE), dispatch.HandlerFunc(handleAllianceCreate))
// }

// //处理仙盟创建
// func handleAllianceCreate(s session.Session, msg interface{}) (err error) {
// 	log.Debug("alliance:处理仙盟创建")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	csAllianceCreate := msg.(*uipb.CSAllianceCreate)
// 	name := csAllianceCreate.GetName()

// 	err = allianceCreate(tpl, name)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"name":     name,
// 				"error":    err,
// 			}).Error("alliance:处理仙盟创建,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 			"name":     name,
// 		}).Debug("alliance:处理仙盟创建,完成")
// 	return nil

// }

// const (
// 	minAllianceNameLen = 1
// 	maxAllianceNameLen = 6
// )

// //仙盟创建
// func allianceCreate(pl player.Player, name string) (err error) {
// 	name = strings.TrimSpace(name)
// 	lenOfName := len([]rune(name))

// 	if lenOfName < minAllianceNameLen && lenOfName > maxAllianceNameLen {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"name":     name,
// 			}).Warn("alliance:处理仙盟创建,名字不合法")
// 		playerlogic.SendSystemMessage(pl, lang.AllianceNameIllegal)
// 		return
// 	}

// 	flag := coredirty.GetDirtyService().IsLegal(name)
// 	if !flag {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"name":     name,
// 			}).Warn("alliance:处理仙盟创建,名字含有脏字")
// 		playerlogic.SendSystemMessage(pl, lang.AllianceNameDirty)
// 		return
// 	}

// 	//判断等级
// 	playerLevel := pl.GetLevel()
// 	minLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceCreateMinLevel)
// 	if playerLevel < minLimit {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":     pl.GetId(),
// 				"allianceName": name,
// 			}).Warn("alliance:处理仙盟创建,玩家等级不足")

// 		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
// 		return
// 	}

// 	//判断消耗品
// 	needItem := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceCreateNeedItem)
// 	needItemCount := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceCreateNeedItemCount)
// 	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*inventoryplayer.PlayerInventoryDataManager)
// 	flag = inventoryManager.HasEnoughItem(needItem, needItemCount)
// 	if !flag {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":     pl.GetId(),
// 				"allianceName": name,
// 			}).Warn("alliance:处理仙盟创建,仙盟令不足")

// 		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 		return
// 	}

// 	//创建仙盟
// 	al, err := alliance.GetAllianceService().CreateAlliance(pl.GetId(), name, pl.GetRole(), pl.GetSex(), pl.GetName(), pl.GetVip())
// 	if err != nil {
// 		return
// 	}
// 	if al == nil {
// 		panic("alliance:创建仙盟应该成功")
// 	}

// 	alliance.GetAllianceService().SyncMemberInfo(pl.GetId(), pl.GetName(), pl.GetSex(), pl.GetLevel(), pl.GetForce(), pl.GetZhuanSheng(), pl.GetLingyuInfo().AdvanceId, pl.GetVip())

// 	//消耗道具
// 	itemReason := commonlog.InventoryLogReasonAllianceCreate
// 	flag = inventoryManager.UseItem(needItem, needItemCount, itemReason, itemReason.String())
// 	if !flag {
// 		panic(fmt.Errorf("alliance: create alliance use item should be ok"))
// 	}

// 	//公告
// 	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
// 	allianceName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(name))
// 	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceCreateNotice), playerName, allianceName)
// 	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
// 	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

// 	//仙盟仓库初始化物品
// 	itemId1 := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDepotInitItemId1)
// 	itemId2 := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDepotInitItemId2)
// 	itemId3 := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDepotInitItemId3)
// 	itemId4 := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDepotInitItemId4)
// 	itemIdList := []int32{itemId1, itemId2, itemId3, itemId4}
// 	for _, itemId := range itemIdList {
// 		itemTemp := item.GetItemService().GetItem(int(itemId))
// 		base := inventorytypes.CreateDefaultItemPropertyDataBase()
// 		propertyData := inventory.CreatePropertyDataInterface(itemTemp.GetItemType(), base)
// 		if itemTemp.IsGoldEquip() {
// 			data := propertyData.(*goldequiptypes.GoldEquipPropertyData)
// 			attrList := itemTemp.GetGoldEquipTemplate().RandomGoldEquipAttr()
// 			data.AttrList = attrList
// 			data.IsHadCountAttr = true
// 		}
// 		itemData := droptemplate.CreateItemData(itemId, 1, 0, itemtypes.ItemBindTypeUnBind)
// 		_, err := alliance.GetAllianceService().SaveInDepot(al.GetAllianceId(), itemData, propertyData)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	//同步道具
// 	inventorylogic.SnapInventoryChanged(pl)

// 	//发送创建成功
// 	scAllianceCreate := pbutil.BuildSCAllianceCreate(al)
// 	pl.SendMsg(scAllianceCreate)
// 	return
// }
