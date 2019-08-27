package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_REALM_INFO_TYPE), dispatch.HandlerFunc(handlerRealmGetInfo))
}

//处理获取天劫塔冲刺信息
func handlerRealmGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取天劫塔冲刺请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityRealmInfo := msg.(*uipb.CSOpenActivityRealmInfo)
	groupId := csOpenActivityRealmInfo.GetGroupId()

	err = getRealmInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取天劫塔冲刺请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取天劫塔冲刺请求完成")

	return
}

//获取天劫塔冲刺请求逻辑
func getRealmInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeRealm
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
			"groupId":  groupId,
		}).Warn("welfare:参数无效,活动时间模板不存在")
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	realmObj := welfareManager.GetOpenActivity(groupId)
	timesList := welfare.GetWelfareService().GetReceiveTimesList(groupId)
	scOpenActivityRealmInfo := pbutil.BuildSCOpenActivityRealmInfo(realmObj, groupId, timesList, startTime, endTime)
	pl.SendMsg(scOpenActivityRealmInfo)
	return
}
