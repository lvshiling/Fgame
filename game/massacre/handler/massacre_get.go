package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/massacre/pbutil"
	playerbshield "fgame/fgame/game/massacre/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MASSACRE_GET_TYPE), dispatch.HandlerFunc(handleMassacreGet))
}

//处理戮仙刃信息
func handleMassacreGet(s session.Session, msg interface{}) (err error) {
	log.Debug("massacre:处理获取戮仙刃消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = massacreGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("massacre:处理获取戮仙刃消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("massacre:处理获取戮仙刃消息完成")
	return nil
}

//获取戮仙刃信息
func massacreGet(pl player.Player) (err error) {
	massacreManager := pl.GetPlayerDataManager(playertypes.PlayerMassacreDataManagerType).(*playerbshield.PlayerMassacreDataManager)
	massacreInfo := massacreManager.GetMassacreInfo()
	scMassacreGet := pbutil.BuildSCMassacreGet(massacreInfo)
	pl.SendMsg(scMassacreGet)
	return
}
