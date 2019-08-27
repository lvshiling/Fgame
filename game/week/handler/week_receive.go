package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/week/pbutil"
	playerweek "fgame/fgame/game/week/player"
	weektemplate "fgame/fgame/game/week/template"
	weektypes "fgame/fgame/game/week/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WEEK_RECEIVE_REW_TYPE), dispatch.HandlerFunc(handlerWeekReceiveRew))
}

//处理领取周卡奖励
func handlerWeekReceiveRew(s session.Session, msg interface{}) (err error) {
	log.Debug("week:处理领取周卡奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSWeekReceiveRew)
	typ := csMsg.GetWeekType()

	weekType := weektypes.WeekType(typ)
	if !weekType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weekType": weekType,
			}).Warn("week:领取周卡奖励请求，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = weekReceive(tpl, weekType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("week:处理领取周卡奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("week:处理购买Week请求完成")

	return
}

//领取周卡奖励请求逻辑
func weekReceive(pl player.Player, weekType weektypes.WeekType) (err error) {
	weekTemp := weektemplate.GetWeekTemplateService().GetWeekTemplate(weekType)
	if weekTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weekType": weekType,
			}).Warn("week:领取周卡奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	weekManager := pl.GetPlayerDataManager(playertypes.PlayerWeekDataManagerType).(*playerweek.PlayerWeekManager)
	weekInfo := weekManager.GetWeekInfo(weekType)
	if weekInfo == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weekType": weekType,
			}).Warn("week:领取周卡奖励请求，未购买周卡")
		playerlogic.SendSystemMessage(pl, lang.WeekNotBuyWeek)
		return
	}

	//是否周卡
	now := global.GetGame().GetTimeService().Now()
	if !weekInfo.IsWeek(now) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weekType": weekType,
			}).Warn("week:领取周卡奖励请求，不是周卡")
		playerlogic.SendSystemMessage(pl, lang.WeekNotBuyWeek)
		return
	}

	dayInt := weekInfo.GetNextCycleDay()
	weekDayTemp := weekTemp.GetCycDayRew(dayInt)
	if weekDayTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weekType": weekType,
				"dayInt":   dayInt,
			}).Warn("week:领取周卡奖励请求，每日奖励模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//是否领取
	if weekInfo.IsReceiveRewards(now) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weekType": weekType,
			}).Warn("week:领取周卡奖励请求，已领取今日奖励")
		playerlogic.SendSystemMessage(pl, lang.WeekHadReceiveRewards)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//背包空间
	totalItemMap := make(map[int32]int32)
	if weekTemp.IsExtralRew(dayInt) {
		totalItemMap = coreutils.MergeMap(totalItemMap, weekTemp.GetExtralRewItemMap())
	}
	totalItemMap = coreutils.MergeMap(totalItemMap, weekDayTemp.GetRewItemMap())

	if !inventoryManager.HasEnoughSlots(totalItemMap) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"totalItemMap": totalItemMap,
			}).Warn("week:领取周卡奖励请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//增加物品
	itemGetReason := commonlog.InventoryLogReasonWeedDayRew
	itemGetReasonText := fmt.Sprintf(itemGetReason.String(), weekType.String())
	flag := inventoryManager.BatchAdd(totalItemMap, itemGetReason, itemGetReasonText)
	if !flag {
		panic("week:week rewards add item should be ok")
	}

	rewSilver := weekDayTemp.RewSilver
	rewBindGold := weekDayTemp.RewBindGold
	rewGold := weekDayTemp.RewGold
	reasonGold := commonlog.GoldLogReasonWeekDayRew
	reasonSilver := commonlog.SilverLogReasonWeekDayRew
	reasonGoldText := fmt.Sprintf(reasonGold.String(), weekType.String())
	reasonSilverText := fmt.Sprintf(reasonSilver.String(), weekType.String())
	flag = propertyManager.AddMoney(int64(rewBindGold), int64(rewGold), reasonGold, reasonGoldText, int64(rewSilver), reasonSilver, reasonSilverText)
	if !flag {
		panic("week:week rewards add RewData should be ok")
	}

	//更新
	flag = weekManager.ReceiveWeekRewards(weekType)
	if !flag {
		panic("week:week rewards add RewData should be ok")
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	rd := propertytypes.CreateRewData(0, 0, rewSilver, rewGold, rewBindGold)
	scMsg := pbutil.BuildSCWeekReceiveRew(weekType, weekInfo, rd, totalItemMap)
	pl.SendMsg(scMsg)
	return
}
