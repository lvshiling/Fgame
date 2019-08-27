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
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_MATCH_CONDTION_FAILED_DEAL_TYPE), dispatch.HandlerFunc(handleTeamMatchCondtionFailedDeal))
}

//处理队伍匹配失败咨询窗决策
func handleTeamMatchCondtionFailedDeal(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理队伍匹配失败咨询窗决策消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamMatchCondtionFailedDeal := msg.(*uipb.CSTeamMatchCondtionFailedDeal)
	result := csTeamMatchCondtionFailedDeal.GetResult()

	err = teamMatchCondtionFailedDeal(tpl, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"result":   result,
				"error":    err,
			}).Error("team:处理队伍匹配失败咨询窗决策消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理队伍匹配失败咨询窗决策消息完成")
	return nil

}

//队伍匹配失败咨询窗决策信息的逻辑
func teamMatchCondtionFailedDeal(pl player.Player, result bool) (err error) {
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
	memberIdList, flag := mananger.GetMatchCondtionFailedList()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
		}).Warn("team:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	defer mananger.ResetMatchCondtionFailedList()
	err = team.GetTeamService().TeamMatchCondtionFailedDeal(pl, result, memberIdList)
	if err != nil {
		return
	}
	return
}
