package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	systemlingyutemplate "fgame/fgame/game/welfare/system/lingyu/template"
	systemlingyutypes "fgame/fgame/game/welfare/system/lingyu/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_SYSTEM_ACTIVATE_TYPE), dispatch.HandlerFunc(handlerSystemActivate))
}

//处理系统激活
func handlerSystemActivate(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理系统激活请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivitySystemActivate)
	groupId := csMsg.GetGroupId()

	err = systemActivate(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理系统激活请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理系统激活请求完成")

	return
}

// 若有新的激活系统，做成注册类型放到common_handler
//系统激活请求逻辑
func systemActivate(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeSystemActivate
	subType := welfaretypes.OpenActivitySystemActivateSubTypeLingYu

	//检验活动
	// checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	// if !checkFlag {
	// 	return
	// }

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:系统激活请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:系统激活请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*systemlingyutypes.SystemLingYuInfo)
	if info.IsActivate {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:系统激活请求，系统已激活")
		playerlogic.SendSystemMessage(pl, lang.OpenActivitySystemHadActivate)
		return
	}

	groupTemp := groupInterface.(*systemlingyutemplate.GroupTemplateSystemLingYu)
	if info.MaxSingleChargeGold < groupTemp.GetActivateCondition() {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"groupId":       groupId,
				"curMaxCharge":  info.MaxSingleChargeGold,
				"needMaxCharge": groupTemp.GetActivateCondition(),
			}).Warn("welfare:系统激活请求，充值条件不满足")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityChargeNotEnoughCondition)
		return
	}

	now := global.GetGame().GetTimeService().Now()
	tempMap := groupTemp.GetOpenTempMap()
	if len(groupTemp.GetOpenTempMap()) != 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("welfare:领域活动激活模板应该只有一条")
		return
	}
	for _, openTemp := range tempMap {
		continuedTime := int64(openTemp.Value2) * int64(common.DAY)
		if info.StartTime+continuedTime < now {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ,
					"subType":  subType,
					"groupId":  groupId,
				}).Warn("welfare:运营活动,不是该玩家领域活动时间")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityArgumentInvalid, fmt.Sprintf("%d", groupId))
			return
		}
	}

	//系统激活
	info.IsActivate = true
	welfareManager.UpdateObj(obj)
	gameevent.Emit(welfareeventtypes.EventTypeSystemActivate, pl, subType)

	scMsg := pbutil.BuildSCOpenActivitySystemActivate(groupId)
	pl.SendMsg(scMsg)
	return
}
