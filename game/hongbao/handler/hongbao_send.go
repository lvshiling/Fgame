package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	hongbaoeventtypes "fgame/fgame/game/hongbao/event/types"
	hongbaologic "fgame/fgame/game/hongbao/logic"
	"fgame/fgame/game/hongbao/pbutil"
	hongbaotemplate "fgame/fgame/game/hongbao/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HONGBAO_SEND_TYPE), dispatch.HandlerFunc(handleHongBaoSend))

}

//红包发送
func handleHongBaoSend(s session.Session, msg interface{}) (err error) {
	log.Debug("hongbao:红包发送")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csHongbaoSend := msg.(*uipb.CSHongbaoSend)
	hongBaoType := itemtypes.ItemHongBaoSubType(csHongbaoSend.GetHongBaoType())
	channelType := chattypes.ChannelType(csHongbaoSend.GetChannel())
	cliArgs := csHongbaoSend.GetArgs()
	countMax := csHongbaoSend.GetCountMax()

	if !hongBaoType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"hongBaoType": int32(hongBaoType),
			}).Warn("hongbao:红包发送,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	if !channelType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"channelType": int32(channelType),
			}).Warn("hongbao:红包发送,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = hongbaoSend(tpl, hongBaoType, channelType, cliArgs, countMax)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"hongBaoType": int32(hongBaoType),
				"error":       err,
			}).Error("hongbao:红包发送,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("hongbao:红包发送完成")
	return nil

}

//红包发送逻辑
func hongbaoSend(pl player.Player, hongBaoType itemtypes.ItemHongBaoSubType, channelType chattypes.ChannelType, cliArgs string, countMax int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeHongBao) {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"hongBaoType": int32(hongBaoType),
				"channelType": int32(channelType),
			}).Warn("hongbao:红包发送，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	curTemplate := hongbaotemplate.GetHongBaoTemplateService().GetHongBaoByTemplateType(hongBaoType)
	if curTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"hongBaoType": int32(hongBaoType),
				"channelType": int32(channelType),
			}).Warn("hongbao:红包发送,模板错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if countMax > curTemplate.CountMax || countMax < curTemplate.CountMin {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"hongBaoType": int32(hongBaoType),
				"channelType": int32(channelType),
				"countMax":    countMax,
			}).Warn("hongbao:红包发送,领取人数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if channelType != chattypes.ChannelTypeWorld && channelType != chattypes.ChannelTypeBangPai {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"hongBaoType": int32(hongBaoType),
				"channelType": int32(channelType),
			}).Warn("hongbao:红包发送,模板错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if channelType == chattypes.ChannelTypeBangPai && pl.GetAllianceId() == 0 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"hongBaoType": int32(hongBaoType),
				"channelType": int32(channelType),
			}).Warn("hongbao:红包发送,不在帮派")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//判断物品
	needItem := curTemplate.ItemId
	needNum := int32(1)
	flag := inventoryManager.HasEnoughItem(needItem, needNum)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":    pl.GetId(),
			"hongBaoType": int32(hongBaoType),
		}).Warn("hongbao:红包发送,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//消耗物品
	reasonText := commonlog.InventoryLogReasonHongBaoSend.String()
	flag = inventoryManager.UseItem(needItem, needNum, commonlog.InventoryLogReasonHongBaoSend, reasonText)
	if !flag {
		panic(fmt.Errorf("hongbao: hongbaosend use item should be ok"))
	}
	inventorylogic.SnapInventoryChanged(pl)

	//塞红包
	hongBaoId, err := hongbaologic.HongBaoPlugAward(pl, hongBaoType, countMax)
	if err != nil {
		return
	}

	//消息事件
	data := hongbaoeventtypes.CreateHongBaoSendEventData(hongBaoId, hongBaoType, channelType, cliArgs)
	gameevent.Emit(hongbaoeventtypes.EventTypeHongBaoSend, pl, data)
	//返回前端
	scMsg := pbutil.BuildSCHongBaoSend(hongBaoId)
	pl.SendMsg(scMsg)
	return
}
