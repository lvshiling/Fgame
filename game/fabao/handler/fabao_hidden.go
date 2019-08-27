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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FABAO_HIDDEN_TYPE), dispatch.HandlerFunc(handleFaBaoHidden))
}

//处理法宝隐藏展示信息
func handleFaBaoHidden(s session.Session, msg interface{}) (err error) {
	log.Debug("fabao:处理法宝隐藏展示信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFaBaoHidden := msg.(*uipb.CSFaBaoHidden)
	hiddenFlag := csFaBaoHidden.GetHidden()

	err = faBaoHidden(tpl, hiddenFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
				"error":      err,
			}).Error("fabao:处理法宝隐藏展示信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"hiddenFlag": hiddenFlag,
		}).Debug("fabao:处理法宝隐藏展示信息完成")
	return nil

}

//法宝隐藏展示的逻辑
func faBaoHidden(pl player.Player, hiddenFlag bool) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	manager.Hidden(hiddenFlag)
	scFaBaoHidden := pbutil.BuildSCFaBaoHidden(hiddenFlag)
	pl.SendMsg(scFaBaoHidden)
	return
}
