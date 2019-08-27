package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_ARENAPVP_PLAYER_DATA_CHANGED_TYPE), dispatch.HandlerFunc(handleArenapvpPlayerDataChanged))
}

//竞技场玩家数据变化
func handleArenapvpPlayerDataChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:竞技场玩家数据变化")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	siMsg := msg.(*crosspb.SIPlayerArenapvpDataChanged)

	if siMsg.PlayerArenapvpData.ReliveTimes != nil {
		//修改玩家竞技场数据
		tpl.SetArenapvpReliveTimes(siMsg.PlayerArenapvpData.GetReliveTimes())
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("arenapvp:竞技场玩家数据变化,完成")
	return nil
}
