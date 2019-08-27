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
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENFA_GET_TYPE), dispatch.HandlerFunc(handleShenfaGet))

}

//处理身法信息
func handleShenfaGet(s session.Session, msg interface{}) (err error) {
	log.Debug("Shenfa:处理获取身法消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = shenfaGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("Shenfa:处理获取身法消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("Shenfa:处理获取身法消息完成")
	return nil

}

//获取身法界面信息逻辑
func shenfaGet(pl player.Player) (err error) {
	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenfaInfo := shenfaManager.GetShenfaInfo()
	shenFaOtherMap := shenfaManager.GetShenfaOtherMap()
	scShenfaGet := pbutil.BuildSCShenfaGet(shenfaInfo, shenFaOtherMap)
	pl.SendMsg(scShenfaGet)
	return
}
