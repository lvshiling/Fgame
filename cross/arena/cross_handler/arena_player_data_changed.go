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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_ARENA_DATA_CHANGED_TYPE), dispatch.HandlerFunc(handleArenaPlayerDataChanged))
}

//竞技场玩家数据变化
func handleArenaPlayerDataChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:竞技场玩家数据变化")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	siPlayerArenaDataChanged := msg.(*crosspb.SIPlayerArenaDataChanged)

	if siPlayerArenaDataChanged.PlayerArenaData.ReliveTime != nil {
		//修改玩家竞技场数据
		tpl.SetArenaReliveTime(siPlayerArenaDataChanged.PlayerArenaData.GetReliveTime())
	}
	if siPlayerArenaDataChanged.PlayerArenaData.WinTime != nil {
		tpl.SetArenaWinTime(siPlayerArenaDataChanged.PlayerArenaData.GetWinTime())
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("arena:竞技场玩家数据变化,完成")
	return nil

}
