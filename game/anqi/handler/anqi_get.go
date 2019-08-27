package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/anqi/pbutil"
	playeranqi "fgame/fgame/game/anqi/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ANQI_GET_TYPE), dispatch.HandlerFunc(handleAnqiGet))
}

//处理暗器信息
func handleAnqiGet(s session.Session, msg interface{}) (err error) {
	log.Debug("anqi:处理获取暗器消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = anqiGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("anqi:处理获取暗器消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("anqi:处理获取暗器消息完成")
	return nil
}

//获取暗器信息
func anqiGet(pl player.Player) (err error) {
	anqiManager := pl.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	anqiInfo := anqiManager.GetAnqiInfo()
	scAnqiGet := pbutil.BuildSCAnqiGet(anqiInfo)
	pl.SendMsg(scAnqiGet)
	return
}
