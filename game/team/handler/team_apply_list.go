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
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_APPLY_GET_TYPE), dispatch.HandlerFunc(handleTeamApplyListGet))
}

//处理获取申请列表信息
func handleTeamApplyListGet(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理获取申请列表消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = teamApplyListGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("team:处理获取申请列表消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理获取申请列表消息完成")
	return nil
}

//获取申请列表信息的逻辑
func teamApplyListGet(pl player.Player) (err error) {

	teamId := pl.GetTeamId()
	if teamId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("team:队伍不存在")
		playerlogic.SendSystemMessage(pl, lang.TeamNoExist)
		return
	}
	applyList, err := team.GetTeamService().GetApplyList(pl)
	if err != nil {
		return
	}
	scTeamApplyGet := pbutil.BuildSCTeamApplyGet(applyList)
	pl.SendMsg(scTeamApplyGet)
	return
}
