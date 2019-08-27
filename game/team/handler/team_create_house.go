package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	crosstypes "fgame/fgame/game/cross/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_TEAM_CREATE_HOUSE_TYPE), dispatch.HandlerFunc(handleTeamCreateHouse))
}

//处理组队创建房间信息
func handleTeamCreateHouse(s session.Session, msg interface{}) (err error) {
	log.Debug("team:处理组队创建房间信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTeamCreateHouse := msg.(*uipb.CSTeamCreateHouse)
	purpose := csTeamCreateHouse.GetPurpose()

	err = teamCreateHouse(tpl, teamtypes.TeamPurposeType(purpose))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("team:处理组队创建房间信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("team:处理组队创建房间信息完成")
	return nil

}

//附近队伍信息的逻辑
func teamCreateHouse(pl player.Player, purpose teamtypes.TeamPurposeType) (err error) {
	if !purpose.Vaild() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"purpose":  purpose,
			}).Warn("team:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//玩家已经在队伍中了
	if pl.GetTeamId() != 0 &&
		pl.GetTeamPurpose() == purpose {
		return
	}

	if pl.GetCrossType() == crosstypes.CrossTypeTeamCopy {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"purpose":  purpose,
			}).Warn("team:您当前已在别的房间中")
		playerlogic.SendSystemMessage(pl, lang.TeamCreateHouseInOther)
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

	//走创建队伍
	if pl.GetTeamId() == 0 {
		teamData, err := team.GetTeamService().CreateTeamByPlayer(pl, purpose)
		if err != nil {
			return err
		}
		teamId := teamData.GetTeamId()
		teamName := teamData.GetTeamName()
		pl.SyncTeam(teamId, teamName, purpose)
		scTeamCreateHouse := pbutil.BuildSCTeamCreateHouse(int32(purpose), teamData)
		pl.SendMsg(scTeamCreateHouse)

		//推送队员信息
		scTeamGet := pbutil.BuildSCTeamGet(teamData, false, pl.GetId())
		pl.SendMsg(scTeamGet)
		return nil
	}

	//切换组队标识
	_, err = team.GetTeamService().TeamChangePurpose(pl, purpose)
	if err != nil {
		return
	}
	return
}
