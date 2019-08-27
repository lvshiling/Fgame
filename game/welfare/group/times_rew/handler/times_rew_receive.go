package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	grouptimesrewtypes "fgame/fgame/game/welfare/group/times_rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_TIMES_REW_RECEIVE_TYPE), dispatch.HandlerFunc(handlerTimesRewReceive))
}

//处理领取次数奖励
func handlerTimesRewReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理次数奖励领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityTimesRewReceive)
	rewId := csMsg.GetRewId()

	err = receiveTimesRew(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理次数奖励领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理次数奖励领取奖励请求完成")

	return
}

//次数奖励领取奖励请求逻辑
func receiveTimesRew(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeGroup
	subType := welfaretypes.OpenActivityGroupSubTypeTimesRew
	timesRewTemp := welfaretemplate.GetWelfareTemplateService().GetTimesRewTemplate(rewId)
	if timesRewTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取次数奖励奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	groupId := timesRewTemp.Group

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取次数奖励奖励请求，活动不存在")
		return
	}

	// vip等级
	if pl.GetVip() < timesRewTemp.VipLevel {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取次数奖励奖励请求，vip等级不足")
		playerlogic.SendSystemMessage(pl, lang.VipLevelToLow)
		return
	}

	//领取条件
	info := obj.GetActivityData().(*grouptimesrewtypes.TimesRewInfo)
	if !info.IsCanReceiveRewards(timesRewTemp.DrawTimes, timesRewTemp.VipLevel) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"rewId":     rewId,
				"needTimes": timesRewTemp.DrawTimes,
			}).Warn("welfare:领取次数奖励奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	//背包空间
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	rewItemMap := timesRewTemp.GetRewItemMap()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	firstTemp := groupInterface.GetFirstOpenTemp()
	newItemDataList := welfarelogic.ConvertToItemData(rewItemMap, firstTemp.GetExpireType(), firstTemp.GetExpireTime())
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemDataList) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"rewItemMap": rewItemMap,
			}).Warn("welfare:领取活动奖励请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//增加物品
	itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
	itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
	flag := inventoryManager.BatchAddOfItemLevel(newItemDataList, itemGetReason, itemGetReasonText)
	if !flag {
		panic("welfare:welfare rewards add item should be ok")
	}

	//更新信息
	info.AddRecord(timesRewTemp.DrawTimes, timesRewTemp.VipLevel)
	welfareManager.UpdateObj(obj)

	//同步资源
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityTimesRewReceive(rewId, rewItemMap, info.RewRecord)
	pl.SendMsg(scMsg)

	return
}
