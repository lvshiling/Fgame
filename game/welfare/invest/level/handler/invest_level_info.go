package handler

//已返还购买价格的废弃活动
/*
import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_INVEST_LEVEL_INFO_TYPE), dispatch.HandlerFunc(handlerLevelGetInfo))
}

//处理获取活动信息
func handlerLevelGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取信息请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityInvestLevelInfo := msg.(*uipb.CSOpenActivityInvestLevelInfo)
	groupId := csOpenActivityInvestLevelInfo.GetGroupId()

	err = getInvestLevelInfo(tpl, groupId)
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
func getInvestLevelInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)

	scMsg := pbutil.BuildSCOpenActivityInvestLevelInfo(obj)
	pl.SendMsg(scMsg)

	return
}

*/
