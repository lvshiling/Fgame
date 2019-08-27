package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arenapvp/pbutil"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENAPVP_GUESS_NOTICE_SETTING_TYPE), dispatch.HandlerFunc(handleArenapvpGuessSetting))
}

//处理竞猜提醒设置
func handleArenapvpGuessSetting(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:处理竞猜提醒设置")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSArenapvpGuessNoticeSetting)
	notice := csMsg.GetNotice()

	err = arenapvpGuessSetting(tpl, notice)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arenapvp:处理竞猜提醒设置,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arenapvp:处理竞猜提醒设置,完成")
	return nil
}

func arenapvpGuessSetting(pl player.Player, notice int32) (err error) {
	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpManager.GuessNoticeSetting(notice)

	scMsg := pbutil.BuildSCArenapvpGuessNoticeSetting(notice)
	pl.SendMsg(scMsg)
	return
}
