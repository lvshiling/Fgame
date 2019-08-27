package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANTI_UNLOAD_TYPE), dispatch.HandlerFunc(handleXianTiUnload))
}

//处理仙体卸下信息
func handleXianTiUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("xianti:处理仙体卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = xianTiUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("xianti:处理仙体卸下信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xianti:处理仙体卸下信息完成")
	return nil

}

//仙体卸下的逻辑
func xianTiUnload(pl player.Player) (err error) {
	xianTiManager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	obj := xianTiManager.GetXianTiInfo()
	if obj.XianTiId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xianti:处理仙体卸下,没有仙体")
		playerlogic.SendSystemMessage(pl, lang.XianTiUnrealNoExist)
		return
	}

	xianTiManager.Unload()
	scXianTiUnload := pbutil.BuildSCXianTiUnload(0)
	pl.SendMsg(scXianTiUnload)
	return
}
