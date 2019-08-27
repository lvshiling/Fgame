package handler

import (
	"fgame/fgame/common/lang"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	hallzhuanshengtypes "fgame/fgame/game/welfare/hall/zhuansheng/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeZhaunSheng, welfare.ReceiveHandlerFunc(handlerZhuanShengReceive))
}

//处理领取转生冲刺
func handlerZhuanShengReceive(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeZhaunSheng
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取转生冲刺奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	groupId := openTemp.Group

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
			}).Warn("welfare:领取转生冲刺奖励请求，活动不存在")
		return
	}

	//领取条件
	needZhuanSheng := openTemp.Value1
	info := obj.GetActivityData().(*hallzhuanshengtypes.ZhuanShengInfo)
	if !info.IsCanReceiveRewards(needZhuanSheng) {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"rewId":          rewId,
				"needZhuanSheng": needZhuanSheng,
				"curZhuanSheng":  info.ZhuanSheng,
				"record":         info.RewRecord,
			}).Warn("welfare:领取转生冲刺奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	info.AddRecord(needZhuanSheng)
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityReceiveRew(rewId, groupId, totalRewData, rewItemMap, info.RewRecord)
	pl.SendMsg(scMsg)
	return
}
