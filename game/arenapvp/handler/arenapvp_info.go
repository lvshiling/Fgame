package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arenapvp/arenapvp"
	"fgame/fgame/game/arenapvp/pbutil"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENAPVP_INFO_TYPE), dispatch.HandlerFunc(handleArenapvpInfo))
}

//处理pvp信息
func handleArenapvpInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:处理pvp信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenapvpInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arenapvp:处理pvp信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arenapvp:处理pvp信息,完成")
	return nil
}

func arenapvpInfo(pl player.Player) (err error) {
	baZhuInfo := arenapvp.GetArenapvpService().GetLastAreanapvpBaZhu()
	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	obj := arenapvpManager.GetPlayerArenapvpObj()
	scMsg := pbutil.BuildSCArenapvpInfo(baZhuInfo, obj)
	pl.SendMsg(scMsg)
	return
}
