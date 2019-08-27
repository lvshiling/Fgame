package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	hallupleveltypes "fgame/fgame/game/welfare/hall/uplevel/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_WELFARE_UPLEVEL_RECEIVE_REW_TYPE), dispatch.HandlerFunc(handlerWelfareUplevelReceive))
}

//处理福利领奖
func handlerWelfareUplevelReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理领取升级福利奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityWelfareUplevelReceiveRew := msg.(*uipb.CSOpenActivityWelfareUplevelReceiveRew)
	rewId := csOpenActivityWelfareUplevelReceiveRew.GetRewId()

	err = welfareUplevelReceiveRew(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理领取升级福利奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理领取升级福利奖励请求完成")

	return
}

//领取福利请求逻辑
func welfareUplevelReceiveRew(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeUpLevel
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取升级福利奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, openTemp.Group)
	if !checkFlag {
		return
	}

	//领取条件
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	rewLevel := openTemp.Value1
	if pl.GetLevel() < rewLevel {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"rewLevel":      rewLevel,
				"curZhuanSheng": propertyManager.GetZhuanSheng(),
			}).Warn("welfare:领取升级福利奖励请求，等级不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	needZhuanSheng := openTemp.Value2
	curZhuanShen := propertyManager.GetZhuanSheng()
	if curZhuanShen < needZhuanSheng {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"needZhuanSheng": needZhuanSheng,
				"curZhuanSheng":  propertyManager.GetZhuanSheng(),
			}).Warn("welfare:领取升级福利奖励请求，转生数不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerZhuanShengTooLow)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, openTemp.Group)
	info := obj.GetActivityData().(*hallupleveltypes.WelfareUplevelInfo)
	if info.IsReceive(rewLevel) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewLevel": rewLevel,
			}).Warn("welfare:领取升级福利奖励请求，已领取过奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新
	info.AddRecord(rewLevel)
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scWelfareUplevelRew := pbutil.BuildSCOpenActivityWelfareUplevelReceiveRew(totalRewData, rewItemMap)
	pl.SendMsg(scWelfareUplevelRew)
	return
}
