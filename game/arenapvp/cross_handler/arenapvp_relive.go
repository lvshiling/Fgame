package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
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
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_RELIVE_TYPE), dispatch.HandlerFunc(handleArenapvpRelive))
}

//处理复活
func handleArenapvpRelive(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:处理跨服pvp复活")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = arenapvpRelive(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arenapvp:处理跨服pvp复活,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arenapvp:处理跨服pvp复活,完成")
	return nil

}

//pvp复活
func arenapvpRelive(pl player.Player) (err error) {
	//竞技场获胜
	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpManager.Relive()
	siMsg := pbutil.BuildSIArenapvpRelive(true)
	pl.SendCrossMsg(siMsg)
	return
}
