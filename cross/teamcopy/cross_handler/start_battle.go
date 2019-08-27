package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/cross/teamcopy/pbutil"
	"fgame/fgame/cross/teamcopy/teamcopy"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_TEAMCOPY_START_BATTLE_TYPE), dispatch.HandlerFunc(handleTeamCopyStartBattle))
}

//处理组队副本开始战斗
func handleTeamCopyStartBattle(s session.Session, msg interface{}) (err error) {
	log.Debug("teamcopy:处理组队副本开始战斗")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siTeamCopyStartBattle := msg.(*crosspb.SITeamCopyStartBattle)
	playerList := pbutil.ConvertFromTeamPlayerList(siTeamCopyStartBattle.GetPlayerList())

	err = teamCopyStartBattle(tpl, playerList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("teamcopy:处理组队副本开始战斗,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("teamcopy:处理组队副本开始战斗,完成")
	return nil

}

//组队副本开始战斗
func teamCopyStartBattle(pl *player.Player, playerList []*teamcopy.BattleTeamMember) (err error) {
	flag := teamcopy.GetTeamCopyService().TeamCopyStartBattle(pl, playerList)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("teamcopy:处理组队副本开始战斗,匹配失败")
		isTeamCopyStartBattle := pbutil.BuildISTeamCopyStartBattle(flag)
		pl.SendMsg(isTeamCopyStartBattle)
		return
	}

	isTeamCopyStartBattle := pbutil.BuildISTeamCopyStartBattle(flag)
	pl.SendMsg(isTeamCopyStartBattle)
	return
}
