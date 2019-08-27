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
	"fgame/fgame/game/week/pbutil"
	playerweek "fgame/fgame/game/week/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WEEK_INFO_TYPE), dispatch.HandlerFunc(handlerWeekGetInfo))
}

//处理获取周卡信息
func handlerWeekGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("week:处理获取周卡请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = getWeekInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("week:处理获取周卡请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("week:处理获取周卡请求完成")

	return
}

//获取周卡请求逻辑
func getWeekInfo(pl player.Player) (err error) {
	weekManager := pl.GetPlayerDataManager(playertypes.PlayerWeekDataManagerType).(*playerweek.PlayerWeekManager)
	weekInfoMap := weekManager.GetWeekInfoMap()
	scWeekInfo := pbutil.BuildSCWeekInfo(weekInfoMap)
	pl.SendMsg(scWeekInfo)
	return
}
