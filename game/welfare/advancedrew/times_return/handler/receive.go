package handler

import (
	"fgame/fgame/common/lang"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	advancedrewtimesreturntypes "fgame/fgame/game/welfare/advancedrew/times_return/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeTimesReturn, welfare.ReceiveHandlerFunc(handlerTimesReturnReceive))
}

//升阶次数返还领取奖励请求逻辑
func handlerTimesReturnReceive(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypeTimesReturn
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取升阶返利奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, openTemp.Group)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(openTemp.Group)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  openTemp.Group,
			}).Warn("welfare:领取升阶返利奖励请求，活动不存在")
		return
	}

	info := obj.GetActivityData().(*advancedrewtimesreturntypes.AdvancedTimesReturnInfo)
	advancedType := welfaretypes.AdvancedType(openTemp.Value1)
	if advancedType != info.RewType {
		log.WithFields(
			log.Fields{
				"playerId":            pl.GetId(),
				"rewId":               rewId,
				"curAdvancedType":     info.RewType,
				"receiveAdvancedType": advancedType,
			}).Warn("welfare:领取升阶返利奖励请求,升级返利类型错误")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityCycleTypeNotEqual)
		return
	}

	//领取条件
	needTimes := openTemp.Value2
	if !info.IsCanReceiveRewards(needTimes) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"rewId":     rewId,
				"curDanNum": info.Times,
				"record":    info.RewRecord,
			}).Warn("welfare:领取升阶返利奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	info.AddRecord(needTimes)
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	// //公告
	// plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	// itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(rewItemMap)
	// if len(itemNameLinkStr) > 0 {
	// 	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(openTemp.Group)
	// 	args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
	// 	link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
	// 	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MergeActivityAdvancedRewardsNotice), plName, itemNameLinkStr, link)
	// 	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	// 	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	// }

	scMsg := pbutil.BuildSCOpenActivityReceiveRew(rewId, openTemp.Group, totalRewData, rewItemMap, info.RewRecord)
	pl.SendMsg(scMsg)

	return
}
