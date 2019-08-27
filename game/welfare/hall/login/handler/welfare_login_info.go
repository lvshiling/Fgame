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
	welfaretypes "fgame/fgame/game/welfare/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_LOGIN_INFO_TYPE), dispatch.HandlerFunc(handlerLoginGetInfo))
}

//处理获取活动信息
func handlerLoginGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取信息请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityLoginInfo := msg.(*uipb.CSOpenActivityLoginInfo)
	groupId := csOpenActivityLoginInfo.GetGroupId()

	err = getLoginInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取信息请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取信息请求完成")

	return
}

//获取信息请求逻辑
func getLoginInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeLogin
	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	dayNum := welfarelogic.CountWelfareLoginDay(pl.GetCreateTime())
	loginObj := welfareManager.GetOpenActivity(groupId)

	scOpenActivityLoginInfo := pbutil.BuildSCOpenActivityLoginInfo(dayNum, loginObj)
	pl.SendMsg(scOpenActivityLoginInfo)

	return
}
