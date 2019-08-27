package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	ringlogic "fgame/fgame/game/ring/logic"
	"fgame/fgame/game/ring/pbutil"
	playerring "fgame/fgame/game/ring/player"
	ringtemplate "fgame/fgame/game/ring/template"
	ringtypes "fgame/fgame/game/ring/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RING_JINGLING_TYPE), dispatch.HandlerFunc(handleRingJingLing))
}

func handleRingJingLing(s session.Session, msg interface{}) (err error) {
	log.Debug("ring: 开始处理特戒净灵请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRingJingLing := msg.(*uipb.CSRingJingLing)
	typ := ringtypes.RingType(csRingJingLing.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"type":     int32(typ),
			}).Warn("ring: 特戒类型不合法")
		return
	}

	err = ringJingLing(tpl, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("ring: 处理特戒净灵请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("ring: 处理特戒净灵请求消息,成功")

	return
}

func ringJingLing(pl player.Player, typ ringtypes.RingType) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeRingJingLing) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("ring: 功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)
	ringObj := ringManager.GetPlayerRingObject(typ)
	if ringObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("ring: 玩家未穿戴该特戒")
		playerlogic.SendSystemMessage(pl, lang.RingNotEquip)
		return
	}

	itemId := ringObj.GetItemId()
	data := ringObj.GetPropertyData()
	ringData := data.(*ringtypes.RingPropertyData)
	level := ringData.JingLingLevel + 1

	jingLingTemp := ringtemplate.GetRingTemplateService().GetRingJingLingTemplate(itemId, level)
	if jingLingTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("ring: 特戒净灵已达到最高")
		playerlogic.SendSystemMessage(pl, lang.RingJingLingAlreadyTop)
		return
	}

	needItemMap := jingLingTemp.GetNeedItemMap()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughItems(needItemMap) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("ring: 所需物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	reason := commonlog.InventoryLogReasonRingJingLing
	reasonText := fmt.Sprintf(reason.String(), typ.String())
	flag := inventoryManager.BatchRemove(needItemMap, reason, reasonText)
	if !flag {
		panic("ring: 消耗物品应该成功")
	}

	// 物品改变推送
	inventorylogic.SnapInventoryChanged(pl)

	success := ringlogic.RingJingLing(ringData.JingLingNum, jingLingTemp)
	if !success {
		level--
	}

	// 刷新数据
	flag = ringManager.RingJingLingSuccess(typ, success)
	if !flag {
		panic("ring: 特戒净灵刷新数据应该成功")
	}

	// 推送属性变化
	ringlogic.RingPropertyChange(pl)
	propertylogic.SnapChangedProperty(pl)

	scRingJingLing := pbutil.BuildSCRingJingLing(success, int32(typ), level)
	pl.SendMsg(scRingJingLing)

	return
}
