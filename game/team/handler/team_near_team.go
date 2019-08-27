package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_NEAR_GET_TYPE), dispatch.HandlerFunc(handleTeamNearGet))
}

//处理附近队伍信息
func handleTeamNearGet(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理附近队伍消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = teamNearGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("team:处理附近队伍消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理附近队伍消息完成")
	return nil

}

//附近队伍信息的逻辑
func teamNearGet(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	mapId := pl.GetScene().MapId()
	teamList := team.GetTeamService().GetNearTeam(pl)

	scTeamNearGet := pbutil.BuildSCTeamNearGet(mapId, teamList)
	pl.SendMsg(scTeamNearGet)
	return
}
