package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/guaji/pbutil"
	playerguaji "fgame/fgame/game/guaji/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_STOP_GUA_JI_TYPE), dispatch.HandlerFunc(handleStopGuaJi))
}

//处理挂机
func handleStopGuaJi(s session.Session, msg interface{}) (err error) {
	log.Info("guaji:处理停止挂机")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = stopGuaJi(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),

				"error": err,
			}).Error("guaji:处理停止挂机,错误")

		return err
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("guaji:处理停止挂机,完成")
	return nil
}

//吞噬
func stopGuaJi(pl player.Player) (err error) {
	guaJiManager := pl.GetPlayerDataManager(playertypes.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)
	guaJiManager.StopGuaJiList()
	scStopGuaJi := pbutil.BuildSCStopGuaJi()
	pl.SendMsg(scStopGuaJi)
	return
}
