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
	"fgame/fgame/game/songbuting/pbutil"
	playersongbuting "fgame/fgame/game/songbuting/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SONGBUTING_GET_TYPE), dispatch.HandlerFunc(handleSongBuTingGet))
}

//处理送不停信息
func handleSongBuTingGet(s session.Session, msg interface{}) (err error) {
	log.Debug("songbuting:处理获取送不停消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = songBuTingGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("songbuting:处理获取送不停消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("songbuting:处理获取送不停消息完成")
	return nil
}

//获取送不停信息的逻辑
func songBuTingGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSongBuTingDataManagerType).(*playersongbuting.PlayerSongBuTingDataManager)
	songBuTingObj := manager.GetSongBuTingObj()
	scSongBuTingChanged := pbutil.BuildSCSongBuTingChanged(songBuTingObj)
	pl.SendMsg(scSongBuTingChanged)
	return
}
