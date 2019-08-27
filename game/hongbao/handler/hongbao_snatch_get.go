package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/hongbao/pbutil"
	playerhongbao "fgame/fgame/game/hongbao/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HONGBAO_SNATCH_GET_TYPE), dispatch.HandlerFunc(handleHongBaoSnatchGet))

}

//抢红包查询
func handleHongBaoSnatchGet(s session.Session, msg interface{}) (err error) {
	log.Debug("hongbao:抢红包查询")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = hongBaoSnatchGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("hongbao:抢红包查询,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("hongbao:抢红包查询完成")
	return nil

}

//红包查询逻辑
func hongBaoSnatchGet(pl player.Player) (err error) {
	// if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeHongBao) {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":  pl.GetId(),
	// 		}).Warn("hongbao:红包查询，功能未开启")
	// 	playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
	// 	return
	// }

	manager := pl.GetPlayerDataManager(playertypes.PlayerHongBaoDataManagerType).(*playerhongbao.PlayerHongBaoDataManager)
	snatchCount := manager.GetSnatchCount()
	scMsg := pbutil.BuildSCHongBaoSnatchGet(snatchCount)
	pl.SendMsg(scMsg)
	return
}
