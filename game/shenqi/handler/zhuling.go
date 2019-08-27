package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shenqieventtypes "fgame/fgame/game/shenqi/event/types"
	shenqilogic "fgame/fgame/game/shenqi/logic"
	"fgame/fgame/game/shenqi/pbutil"
	playershenqi "fgame/fgame/game/shenqi/player"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENQI_ZHULING_TYPE), dispatch.HandlerFunc(handleShenQiZhuLing))
}

//处理注灵
func handleShenQiZhuLing(s session.Session, msg interface{}) (err error) {
	log.Debug("zhuling:处理装备槽注灵")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csShenqiZhuling := msg.(*uipb.CSShenqiZhuling)
	typInt := csShenqiZhuling.GetShenQiType()
	subTypeInt := csShenqiZhuling.GetSubType()
	slotIdInt := csShenqiZhuling.GetSlotId()
	auto := csShenqiZhuling.GetAuto()
	typ := shenqitypes.ShenQiType(typInt)
	subType := shenqitypes.QiLingType(subTypeInt)

	//参数不对
	if !typ.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typInt":   typInt,
			}).Warn("zhuling:器灵类型,错误")
		return
	}
	if !subType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"subTypeInt": subTypeInt,
			}).Warn("zhuling:器灵类型,错误")
		return
	}
	slotId := shenqitypes.CreateQiLingSubType(subType, slotIdInt)
	if !slotId.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"subTypeInt": subTypeInt,
				"slotIdInt":  slotIdInt,
			}).Warn("zhuling:器灵类型,错误")
		return
	}

	err = shenQiZhuLing(tpl, typ, subType, slotId, auto)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("zhuling:处理装备槽注灵,错误")

		return err
	}
	log.Debug("zhuling:处理装备槽注灵,完成")
	return nil
}

//注灵
func shenQiZhuLing(pl player.Player, typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, slotId shenqitypes.QiLingSubType, auto bool) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShenQiZhuLing) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("zhuling:处理装备槽注灵,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	shenQiManager := pl.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	nextTemp := shenQiManager.GetNextZhuLingTemplate(typ, subType, slotId)
	if nextTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"subType":  int32(subType),
				"slotId":   slotId.String(),
			}).Warn("zhuling:处理装备槽注灵,已经最高级")
		playerlogic.SendSystemMessage(pl, lang.ShenQiSlotLevelMax)
		return
	}

	slotObj := shenQiManager.GetShenQiQiLingMapByArg(typ, subType, slotId)
	if slotObj.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"subType":  int32(subType),
				"slotId":   slotId.String(),
			}).Warn("zhuling:处理装备槽注灵,没有器灵")
		playerlogic.SendSystemMessage(pl, lang.ShenQiNotInlayQiLing)
		return
	}

	itemTemplate := item.GetItemService().GetItem(int(slotObj.ItemId))
	if itemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"subType":  int32(subType),
				"slotId":   slotId.String(),
			}).Warn("zhuling:处理装备槽注灵,物品模板错误")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	if int32(itemTemplate.GetQualityType()) < slotObj.Level {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"subType":  int32(subType),
				"slotId":   slotId.String(),
			}).Warn("zhuling:处理装备槽注灵,器灵品质不足")
		playerlogic.SendSystemMessage(pl, lang.ShenQiQiLingQualityNoEnough)
		return
	}

	needLingQiNum := int64(nextTemp.NeedZhuLing)
	shenQiObj := shenQiManager.GetShenQiOjb()
	if needLingQiNum > shenQiObj.LingQiNum {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"subType":  int32(subType),
				"slotId":   slotId.String(),
			}).Warn("zhuling:处理装备槽注灵,灵气值不足")
		playerlogic.SendSystemMessage(pl, lang.ShenQiLingQiNotEnough)
		return
	}

	flag := shenQiManager.SubLingQiNum(needLingQiNum)
	if !flag {
		panic(fmt.Errorf("zhuling:消耗灵气值应该成功"))
	}

	useItemTemp := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeLingQi)
	eventdata := shenqieventtypes.CreatePlayerShenQiUseItemEventData(int32(useItemTemp.Id), int32(needLingQiNum))
	gameevent.Emit(shenqieventtypes.EventTypeShenQiUseItem, pl, eventdata)

	//进阶判断
	sucess, pro, _, addTimes, _ := shenqilogic.ZhuLingJudge(pl, slotObj.UpNum, slotObj.UpPro, nextTemp)
	befLev := slotObj.Level
	shenQiManager.ZhuLingAdvanced(typ, subType, slotId, pro, addTimes, sucess)
	if sucess {
		//同步改变
		shenqilogic.ShenQiPropertyChanged(pl)
		//日志
		logReason := commonlog.ShenQiLogReasonRelatedUpLevel
		reasonText := fmt.Sprintf(logReason.String(), typ.String(), slotId.String(), commontypes.AdvancedTypeLingQi.String())
		logData := shenqieventtypes.CreatePlayerShenQiRelatedUpLevelLogEventData(befLev, slotObj.Level, logReason, reasonText)
		gameevent.Emit(shenqieventtypes.EventTypeShenQiRelatedUpLevelLog, pl, logData)
	}

	//注灵成功
	scMsg := pbutil.BuildSCShenQiZhuling(slotObj, shenQiObj.LingQiNum, auto)
	pl.SendMsg(scMsg)
	return
}
