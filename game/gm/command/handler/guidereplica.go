package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	guidereplicalogic "fgame/fgame/game/guidereplica/logic"
	"fgame/fgame/game/guidereplica/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	questtemplate "fgame/fgame/game/quest/template"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeGuideReplica, command.CommandHandlerFunc(handleGuideReplica))
}

func handleGuideReplica(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	log.Debug("gm:处理进入引导副本")
	if len(c.Args) < 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	questIdStr := c.Args[0]

	questId, err := strconv.ParseInt(questIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"mapId": questId,
				"error": err,
			}).Warn("gm:处理引导副本任务,类型questId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	err = enterGuideReplica(pl, int32(questId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"questId": questId,
				"error":   err,
			}).Warn("gm:处理跳转引导副本,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":      pl.GetId(),
			"questId": questId,
		}).Debug("gm:处理跳转引导副本,完成")
	return
}

func enterGuideReplica(pl player.Player, questId int32) (err error) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("guidereplica:引导副本挑战请求，任务不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	guideReplicaTemplate := questTemplate.GetGuideReplicaTemplate()
	if guideReplicaTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("guidereplica:引导副本挑战请求，引导副本不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	//进入场景
	flag := guidereplicalogic.PlayerEnterGuideReplica(pl, guideReplicaTemplate, questId)
	if !flag {
		panic("enter guidereplica scene should be ok!")
	}
	scMsg := pbutil.BuildSCGuideReplicaChallenge()
	pl.SendMsg(scMsg)
	return
}
