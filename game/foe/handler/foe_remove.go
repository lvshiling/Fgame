package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/foe/pbutil"
	playerfoe "fgame/fgame/game/foe/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FOE_REMOVE_TYPE), dispatch.HandlerFunc(handleFoeDelete))
}

//处理移除仇人
func handleFoeDelete(s session.Session, msg interface{}) error {
	log.Debug("foe:处理仇人移除")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFoeRemove := msg.(*uipb.CSFoeRemove)
	foeId := csFoeRemove.GetFoeId()
	err := foeRemove(tpl, foeId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"foeId":    foeId,
				"error":    err,
			}).Error("foe:处理仇人移除,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("foe:处理仇人移除,完成")
	return nil

}

//处理仇人移除
func foeRemove(pl player.Player, foeId int64) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerFoeDataManagerType).(*playerfoe.PlayerFoeDataManager)
	flag := manager.IsFoe(foeId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"foeId":    foeId,
			}).Warn("foe:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager.RemoveFoe(foeId)
	scFoeRemove := pbutil.BuildSCFoeRemove(foeId)
	pl.SendMsg(scFoeRemove)
	return
}
