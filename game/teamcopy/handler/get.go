package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/teamcopy/pbutil"
	playerteamcopy "fgame/fgame/game/teamcopy/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAMCOPY_ALL_GET_TYPE), dispatch.HandlerFunc(handleTeamCopyAllGet))
}

//处理获取组队副本信息
func handleTeamCopyAllGet(s session.Session, msg interface{}) (err error) {
	log.Debug("teamcopy:处理获取组队副本信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = teamCopyAllGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("teamcopy:处理获取组队副本信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("teamcopy:处理获取组队副本信息完成")
	return nil
}

//处理组队副本界面信息逻辑
func teamCopyAllGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerTeamCopyDataManagerType).(*playerteamcopy.PlayerTeamCopyDataManager)
	teamCopyMap := manager.GetTeamCopyMap()
	scTeamCopyAllGet := pbutil.BuildSCTeamCopyAllGet(teamCopyMap)
	pl.SendMsg(scTeamCopyAllGet)
	return
}
