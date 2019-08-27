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
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	playerteam "fgame/fgame/game/team/player"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_INVITE_ALL_TYPE), dispatch.HandlerFunc(handleTeamNearInviteAll))
}

//处理一键邀请信息
func handleTeamNearInviteAll(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理一键邀请消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = teamNearInviteAll(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("team:处理一键邀请消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理一键邀请消息完成")
	return nil

}

//一键邀请信息的逻辑
func teamNearInviteAll(pl player.Player) (err error) {
	mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	//操作过于频繁
	sucess := mananger.IfCanInviteAllTime()
	if !sucess {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("team:操作过于频繁")
		playerlogic.SendSystemMessage(pl, lang.CommonOperFrequent)
		return
	}

	playerList := teamlogic.GetNearPlayers(pl)
	err = team.GetTeamService().TeamInviteAll(pl, playerList)
	if err != nil {
		return
	}

	//更新时间
	mananger.TeamInviteAllTime()
	scTeamInviteAll := pbutil.BuildSCTeamInviteAll()
	pl.SendMsg(scTeamInviteAll)
	return
}
