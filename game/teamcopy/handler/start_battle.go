package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/team"
	"fgame/fgame/game/teamcopy/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAMCOPY_START_BATTLE_TYPE), dispatch.HandlerFunc(handleTeamStartBattle))
}

//处理开始战斗信息
func handleTeamStartBattle(s session.Session, msg interface{}) (err error) {
	log.Debug("teamcopy:处理开始战斗信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = teamStartBattle(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("teamcopy:处理开始战斗信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("teamcopy:处理开始战斗信息完成")
	return nil
}

//处理组队副本界面信息逻辑
func teamStartBattle(pl player.Player) (err error) {
	teamId := pl.GetTeamId()
	if teamId == 0 {
		return
	}
	err = team.GetTeamService().TeamCopyStartBattle(pl)
	if err != nil {
		return
	}
	scTeamCopyStartBattle := pbutil.BuildSCTeamCopyStartBattle()
	pl.SendMsg(scTeamCopyStartBattle)
	return
}
