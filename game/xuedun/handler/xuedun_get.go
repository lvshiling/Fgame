package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/xuedun/pbutil"
	playerxuedun "fgame/fgame/game/xuedun/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XUEDUN_GET_TYPE), dispatch.HandlerFunc(handleXueDunGet))
}

//处理血盾信息
func handleXueDunGet(s session.Session, msg interface{}) (err error) {
	log.Debug("xuedun:处理获取血盾消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = xueDunGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("xuedun:处理获取血盾消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xuedun:处理获取血盾消息完成")
	return nil

}

//获取血盾界面信息的逻辑
func xueDunGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	xueDunInfo := manager.GetXueDunInfo()
	scXueDunGet := pbutil.BuildSCXueDunGet(xueDunInfo)
	pl.SendMsg(scXueDunGet)
	return
}
