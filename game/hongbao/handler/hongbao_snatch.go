package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	hongbaoeventtypes "fgame/fgame/game/hongbao/event/types"
	"fgame/fgame/game/hongbao/hongbao"
	"fgame/fgame/game/hongbao/pbutil"
	playerhongbao "fgame/fgame/game/hongbao/player"
	hongbaotemplate "fgame/fgame/game/hongbao/template"
	hongbaotypes "fgame/fgame/game/hongbao/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HONGBAO_SNATCH_TYPE), dispatch.HandlerFunc(handleHongBaoSnatch))

}

//抢红包
func handleHongBaoSnatch(s session.Session, msg interface{}) (err error) {
	log.Debug("hongbao:抢红包")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csHongbaoSnatch := msg.(*uipb.CSHongbaoSnatch)
	hongBaoId := csHongbaoSnatch.GetHongBaoId()
	channelTypeInt := csHongbaoSnatch.GetChannel()

	channelType := chattypes.ChannelType(channelTypeInt)
	if !channelType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"channelType": int32(channelType),
			}).Warn("hongbao:抢红包,频道参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = hongBaoSnatch(tpl, hongBaoId, channelType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"hongBaoId": hongBaoId,
				"error":     err,
			}).Error("hongbao:抢红包,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("hongbao:抢红包完成")
	return nil

}

//抢红包逻辑
func hongBaoSnatch(pl player.Player, hongBaoId int64, channelType chattypes.ChannelType) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeHongBao) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"hongBaoId": hongBaoId,
			}).Warn("hongbao:抢红包，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	hongBaoService := hongbao.GetHongBaoService()
	hongBaoObj := hongBaoService.GetHongBaoObj(hongBaoId)
	if hongBaoObj == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"hongBaoId": hongBaoId,
			}).Warn("hongbao:抢红包，红包id错误")
		playerlogic.SendSystemMessage(pl, lang.HongBaoExpire)
		return
	}

	curTemplate := hongbaotemplate.GetHongBaoTemplateService().GetHongBaoByTemplateType(hongBaoObj.GetHongBaoType())
	if curTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"hongBaoId": hongBaoId,
			}).Warn("hongbao:抢红包, 模板错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//抢红包条件
	if !curTemplate.CheckNeedCondition(pl.GetVip(), pl.GetZhuanSheng()) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"hongBaoId": hongBaoId,
			}).Warn("hongbao:抢红包, 不满足抢红包条件")
		playerlogic.SendSystemMessage(pl, lang.HongBaoSnatchNeedConditionLimit)
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerHongBaoDataManagerType).(*playerhongbao.PlayerHongBaoDataManager)
	if manager.IsSnatchCountReachLimit() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"hongBaoId": hongBaoId,
			}).Warn("hongbao:抢红包, 每日抢红包次数不足")
		playerlogic.SendSystemMessage(pl, lang.HongBaoSnatchCountReachedLimit)
		return
	}

	if hongBaoObj.GetHongBaoType() == itemtypes.ItemHongBaoSubTypeZhenXi {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		if inventoryManager.GetEmptySlots(inventorytypes.BagTypePrim) == 0 {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"hongBaoId": hongBaoId,
				}).Warn("hongbao:抢红包, 背包空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
	}

	keepTime := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeHongBaoKeepTime))
	result, nowCount := hongBaoService.SnatchHongBao(hongBaoObj, pl, keepTime)
	snatchCount := manager.GetSnatchCount()
	if result == hongbaotypes.HongBaoResultTypeSucceed {
		snatchCount = manager.AddSnatchCount()
		//给奖励
		awardList := hongBaoObj.GetAwardList()
		awardInfo := awardList[nowCount]

		goldReason := commonlog.GoldLogReasonHongBaoSnatch
		silverReason := commonlog.SilverLogReasonHongBaoSnatch
		itemGetReason := commonlog.InventoryLogReasonHongBaoSnatch
		levelReason := commonlog.LevelLogReasonHongBaoSnatch
		goldReasonText := goldReason.String()
		silverReasonText := silverReason.String()
		itemGetReasonText := itemGetReason.String()
		levelReasonText := levelReason.String()
		dropData := droptemplate.CreateItemData(awardInfo.ItemId, awardInfo.ItemCnt, awardInfo.Level, itemtypes.ItemBindTypeBind)
		flag, err := droplogic.AddItem(pl, dropData, goldReason, goldReasonText, silverReason, silverReasonText, itemGetReason, itemGetReasonText, levelReason, levelReasonText)
		if err != nil {
			return err
		}
		if !flag {
			panic("hongbao:add award should be success")
		}
		//同步
		propertylogic.SnapChangedProperty(pl)
		inventorylogic.SnapInventoryChanged(pl)
	}

	//返回前端
	endTime := hongBaoObj.GetCreateTime() + keepTime
	scMsg := pbutil.BuildSCHongBaoSnatch(hongBaoObj, endTime, int32(result), snatchCount)
	pl.SendMsg(scMsg)
	//消息事件
	gameevent.Emit(hongbaoeventtypes.EventTypeHongBaoThanks, pl, channelType)

	return
}
