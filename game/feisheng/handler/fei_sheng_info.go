package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/feisheng/pbutil"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_INFO_TYPE), dispatch.HandlerFunc(handleFeiShengInfo))
}

//处理飞升信息
func handleFeiShengInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理飞升信息消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = feiShengInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("feisheng:处理飞升信息消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("feisheng:处理飞升信息消息完成")
	return nil

}

//飞升信息界面逻辑
func feiShengInfo(pl player.Player) (err error) {

	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiShengInfo := feiManager.GetFeiShengInfo()
	scMsg := pbutil.BuildSCFeiShengInfo(feiShengInfo)
	pl.SendMsg(scMsg)
	return
}
