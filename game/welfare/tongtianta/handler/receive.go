package handler

import (
	"fgame/fgame/common/lang"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	tongtiantatemplate "fgame/fgame/game/welfare/tongtianta/template"
	tongtiantatypes "fgame/fgame/game/welfare/tongtianta/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeLingTong, welfare.ReceiveHandlerFunc(tongTianTaReceiveReward))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeMingGe, welfare.ReceiveHandlerFunc(tongTianTaReceiveReward))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeTuLong, welfare.ReceiveHandlerFunc(tongTianTaReceiveReward))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeShengHen, welfare.ReceiveHandlerFunc(tongTianTaReceiveReward))

	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeZhenFa, welfare.ReceiveHandlerFunc(tongTianTaReceiveReward))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeBaby, welfare.ReceiveHandlerFunc(tongTianTaReceiveReward))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeDianXing, welfare.ReceiveHandlerFunc(tongTianTaReceiveReward))
}

func tongTianTaReceiveReward(pl player.Player, rewId int32) (err error) {
	playerId := pl.GetId()

	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取活动目标奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	groupId := openTemp.Group
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:处理运营活动信息请求，时间模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	typ := timeTemp.GetOpenType()
	subType := timeTemp.GetOpenSubType()

	force := openTemp.Value1
	goldNum := openTemp.Value2

	// 校验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		return
	}
	info := obj.GetActivityData().(*tongtiantatypes.TongTianTaInfo)

	// 判断充值金额是否足够
	if !info.IsEnoughChargeNum(goldNum) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"groupId":  groupId,
				"force":    force,
			}).Warn("welfare: 没有资格领取该奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityTongTianTaNotReachGoal, fmt.Sprintf("%d", goldNum))
		return
	}

	groupTempI := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupTempI == nil {
		return
	}
	groupTemp := groupTempI.(*tongtiantatemplate.GroupTemplateTongTianTa)
	nearForceTemp := groupTemp.GetTongTianTaTemplateByNearForce(info.MinForce)
	nearForce := int32(0)
	if nearForceTemp != nil {
		nearForce = nearForceTemp.Value1
	}
	// 判断是否有资格领取奖励
	if force < nearForce || force > info.MaxForce {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"groupId":  groupId,
				"force":    force,
			}).Warn("welfare: 没有资格领取该奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityTongTianTaNotReceive)
		return
	}

	// 判断是否已经领取该奖励
	if info.IsAlreadyReceiveByForce(force) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"groupId":  groupId,
				"force":    force,
			}).Warn("welfare: 已经领取该奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityTongTianTaAlreadyReceive)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	// 推送改变
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	// 刷新数据
	info.ReceiveSuccess(force)
	welfareManager.UpdateObj(obj)

	scOpenActivityReceiveRew := pbutil.BuildSCOpenActivityReceiveRew(rewId, openTemp.Group, totalRewData, rewItemMap, info.Record)
	pl.SendMsg(scOpenActivityReceiveRew)

	return
}
