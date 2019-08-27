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
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	xiuxianbooklogic "fgame/fgame/game/welfare/xiuxianbook/logic"
	xiuxianbooktypes "fgame/fgame/game/welfare/xiuxianbook/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipStrength, welfare.ReceiveHandlerFunc(handleXiuxianBookReceive))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipOpenLight, welfare.ReceiveHandlerFunc(handleXiuxianBookReceive))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipUpStar, welfare.ReceiveHandlerFunc(handleXiuxianBookReceive))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeLingTong, welfare.ReceiveHandlerFunc(handleXiuxianBookReceive))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeDianXing, welfare.ReceiveHandlerFunc(handleXiuxianBookReceive))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeShenQi, welfare.ReceiveHandlerFunc(handleXiuxianBookReceive))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillXinFa, welfare.ReceiveHandlerFunc(handleXiuxianBookReceive))
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillDiHun, welfare.ReceiveHandlerFunc(handleXiuxianBookReceive))
}

func handleXiuxianBookReceive(pl player.Player, rewId int32) (err error) {
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:修仙典籍领取请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	groupId := openTemp.Group
	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:修仙典籍领取请求，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	//刷新一下数据先
	welfareManager.RefreshActivityDataByGroupId(groupId)
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
	// 校验活动
	checkFlag := welfarelogic.CheckGroupId(pl, timeTemp.GetOpenType(), timeTemp.GetOpenSubType(), groupId)
	if !checkFlag {
		return
	}

	obj := welfareManager.GetOpenActivityIfNotCreate(timeTemp.GetOpenType(), timeTemp.GetOpenSubType(), groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Error("welfare:修仙典籍领取请求，活动对象不存在")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityObjectNotExist)
		return
	}
	info := obj.GetActivityData().(*xiuxianbooktypes.XiuxianBookInfo)

	can, curLevel, needLevel := xiuxianbooklogic.IsCanReceiceRew(obj, openTemp)
	if !can {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"groupId":   groupId,
				"needLevel": needLevel,
				"curLevel":  curLevel,
			}).Warn("welfare:修仙典籍领取请求，条件不满足")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//增加记录
	info.AddReceiveRecord(needLevel)
	welfareManager.UpdateObj(obj)

	//领取奖励
	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityReceiveRew(rewId, groupId, totalRewData, rewItemMap, info.HasReceiveRecord)
	pl.SendMsg(scMsg)
	return
}
