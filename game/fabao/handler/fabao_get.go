package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/fabao/pbutil"
	playerfabao "fgame/fgame/game/fabao/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FABAO_GET_TYPE), dispatch.HandlerFunc(handleFaBaoGet))

}

//处理法宝信息
func handleFaBaoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("fabao:处理获取法宝消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = faBaoGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("fabao:处理获取法宝消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fabao:处理获取法宝消息完成")
	return nil

}

//获取法宝界面信息逻辑
func faBaoGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	faBaoInfo := manager.GetFaBaoInfo()
	faBaoOtherMap := manager.GetFaBaoOtherMap()
	scFaBaoGet := pbutil.BuildSCFaBaoGet(faBaoInfo, faBaoOtherMap)
	pl.SendMsg(scFaBaoGet)
	return
}
