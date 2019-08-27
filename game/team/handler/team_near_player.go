package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_NEAR_PLAYER_GET_TYPE), dispatch.HandlerFunc(handleTeamNearPlayerGet))
}

//处理玩家信息
func handleTeamNearPlayerGet(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理玩家消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = teamNearPlayerGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("team:处理玩家消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理玩家消息完成")
	return nil

}

//玩家信息的逻辑
func teamNearPlayerGet(pl player.Player) (err error) {
	nearPlayerList := teamlogic.GetNearPlayers(pl)
	scTeamNearPlayerGet := pbutil.BuildSCTeamNearPlayerGet(nearPlayerList)
	pl.SendMsg(scTeamNearPlayerGet)
	return
}
