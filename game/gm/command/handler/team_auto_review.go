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
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTeamAutoReview, command.CommandHandlerFunc(handleTeamAutoReview))

}

func handleTeamAutoReview(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	reviewStr := c.Args[0]
	review, err := strconv.ParseInt(reviewStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"reviewStr": reviewStr,
				"error":     err,
			}).Warn("gm:处理设置队伍审核,review不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	teamId := pl.GetTeamId()
	if teamId == 0 {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"teamId": teamId,
			}).Warn("gm:处理设置队伍审核,review不是数字")
		return
	}
	autoReview := false
	if review != 0 {
		autoReview = true
	}

	err = team.GetTeamService().AutoReviewChoose(pl, autoReview)
	if err != nil {
		return
	}

	scTeamAutoReview := pbutil.BuildSCTeamAutoReview(autoReview)
	pl.SendMsg(scTeamAutoReview)
	return
}
