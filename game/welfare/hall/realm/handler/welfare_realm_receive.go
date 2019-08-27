package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	hallrealmtypes "fgame/fgame/game/welfare/hall/realm/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_REALM_REWARDS_TYPE), dispatch.HandlerFunc(handlerRealmReceive))
}

//处理天劫塔冲刺
func handlerRealmReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理天劫塔冲刺领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityRealmRewards := msg.(*uipb.CSOpenActivityRealmRewards)
	rewId := csOpenActivityRealmRewards.GetRewId()

	err = receiveRealm(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理天劫塔冲刺领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理天劫塔冲刺领取奖励请求完成")

	return
}

//天劫塔冲刺领取奖励请求逻辑
func receiveRealm(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeRealm
	addTimes := int32(1)
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:天劫塔冲刺奖励请求，模板不存在")
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
			}).Warn("welfare:天劫塔冲刺奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*hallrealmtypes.WelfareRealmChallengeInfo)
	needLevel := openTemp.Value1

	//领取条件
	if !info.IsCanReceiveRewards(needLevel) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"needLevel": needLevel,
				"rewId":     rewId,
			}).Warn("welfare:天劫塔冲刺奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	// 总次数限制
	timesMax := openTemp.Value2
	rewLevel := openTemp.Value1
	if !welfare.GetWelfareService().IsHadReceiveTimes(openTemp.Group, rewLevel, timesMax, addTimes) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"needLevel": needLevel,
				"rewId":     rewId,
			}).Warn("welfare:天劫塔冲刺奖励请求，全服次数已领完")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityGlobalNotTimesReceive)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	info.AddRecord(needLevel)
	welfareManager.UpdateObj(obj)
	welfare.GetWelfareService().AddReceiveTimes(openTemp.Group, needLevel, addTimes)

	//公告
	itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(rewItemMap)
	if len(itemNameLinkStr) > 0 {
		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(openTemp.Group)
		plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		realmLevel := coreutils.FormatColor(chattypes.ColorTypeModuleName, fmt.Sprintf("天劫塔%d层", needLevel))
		args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
		link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityWelfareRealmRewardsNotice), plName, realmLevel, itemNameLinkStr, link)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	timesList := welfare.GetWelfareService().GetReceiveTimesList(openTemp.Group)
	scOpenActivityRealmRewards := pbutil.BuildSCOpenActivityRealmRewards(totalRewData, timesList, rewItemMap)
	pl.SendMsg(scOpenActivityRealmRewards)

	return
}
