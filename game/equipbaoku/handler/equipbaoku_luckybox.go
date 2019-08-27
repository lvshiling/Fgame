package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	commontypes "fgame/fgame/game/common/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	equipbaokueventtypes "fgame/fgame/game/equipbaoku/event/types"
	"fgame/fgame/game/equipbaoku/pbutil"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	equipbaokutemplate "fgame/fgame/game/equipbaoku/template"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_LUCKY_BOX_TYPE), dispatch.HandlerFunc(handleEquipBaoKuLuckyBox))

}

//装备宝库幸运宝箱
func handleEquipBaoKuLuckyBox(s session.Session, msg interface{}) (err error) {
	log.Debug("equipbaoku:装备宝库幸运宝箱")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csEquipbaokuLuckyBox := msg.(*uipb.CSEquipbaokuLuckyBox)
	typ := equipbaokutypes.BaoKuType(csEquipbaokuLuckyBox.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Warn("equipbaoku:处理获取宝库日志请求,宝库类型不合法")
		return
	}

	err = equipBaoKuLuckyBox(tpl, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("equipbaoku:处理幸运宝箱,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("equipbaoku:处理幸运宝箱完成")
	return nil

}

//装备宝库幸运宝箱逻辑
func equipBaoKuLuckyBox(pl player.Player, typ equipbaokutypes.BaoKuType) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeEquipBaoKu) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("baokujifen:幸运宝箱错误,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}
	equipBaoKuManager := pl.GetPlayerDataManager(playertypes.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
	equipBaoKuTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetEquipBaoKuByLevAndZhuanNum(pl.GetLevel(), pl.GetZhuanSheng(), typ)
	if equipBaoKuTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:幸运宝箱错误,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//幸运值是否足够
	equipBaoKuObj := equipBaoKuManager.GetEquipBaoKuObj(typ)
	luckyPoints := equipBaoKuObj.GetLuckyPoints()
	num := int32(luckyPoints / equipBaoKuTemplate.NeedXingYunZhi)
	needLuckyPoint := num * equipBaoKuTemplate.NeedXingYunZhi
	if num == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:幸运宝箱错误,参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//掉落
	var rewList []*droptemplate.DropItemData
	for num > 0 {
		dropData := droptemplate.GetDropTemplateService().GetDropBaoKuItemLevel(equipBaoKuTemplate.ScriptXingYun)
		if dropData != nil {
			rewList = append(rewList, dropData)
		}
		num -= 1
	}

	var rewItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(rewList) > 0 {
		rewItemList, resMap = droplogic.SeperateItemDatas(rewList)
	}
	//背包空间
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlotsOfItemLevel(rewItemList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:幸运宝箱错误,背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗
	equipBaoKuManager.SubEquipBaoKuLuckyPoints(needLuckyPoint, typ)
	//宝库幸运值日志
	luckyReason := commonlog.EquipBaoKuLogReasonLuckyPointsChange
	luckyReasonText := fmt.Sprintf(luckyReason.String(), typ.GetBaoKuName(), commontypes.ChangeTypeUse.String())
	luckyData := equipbaokueventtypes.CreatePlayerEquipBaoKuLuckyPointsLogEventData(luckyPoints, equipBaoKuObj.GetLuckyPoints(), rewList, luckyReason, luckyReasonText, typ)
	gameevent.Emit(equipbaokueventtypes.EventTypeEquipBaoKuLuckyPointsLog, pl, luckyData)
	//增加掉落
	if len(resMap) > 0 {
		goldReason := commonlog.GoldLogReasonEquipBaoKuLuckyBoxGet
		silverReason := commonlog.SilverLogReasonEquipBaoKuLuckyBoxGet
		levelReason := commonlog.LevelLogReasonEquipBaoKuLuckyBoxGet
		err = droplogic.AddRes(pl, resMap, goldReason, goldReason.String(), silverReason, silverReason.String(), levelReason, levelReason.String())
		if err != nil {
			return
		}
	}

	if len(rewItemList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonEquipBaoKuLuckyBoxGet
		flag := inventoryManager.BatchAddOfItemLevel(rewItemList, itemGetReason, itemGetReason.String())
		if !flag {
			panic("equipbaoku:增加物品应该成功")
		}
	}

	for _, itemData := range rewList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		//稀有道具公告
		inventorylogic.PrecioustemBroadcast(pl, itemId, num, lang.InventoryEquipBaoKuLuckyBoxItemNotice)
	}

	//同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	newLuckyPoints := equipBaoKuManager.GetEquipBaoKuObj(typ).GetLuckyPoints()
	scEquipBaoKuLuckyBox := pbutil.BuildSCEquipBaoKuLuckyBox(rewList, newLuckyPoints, int32(typ))
	pl.SendMsg(scEquipBaoKuLuckyBox)
	return
}
