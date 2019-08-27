package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/funcopen/funcopen"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
	teamtypes "fgame/fgame/game/team/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_HOUSES_GET_TYPE), dispatch.HandlerFunc(handleTeamHousesGet))
}

//处理组队标识所有队伍信息
func handleTeamHousesGet(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理组队标识所有队伍信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamHousesGet := msg.(*uipb.CSTeamHousesGet)
	purpose := csTeamHousesGet.GetPurpose()

	err = teamHousesGet(tpl, teamtypes.TeamPurposeType(purpose))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("team:处理组队标识所有队伍信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理组队标识所有队伍信息完成")
	return nil

}

//附近队伍信息的逻辑
func teamHousesGet(pl player.Player, purpose teamtypes.TeamPurposeType) (err error) {
	if !purpose.Vaild() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"purpose":  purpose,
			}).Warn("team:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if pl.GetTeamId() != 0 && pl.GetTeamPurpose() == purpose {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"purpose":  purpose,
			}).Warn("team:您当前的队伍已是房间信息")
		playerlogic.SendSystemMessage(pl, lang.TeamSelfPuroseIsHouseInfo)
		return
	}

	openType := purpose.GetFuncOpenType()
	funcOpenTemplate := funcopen.GetFuncOpenService().GetFuncOpenTemplate(openType)
	if funcOpenTemplate == nil {
		return
	}

	if pl.GetLevel() < funcOpenTemplate.OpenedLevel {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"purpose":  purpose,
			}).Warn("team:功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	teamList := team.GetTeamService().GetTeamsByPurpose(pl, purpose)
	scTeamHousesGet := pbutil.BuildSCTeamHousesGet(int32(purpose), teamList)
	pl.SendMsg(scTeamHousesGet)
	return
}
