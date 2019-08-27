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

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_MATCH_CONDTION_PREPARE_DEAL_TYPE), dispatch.HandlerFunc(handleTeamMatchCondtionPrepareDeal))
}

//处理队员准备决策决策
func handleTeamMatchCondtionPrepareDeal(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理队员准备决策决策消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamMatchCondtionPrepareDeal := msg.(*uipb.CSTeamMatchCondtionPrepareDeal)
	result := csTeamMatchCondtionPrepareDeal.GetResult()

	err = teamMatchCondtionPrepareDeal(tpl, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"result":   result,
				"error":    err,
			}).Error("team:处理队员准备决策决策消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理队员准备决策决策消息完成")
	return nil

}

//队员准备决策决策信息的逻辑
func teamMatchCondtionPrepareDeal(pl player.Player, result bool) (err error) {
	teamId := pl.GetTeamId()
	if teamId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
		}).Warn("team:队伍不存在")
		playerlogic.SendSystemMessage(pl, lang.TeamNoExist)
		return
	}
	mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	flag := mananger.IsExistMatchCondtionFailed()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
		}).Warn("team:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	defer mananger.ResetMatchCondtionFailed()
	err = team.GetTeamService().TeamMatchCondtionPrepareDeal(pl, result)
	if err != nil {
		return
	}

	return
}
