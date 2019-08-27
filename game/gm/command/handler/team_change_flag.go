package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/team/pbutil"
	team "fgame/fgame/game/team/team"
	teamtypes "fgame/fgame/game/team/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTeamChangePurpose, command.CommandHandlerFunc(handleTeamChangePurpose))

}

func handleTeamChangePurpose(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	purposeStr := c.Args[0]
	purposeInt64, err := strconv.ParseInt(purposeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"purposeStr": purposeStr,
				"error":      err,
			}).Warn("gm:处理设置队伍标识,flag不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	purpose := teamtypes.TeamPurposeType(purposeInt64)
	if !purpose.Vaild() {
		return
	}
	teamId := pl.GetTeamId()
	if teamId != 0 && pl.GetTeamPurpose() == purpose {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"teamId":  teamId,
				"purpose": purpose,
			}).Warn("gm:处理组队改变队伍标识")
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
