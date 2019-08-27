package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	welfarelogic "fgame/fgame/game/welfare/logic"
	madeexptemplate "fgame/fgame/game/welfare/made/exp/template"
	madeexptypes "fgame/fgame/game/welfare/made/exp/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_MADE_RES_TYPE), dispatch.HandlerFunc(handlerMadeRes))
}

//处理炼制资源
func handlerMadeRes(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理炼制资源请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityMadeRes)
	groupId := csMsg.GetGroupId()

	err = madeRes(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理炼制资源请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理炼制资源请求完成")

	return
}

//炼制资源请求逻辑
func madeRes(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeMade
	subType := welfaretypes.OpenActivityMadeSubTypeResource

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	madeTemp := welfaretemplate.GetWelfareTemplateService().GetMadeTemplate(groupId, pl.GetLevel())
	if madeTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"level":    pl.GetLevel(),
			}).Warn("welfare:炼制资源请求，炼制模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:炼制资源请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	err = welfareManager.RefreshActivityDataByGroupId(groupId)
	if err != nil {
		return
	}
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*madeexptypes.MadeInfo)

	// 次数
	if madeTemp.IsFullTimes(info.Times) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"curTimes": info.Times,
				"level":    pl.GetLevel(),
			}).Warn("welfare:炼制资源请求，炼制次数不足")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityMadeResTimesNotEnough)
		return
	}

	// 等级
	groupTemp := groupInterface.(*madeexptemplate.GroupTemplateMadeExp)
	if pl.GetLevel() < groupTemp.GetMadeLevelLimit() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"level":    pl.GetLevel(),
			}).Warn("welfare:炼制资源请求，等级不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//元宝是否足够
	needGold := int64(madeTemp.GetNeedCost(info.Times))
	if needGold < 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"needGold": needGold,
			}).Warn("welfare:炼制资源请求，次数不足")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(needGold, false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("welfare:炼制资源请求，当前元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗元宝
	if needGold > 0 {
		goldReason := commonlog.GoldLogReasonActivityCost
		goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
		flag := propertyManager.CostGold(needGold, false, goldReason, goldReasonText)
		if !flag {
			panic("welfare: made resource use gold should be ok")
		}
	}

	exp := int64(madeTemp.Exp)
	expPoint := int64(madeTemp.ExpPoint)
	levelReason := commonlog.LevelLogReasonOpenActivityMadeRes
	if expPoint > 0 {
		propertyManager.AddExpPoint(expPoint, levelReason, levelReason.String())
	}
	if exp > 0 {
		propertyManager.AddExp(exp, levelReason, levelReason.String())
	}

	info.Times += 1
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCOpenActivityMadeRes(groupId, info.Times)
	pl.SendMsg(scMsg)
	return
}
