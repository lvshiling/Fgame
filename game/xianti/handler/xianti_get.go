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
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANTI_GET_TYPE), dispatch.HandlerFunc(handleXianTiGet))
}

//处理仙体信息
func handleXianTiGet(s session.Session, msg interface{}) (err error) {
	log.Debug("xianti:处理获取仙体消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = xianTiGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("xianti:处理获取仙体消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xianti:处理获取仙体消息完成")
	return nil

}

//获取仙体界面信息的逻辑
func xianTiGet(pl player.Player) (err error) {
	xianTiManager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xianTiInfo := xianTiManager.GetXianTiInfo()
	xianTiOtherMap := xianTiManager.GetXianTiOtherMap()
	scXianTiGet := pbutil.BuildSCXianTiGet(xianTiInfo, xianTiOtherMap)
	pl.SendMsg(scXianTiGet)
	return
}
