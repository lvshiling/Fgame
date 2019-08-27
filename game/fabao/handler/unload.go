package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/fabao/pbutil"
	playerfabao "fgame/fgame/game/fabao/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FABAO_UNLOAD_TYPE), dispatch.HandlerFunc(handleFaBaoUnload))
}

//处理法宝卸下信息
func handleFaBaoUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("fabao:处理法宝卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = faBaoUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("fabao:处理法宝卸下信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fabao:处理法宝卸下信息完成")
	return nil

}

//法宝卸下的逻辑
func faBaoUnload(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	obj := manager.GetFaBaoInfo()
	if obj.GetFaBaoId() == 0 {
		playerlogic.SendSystemMessage(pl, lang.FaBaoUnrealNoExist)
		return
	}
	manager.Unload()
	scFaBaoUnload := pbutil.BuildSCFaBaoUnload(0)
	pl.SendMsg(scFaBaoUnload)
	return
}
