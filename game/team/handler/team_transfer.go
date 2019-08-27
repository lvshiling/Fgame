package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_TRANSFER_CAPTAIN_TYPE), dispatch.HandlerFunc(handleTeamTransferCaptain))
}

//处理玩家转让队长信息
func handleTeamTransferCaptain(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理玩家转让队长消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamTransferCaptain := msg.(*uipb.CSTeamTransferCaptain)
	captainId := csTeamTransferCaptain.GetCaptainId()

	err = teamTransferCaptain(tpl, captainId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"captainId": captainId,
				"error":     err,
			}).Error("team:处理玩家转让队长消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理玩家转让队长消息完成")
	return nil

}

//玩家转让队长信息的逻辑
func teamTransferCaptain(pl player.Player, captainId int64) (err error) {
	playerId := pl.GetId()
	if playerId == captainId {
		log.WithFields(log.Fields{
			"playerId":  playerId,
			"captainId": captainId,
		}).Warn("team:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	teamId := pl.GetTeamId()
	if teamId == 0 {
		log.WithFields(log.Fields{
			"playerId":  playerId,
			"captainId": captainId,
		}).Warn("team:队伍不存在")
		playerlogic.SendSystemMessage(pl, lang.TeamNoExist)
		return
	}

	_, err = team.GetTeamService().TransferCaptain(pl, captainId)
	if err != nil {
		return
	}
	return
}
