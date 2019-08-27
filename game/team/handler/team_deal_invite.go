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
	teamtypes "fgame/fgame/game/team/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_INVITE_RESULT_TYPE), dispatch.HandlerFunc(handleTeamInviteResult))
}

//处理被邀请玩家决策信息
func handleTeamInviteResult(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理获取被邀请玩家决策消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamInviteResult := msg.(*uipb.CSTeamInviteResult)
	typ := teamtypes.TeamInviteType(csTeamInviteResult.GetTyp())
	result := csTeamInviteResult.GetResult()
	id := csTeamInviteResult.GetId()

	err = teamInviteResult(tpl, typ, teamtypes.TeamResultType(result), id)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"result":   result,
				"id":       id,
				"error":    err,
			}).Error("team:处理获取被邀请玩家决策消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理获取被邀请玩家决策消息完成")
	return nil

}

//获取被邀请玩家决策信息的逻辑
func teamInviteResult(pl player.Player, typ teamtypes.TeamInviteType, result teamtypes.TeamResultType, id int64) (err error) {
	if !result.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"result":   result,
			"id":       id,
		}).Warn("team:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if pl.GetTeamId() != 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
			"result":   result,
			"id":       id,
		}).Warn("team:玩家已经在队伍中")
		playerlogic.SendSystemMessage(pl, lang.TeamPlayerInTeam)
		return
	}

	mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	flag := mananger.IfExistInvite(typ, id)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
			"result":   result,
			"id":       id,
		}).Warn("team:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	//删除邀请处理
	defer mananger.TeamInvitedDeal(typ, id)
	_, err = team.GetTeamService().InvitedPlayerChoose(pl, typ, result, id)
	if err != nil {
		return
	}
	return
}
