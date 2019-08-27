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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_TEAMCOPY_BATTLE_RESULT_TYPE), dispatch.HandlerFunc(handleTeamCopyBattleResult))
}

//处理组队副本战斗结果
func handleTeamCopyBattleResult(s session.Session, msg interface{}) (err error) {
	log.Debug("teamcopy:处理组队副本战斗结果")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = teamCopyBattleResult(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("teamcopy:处理组队副本战斗结果,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("teamcopy:处理组队副本战斗结果,完成")
	return nil

}

//组队副本战斗结果
func teamCopyBattleResult(pl *player.Player) (err error) {
	return
}
