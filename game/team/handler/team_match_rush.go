package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	playerteam "fgame/fgame/game/team/player"
	"fgame/fgame/game/team/team"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_MATCH_RUSH_START_TYPE), dispatch.HandlerFunc(handleTeamMatchRushStart))
}

//处理队员催处开始
func handleTeamMatchRushStart(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理队员催处开始消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = teamMatchRushStart(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("team:处理队员催处开始消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理队员催处开始消息完成")
	return nil

}

//队员催处开始信息的逻辑
func teamMatchRushStart(pl player.Player) (err error) {
	teamId := pl.GetTeamId()
	if teamId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("team:队伍不存在")
		playerlogic.SendSystemMessage(pl, lang.TeamNoExist)
		return
	}
	mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	flag := mananger.IsRushTimeInCd()
	if flag {
		leftTimeMs := mananger.RushTimeLeftTime()
		leftTime := timeutils.MillisecondToSecondCeil(leftTimeMs)
		leftTimeStr := fmt.Sprintf("%d", leftTime)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("team:催处开始,cd")
		playerlogic.SendSystemMessage(pl, lang.MiscExitKaSiCd, leftTimeStr)
		return
	}
	err = team.GetTeamService().TeamMatchRush(pl)
	if err != nil {
		return
	}
	return
}
